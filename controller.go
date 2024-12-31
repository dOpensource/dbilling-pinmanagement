package dbilling_pinmanagement

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func PinGetHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		var response_data gin.H

		response_data = gin.H{
			//"pin": database.pinmanagement.Get(&query_params),
			"pins": GetPinInfo(db, "1012"),
		}

		accept_hdr := c.GetHeader("Accept")
		if strings.Contains(accept_hdr, "application/json") {
			c.JSON(http.StatusOK, response_data)
		} else if strings.Contains(accept_hdr, "text/html") {
			c.HTML(http.StatusOK, "routes", response_data)
		} else {
			panic(fmt.Errorf("unsupported response formast"))
		}
	}
}
