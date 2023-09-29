package domain

import (
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/lncitador/alura-flix-backend/pkg/errors"
	"time"
)

type VideoDto struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	URL         string     `json:"url"`
	CategoryID  string     `json:"categoryId"`
	Category    *Categoria `json:"category,omitempty"`
	UsuarioID   string     `json:"usuarioId"`
	Usuario     *Usuario   `json:"usuario,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

// MapFrom maps VideoDto to Video struct
func (d VideoDto) MapFrom() (*Video, errors.Error) {
	id, err := vo.NewUniqueEntityID(&d.ID)
	if err != nil {
		return nil, errors.NewErrorByValidation(err)
	}

	newUrl, err := vo.NewURL(d.URL)
	if err != nil {
		return nil, errors.NewErrorByValidation(err)
	}

	base := Base{
		ID:        id,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}

	categoriaId, err := vo.NewUniqueEntityID(&d.CategoryID)
	if err != nil {
		return nil, errors.NewErrorByValidation(err)
	}

	return &Video{
		Base:        base,
		Title:       d.Title,
		Description: d.Description,
		URL:         newUrl,
		CategoryID:  categoriaId,
		Category:    d.Category,
	}, nil
}

func VideosToDto(videos []Video) *[]VideoDto {
	var videosDto []VideoDto

	for _, video := range videos {
		videosDto = append(videosDto, *video.MapTo())
	}

	return &videosDto
}
