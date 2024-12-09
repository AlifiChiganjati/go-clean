package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Status struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type SingleResponse struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}

type PagedResponse struct {
	Status Status        `json:"status"`
	Data   []interface{} `json:"data"`
	Paging any           `json:"paging"`
}

func SendCreateResponse(ctx *gin.Context, description string, data interface{}) {
	ctx.JSON(http.StatusCreated, SingleResponse{
		Status: Status{
			Code:        http.StatusCreated,
			Description: description,
		},
		Data: data,
	})
}

func SendSingleResponse(ctx *gin.Context, description string, data interface{}) {
	ctx.JSON(http.StatusOK, SingleResponse{
		Status: Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data: data,
	})
}

func SendErrorResponse(ctx *gin.Context, code int, description string) {
	ctx.JSON(code, SingleResponse{
		Status: Status{
			Code:        code,
			Description: description,
		},
	})
}

func SendPagedResponse(ctx *gin.Context, description string, data []interface{}, paging any) {
	ctx.JSON(http.StatusOK, PagedResponse{
		Status: Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data:   data,
		Paging: paging,
	})
}
