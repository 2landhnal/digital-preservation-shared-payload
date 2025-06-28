package middleware

import (
	"net/http"

	"github.com/2landhnal/digital-preservation-shared-payload/util"

	"github.com/gin-gonic/gin"
)

func ExtractUser(ctx *gin.Context) {
	// extract user_email from header
	userEmail := ctx.GetHeader(util.HeaderUserEmail)
	if userEmail == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	ctx.Set(util.HeaderUserEmail, userEmail)

	userRole := ctx.GetHeader(util.HeaderUserRole)
	if userRole == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		ctx.Abort()
		return
	}
	ctx.Set(util.HeaderUserRole, userRole)
}
