package http

import (
	"fmt"
	"github.com/lncitador/alura-flix-backend/config"
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/impl"
	"github.com/lncitador/alura-flix-backend/internal/delivery/http/categorias"
	"github.com/lncitador/alura-flix-backend/internal/delivery/http/videos"
)

type Routes struct {
	*videos.VideoHandlers
	*categorias.CategoriaHandlers
}

func InitRoutes(config *config.Config) *Routes {
	router := config.GetRouter()
	api := router.Group("/api")

	db := config.GetDb()
	videoRepository := impl.NewVideoRepository(db)
	categoriasRepository := impl.NewCategoriaRepository(db)

	return &Routes{
		VideoHandlers:     videos.NewVideoHandlers(api, videoRepository),
		CategoriaHandlers: categorias.NewCategoriaHandlers(api, categoriasRepository),
	}
}

func (r *Routes) Register() {
	fmt.Print("Registering routes...\n")

	r.VideoHandlers.Register()
	r.CategoriaHandlers.Register()
}
