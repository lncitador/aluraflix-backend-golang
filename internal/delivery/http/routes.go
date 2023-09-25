package http

import (
	"fmt"
	"github.com/lncitador/alura-flix-backend/config"
	"github.com/lncitador/alura-flix-backend/internal/delivery/http/categorias"
	"github.com/lncitador/alura-flix-backend/internal/delivery/http/videos"
)

type Routes struct {
	*videos.VideoHandlers
	*categorias.CategoriaHandlers
}

func InitRoutes(config *config.Config) *Routes {
	return &Routes{
		VideoHandlers:     videos.NewVideoHandlers(config),
		CategoriaHandlers: categorias.NewCategoriaHandlers(config),
	}
}

func (r *Routes) Register() {
	fmt.Print("Registering routes...\n")

	r.VideoHandlers.Register()
	r.CategoriaHandlers.Register()
}
