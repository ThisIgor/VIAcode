package article

import (
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/gin-gonic/gin"
	"knowledge/database"
	"net/http"
	"strconv"
)

/*
	POST - принимает json
	{
		fullName     string
		description  string
		heading_id   int
		group_id	 int
	}
*/
func AddArticleType(c *gin.Context) {
	articleType := &ArticleType{}
	err := c.BindJSON(&articleType)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	res := CreateArticleType(articleType.FullName, articleType.Description, 0, 0)
	if res.Error != nil {
		c.JSON(http.StatusAlreadyReported, gin.H{"status": http.StatusAlreadyReported, "message": res.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Article type success created", "data": articleType})
}

/*
	POST - На вход json
	{
		blackword     	string
		article_type_id int
		content       	string
		permissions   uint TODO:: под вопросом
	}
*/
func CreateArticle(c *gin.Context) {
	article := &Article{}
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	res, id := InsertArticle(article)
	if res != nil {
		c.JSON(http.StatusAlreadyReported, gin.H{"status": http.StatusAlreadyReported, "message": res})
		return
	}

	doc := Document{DBUid: id, Bkackword: article.Bkackword, ArticleTypeID: article.ArticleTypeID, Permissions: article.Permissions, Content: article.Content}
	CreateDocument(doc)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Article success created", "data": article})
}

func UpdateArticle(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("articleid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	db := database.Connect()
	defer db.Close()

	article := &Article{}
	res1 := db.First(&article, articleId)
	if res1.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Record not found", "data": article})
		return
	}

	err = c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params", "data": article})
		return
	}
	res := db.Save(&article)
	if res.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Article not found", "data": article})
		return
	}
	UpdateDocument(0, article)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Article success updated", "data": article})
}

func DeleteArticle(c *gin.Context) {
	db := database.Connect()
	defer db.Close()

	articleId, err := strconv.Atoi(c.Param("articleid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	article := &Article{}
	res1 := db.First(&article, articleId)
	if res1.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": "Record not found", "data": article})
		return
	}
	res := db.Where("id = ?", articleId).Delete(&Article{})
	if res.Error != nil {
		c.JSON(500, gin.H{"status": 500, "message": "Something went wrong"})
		return
	}

	DeleteDocument(articleId)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Article success deleted"})
}

/*
	GET - на вход query параметры по которым строится WHERE запроса
*/
func GetAllArticles(c *gin.Context) {
	// TODO:: Принимаем get параметры и формируем запрос в базу, когда будет ясно какие параметры
	db := database.Connect()
	defer db.Close()

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "100"))

	var articles []Article
	query := db.Find(&articles)
	pagination.Paging(&pagination.Param{
		DB:      query,
		Page:    page,
		Limit:   limit,
		OrderBy: []string{"id desc"},
	}, &articles)

	if len(articles) <= 0 {
		var res []string
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "data": res})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": articles})
}

/*
	GET - Возвращае article по id
*/
func GetArticle(c *gin.Context) {
	articleId, err := strconv.Atoi(c.Param("articleid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	article, db := FetchArticle(articleId)
	if db.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{"status": http.StatusNoContent, "message": db.Error.Error(), "data": article})
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": article})
}

/*
	GET - принимает  строку которую надо искать
*/
func FindArticle(c *gin.Context) {
	query := c.Param("query")

	articles, res := FindDocument(query)

	if res != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": res})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Articles found", "data": articles})
}
