package repositories

import "github.com/lncitador/alura-flix-backend/internal/domain"

type VideoRepositoryContract RepositoryContract[domain.Video, *domain.VideoQuery]
