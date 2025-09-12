package restapi

import (
	"unchained/server/config"
	"unchained/server/external/rest_api/middleware"
	"unchained/server/internal/transaction"
	"unchained/server/tools/logger"
	"unchained/server/uimport"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	ui         *uimport.UsecaseImport
	router     *gin.RouterGroup
	config     *config.Config
	log        *logger.Logger
	middleware *middleware.AuthMiddleware
	sm         transaction.SessionManager
}

func NewAuthHandler(
	ui *uimport.UsecaseImport,
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
		group.GET(
			"/sms_session",
			handler.GetAuthData,
		)
	}
}

func (h *AuthHandler) GetSmsSessionData(gctx *gin.Context)
