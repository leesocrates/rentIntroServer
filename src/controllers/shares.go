package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ShareController struct {

}

func (*ShareController) Get(c *gin.Context) {


	c.Status(http.StatusOK)
}

func (*ShareController) Put(c *gin.Context) {


	c.Status(http.StatusOK)
}