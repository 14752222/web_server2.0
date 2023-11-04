package controller

// @BasePath /api/v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"time"
	"web_server_2.0/api/repository"
	"web_server_2.0/config"
	"web_server_2.0/types"
	"web_server_2.0/utils"
)

type LoginController struct {
	Result     *BaseController
	Env        *config.Env
	Db         *gorm.DB
	Repository *repository.LoginRepository
	Redis      *redis.Client
}

// Login 登录
// @Summary 登录
// @Tags 登录/注册
// @version 1.0
// @param email query string true "邮箱"
// @param password query string true "密码"
// @router /api/v1/login [post]
// @Accept json
// @Produce json
// @Success 200 {string} success
func (lc *LoginController) Login(ctx *gin.Context) {
	var login types.LoginRequest
	fmt.Println(ctx.Request.RequestURI)
	if err := ctx.ShouldBindJSON(&login); err != nil {
		ctx.JSON(400, gin.H{
			"message": err.Error(),
		})
		lc.Result.SendError(ctx, 0, err.Error(), nil)
		return
	}
	ok, user := lc.Repository.IsUserExist(lc.Db, login.Email)

	if !ok {
		//ctx.JSON(400, gin.H{
		//	"message": "用户不存在",
		//	"code":    "400",
		//})
		lc.Result.SendError(ctx, -1, "用户不存在", nil)
		return
	}
	if !utils.EqualMd5(user.Password, login.Password) {
		//ctx.JSON(400, gin.H{
		//	"message": "密码错误",
		//	"code":    "400",
		//})
		lc.Result.SendError(ctx, 0, "密码错误", nil)
		return
	}

	token, err := utils.CreateToken(user, lc.Env.SecretKey)
	if err != nil {
		//ctx.JSON(400, gin.H{
		//	"message": `服务端错误`,
		//})
		lc.Result.SendError(ctx, -1, `服务端错误`, nil)
		return
	}
	//ctx.JSON(200, gin.H{
	//	"message": "success",
	//	"code":    "200",
	//	"token":   token,
	//})
	lc.Result.SendSuccess(ctx, 1, "success", token)
	return

}

// Register 注册
// @Summary 注册
// @Tags 登录/注册
// @version 1.0
// @param email query string true "email" 邮箱
// @param password query string true "password" 密码
// @Success 200 {string} string 注册成功
// @Router /api/v1/login/register [post]
func (lc *LoginController) Register(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover", err)
		}
	}()

	var request types.RegisterRequest
	err := ctx.ShouldBind(&request)
	//获取参数
	fmt.Println(`1 request:`, request)
	if err != nil {
		fmt.Println(`err:`, err)
		lc.Result.SendError(ctx, 0, `参数错误`, nil)
		return
	}

	if request.Password != request.RPassword {
		lc.Result.SendError(ctx, 0, "两次密码不一致", nil)
		return
	}

	// check if user exists
	isUserExist, _ := lc.Repository.IsUserExist(lc.Db, request.Email)

	if isUserExist {
		lc.Result.SendError(ctx, -1, "用户已存在", nil)
		return
	}

	authCode, err := lc.Redis.Get(ctx.Request.Context(), request.Email).Result()

	if err != nil {
		lc.Result.SendError(ctx, 0, "验证码已过期", nil)
		return
	}

	if authCode != request.AuthCode {
		lc.Result.SendError(ctx, 0, "验证码错误", nil)
		return
	}

	// create user
	ok, user := lc.Repository.CreateUser(lc.Db, request.Email, utils.CreateMd5(request.Password))

	if !ok {
		lc.Result.SendError(ctx, -1, "注册失败", nil)
		return
	}

	// create token
	token, err := utils.CreateToken(user, lc.Env.SecretKey)
	if err != nil {
		lc.Result.SendError(ctx, -1, "服务端错误", nil)
		return
	}

	lc.Result.SendSuccess(ctx, 1, "注册成功", gin.H{
		"token": token,
		"user":  user,
	})
}

// SendEmail 发送邮件
// @Summary 发送邮件
// @Tags 登录/注册
// @Accept  json
// @Produce  json
// @Param email query string true "邮箱"
// @Param authCode query string true "验证码"
// @Success 200 {object} types.RegisterResponse
// @Failure 400 {object} types.ErrorResponse
// @Router /api/v1/login/send_email [post]
func (lc *LoginController) SendEmail(ctx *gin.Context) {
	var request types.SendEmailRequest
	err := ctx.ShouldBind(&request)
	fmt.Println(`request:`, request)
	if err != nil {
		//ctx.JSON(400, types.ErrorResponse{Message: err.Error()})
		lc.Result.SendError(ctx, 0, err.Error(), nil)
		return
	}
	fmt.Println(request)
	// check if user exists
	isUserExist, _ := lc.Repository.IsUserExist(lc.Db, request.Email)

	if isUserExist {
		//ctx.JSON(400, types.ErrorResponse{Message: "用户已存在"})
		lc.Result.SendError(ctx, -1, "用户已存在", nil)
		return
	}

	//fake send email
	code := utils.CreateCheckCode(6, false)
	err = utils.SendEmail(request.Email, "", code)

	if err != nil {
		//ctx.JSON(400, types.ErrorResponse{Message: err.Error()})
		lc.Result.SendError(ctx, 0, err.Error(), nil)
		return
	}

	// save code to redis
	context := ctx.Request.Context()
	err = lc.Redis.Set(context, request.Email, code, time.Minute*3).Err()

	if err != nil {
		lc.Result.SendError(ctx, 0, err.Error(), nil)
		return
	}

	//ctx.JSON(200, types.SuccessResponse{Message: "发送成功"})
	lc.Result.SendSuccess(ctx, 1, "发送成功", nil)

}

// AuthCodeLogin 验证码登录
// @Summary 验证码登录
// @Tags 登录/注册
// @Accept  json
// @Produce  json
// @Param email query string true "邮箱"
// @Param authCode query string true "验证码"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Router /api/v1/login/auth_code_login [post]
func (lc *LoginController) AuthCodeLogin(ctx *gin.Context) {
	var request types.SendEmailRequest
	err := ctx.ShouldBind(&request)
	fmt.Println(`request:`, request)
	if err != nil {
		ctx.JSON(400, types.ErrorResponse{Message: err.Error()})
		return
	}
	fmt.Println(request)
	// check if user exists
	isUserExist, _ := lc.Repository.IsUserExist(lc.Db, request.Email)

	if !isUserExist {
		ctx.JSON(400, types.ErrorResponse{Message: "用户不存在"})
		return
	}

	//fake send email
	code := utils.CreateCheckCode(6, false)
	fmt.Println(`code:`, code)
	err = utils.SendEmail(request.Email, "", code)

	if err != nil {
		ctx.JSON(400, types.ErrorResponse{Message: err.Error()})
		return
	}

	// save code to redis
	context := ctx.Request.Context()
	err = lc.Redis.Set(context, request.Email, code, time.Minute*3).Err()

	key, err := lc.Redis.Get(context, request.Email).Result()
	fmt.Println("key ", key, err)

	if err != nil {
		ctx.JSON(400, types.ErrorResponse{Message: err.Error()})
		return
	}

	ctx.JSON(200, types.SuccessResponse{Message: "发送成功"})
}
