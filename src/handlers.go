package speka_space

import (
	gen "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/openapi"
	"net/http"
)

func (s Server) CreateSpecification(w http.ResponseWriter, r *http.Request) { return }

func (s Server) ReadSpecification(w http.ResponseWriter, r *http.Request, id string) {
	s.logger.Debug().Msg("Handle ReadSpecification request")
	author := "chistopat"
	title := "my_super_spec"
	response := gen.Specification{
		Author: &author,
		Name:   &title,
	}
	openapi.Resp(w, http.StatusOK, response)
}

func (s Server) RetrieveSpecifications(w http.ResponseWriter, r *http.Request, params gen.RetrieveSpecificationsParams) {
	return
}
