package csv2json

import (
	"reflect"
	"testing"
)

func TestToMap(t *testing.T) {
	type args struct {
		header []string
		item   []string
	}

	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			"success_same_length",
			args{
				header: []string{"name", "id", "age", "height", "is_active", "created_at"},
				item:   []string{"John", "1", "20", "1.80", "true", "2020-01-01T00:00:00Z"},
			},
			map[string]string{
				"name":       "John",
				"id":         "1",
				"age":        "20",
				"height":     "1.80",
				"is_active":  "true",
				"created_at": "2020-01-01T00:00:00Z",
			},
			false,
		},
		{
			"success_different_length",
			args{
				header: []string{"name", "id", "age", "height", "is_active", "created_at", "extra"},
				item:   []string{"John", "1", "20", "1.80", "true", "2020-01-01T00:00:00Z"},
			},
			map[string]string{
				"name":       "John",
				"id":         "1",
				"age":        "20",
				"height":     "1.80",
				"is_active":  "true",
				"created_at": "2020-01-01T00:00:00Z",
			},
			false,
		},
		{
			"failure",
			args{
				header: []string{"name", "id", "age", "height", "is_active"},
				item:   []string{"John", "1", "20", "1.80", "true", "2020-01-01T00:00:00Z"},
			},
			nil,
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMap(tt.args.header, tt.args.item)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
