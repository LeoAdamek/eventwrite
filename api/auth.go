package api

import (
	"net/http"

	"bitbucket.org/mr-zen/eventwrite/auth"
	"github.com/gin-gonic/gin"
)

func (a *API) auth(c *gin.Context) {

	apiKey := c.Request.Header.Get("X-Api-Key")

	if apiKey == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	key, err := auth.GetAPIKey(apiKey)

	if err != nil || key == nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Set("source_id", key.SourceID)

	c.Next()

}
