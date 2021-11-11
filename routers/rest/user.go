package rest

import (
	"context"
	"fmt"
	"gin-simple-base/pkg/e"
	"github.com/prometheus/common/log"
	"time"

	//"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	//"github.com/unknwon/com"
	"gin-simple-base/application"
	"gin-simple-base/routers/typespec"
	"net/http"
)

//ShouldBindBodyWith
//ShouldBindJSON
//ShouldBindXML
//ShouldBindQuery
//ShouldBindYAML
//ShouldBindHeader
//ShouldBindBodyWith
//GetRawData
//GetHeader
//Header
//Status
//SaveUploadedFile
func AddUser(c *gin.Context) {
	/*
		name := c.Query("name")
		state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
		createdBy := c.Query("created_by")

		valid := validation.Validation{}
		valid.Required(name, "name").Message("名称不能为空")
		valid.MaxSize(name, 100, "name").Message("名称最长为100字符")
		valid.Required(createdBy, "created_by").Message("创建人不能为空")
		valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	*/
	var req typespec.AddUserRequest
	var resp typespec.AddUserResponse

	code := e.SUCCESS

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.BIND_PARAMS_FAIL,
			"msg":  err.Error(),
		})
		return
	}
	fmt.Printf("req:%+v", req)

	//参数校验
	//err = filter.AddUserFilter(req)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": e.BIND_PARAMS_FAIL,
	//		"user": "",
	//		"pwd":  "",
	//	})
	//	return
	//}

	app := &application.User{}
	//ctx := server.NewContext(context.Background(), c)
	if err := app.AddUser(context.TODO(), &req, &resp); err != nil {
		log.Errorf("call app.AddUser trigger err: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": e.BIND_PARAMS_FAIL,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": resp,
	})
}

func GetUserList(c *gin.Context) {
	var req typespec.GetUserListRequest
	var resp typespec.GetUserListResponse

	code := e.SUCCESS

	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.BIND_PARAMS_FAIL,
			"msg":  err.Error(),
		})
		return
	}
	fmt.Printf("req:%+v\n", req)

	//参数校验

	app := &application.User{}
	//ctx := server.NewContext(context.Background(), c)
	if err := app.GetUserList(context.TODO(), &req, &resp); err != nil {
		log.Errorf("call app.GetUserList trigger err: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": e.BIND_PARAMS_FAIL,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": resp,
	})
}

func GetUser(c *gin.Context) {
	var req typespec.GetUserRequest
	var resp typespec.GetUserResponse

	code := e.SUCCESS

	go func() {
		fmt.Printf("xxxxxxxxxstart\n")
		time.Sleep(time.Second * 10)
		fmt.Printf("xxxxxxxxxend\n")
	}()
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": e.BIND_PARAMS_FAIL,
			"msg":  err.Error(),
		})
		return
	}

	//参数校验

	app := &application.User{}
	//ctx := server.NewContext(context.Background(), c)
	if err := app.GetUser(context.TODO(), &req, &resp); err != nil {
		log.Errorf("call app.GetUserList trigger err: %v", err)
		c.JSON(http.StatusOK, gin.H{
			"code": e.BIND_PARAMS_FAIL,
			"msg":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": resp,
	})
}
