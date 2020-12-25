package articleAttribute

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
	Post
	Привязка атрибутов(ArticleAttribute) к статье (Article)
*/
func BindAttributeToArticle(c *gin.Context) {
	data := &RelArticleAttribute{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	AddArticleAttrib(data.AttributeID, data.ArticleID, data.Value)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Attribute added"})
	//	c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Attribute not accessible for this user"})
}

func GetArticleAttribute(c *gin.Context) {
	attributeid, err := strconv.Atoi(c.Param("attrid"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}

	err, name, value := GetArticleAttrib(uint(attributeid))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Attribute not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Attribute found", "name": name, "value": value})
	}
	//	c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Attribute not accessible for this user"})
}

// PUT
func ChangeArticleAttribute(c *gin.Context) {
	attributeid, err := strconv.Atoi(c.Param("attrid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	data := &ArticleAttribute{}
	err = c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	attrib, db := ChangeArticleAttrib(uint(attributeid), data)
	if db.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": db.Error.Error(), "data": attrib})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Attribute changed", "data": attrib})
	}
	//	c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Attribute not accessible for this user"})
}

// DELETE
func RemoveArticleAttribute(c *gin.Context) {
	attributeid, err := strconv.Atoi(c.Param("attrid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	attrrelation, db := DeleteArticleAttrib(uint(attributeid))
	if db.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": db.Error.Error(), "data": attrrelation})

	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Attribute deleted", "data": attrrelation})
	}
	//	c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Attribute not accessible for this user"})
}

// GET
func GetAllArticleAttributes(c *gin.Context) {
	articleid, err := strconv.Atoi(c.Param("articleid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	attributes, db := GetAllArticleAttrib(uint(articleid))
	if db.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Attributes not found", "data": attributes})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Attributes found", "data": attributes})
	}
	//	c.JSON(http.StatusForbidden, gin.H{"status": http.StatusForbidden, "message": "Attribute group not accessible for this user"})
}

/*
	POST - принимает на вход json
	{name: "String"}
*/
func CreateArticleAttribute(c *gin.Context) {
	articleAttribute := &ArticleAttribute{}
	err := c.BindJSON(&articleAttribute)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	err = CreateArticleAttrib(articleAttribute)
	if err != nil {
		c.JSON(http.StatusAlreadyReported, gin.H{"status": http.StatusAlreadyReported, "message": "Could not create attribute"})
	} else {
		c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Attribute created"})
	}
}

func ChangeArticleAttributeName(c *gin.Context) {
	attributeid, err := strconv.Atoi(c.Param("attributeid"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	attributename := c.Param("attributename")
	err = ChangeArticleAttribName(uint(attributeid), attributename)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "Attribute not found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Attribute changed"})
	}
}
