package transactions

import (
	"github.com/garyjdn/djubli-api-go/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	users, err := services.TransactionService.GetTransaction()
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users.Marshall())
}
