package value_objects

import (
	"github.com/google/uuid"
	"reflect"
	"testing"
)

func TestNewUniqueEntityID(t *testing.T) {
	type args struct {
		value *string
	}

	uid := uuid.New()
	uidStr := uid.String()
	idInvalido := "id inválido"

	tests := []struct {
		name    string
		args    args
		want    *UniqueEntityID
		wantErr bool
	}{
		{"Novo Id", args{value: nil}, nil, false},
		{"ID válido", args{value: &uidStr}, &UniqueEntityID{value: uid}, false},
		{"ID inválido", args{value: &idInvalido}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUniqueEntityID(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUniqueEntityID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want != nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUniqueEntityID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
