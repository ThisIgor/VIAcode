package rbac

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"knowledge/database"
	"net/http"
	"reflect"
	"strconv"
)

func InitRoles() {
	//	CreateRole("Администратор", RoleAdmin)
	//	CreateRole("Редактор", RoleDocMaster)
}

func parceJSON(c *gin.Context, db *gorm.DB) (*Role, error) {
	ferole := &FERole{}
	err := c.BindJSON(&ferole)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return nil, err
	}

	// Выбираем все присланные разрешения
	permissions := []Permission{}
	var pids string
	for _, value := range ferole.PermissionsID {
		pids += strconv.Itoa(int(value))
		pids += ","
	}
	err = db.Model(&Permission{}).Where("id IN (" + pids[:len(pids)-1] + ")").Find(&permissions).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return nil, err
	}

	// Делае OR для всех разрешениям по категориям
	catpermissions := make(map[uint]uint64)
	for _, value := range permissions {
		catpermissions[value.CategoryID] |= value.Permission
	}

	// Готовим роль
	rl := &Role{}
	rl.Name = ferole.Name
	rl.Permissions = make([]RolePermisson, len(catpermissions))
	rl.PermissionsID = pids[:len(pids)-1]

	keys := reflect.ValueOf(catpermissions).MapKeys()
	for i := int(0); i < len(catpermissions); i++ {
		rl.Permissions[i].CategoryID = uint(keys[i].Uint())
		rl.Permissions[i].Permissions = catpermissions[uint(keys[i].Uint())]
	}

	ba, err := json.Marshal(rl.Permissions)
	if err != nil {
		return nil, err
	}
	rl.PermissionsStr = string(ba)

	return rl, nil
}

func AddRole(c *gin.Context) {
	db := database.Connect()
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Can't open database"})
		return
	}
	defer db.Close()

	rl, err := parceJSON(c, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}

	err = db.Model(&Role{}).Create(&rl).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Can't add role"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Role added"})
	}
}

func ChangeRole(c *gin.Context) {
	roleid, err := strconv.Atoi(c.Param("roleid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}

	db := database.Connect()
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Can't open database"})
		return
	}
	defer db.Close()

	rl, err := parceJSON(c, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}

	err = db.Model(&Role{}).Where("id = ?", roleid).Update(&rl).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Can't change role"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Role changed", "data": rl})
	}
}

type PermissionsforFront struct {
	ID   uint
	Name string
}

type RolesForFront struct {
	ID          uint
	Name        string
	Permissions []PermissionsforFront
}

func GetRoles(c *gin.Context) {
	roleid, paramerr := strconv.Atoi(c.Param("roleid"))
	db := database.Connect()
	if db.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Can't open database"})
		return
	}
	defer db.Close()

	roles := []Role{}
	var err error
	if paramerr != nil {
		err = db.Model(&Role{}).Find(&roles).Error
	} else {
		err = db.Model(&Role{}).Where("id = ?", roleid).Find(&roles).Error
	}
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Can't find roles"})
	} else {
		rolesforfront := make([]RolesForFront, len(roles))
		for i := int(0); i < len(roles); i++ {
			rolesforfront[i].ID = roles[i].ID
			rolesforfront[i].Name = roles[i].Name
			perms := []Permission{}
			err = db.Model(&Permission{}).Where("id IN (" + roles[i].PermissionsID + ")").Find(&perms).Error
			for _, permid := range perms {
				rolesforfront[i].Permissions = append(rolesforfront[i].Permissions, PermissionsforFront{ID: permid.ID, Name: permid.Name})
			}
		}

		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Roles acquired", "data": rolesforfront})
	}
}
