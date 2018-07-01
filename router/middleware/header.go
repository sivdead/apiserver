package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// NoCache is a middleware function that appends Request.Header.Adds
// to prevent the client from caching the HTTP response.
func NoCache(c *gin.Context) {
	c.Request.Header.Add("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Request.Header.Add("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Request.Header.Add("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

// Options is a middleware function that appends Request.Header.Adds
// for options requests and aborts then exits the middleware
// chain and ends the request.
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Request.Header.Add("Access-Control-Allow-Origin", "*")
		c.Request.Header.Add("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Request.Header.Add("Access-Control-Allow-Request.Header.Adds", "authorization, origin, content-type, accept")
		c.Request.Header.Add("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Request.Header.Add("Content-Type", "application/json")
		c.AbortWithStatus(200)
	}
}

// Secure is a middleware function that appends security
// and resource access Request.Header.Adds.
func Secure(c *gin.Context) {
	c.Request.Header.Add("Access-Control-Allow-Origin", "*")
	c.Request.Header.Add("X-Frame-Options", "DENY")
	c.Request.Header.Add("X-Content-Type-Options", "nosniff")
	c.Request.Header.Add("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Request.Header.Add("Strict-Transport-Security", "max-age=31536000")
	}

	// Also consider adding Content-Security-Policy Request.Header.Adds
	// c.Request.Header.Add("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
}
