package value_objects

import (
	"database/sql/driver"
	"fmt"
	"github.com/google/uuid"
	"github.com/lncitador/alura-flix-backend/internal/domain/value-objects/errors"
	. "github.com/lncitador/alura-flix-backend/pkg/errors"
)

type UniqueEntityID ValueObject[uuid.UUID]

func NewUniqueEntityID(value *string) (*UniqueEntityID, Error) {
	if value == nil {
		return &UniqueEntityID{value: uuid.New()}, nil
	}

	id, err := uuid.Parse(*value)
	if err != nil {
		return nil, NewErrorByValidation(err)
	}

	return &UniqueEntityID{value: id}, nil
}

func (id *UniqueEntityID) Equals(value ValueObjectContract[uuid.UUID]) bool {
	return id.value == value.GetValue()
}

func (id *UniqueEntityID) GetValue() uuid.UUID {
	return id.value
}

func (id *UniqueEntityID) ToString() string {
	return id.value.String()
}

// Scan implements the Scanner interface.
func (id *UniqueEntityID) Scan(value interface{}) error {
	if value == nil {
		return fmt.Errorf(errors.VALOR_NULO)
	}

	switch v := value.(type) {
	case []byte:
		parsed, err := uuid.ParseBytes(v)
		if err != nil {
			return fmt.Errorf("Erro ao fazer a varredura do valor do banco de dados para UniqueEntityID: %v", err)
		}
		id.value = parsed
	case string:
		parsed, err := uuid.Parse(v)
		if err != nil {
			return fmt.Errorf("Erro ao fazer a varredura do valor do banco de dados para UniqueEntityID: %v", err)
		}
		id.value = parsed
	default:
		return fmt.Errorf(errors.TIPO_INVALIDO)
	}

	return nil
}

// Value implements the driver Valuer interface.
func (id *UniqueEntityID) Value() (driver.Value, error) {
	return id.value.String(), nil
}
