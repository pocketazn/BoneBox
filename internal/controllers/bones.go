package controllers

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/pocketazn/BoneBox/internal/configuration"
	"github.com/pocketazn/BoneBox/internal/framework"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type V1BoneController struct {
	config *configuration.AppConfig
}

func NewV1BoneController(c *configuration.AppConfig) V1BoneController {
	return V1BoneController{
		config: c,
	}
}

func (b *V1BoneController) RegisterRoutes(v1 *mux.Router) {
	v1.Path("/bones").Name("CreateBones").Handler(http.HandlerFunc(b.CreateBone)).Methods(http.MethodPost)
}

func (b *V1BoneController) CreateBone(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	log.Info("Validate Bone Request")
	bone := framework.Bone{
		BoneBase: framework.BoneBase{
			Name: "a name",
		},
	}

	respondModel(ctx, w, http.StatusCreated, &bone)
}
