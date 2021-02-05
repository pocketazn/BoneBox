package framework

import (
	"context"
)

type BoneRepository interface {
	CreateBone(ctx context.Context, bone BoneBase) (Bone, error)
}

type DataAccessor interface {
	CreateBone(ctx context.Context, bone BoneBase) (Bone, error)
}

