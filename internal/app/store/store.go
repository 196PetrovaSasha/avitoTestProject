package store

type Store interface {
	Segment() SegmentRepository
}
