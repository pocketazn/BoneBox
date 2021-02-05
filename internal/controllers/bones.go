package controllers

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pocketazn/BoneBox/internal/configuration"
	"github.com/pocketazn/BoneBox/internal/framework"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type V1BoneController struct {
	config *configuration.AppConfig
	dataAccess framework.DataAccessor
}

func NewV1BoneController(c *configuration.AppConfig, d framework.DataAccessor) V1BoneController {
	return V1BoneController{
		config: c,
		dataAccess: d,
	}
}

func (b *V1BoneController) RegisterRoutes(v1 *mux.Router) {
	v1.Path("/bones").Name("CreateBones").Handler(http.HandlerFunc(b.CreateBone)).Methods(http.MethodPost)
}

func (b *V1BoneController) CreateBone(w http.ResponseWriter, r *http.Request) {
	var ctx context.Context
	var bone framework.BoneBase

	err := json.NewDecoder(r.Body).Decode(&bone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Info("Validate Create Params")

	createdBone, err := b.dataAccess.CreateBone(ctx, bone)
	if err != nil {
		// TODO HANDLE THIS
	}

	respondModel(ctx, w, http.StatusCreated, &createdBone)
}
