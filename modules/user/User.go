package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var CurrentSession Session

/*
	GET - получение информции о всех пользователелях
*/
func GetUserlist(c *gin.Context) {

	users := GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Users found", "data": users})
	//	c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Users not accessible for this user"})
}

/*
	GET - получение информции о пользователеле
	принимает на вход параметр
	{userid: "uint"}
*/
func GetUser(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}

	user, db := GetUserFromDB(uint(userid))
	if db.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found", "data": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User found", "data": user})
	}
	//	c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Attribute group not accessible for this user"})
}

/*
	PUT - изменение информции о пользователеле
	принимает на вход json
	{GetUserData struct}
*/
func ChangeUserData(c *gin.Context) {
	userid, err := strconv.Atoi(c.Param("userid"))
	user := &User{}
	err = c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}

	newuser, db := ChangeUser(uint(userid), *user)
	if db.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found", "data": newuser})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User data chaged", "data": newuser})
	}
}

/*
	DELETE - удаление пользователеля
	принимает на вход параметр
	{userid: "uint"}
*/
func DeleteUser(c *gin.Context) {

	userid, err := strconv.Atoi(c.Param("userid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}

	user, db := DeleteUserFromDB(uint(userid))
	if db.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "User not found", "data": user})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User killed and eaten", "data": user})
	}
}
