package controller

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"liaoBa/dao/mysql"
	"liaoBa/logic"
	"liaoBa/models"
)

// SignUpHandler 注册
func SignUpHandler(c *gin.Context) {
	// 1. 获取参数和参数校验
	//var p models.ParamSignUp
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 请求参数有错误，返回响应
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
		//c.JSON(http.StatusOK, gin.H{
		//	//"msg": "请求参数有误",
		//	"msg": removeTopStruct(err.Translate(trans)), // 翻译错误
		//})
		return
	}
	// 参数校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("SignUp with invalid param")
	//	// 请求参数有错误，返回相应
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}
	// 输出看一下拿到的json参数
	//fmt.Println(*p)
	// 2. 业务处理logic层，也叫service层
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		// 判断err是不是validator.ValidationErrors类型
		err, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		// 请求参数有错误，返回响应
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(err.Translate(trans)))
		//c.JSON(http.StatusOK, gin.H{
		//	//"msg": "请求参数有误",
		//	"msg": removeTopStruct(err.Translate(trans)), // 翻译错误
		//})
		return
	}
	//业务处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("login.login failed", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNoExist) {
			ResponseError(c, CodeUserNoExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}

	// 返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
