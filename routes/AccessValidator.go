package routes

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"knowledge/database"
	"knowledge/modules/rbac"
	"net/http"
)

const (
	GetUserListMask    = 0b00000001
	GetUserMask        = 0b00000010
	ChangeUserDataMask = 0b00000100
	DeleteUserMask     = 0b00001000

	ArticleAttributeListMask   = 0b00000001
	ArticleAttributeMask       = 0b00000010
	ArticleAttributeCreateMask = 0b00000100
	ArticleAttributeChangeMask = 0b00001000
	ArticleAttributeDeleteMask = 0b00010000
	ArticleAttributeLinkMask   = 0b00100000

	ArticleCreateMask         = 0b00000001
	ArticleListMask           = 0b00000010
	ArticleMask               = 0b00000100
	ArticleChangeArticlesMask = 0b00001000
	ArticledeleteArticlesMask = 0b00010000
	ArticleChangeMask         = 0b00100000
	ArticleDeleteMask         = 0b01000000
	ArticleFindarticleMask    = 0b10000000

	ArticleTypeCreateMask = 0b00000001

	RoleCreateMask     = 0b00000001
	RoleChangeDataMask = 0b00000010
	RoleGetDataMask    = 0b00000100
)

func ValidateAccess(c *gin.Context) {
	CurrentSession.RoleID = 130

	db := database.Connect()

	rl := &rbac.Role{}
	err := db.Model(&rbac.Role{}).Where("id = ?", CurrentSession.RoleID).First(&rl).Error
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(rl.PermissionsStr), &rl.Permissions)
	if err != nil {
		return
	}

	rt := Route{Type: c.Request.Method, Path: c.Request.RequestURI}
	for _, value := range rl.Permissions {
		right, exist := AccessRights[value.CategoryID]
		if exist == false {
			continue
		}
		permission, exist := right[rt]
		if exist == false {
			continue
		}
		if permission&value.Permissions != 0 {
			c.Next()
			return
		}
	}
	c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Forbidden"})
}

type Route struct {
	Type, Path string
}

var AsseccUserGroup = map[Route]uint64{
	Route{Type: "GET", Path: UserMapRoute + UserGetUserListRoute}:    GetUserListMask,
	Route{Type: "GET", Path: UserMapRoute + UserGetUserRoute}:        GetUserMask,
	Route{Type: "PUT", Path: UserMapRoute + UserChangeUserDataRoute}: ChangeUserDataMask,
	Route{Type: "DELETE", Path: UserMapRoute + UserDeleteUserRoute}:  DeleteUserMask,
}

var ArticleAttributeGroup = map[Route]uint64{
	Route{Type: "GET", Path: ArticleAttributeMapRoute + ArticleAttributeListRoute}:      ArticleAttributeListMask,
	Route{Type: "GET", Path: ArticleAttributeMapRoute + ArticleAttributeRoute}:          ArticleAttributeMask,
	Route{Type: "POST", Path: ArticleAttributeMapRoute + ArticleAttributeCreateRoute}:   ArticleAttributeCreateMask,
	Route{Type: "PUT", Path: ArticleAttributeMapRoute + ArticleAttributeChangeRoute}:    ArticleAttributeChangeMask,
	Route{Type: "DELETE", Path: ArticleAttributeMapRoute + ArticleAttributeDeleteRoute}: ArticleAttributeDeleteMask,
	Route{Type: "POST", Path: ArticleAttributeMapRoute + ArticleAttributeLinkRoute}:     ArticleAttributeLinkMask,
}

var ArticleGroup = map[Route]uint64{
	Route{Type: "POST", Path: ArticleMapRoute + ArticleCreateRoute}:           ArticleCreateMask,
	Route{Type: "GET", Path: ArticleMapRoute + ArticleListRoute}:              ArticleListMask,
	Route{Type: "GET", Path: ArticleMapRoute + ArticleRoute}:                  ArticleMask,
	Route{Type: "PUT", Path: ArticleMapRoute + ArticleChangeArticlesRoute}:    ArticleChangeArticlesMask,
	Route{Type: "DELETE", Path: ArticleMapRoute + ArticledeleteArticlesRoute}: ArticledeleteArticlesMask,
	Route{Type: "PUT", Path: ArticleMapRoute + ArticleChangeRoute}:            ArticleChangeMask,
	Route{Type: "DELETE", Path: ArticleMapRoute + ArticleDeleteRoute}:         ArticleDeleteMask,
	Route{Type: "POST", Path: ArticleMapRoute + ArticleFindarticleRoute}:      ArticleFindarticleMask,
}

var AtricleTypeGroup = map[Route]uint64{
	Route{Type: "POST", Path: ArticleTypeMapRoute + ArticleTypeCreateRoute}: ArticleTypeCreateMask,
}

var RoleGroup = map[Route]uint64{
	Route{Type: "POST", Path: RoleMapRoute + RoleCreateRoute}:    RoleCreateMask,
	Route{Type: "PUT", Path: RoleMapRoute + RoleChangeDataRoute}: RoleChangeDataMask,
	Route{Type: "GET", Path: RoleMapRoute + RoleGetDataRoute}:    RoleGetDataMask,
}

var AccessRights = map[uint]map[Route]uint64{
	UserGroupID:             AsseccUserGroup,
	ArticleAttributeGroupID: ArticleAttributeGroup,
	ArticleGroupID:          ArticleGroup,
	AtricleTypeGroupID:      AtricleTypeGroup,
	RoleGroupID:             RoleGroup,
}

func InitPermissionCategories() {
	cats := []rbac.PermissionCategory{{CategoryID: UserGroupID, Name: "Пользователи"},
		{CategoryID: ArticleAttributeGroupID, Name: "Аттрибуты статей"},
		{CategoryID: ArticleGroupID, Name: "Статьи"},
		{CategoryID: AtricleTypeGroupID, Name: "Группы статей"},
		{CategoryID: RoleGroupID, Name: "Роли"}}

	db := database.Connect()
	defer db.Close()

	db.Model(&rbac.PermissionCategory{}).Delete(&rbac.PermissionCategory{})
	for _, cat := range cats {
		db.Create(&cat)
	}
}

func InitPermissions() {
	db := database.Connect()
	defer db.Close()

	db.Model(&rbac.Permission{}).Delete(&rbac.Permission{})

	for arkey, arval := range AccessRights {
		for _, permval := range arval {
			perm := rbac.Permission{CategoryID: uint(arkey), Name: fmt.Sprintf("%v : %v", uint(arkey), permval), Permission: permval}
			err := db.Model(&rbac.Permission{}).Create(&perm).Error
			if err != nil {
				return // TODO:: Добавить логировагие
			}
		}
	}
}
