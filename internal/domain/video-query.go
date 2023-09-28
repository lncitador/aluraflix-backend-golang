package domain

import (
	"fmt"
	vo "github.com/lncitador/alura-flix-backend/internal/domain/value-objects"
	"github.com/lncitador/alura-flix-backend/pkg/errors"
	"net/http"
	"strconv"
	"strings"
)

type VideoQuery struct {
	usuarioId *vo.UniqueEntityID
	search    *string
	page      *int
	limit     *int
	total     *int64
}

func (q *VideoQuery) UsuarioID() *vo.UniqueEntityID {
	return q.usuarioId
}

func (q *VideoQuery) SetUsuarioID(value string) *errors.Error {
	if value != "" {
		id, err := vo.NewUniqueEntityID(&value)
		if err != nil {
			return errors.NewErrorByUnauthorized(err)
		}

		q.usuarioId = id

		return nil
	}

	return errors.NewError(http.StatusUnprocessableEntity, "Unauthorized", "usuario_id is required")
}

func (q *VideoQuery) SetSearch(value string) {
	if value != "" {
		value = strings.ToLower(value)
		value = fmt.Sprintf("%%%s%%", value)
		q.search = &value
	}
}

func (q *VideoQuery) Search() *string {
	return q.search
}

func (q *VideoQuery) SetPage(value string) *errors.Error {
	if value != "" {
		if page, err := strconv.Atoi(value); err != nil {
			return errors.NewError(http.StatusUnprocessableEntity, "Validation error", "page must be a number")
		} else if page < 1 {
			return errors.NewError(http.StatusUnprocessableEntity, "Validation error", "page must be greater than 0")
		} else {
			q.page = &page

			return nil
		}
	}

	return nil
}

func (q *VideoQuery) Page() *int {
	return q.page
}

func (q *VideoQuery) SetLimit(value string) *errors.Error {
	if value != "" {
		if limit, err := strconv.Atoi(value); err != nil {
			return errors.NewError(http.StatusUnprocessableEntity, "Validation error", "limit must be a number")
		} else if limit < 10 {
			return errors.NewError(http.StatusUnprocessableEntity, "Validation error", "limit must be greater than 10")
		} else if limit > 100 {
			return errors.NewError(http.StatusUnprocessableEntity, "Validation error", "limit must be less than 100")
		} else {
			q.limit = &limit

			return nil
		}
	}

	return nil
}

func (q *VideoQuery) Limit() *int {
	if q.limit == nil {
		limit := 10
		q.limit = &limit
	}

	return q.limit
}

func (q *VideoQuery) Total() *int64 {
	if q.total == nil {
		total := int64(0)
		q.total = &total
	}

	return q.total
}

func (q *VideoQuery) SetTotal(value int64) {
	q.total = &value
}
