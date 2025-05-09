package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "GetNotes working!"})
}

func CreateNote(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "CreateNote working!"})
}
