package types

import (
	"github.com/golang-jwt/jwt/v5"
	data "web_server_2.0/data/mysql"
)

type RegisterRequest struct {
	Email     string `form:"email" binding:"required,email"`
	Password  string `form:"password" binding:"required"`
	RPassword string `form:"rPassword" binding:"required"`
	AuthCode  string `form:"authCode" binding:"required"`
}

type SendEmailRequest struct {
	Email string `form:"email" binding:"required,email"`
}

type Data struct {
	Token    string      `json:"token"`
	UserInfo interface{} `json:"userInfo"`
}

type RegisterResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Data
}

type JWT struct {
	User data.SysUser `json:"user"`
	jwt.RegisteredClaims
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}
