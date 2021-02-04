package framework

import (
	"context"
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	"time"
)

const (
	tableName  = "bones"
	schemaName = "bonebox"
	BonesTable = schemaName + "." + tableName
)

type BoneRepo struct {
	db sql.DB
}

type Bone struct {
	Id uint64 `json:"id" db:"id"`
	BoneBase
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	ArchivedAt *time.Time
}

type BoneBase struct {
	Name          string `json:"name" db:"name"`
	Description   string `json:"description" db:"description"`
	ExternalLabel string `json:"external_label" db:"external_label"`
}

var columnsBase = []string{
	"name",
	"description",
	"external_label",
}

func NewBoneRepository(db sql.DB) BoneRepo {
	return BoneRepo{
		db: db,
	}
}

func (r *BoneRepo) CreateBone(ctx context.Context, bone Bone) (Bone, error) {
	var b Bone

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	createBuilder := psql.Insert(BonesTable)
	query, args, err := createBuilder.Columns(columnsBase...).
		Values(
			bone.Name,
			bone.Description,
			bone.ExternalLabel,
		).
		ToSql()
	if err != nil {
		//TODO create custom error
		return b, err
	}

	rows := r.db.QueryRow(query, args...)
	if rows.Err() != nil {
		rows.Scan(bone)
	}

	return b, nil
}
