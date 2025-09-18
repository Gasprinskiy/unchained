package restapi

import (
	"net/http"
	"unchained/server/config"
	"unchained/server/external/rest_api/middleware"
	"unchained/server/internal/entity/auth"
	"unchained/server/internal/entity/global"
	"unchained/server/internal/transaction"
	"unchained/server/tools/gin_gen"
	"unchained/server/tools/logger"
	"unchained/server/uimport"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	ui         *uimport.Usecase
	router     *gin.RouterGroup
	config     *config.Config
	log        *logger.Logger
	middleware *middleware.AuthMiddleware
	sm         transaction.SessionManager
}

func NewAuthHandler(
	ui *uimport.Usecase,
	router *gin.RouterGroup,
	config *config.Config,
	log *logger.Logger,
	middleware *middleware.AuthMiddleware,
	sm transaction.SessionManager,
) {
	handler := AuthHandler{
		ui,
		router,
		config,
		log,
		middleware,
		sm,
	}

	group := handler.router.Group("/auth")

	{
		group.POST(
			"/verification_code",
			handler.CreateVerificationCode,
		)

		group.POST(
			"/verify_code",
		)
	}
}

func (h *AuthHandler) CreateVerificationCode(gctx *gin.Context) {
	var params auth.CreateVerificationCodeParam
	if err := gctx.ShouldBindJSON(&params); err != nil {
		gin_gen.HandleError(gctx, global.ErrInternalError)
		return
	}

	if err := h.ui.Auth.CreateVerificationCode(gctx, params); err != nil {
		gin_gen.HandleError(gctx, err)
		return
	}

	gctx.JSON(http.StatusCreated, gin.H{"success": true})
}
