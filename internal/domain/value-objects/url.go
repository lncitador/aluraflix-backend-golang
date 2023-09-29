package value_objects

import (
	"database/sql/driver"
	"fmt"
	"github.com/lncitador/alura-flix-backend/internal/domain/value-objects/errors"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
	"net/http"
	"regexp"
)

type URL ValueObject[string]

func NewURL(value string) (*URL, Error) {
	rgx := `^(http|https):\/\/[a-zA-Z0-9]+([\-\.]{1}[a-zA-Z0-9]+)*\.[a-zA-Z]{2,5}`

	if len(value) == 0 {
		return nil, NewError(http.StatusBadRequest, errors.URL_VAZIO, "")
	}

	match, err := regexp.MatchString(rgx, value)
	if err != nil {
		return nil, NewErrorByBadRequest(err)
	}

	if !match {
		return nil, NewError(http.StatusBadRequest, errors.URL_INVALIDO, "")
	}

	return &URL{value: value}, nil
}

func (u *URL) GetValue() string {
	return u.value
}

func (u *URL) ToString() string {
	return u.value
}

// Scan implements the Scanner interface.
func (u *URL) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf(errors.VALOR_NULO)
	}

	switch v := value.(type) {
	case []byte:
		u.value = string(v)
	case string:
		u.value = v
	default:
		return fmt.Errorf(errors.TIPO_INVALIDO)
	}

	return nil
}

// Value implements the driver Valuer interface.
func (u *URL) Value() (driver.Value, error) {
	return u.value, nil
}
