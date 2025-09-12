package middleware

import (
	"unchained/server/internal/auth"
	"unchained/server/internal/entity/global"
	"unchained/server/internal/entity/jwt_auth"
	"unchained/server/tools/gin_gen"

	"github.com/gin-gonic/gin"
)

const (
	userDataKey = "user_data"
)

type AuthMiddleware struct {
	jwtAuth *auth.Jwt
}

func NewAuthMiddleware(jwtAuth *auth.Jwt) *AuthMiddleware {
	return &AuthMiddleware{jwtAuth}
}

func (m *AuthMiddleware) CheckAccesToken() gin.HandlerFunc {
	return func(gctx *gin.Context) {
		token, err := gctx.Cookie("access_token")
		if err != nil {
			gin_gen.HandleError(gctx, global.ErrPermissionDenied)
			gctx.Abort()
			return
		}

		claims, err := m.jwtAuth.ParseToken(token)
		if err != nil {
			gin_gen.HandleError(gctx, err)
			gctx.Abort()
			return
		}

		gctx.Set(userDataKey, claims)
		gctx.Next()
	}
}

func (m *AuthMiddleware) GetUserData(gctx *gin.Context) (jwt_auth.Claims, error) {
	var zero jwt_auth.Claims

	jwtClaims, exists := gctx.Get(userDataKey)
	if !exists {
		return zero, global.ErrInternalError
	}

	claimsData, ok := jwtClaims.(jwt_auth.Claims)
	if !ok {
		return zero, global.ErrInternalError
	}

	return claimsData, nil
}
