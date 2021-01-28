package framework

import (
	"context"
	_ "github.com/Masterminds/squirrel"
)

type BoneRepoType struct {
}

type Bone struct {
	name string
}

func NewBoneRepository() BoneRepoType {
	return BoneRepoType{}
}

func (r *BoneRepoType) CreateBone(ctx context.Context, bone Bone) (Bone, error) {
	var b Bone
	return b, nil
}
