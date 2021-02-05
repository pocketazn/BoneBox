package framework

import (
	"context"
	"database/sql"
)

// DataAccess stores multiple DB structs
type DataAccess struct {
	DB    BoneRepo
}

// NewDataAccess initializes the data fetch layer for framework
func NewDataAccess(db sql.DB) DataAccess {
	br := NewBoneRepository(db)
	return DataAccess{
		DB: br,
	}
}

func (d *DataAccess) CreateBone(ctx context.Context, bone BoneBase) (Bone, error) {
	b, err := d.DB.CreateBone(ctx, bone)
	if err != nil {
		return b, err
	}
	return b, nil
}