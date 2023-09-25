package domain

import (
	"fmt"
	"strings"
)

type VideoQuery struct {
	search *string
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
