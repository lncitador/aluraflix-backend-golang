package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type VideoQuery struct {
	search *string
	page   *int
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
	page := 1
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

	q.page = &page

	return nil
}

func (q *VideoQuery) Page() *int {
	if q.page == nil {
		page := 1
		q.page = &page
	}

	return q.page
}
