package proxy

import (
	"bytes"
	"hung1299/go-projects/caching-proxy/global"
	"hung1299/go-projects/caching-proxy/internal/cache"
	"hung1299/go-projects/caching-proxy/internal/http/proxy"
	"hung1299/go-projects/caching-proxy/internal/http/response"
	"io"
	"time"

	"github.com/gin-gonic/gin"
)

func InitProxyRouter(r *gin.Engine) {
	r.NoRoute(handlerProxyRoute)
}

func handlerProxyRoute(c *gin.Context) {
	key := c.Request.URL.Path

	i := cache.Get(key)

	if i != nil {
		resp := response.ToResponse(i.Value)
		resp.Header.Set("X-Cache", "HIT")
		response.Copy(c, resp)
		return
	}

	resp := proxy.Forward(c.Request, global.Config.Origin)
	r, b := response.ToBytes(resp)
	cache.Set(key, r, time.Hour)

	resp.Body = io.NopCloser(bytes.NewBuffer(b))
	resp.Header.Set("X-Cache", "MISS")
	response.Copy(c, resp)
}
