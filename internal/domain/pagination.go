package domain

import (
	"fmt"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
	"net/url"
	"strings"
)

type Pagination[Model any] struct {
	Data      []Model `json:"data,omitempty"`
	FirstPage string  `json:"firstPage,omitempty"`
	LastPage  *string `json:"lastPage,omitempty"`
	PrevPage  *string `json:"prevPage,omitempty"`
	NextPage  string  `json:"nextPage,omitempty"`
	TotalPage int64   `json:"totalPage,omitempty"`
}

type PaginationContract interface {
	Page() *int
	Limit() *int
	Total() *int64
}

func (p *Pagination[Model]) Paginate(URL string, data *[]Model, config PaginationContract) Error {
	if config.Page() != nil {
		limit := *config.Limit()
		total := *config.Total()

		p.TotalPage = total / int64(limit)

		if total%int64(limit) > 0 {
			p.TotalPage++
		}

		if *config.Page() > 1 {
			prevPage, err := url.Parse(URL)
			if err != nil {
				return NewError(500, "Erro ao gerar paginação", err.Error())
			}

			query := prevPage.Query()
			query.Set("page", fmt.Sprintf("%d", *config.Page()-1))

			p.PrevPage = prepare(prevPage, query)
		}

		if *config.Page() < int(p.TotalPage) {
			nextPage, _ := url.Parse(URL)
			query := nextPage.Query()
			query.Set("page", fmt.Sprintf("%d", *config.Page()+1))

			p.NextPage = *prepare(nextPage, query)
		}

		{
			firstPage, _ := url.Parse(URL)
			query := firstPage.Query()
			query.Set("page", "1")

			p.FirstPage = *prepare(firstPage, query)
		}

		{
			lastPage, _ := url.Parse(URL)
			query := lastPage.Query()
			query.Set("page", fmt.Sprintf("%d", p.TotalPage))

			p.LastPage = prepare(lastPage, query)
		}

		p.Data = *data

		return nil
	}

	return NewError(500, "Erro ao gerar paginação", "O parâmetro page não foi informado")
}

func prepare(url *url.URL, query url.Values) *string {
	if strings.HasSuffix(url.Path, "/") {
		path := url.Path[:len(url.Path)-1]
		url.Path = path
	}

	url.RawQuery = query.Encode()
	path := url.String()

	return &path
}
