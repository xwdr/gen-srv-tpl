package tmpl

var MainSRV = `package main

import (
	"fmt"
	"github.com/{{.Group}}/{{.Project}}/routers"
	"github.com/{{.Group}}/{{.Project}}/models"
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func init() {
	models.Setup()
}

func run() {
	app := gin.New()
	routers.Service(app)

	readTimeout := time.Duration(60)*time.Second
	writeTimeout := time.Duration(60)*time.Second
	endPoint := fmt.Sprintf(":%d", 8080)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        app,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	logrus.Printf("[info] start http server listening %s", endPoint)
	server.ListenAndServe()
}

func main() {
	run() 
}

`
