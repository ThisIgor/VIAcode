package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"knowledge/database"
	"net/http"
	"time"
)

// TODO:: Переписать с использованием JWT
func Signup(c *gin.Context) {
	user := &User{}
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	// хешируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword)

	db := database.Connect()
	res := CreateUser(user)
	if res.Error != nil {
		c.JSON(http.StatusConflict, gin.H{"status": http.StatusConflict, "message": res.Error})
		_ = db.Close()
		return
	}
	_ = db.Close()
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "User created success id: ", "resourceId": user.ID})
}

func Signin(c *gin.Context) {
	user := &User{}
	userDb := &User{}
	err := c.BindJSON(&user) // Парсим запрос в структуру User
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	// Ищем пользователя по login
	db := database.Connect()
	if err := db.Where("login = ?", user.Login).First(&userDb).Error; gorm.IsRecordNotFoundError(err) {
		//TODO:: затереть куку с сессией на всякий
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Login or password is incorrect"})
		_ = db.Close()
		return
	}
	// сравним пароль и хеш
	if err = bcrypt.CompareHashAndPassword([]byte(userDb.Password), []byte(user.Password)); err != nil {
		_ = db.Close()
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Login or password is incorrect"})
		return
	}
	//	session = &Session{}
	// Если сессия для этого пользователя уже есть, обновим её или создадим новую
	db.Where("user_id = ?", userDb.ID).First(&CurrentSession)
	st, _ := uuid.NewV4() // Создали новый рандомный токен сессии
	CurrentSession.UserID = uint(userDb.ID)
	CurrentSession.RoleID = uint(userDb.RoleID)
	CurrentSession.Token = st.String()
	CurrentSession.Expires = time.Now().Add(time.Minute * 10) // время жизни 10 мин
	res := db.Save(&CurrentSession)

	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
		_ = db.Close()
		return
	}
	c.SetCookie("session_token", CurrentSession.Token, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Login successful"})
}

func ValidateSession(c *gin.Context) {
	// Обязательно говорим браузеру не кешировать запрос, а то авторизация не будет работать
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	cookieToken, err := c.Cookie("session_token")
	if err != nil {
		fmt.Println("Cookie parse error", err)
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Unauthorized"})
		c.Abort()
		return
	}
	session := &Session{}
	db := database.Connect()
	err = db.Where("token = ?", cookieToken).First(&session).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) { // сессия не найдена
			c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Unauthorized"})
			c.Abort()
			return
		} else { // Любая другая ошибка
			fmt.Println(err)
			c.Abort()
			return
		}
	}

	// Сессия нашлась, проверим время жизни
	if time.Now().Unix() > session.Expires.Unix() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Session expired"})
		c.Abort()
		return
	}
	c.Next() // Сессия есть и не протухла, идём дальше
}
