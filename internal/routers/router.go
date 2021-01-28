package routers

import (
	"github.com/gorilla/mux"
	"github.com/pocketazn/BoneBox/internal/configuration"
	"github.com/pocketazn/BoneBox/internal/controllers"
	"github.com/pocketazn/BoneBox/internal/framework"
	"github.com/urfave/negroni"
	"net/http"
)

type V1Router struct {
	config   *configuration.AppConfig
	boneRepo framework.BoneRepoType
}

func NewV1Router(c *configuration.AppConfig, cRepo framework.BoneRepoType) V1Router {
	return V1Router{
		config:   c,
		boneRepo: cRepo,
	}
}

//InitRoutes initialize all routes
func (v *V1Router) Register(root *mux.Router) {


	r := root.PathPrefix("/v1").Subrouter()


	var bonesController controllers.V1BoneController

	bonesController.RegisterRoutes(r)


	root.PathPrefix("/docs").Handler(&docsHander{
		root: v.config.DocsPath,
		staticHandler: negroni.New(&negroni.Static{
			Dir:       http.Dir(v.config.DocsPath),
			Prefix:    "/docs",
			IndexFile: "index.html",
		}),
	})
}

type docsHander struct {
	root          string
	staticHandler *negroni.Negroni
}

func (h *docsHander) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "text/html")
	h.staticHandler.ServeHTTP(resp, req)
}