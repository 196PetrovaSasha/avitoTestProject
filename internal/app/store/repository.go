package store

import (
	"avitoTest/internal/app/model"
	"container/list"
)

type SegmentRepository interface {
	Create(segment *model.Segment) error
	FindBySegmentName(string) (*model.Segment, error)
	DeleteSegment(string) error
	UpdateClient(list.List, list.List, int) error
	TakeSegments(int) error
}
