package controllers

import (
	"github.com/gorilla/mux"
	"github.com/pocketazn/BoneBox/internal/configuration"
	"net/http"
)

type V1BoneController struct {
	config     *configuration.AppConfig
}

func NewV1BoneController(c *configuration.AppConfig) V1BoneController {
	return V1BoneController{
		config:     c,
	}
}

func (b * V1BoneController) RegisterRoutes(v1 *mux.Router)  {
	v1.Path("/bones").Name("CreateBones").Handler(http.HandlerFunc(b.CreateBone)).Methods(http.MethodPost)
}

func (b * V1BoneController) CreateBone(w http.ResponseWriter, r *http.Request) {
}