package api

import "github.com/gin-gonic/gin"

func (a *API) auth(c *gin.Context) {

	c.Request.Header.Get("X-Api-Key")

}
