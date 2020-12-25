package migrations

import (
	"knowledge/database"
	"knowledge/modules/article"
	"knowledge/modules/articleAttribute"
	"knowledge/modules/editor"
	"knowledge/modules/exchangeTask"
	"knowledge/modules/rbac"
	"knowledge/modules/storage"
	"knowledge/modules/user"
)

func MigrateAll() {
	db := database.Connect() // Коннект к базе и накатывание миграций
	db.AutoMigrate(&articleAttribute.ArticleAttribute{}, &articleAttribute.ArticleAttribute{},
		&editor.EditorType{}, &editor.Editor{},
		&rbac.Role{}, &user.User{}, &user.Session{},
		&rbac.PermissionCategory{}, &rbac.Permission{},
		&article.ArticleType{}, &article.Article{}, &article.ArticleTemplate{}, article.ArticleTypeGroup{},
		&storage.Upload{},
		&exchangeTask.EditorialOffer{}, &exchangeTask.EditorialType{}, &exchangeTask.EditorialOffice{}, &exchangeTask.EditorialMart{})

	data := db.Model(&article.Article{}).AddForeignKey("article_type_id", "article_types(id)", "RESTRICT", "RESTRICT")
	//data = data.Model(&article.ArticleAttribute{}).AddForeignKey("article_type_id", "articles_types(id)", "RESTRICT", "RESTRICT")
	//data = data.Model(&article.ArticleAttribute{}).AddForeignKey("article_id", "articles(id)", "RESTRICT", "RESTRICT")
	data = data.Model(&article.ArticleTemplate{}).AddForeignKey("article_type_id", "article_types(id)", "RESTRICT", "RESTRICT")

	data = data.Model(&article.ArticleTemplate{}).AddForeignKey("article_type_id", "article_types(id)", "RESTRICT", "RESTRICT")
	data = data.Model(&article.ArticleTemplate{}).AddForeignKey("article_type_id", "article_types(id)", "RESTRICT", "RESTRICT")

	data = data.Model(&editor.Editor{}).AddForeignKey("editor_type", "editor_types(id)", "RESTRICT", "RESTRICT")
	data = data.Model(&editor.Editor{}).AddForeignKey("article_id", "article_types(id)", "RESTRICT", "RESTRICT")

	data = data.Model(&exchangeTask.EditorialOffer{}).AddForeignKey("editorial_type", "editorial_types(id)", "RESTRICT", "RESTRICT")
	data = data.Model(&exchangeTask.EditorialOffer{}).AddForeignKey("editorial_type", "editorial_types(id)", "RESTRICT", "RESTRICT")
	data = data.Model(&exchangeTask.EditorialOffer{}).AddForeignKey("editorial_mart", "editorial_marts(id)", "RESTRICT", "RESTRICT")
	data = data.Model(&exchangeTask.EditorialOffer{}).AddForeignKey("article_type", "article_types(id)", "RESTRICT", "RESTRICT")
	data = data.Model(&exchangeTask.EditorialOffer{}).AddForeignKey("editorial_office_id", "editorial_offices(id)", "RESTRICT", "RESTRICT")
	data = data.Model(&exchangeTask.EditorialOffer{}).AddForeignKey("editor_id", "editors(id)", "RESTRICT", "RESTRICT")

	data = data.Model(&user.User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	data = data.Model(&user.Session{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")

}
