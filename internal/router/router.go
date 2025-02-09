package router

import (
	"github.com/gin-gonic/gin"
	"github.com/libaishwarya/mock-aws-ses-go/internal/server/ses"
	"github.com/libaishwarya/mock-aws-ses-go/internal/store/inmemory"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	inMemoryStore := inmemory.NewInMemoryStore()

	ses.AttachRoutes(r, inMemoryStore)

	return r
}
