package model

import (
	"container/list"
	validation "github.com/go-ozzo/ozzo-validation"
)

type Segment struct {
	Name    string    `json:"segment_name"`
	ID      int       `json:"segment_id"`
	Clients list.List `json:"clients_in_segment"`
}

func (s *Segment) Validate() error {
	return validation.ValidateStruct(s, validation.Field(&s.Name, validation.Required))
}
