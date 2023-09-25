package value_objects

import (
	"reflect"
	"testing"
)

func TestNewURL(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *URL
		wantErr bool
	}{
		{"URL válida", args{value: "https://www.google.com"}, &URL{value: "https://www.google.com"}, false},
		{"URL inválida", args{value: "www.google.com"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewURL(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
