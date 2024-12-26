package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Status struct {
		Code        int    `json:"code"`
		Description string `json:"description"`
	}

	SingleResponse struct {
		Status Status      `json:"status"`
		Data   interface{} `json:"data"`
	}

	PagedResponse struct {
		Status Status      `json:"status"`
		Data   interface{} `json:"data"`
		Paging interface{} `json:"paging"`
	}
)

func SendCreateResponse(c *gin.Context, description string, data interface{}) {
	c.JSON(http.StatusCreated, SingleResponse{
		Status: Status{
			Code:        http.StatusCreated,
			Description: description,
		},
		Data: data,
	})
}

func SendSingleResponse(c *gin.Context, description string, data interface{}) {
	c.JSON(http.StatusOK, SingleResponse{
		Status: Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data: data,
	})
}

func SendErrorResponse(c *gin.Context, code int, description string) {
	c.JSON(code, SingleResponse{
		Status: Status{
			Code:        code,
			Description: description,
		},
	})
}

func SendPagedResponse(c *gin.Context, description string, data interface{}, paging interface{}) {
	c.JSON(http.StatusOK, PagedResponse{
		Status: Status{
			Code:        http.StatusOK,
			Description: description,
		},
		Data:   data,
		Paging: paging,
	})
}
