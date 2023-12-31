package usecases

import (
	"github.com/lncitador/alura-flix-backend/internal/application/repositories/inmemory"
	"github.com/lncitador/alura-flix-backend/internal/domain"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

type VideoSut struct {
	repo      *inmemory.VideoRepository
	useCase   *VideosUseCase
	constants domain.VideoInput
}

var invalidUrl = "www.youtube"

func setupVideoSut() *VideoSut {
	relation := inmemory.NewCategoriaRepository()
	repo := inmemory.NewVideoRepository(relation)

	useCase := NewVideosUseCase(repo)

	titulo := "Título do vídeo"
	descricao := "Descrição do vídeo"
	url := "https://www.youtube.com/watch?v=123456789"

	name := "Name da categoria"
	hexColor := "#FFFFFF"

	var usuarioId string

	{
		name := "John Doe"
		email := "doejoe@test.com"
		password := "123456"

		usuario, _ := domain.NewUsuario(domain.UsuarioInput{
			Name:     &name,
			Email:    &email,
			Password: &password,
		})

		usuarioId = usuario.ID.ToString()
	}

	categoria, _ := domain.NewCategoria(domain.CategoriaInput{
		Name:      &name,
		Color:     &hexColor,
		UsuarioID: &usuarioId,
	})

	categoriaId := categoria.ID.ToString()

	if err := relation.Create(*categoria); err != nil {
		log.Fatal(err)
	}

	return &VideoSut{
		repo:    repo,
		useCase: useCase,
		constants: domain.VideoInput{
			Title:       &titulo,
			Description: &descricao,
			URL:         &url,
			CategoryID:  &categoriaId,
			UsuarioID:   &usuarioId,
		},
	}
}

func TestVideosUseCase_Create(t *testing.T) {
	sut := setupVideoSut()
	data := sut.constants

	t.Run("should create a video", func(t *testing.T) {
		video, err := sut.useCase.Create(data)
		assert.Nil(t, err)
		assert.NotNil(t, video)

		assert.Equal(t, *data.Title, video.Title)
		assert.Equal(t, *data.Description, video.Description)
		assert.Equal(t, *data.URL, video.URL)
	})

	t.Run("should not create a video with invalid data", func(t *testing.T) {
		data := domain.VideoInput{
			Title:       nil,
			Description: nil,
			URL:         nil,
		}

		video, err := sut.useCase.Create(data)
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})

	t.Run("should not create a video with invalid url", func(t *testing.T) {

		data := domain.VideoInput{
			Title:       data.Title,
			Description: data.Description,
			URL:         &invalidUrl,
		}

		video, err := sut.useCase.Create(data)
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})
}

func TestVideosUseCase_Delete(t *testing.T) {
	sut := setupVideoSut()

	t.Run("should delete a video", func(t *testing.T) {
		video, err := sut.useCase.Create(sut.constants)

		id, err := vo.NewUniqueEntityID(&video.ID)
		assert.Nil(t, err)
		assert.NotNil(t, video)

		err = sut.useCase.Delete(id)
		assert.Nil(t, err)
	})

	t.Run("should not delete a video with invalid id", func(t *testing.T) {
		err := sut.useCase.Delete(nil)
		assert.NotNil(t, err)
	})
}

func TestVideosUseCase_FindAll(t *testing.T) {
	sut := setupVideoSut()
	query := domain.VideoQuery{}

	t.Run("should find all videos", func(t *testing.T) {
		videos, err := sut.useCase.FindAll(&query)
		assert.Nil(t, err)
		assert.NotNil(t, videos)
	})

	t.Run("should find all videos with search", func(t *testing.T) {
		query.SetSearch("Título")
		videos, err := sut.useCase.FindAll(&query)
		assert.Nil(t, err)
		assert.NotNil(t, videos)
	})
}

func TestVideosUseCase_FindById(t *testing.T) {
	sut := setupVideoSut()

	t.Run("should find a video by id", func(t *testing.T) {
		video, err := sut.useCase.Create(sut.constants)

		id, err := vo.NewUniqueEntityID(&video.ID)
		assert.Nil(t, err)
		assert.NotNil(t, video)

		video, err = sut.useCase.FindById(id)
		assert.Nil(t, err)
		assert.NotNil(t, video)
	})

	t.Run("should not find a video with invalid id", func(t *testing.T) {
		video, err := sut.useCase.FindById(nil)
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})

	t.Run("should not find a video with id not found", func(t *testing.T) {
		id, _ := vo.NewUniqueEntityID(nil)

		video, err := sut.useCase.FindById(id)
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})
}

func TestVideosUseCase_Update(t *testing.T) {
	sut := setupVideoSut()

	t.Run("should update a video", func(t *testing.T) {
		video, err := sut.useCase.Create(sut.constants)

		id, err := vo.NewUniqueEntityID(&video.ID)
		assert.Nil(t, err)
		assert.NotNil(t, video)

		newTitle := "Novo título"

		data := domain.VideoInput{
			Title: &newTitle,
		}

		video, err = sut.useCase.Update(id, data)
		assert.Nil(t, err)
		assert.NotNil(t, video)
		assert.Equal(t, *data.Title, video.Title)
	})

	t.Run("should not update a video with invalid id", func(t *testing.T) {
		video, err := sut.useCase.Update(nil, sut.constants)
		assert.NotNil(t, err)
		assert.Nil(t, video)
	})

	t.Run("should not update a video with invalid data", func(t *testing.T) {
		video, _ := sut.useCase.Create(sut.constants)
		id, _ := vo.NewUniqueEntityID(&video.ID)
		_, err := sut.useCase.Update(id, domain.VideoInput{
			URL: &invalidUrl,
		})

		assert.NotNil(t, err)
	})
}
