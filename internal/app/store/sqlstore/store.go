package sqlstore

import (
	"avitoTest/internal/app/store"
	"database/sql"
	_ "github.com/lib/pq"
)

type Store struct {
	db                *sql.DB
	segmentRepository *SegmentRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Segment() store.SegmentRepository {
	if s.segmentRepository != nil {
		return s.segmentRepository
	}

	s.segmentRepository = &SegmentRepository{
		store: s,
	}
	return s.segmentRepository
}
