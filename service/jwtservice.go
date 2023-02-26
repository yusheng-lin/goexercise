package service

import (
	"fmt"
	"goexercise/config"
	"goexercise/models"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type JWTService struct {
	config *config.Config
}

func NewJWTService(config *config.Config) *JWTService {
	return &JWTService{
		config: config,
	}
}

func (svc *JWTService) GenToken(account string) (string, error) {
	c := models.JWTClaim{
		Account: account,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(svc.config.JWT_EXPIRE)).Unix(),
			Issuer:    "goexercise",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(svc.config.JWT_SECRET))
}

func (svc *JWTService) ParseToken(token string) (*models.JWTClaim, error) {
	jwt, err := jwt.ParseWithClaims(token, &models.JWTClaim{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(svc.config.JWT_SECRET), nil
	})
	if err == nil && jwt != nil {
		if claim, ok := jwt.Claims.(*models.JWTClaim); ok && jwt.Valid {
			return claim, nil
		}
	}
	log.Error().Err(err)
	return nil, err
}

func (svc *JWTService) AuthRequired(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "Authorization is null in Header",
		})
		c.Abort()
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "Format of Authorization is wrong",
		})
		c.Abort()
		return
	}
	claims, err := svc.ParseToken(parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": fmt.Sprint(err),
		})
		c.Abort()
		return
	}
	c.Set("account", claims.Account)
	c.Next()
}
