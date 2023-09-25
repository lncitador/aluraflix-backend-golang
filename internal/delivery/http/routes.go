package http

import (
	"fmt"
	"github.com/lncitador/alura-flix-backend/config"
	"github.com/lncitador/alura-flix-backend/internal/delivery/http/videos"
)

type Routes struct {
	*videos.VideoHandlers
}

func InitRoutes(config *config.Config) *Routes {
	return &Routes{
		VideoHandlers: videos.NewVideoHandlers(config),
	}
}

func (r *Routes) Register() {
	fmt.Print("Registering routes...\n")

	r.VideoHandlers.Register()
}
