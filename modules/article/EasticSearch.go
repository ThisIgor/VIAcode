package article

import (
	"context"
	"encoding/json"
	"github.com/olivere/elastic/v7"
	"knowledge/config"
	"strconv"
)

type Document struct {
	DBUid         uint   `json:"dbuid"`
	Bkackword     string `json:"blackword"`
	ArticleTypeID uint   `json:"article_type_id"`
	Permissions   uint   `json:"permissions"`
	Content       string `json:"content"`
}

const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"article":{
			"properties":{
				"dbid": {
					"type":"integer"
				}
				"blackword":{
					"type":"keyword"
				},
				"article_type_id":{
					"type":"integer"
				},
				"permissions":{
					"type":"integer"
				},
				"content":{
					"type":"text"
					"store": true,
					"fielddata": true				
				}
			}
		}
	}
}
`

func CreateDocument(document Document) error {
	client, err := elastic.NewClient(elastic.SetURL(config.Config.ElasticSearch.Address))
	if err != nil {
		return err
	}

	exists, err := client.IndexExists("content1").Do(context.Background())
	if err != nil {
		return err
	}

	if exists == false {
		_, err = client.CreateIndex("content1").BodyString(mapping).Do(context.Background())
		if err != nil {
			return err
		}
	}

	_, err = client.Index().
		Index("content").
		Type("article").
		Id(strconv.Itoa(int(document.DBUid))).
		BodyJson(document).
		Do(context.Background())
	if err != nil {
		return err
	}

	_, err = client.Flush().Index("content").Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func FindDocument(query string) (error, map[uint]Article) {
	client, err := elastic.NewClient(elastic.SetURL(config.Config.ElasticSearch.Address))
	if err != nil {
		return err, nil
	}

	blackwordQuery := elastic.NewMatchQuery("blackword", query)
	contentQuery := elastic.NewMatchQuery("content", query)
	searchResult, err := client.Search().
		Index("content").
		Type("article").
		Query(blackwordQuery).
		Query(contentQuery).
		Do(context.Background())
	if err != nil {
		return err, nil
	}

	documents := make(map[uint]Article)
	if searchResult.Hits.TotalHits.Value > 0 {
		for _, hit := range searchResult.Hits.Hits {
			var document Document
			err := json.Unmarshal(hit.Source, &document)
			if err != nil {
				continue
			}

			article := Article{Bkackword: document.Bkackword, ArticleTypeID: document.ArticleTypeID, Permissions: document.Permissions, Content: document.Content}
			documents[document.DBUid] = article
		}
	}

	return nil, documents
}

func UpdateDocument(id uint, document *Article) error {
	client, err := elastic.NewClient(elastic.SetURL(config.Config.ElasticSearch.Address))
	if err != nil {
		return err
	}

	_, err = client.Update().
		Index("content").
		Type("article").
		Id(strconv.Itoa(int(id))).
		Upsert(map[string]interface{}{"Bkackword": document.Bkackword, "ArticleTypeID": document.ArticleTypeID, "Permissions": document.Permissions, "Content": document.Content}).
		Do(context.Background())

	return err
}

func DeleteDocument(id int) error {
	client, err := elastic.NewClient(elastic.SetURL(config.Config.ElasticSearch.Address))
	if err != nil {
		return err
	}

	_, err = client.Delete().
		Index("content").
		Type("article").
		Id(strconv.Itoa(id)).
		Do(context.Background())

	return err
}
