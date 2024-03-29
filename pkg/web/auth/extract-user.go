package auth

import (
	"github.com/ethereum/agentCard/config"
	"github.com/ethereum/agentCard/pkg/log"
	"github.com/ethereum/agentCard/pkg/util"
	"github.com/ethereum/agentCard/pkg/web"
	"github.com/gin-gonic/gin"
	"strings"
)

var logger = log.C.Logger()

func MustExtractUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken string
		cookie, err := c.Cookie("access_token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			web.BadRes(c, util.ErrTokenInvalid)
			logger.Error().Msg("no accessToken")
			return
		}

		sub, er := ValidateToken(accessToken, config.Conf.Access.Pub)
		if er != nil {
			web.BadRes(c, er)
			logger.Error().Msg(er.LStr())
			return
		}

		c.Set("currentUser", sub)
		c.Next()
	}
}

func SilentExtractUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accessToken string
		cookie, err := c.Cookie("access_token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			accessToken = fields[1]
		} else if err == nil {
			accessToken = cookie
		}

		if accessToken == "" {
			logger.Debug().Msg("no accessToken")
			return
		}

		sub, er := ValidateToken(accessToken, config.Conf.Access.Pub)
		if er != nil {
			logger.Warn().Msg(er.LStr())
			return
		}

		c.Set("currentUser", sub)
		c.Next()
	}
}
