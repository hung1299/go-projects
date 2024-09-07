package proxy

import (
	"hung1299/go-projects/caching-proxy/global"
	"hung1299/go-projects/caching-proxy/internal/proxy"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitProxyRouter(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		resp := proxy.Forward(c.Request, global.Config.Origin)

		forwardResponse(c, resp)
	})
}

func forwardResponse(c *gin.Context, resp *http.Response) {
	c.Status(resp.StatusCode)

	for k, vs := range resp.Header {
		for _, v := range vs {
			c.Header(k, v)
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Writer.Write(body)
}
