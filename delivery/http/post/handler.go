package post

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sharing-vision-2021/common"
	"sharing-vision-2021/config/validator"
	"sharing-vision-2021/delivery/http/post/model"
)

// Add godoc
// @Summary Add Post
// @Description This endpoint for Add Post
// @Tags Post
// @Accept  json
// @Produce  json
// @Param services body dto.RequestUser true "Login Authentication"
// @Success 200 {object} models.JSONResponsesSwaggerSucceed
// @Router /article/ [post]
func (c *controller) Add(ctx *gin.Context) {
	bodyRequest := new(model.AddArticle)
	if err := ctx.BindJSON(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.ErrorResponse(err.Error()))
		return
	}

	if err := validator.ValidateStruct(bodyRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, common.BadRequestResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, common.SuccessResponseWithData(bodyRequest, "success"))
}
