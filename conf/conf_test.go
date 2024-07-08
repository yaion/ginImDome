package conf

import (
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "config init!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Init()
			if Conf == nil {
				t.Errorf("conf is nil")
			}
		})
	}
}
