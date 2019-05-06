package idleaf

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonRspErr(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
	c.Abort()
}

func GenDomainId(c *gin.Context) {
	domain := c.Param("domain")
	if domain == "" {
		c.JSON(http.StatusOK, gin.H{"code": ErrInternal, "id": 0, "msg": "domain lost"})
	} else {
		if id, err := idLeaf.GenId(domain); err != nil {
			c.JSON(http.StatusOK, gin.H{"code": ErrInternal, "id": 0, "msg": err.Error()})
		} else {
			c.JSON(http.StatusOK, gin.H{"code": ErrOK, "msg": "succ", "id": id})
		}
	}
}

func InitRouter() *gin.Engine {
	router := gin.Default()
	v1g := router.Group("/v1/")
	v1g.GET("gen/:domain", GenDomainId)
	return router
}
