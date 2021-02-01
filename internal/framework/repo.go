package framework

import (
	"context"
	_ "github.com/Masterminds/squirrel"
)

type BoneRepoType struct {
}

type Bone struct {
	Id uint64 `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func NewBoneRepository() BoneRepoType {
	return BoneRepoType{}
}

func (r *BoneRepoType) CreateBone(ctx context.Context, bone Bone) (Bone, error) {
	var b Bone
	return b, nil
}
