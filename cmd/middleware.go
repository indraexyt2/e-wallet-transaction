package cmd

import (
	"e-wallet-transaction/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (d *Dependency) MiddlewareValidateToken(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	auth := c.Request.Header.Get("Authorization")
	if auth == "" {
		log.Error("Authorization header is required")
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"Authorization header is required",
			nil,
		)
		c.Abort()
		return
	}

	tokenData, err := d.External.ValidateToken(c.Request.Context(), auth)
	if err != nil {
		log.Error("failed to validate token", err)
		helpers.SendResponseHTTP(
			c,
			http.StatusUnauthorized,
			false,
			"failed to validate token",
			nil,
		)
		c.Abort()
		return
	}

	tokenData.Token = auth

	c.Set("token", tokenData)
	c.Next()
}
