package router

import (
	"github.com/gin-gonic/gin"
	"github.com/libaishwarya/mock-aws-ses-go/internal/server/ses"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	ses.AttachRoutes(r)

	return r
}
