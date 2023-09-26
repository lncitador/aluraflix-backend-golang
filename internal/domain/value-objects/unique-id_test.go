package value_objects

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUniqueEntityID(t *testing.T) {
	uid := uuid.New()
	uidStr := uid.String()
	idInvalido := "id inválido"

	t.Run("should create a new unique entity id", func(t *testing.T) {
		id, err := NewUniqueEntityID(nil)
		assert.Nil(t, err)
		assert.NotNil(t, id)
	})

	t.Run("should not create a new unique entity id with invalid id", func(t *testing.T) {
		id, err := NewUniqueEntityID(&idInvalido)
		assert.NotNil(t, err)
		assert.Nil(t, id)
	})

	t.Run("should be able to convert a unique entity id to string", func(t *testing.T) {
		id, _ := NewUniqueEntityID(&uidStr)
		assert.Equal(t, uidStr, id.ToString())
	})
}
