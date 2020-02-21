package http_server

import (
	"github.com/gin-gonic/gin"
	"iot-demo/pkg/auth"
	"strings"
)

type DeviceTokenParser interface {
	Parse(authToken auth.Token) (*auth.DeviceCredential, error)
}

func deviceAuthParserHandler(parser DeviceTokenParser) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("authorization")
		if header == "" {
			c.Next()
			return
		}

		split := strings.SplitN(header, " ", 2)
		if len(split) != 2 || split[0] != "Bearer" {
			c.Next()
			return
		}

		token := auth.Token(strings.TrimSpace(split[1]))
		authInfo, err := parser.Parse(token)

		if err != nil {
			c.String(401, "bad authentication token")
			c.Abort()
			return
		}

		c.Set("auth_info", authInfo)
		c.Next()
	}
}
