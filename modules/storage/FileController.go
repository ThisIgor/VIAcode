package storage

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"knowledge/config"
	"net/http"
	"os"
	"strconv"
)

type UploadFromFrontend struct {
	Comment        string
	FileName       string // Только имя
	DocumentNumber string
	Content        []byte
}

type EditFromFrontend struct {
	FileId  uint
	Comment string
}

type GetInfoFromFrontend struct {
	FileId uint
}

type GetInfoFromFrontendResult struct {
	FileId           uint
	Comment          string
	Path             string
	Filesize         int64
	DocumentNumber   string
	ModificationTime int64
}

//http://locahost:3000/uploadmedia?body= { "Comment": "111",  "FileName": "file.txt", "DocumentNumber": "1111", "Content": "Content"}
func UploadMedia(c *gin.Context) {
	data := &UploadFromFrontend{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}

	uid, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Bad requests params"})
		return
	}

	filename := config.Config.MediaStorage.RootPath + data.DocumentNumber + uid.String() + data.FileName
	os.MkdirAll(config.Config.MediaStorage.RootPath+data.DocumentNumber, os.ModeDir)

	attachment, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Can not create file"})
		return
	}
	defer attachment.Close()

	_, err = attachment.Write(data.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError, "message": "Can not create file"})
		return
	}

	u := Upload{Path: filename, Comment: data.Comment, DocumentNumber: data.DocumentNumber}
	CreateFileInfo(&u)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "File uploaded"})
}

//http://locahost:3000/getinfofromfrontend?body={"FileId": "11"}
func GetMediaInfo(c *gin.Context) {
	mediaid, err := strconv.Atoi(c.Param("articleid"))

	//	data := &GetInfoFromFrontend{}
	//	err := c.BindJSON(&data)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}

	filedata := GetFileInfo(uint(mediaid))
	fileInfo, err := os.Stat(filedata.Path)

	c.JSON(http.StatusOK,
		gin.H{"status": http.StatusOK,
			"fieldid":          mediaid,
			"comment":          filedata.Comment,
			"path":             filedata.Path,
			"filesize":         fileInfo.Size(),
			"documentnumber":   filedata.DocumentNumber,
			"modificationtime": fileInfo.ModTime().Unix()})
}

func UpdateMedia(c *gin.Context) {
	data := &EditFromFrontend{}
	err := c.BindJSON(&data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Bad requests params"})
		return
	}
	UpdateFileInfo(data.FileId, data.Comment)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "File updated"})
}
