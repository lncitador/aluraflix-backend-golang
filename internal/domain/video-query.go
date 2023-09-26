package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type VideoQuery struct {
	search *string
	page   *int
	limit  *int
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

func (q *VideoQuery) SetPage(value string) error {
	if value != "" {
		if page, err := strconv.Atoi(value); err != nil {
			return fmt.Errorf("page must be a number")
		} else if page < 1 {
			return fmt.Errorf("page must be greater than 0")
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

func (q *VideoQuery) SetLimit(value string) error {
	if value != "" {
		if limit, err := strconv.Atoi(value); err != nil {
			return fmt.Errorf("limit must be a number")
		} else if limit < 10 {
			return fmt.Errorf("limit must be greater than 10")
		} else if limit > 100 {
			return fmt.Errorf("limit must be less than 100")
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
