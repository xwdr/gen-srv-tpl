package tmpl

var Common = `package controllers

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code      int         {{.Quote}}json:"code"{{.Quote}}
	ErrorCode int         {{.Quote}}json:"errorCode,omitempty"{{.Quote}}
	ErrorDesc string      {{.Quote}}json:"errorDesc,omitempty"{{.Quote}}
	Data      interface{} {{.Quote}}json:"data,omitempty"{{.Quote}}
}

func (g *Gin) ResponseOk(httpCode int, data interface{}) {
	var rsp Response
	rsp.Code = 0 // 成功
	rsp.ErrorCode = 0
	rsp.ErrorDesc = "ok"
	rsp.Data = data
	g.C.JSON(httpCode, &rsp)
}

func (g *Gin) ResponseFail(httpCode, errCode int, data interface{}) {
	var rsp Response
	rsp.Code = 1 // 失败
	rsp.ErrorCode = errCode
	rsp.ErrorDesc = "fail"
	rsp.Data = data
	g.C.JSON(httpCode, &rsp)
}

type User struct {
	ID      int64   {{.Quote}}json:"id"{{.Quote}}
	Name    string  {{.Quote}}json:"name"{{.Quote}}
	Address string  {{.Quote}}json:"address"{{.Quote}}
	Age     uint8   {{.Quote}}json:"age"{{.Quote}}
}

func (u *User) Invalid() bool {
	if len(u.Name) == 0 {
		return true
	}

	return false
}

type GetUsersRsp struct {
	User      User  {{.Quote}}json:"info"{{.Quote}}
}

type CreateUserReq struct {
	User
}
type CreateUserRsp struct {
	Id int64   {{.Quote}}json:"id"{{.Quote}}
}

`

var Controllers = `package controllers

import (
	"github.com/{{.Group}}/{{.Project}}/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context) {
	var (
		rsp GetUsersRsp
		g = Gin{C: c}
	)
	id,_ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		g.ResponseFail(http.StatusOK, 400, nil)
		return 
	}

	po := models.NewUser()
	po.Id = int64(id)
	if err := po.Query(); err != nil {
		logrus.Errorf("CreateUser po.Query err: %v", err)
		g.ResponseFail(http.StatusOK, 400, nil)
		return
	}
	
	rsp.User.ID, rsp.User.Name = po.Id, po.Name
	rsp.User.Address, rsp.User.Age = po.Address, po.Age
	g.ResponseOk(http.StatusOK, rsp)
	return 
}

func CreateUser(c *gin.Context) {
	var (
		req CreateUserReq
		rsp CreateUserRsp
		g = Gin{C: c}
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		logrus.Errorf("CreateUser ShouldBindJSON err: %v", err)
		g.ResponseFail(http.StatusOK, 400, nil)
		return
	}
	if req.User.Invalid() {
		g.ResponseFail(http.StatusOK, 400, nil)
		return 
	}

	po := models.NewUser()
	po.Name, po.Address, po.Age = req.Name, req.Address, req.Age
	if err := po.Create(); err != nil {
		logrus.Errorf("CreateUser po.Create err: %v", err)
		g.ResponseFail(http.StatusOK, 400, nil)
		return
	}

	rsp.Id = po.Id
	g.ResponseOk(http.StatusOK, rsp)
	return 
}

`
