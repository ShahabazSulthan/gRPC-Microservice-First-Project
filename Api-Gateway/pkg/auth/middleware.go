package auth

import (
	"api-gateway/pkg/auth/pb"
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization"})
		return
	}
	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing authorization"})
		return
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})
	fmt.Println("res", res)

	if err != nil || res.Status != http.StatusOK {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization token"})
		return
	}

	ctx.Set("userId", res.UserId)

	if res.Role == "admin" {
		ctx.Set("userRole", "admin")
	} else if res.Role == "user" {
		ctx.Set("userRole", "user")
	}
	fmt.Println("User role:", ctx.GetString("userRole"))

	ctx.Next()
}
