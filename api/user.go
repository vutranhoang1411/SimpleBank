package api

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/vutranhoang1411/SimpleBank/db/sqlc"
	"github.com/vutranhoang1411/SimpleBank/util"
)
var(
	WrongLoginInfo=errors.New("Wrong email or password");
)
type createUserRequest struct{
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=8"`
}
func (server *Server) createUser(ctx *gin.Context){
	var reqBody createUserRequest;
	if err:=ctx.ShouldBind(&reqBody);err!=nil{
		ctx.JSON(http.StatusBadRequest,handleError(err))
		return;
	}
	password,err:=util.HashPassword(reqBody.Password);
	if (err!=nil){
		ctx.JSON(http.StatusInternalServerError,handleError(err));
		return
	}
	
	user,err:=server.store.CreateUser(ctx,db.CreateUserParams{
		ID:uuid.NewString()[2:18],
		Name:reqBody.Name,
		Email:reqBody.Email,
		Password: password,
	})
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError,handleError(err));
		return
	}
	ctx.JSON(http.StatusOK,user)
}
type loginRequest struct{
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required,min=8"`

}
type loginResponse struct{
	Email string `json:"email"`
	AccessToken string `json:"access_token"`
	AccessTokenExpiredAt time.Time `json:"access_token_expired_at"`
}
func (server *Server)userLogin(ctx *gin.Context){
	//get request body
	var reqBody loginRequest;
	if err:=ctx.ShouldBind(&reqBody);err!=nil{
		ctx.JSON(http.StatusBadRequest,handleError(err));
		return;
	}

	user,err:=server.store.GetUserByEmail(ctx,reqBody.Email);
	if (err!=nil){
		if err==sql.ErrNoRows{
			ctx.JSON(http.StatusNotFound,handleError(WrongLoginInfo));
			return
		}
		ctx.JSON(http.StatusInternalServerError,handleError(err))
		return
	}
	if !util.CheckPasswordHash(user.Password,reqBody.Password){
		ctx.JSON(http.StatusBadRequest,handleError(WrongLoginInfo));
		return;
	}
	accessToken,payload,err:=server.maker.CreateToken(user.Name,server.config.TokenDuration);
	if (err!=nil){
		ctx.JSON(http.StatusInternalServerError,handleError(err))
		return
	}
	rsp:=loginResponse{
		Email:user.Email,
		AccessToken: accessToken,
		AccessTokenExpiredAt: payload.ExpiredAt,
	}
	ctx.JSON(http.StatusOK,rsp);
}
