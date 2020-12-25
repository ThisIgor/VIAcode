package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	conf "knowledge/config"
	"knowledge/modules/article"
	"knowledge/modules/articleAttribute"
	"knowledge/modules/base"
	"knowledge/modules/rbac"
	"knowledge/modules/user"
	logging "logger"
	"net/http"
)

const (
	// Routing
	RootRout = "/api/encyclopedia"

	UserGroupRoute          = "/user"
	UserGetUserListRoute    = "/"
	UserGetUserRoute        = "/:userid"
	UserChangeUserDataRoute = "/:userid"
	UserDeleteUserRoute     = "/:userid"

	ArticleAttributeGroupRoute  = "/articleattribute"
	ArticleAttributeListRoute   = "/:articleid"
	ArticleAttributeRoute       = "/:articleid/:attrid"
	ArticleAttributeCreateRoute = "/"
	ArticleAttributeChangeRoute = "/:attrid"
	ArticleAttributeDeleteRoute = "/:attrid"
	ArticleAttributeLinkRoute   = "/bindattribute"

	ArticleRootRoute           = "/article"
	ArticleCreateRoute         = "/"
	ArticleListRoute           = "/"
	ArticleRoute               = "/:articleid"
	ArticleChangeRoute         = "/:articleid"
	ArticleDeleteRoute         = "/:articleid"
	ArticleFindarticleRoute    = "/find"
	ArticleChangeArticlesRoute = "/"
	ArticledeleteArticlesRoute = "/"

	ArticleTypeRootRoute   = "/articletype"
	ArticleTypeCreateRoute = "/"

	RoleGroupRoute      = "/role"
	RoleCreateRoute     = "/"
	RoleChangeDataRoute = "/:roleid"
	RoleGetRolesRoute   = "/:roleid"
	//	RoleGetAllPernissionsRoute = "/permissions/"

	RoleGetDataRoute = "/:roleid"

	// Role mapping
	UserMapRoute             = RootRout + UserGroupRoute
	ArticleAttributeMapRoute = RootRout + ArticleAttributeGroupRoute
	ArticleMapRoute          = RootRout + ArticleRootRoute
	ArticleTypeMapRoute      = RootRout + ArticleTypeRootRoute
	RoleMapRoute             = RootRout + RoleGroupRoute
	// Access groups ID
	UserGroupID             = 1
	ArticleAttributeGroupID = 2
	ArticleGroupID          = 3
	AtricleTypeGroupID      = 4
	RoleGroupID             = 5
)

func Init() {
	//	InitPermissionCategories()
	//	InitPermissions()

	router := gin.Default()
	config := cors.DefaultConfig()      // Настройка cors, пока для простоты *
	config.AllowOrigins = []string{"*"} //FIXME:: Заменить на нужные хосты
	router.Use(cors.New(config))
	router.GET("/", base.IndexAction)
	router.POST("/signin", user.Signin)
	router.POST("/signup", user.Signup)
	enc := router.Group(RootRout) //.Use(user.ValidateSession).Use(ValidateAccess)
	{
		roleAttr := enc.Group(RoleGroupRoute)
		{
			roleAttr.POST(RoleCreateRoute, rbac.AddRole)       // Создание новой роли
			roleAttr.PUT(RoleChangeDataRoute, rbac.ChangeRole) // Изммененме роли
			roleAttr.GET(RoleGetRolesRoute, rbac.GetRoles)     // Получение списка ролей

			//			roleAttr.GET(RoleGetDataRoute, rbac.GetRole)       // Получение ролей и прав
		}
		userGroup := enc.Group(UserGroupRoute)
		{			
			userGroup.GET(UserGetUserListRoute, user.GetUserlist)       // Список всех пользователей
			userGroup.GET(UserGetUserRoute, user.GetUser)               // Данные пользователя
			userGroup.PUT(UserChangeUserDataRoute, user.ChangeUserData) // Изменение данных пользователя
			userGroup.DELETE(UserDeleteUserRoute, user.DeleteUser)      // Удаление пользователя
		}
		artAttr := enc.Group(ArticleAttributeGroupRoute)
		{
			artAttr.GET(ArticleAttributeListRoute, articleAttribute.GetAllArticleAttributes)     // Список атрибутов статьи
			artAttr.GET(ArticleAttributeRoute, articleAttribute.GetArticleAttribute)             // Конкретный аттрибут статьи по id
			artAttr.POST(ArticleAttributeCreateRoute, articleAttribute.CreateArticleAttribute)   // Создание нового атрибута
			artAttr.PUT(ArticleAttributeChangeRoute, articleAttribute.ChangeArticleAttribute)    // Обновление атрибута
			artAttr.DELETE(ArticleAttributeDeleteRoute, articleAttribute.RemoveArticleAttribute) // Удаление атрибута
			//artAttr.POST(ArticleAttributeLinkRoute, articleAttribute.BindAttributeToArticle) // Привязка атрибута к статье
		}
		art := enc.Group(ArticleRootRoute)
		{
			//art.POST(ArticleCreateRoute, article.CreateArticle) // Создание статьи
			art.GET(ArticleListRoute, article.GetAllArticles) // Список статей
			art.GET(ArticleRoute, article.GetArticle)         // Конкретная статья
			art.PUT(ArticleChangeArticlesRoute, MethodNotAllowed)
			art.DELETE(ArticledeleteArticlesRoute, MethodNotAllowed)

			art.PUT(ArticleChangeRoute, article.UpdateArticle)    // Обновить статью
			art.DELETE(ArticleDeleteRoute, article.DeleteArticle) // Удалить статью

			art.POST(ArticleFindarticleRoute, article.FindArticle) // Найти статьи
		}
		artType := enc.Group(ArticleTypeRootRoute)
		{
			artType.POST(ArticleTypeCreateRoute, article.AddArticleType) // Создание типа статьи
			//artType.POST("/", article.DeleteArticleType) 					// Удаление типа статьи
		}
	}

	var err error = nil
	if conf.Config.Routing.Type == "http" {
		err = router.Run(conf.Config.Routing.Http)
	} else {
		err = router.RunTLS(conf.Config.Routing.Https, conf.Config.Routing.CertFile, conf.Config.Routing.KeyFile)
	}
	if err != nil {
		logging.DefaultLog.Critical("Can't start server. %v.", err)
	}
}

func MethodNotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, gin.H{"status": http.StatusMethodNotAllowed})
}

var CurrentSession user.Session
