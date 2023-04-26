package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nurhidaylma/dts-secure-go-app/database"
	"github.com/nurhidaylma/dts-secure-go-app/helpers"
	"github.com/nurhidaylma/dts-secure-go-app/models"
	"github.com/nurhidaylma/dts-secure-go-app/response"
)

var (
	appJSON = "application/json"
)

// @BasePath /api/v1

// RegisterUser godoc
// @Summary Register user
// @Schemes
// @Description Register user by email and password
// @Accept json
// @Produce json
// @Param data body response.RegisterUser true "Sample payload"
// @Created 201 {object} response.JSONCreatedResult{data=response.RegisterUserResult,code=int,message=string}
// @Failure 400 {object} response.JSONBadReqResult{data=object,code=int,message=string}
// @Failure 500 {object} response.JSONIntServerErrReqResult{data=object,code=int,message=string}
// @Router /users/register [post]
func RegisterUser(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	user := models.User{}

	if contentType == appJSON {
		c.ShouldBindJSON(&user)
	} else {
		c.ShouldBind(&user)
	}

	err := db.Debug().Create(&user).Error

	if err != nil {
		response.FailResponse(c, 500, err.Error())

		return
	}

	response.CreatedResponse(c, response.RegisterUserResult{
		ID:       user.ID,
		FullName: user.FullName,
		Email:    user.Email,
	})

}
