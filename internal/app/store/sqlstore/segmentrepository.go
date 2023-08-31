package sqlstore

import (
	"avitoTest/internal/app/model"
	"avitoTest/internal/app/store"
	"container/list"
	"database/sql"
	"errors"
)

type SegmentRepository struct {
	store *Store
}

func (r *SegmentRepository) Create(u *model.Segment) error {

	return r.store.db.QueryRow(
		"INSERT INTO  segments (name) VALUES ($1) RETURNING id",
		u.Name,
	).Scan(u.ID)
}

func (r *SegmentRepository) FindBySegmentName(name string) (*model.Segment, error) {
	u := &model.Segment{}
	if err := r.store.db.QueryRow(`SELECT name, clients FROM segments WHERE name = $1`, name).Scan(&u.Name, &u.Clients); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}
	return u, nil
}

func (r *SegmentRepository) DeleteSegment(name string) error {
	if _, err := r.FindBySegmentName(name); err != nil {
		return err
	}
	r.store.db.QueryRow("DELETE FROM segments WHERE name = $1", name)
	return nil
}

func (r *SegmentRepository) UpdateClient(segmentsforadd list.List, segmentsfordelete list.List, id int) error {
	return nil
}

func (r *SegmentRepository) TakeSegments(id int) error { return nil }
