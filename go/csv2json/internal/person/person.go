package person

import "time"

type Person struct {
	Id        int       `json:"id" mapstructure:"id"`
	Name      string    `json:"name" mapstructure:"name"`
	Age       int       `json:"age" mapstructure:"age"`
	Height    float64   `json:"height" mapstructure:"height"`
	IsActive  bool      `json:"is_active" mapstructure:"is_active"`
	CreatedAt time.Time `json:"created_at" mapstructure:"created_at"`
}
