package speka_space

import (
	gen "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/openapi"
	"net/http"
)

func (s Server) CreateSpecification(w http.ResponseWriter, r *http.Request) { return }

func (s Server) ReadSpecification(w http.ResponseWriter, r *http.Request, id int) {
	s.logger.Debug().Msg("Handle ReadSpecification request")
	response, err := s.repo.Client.Specification.Get(r.Context(), id)
	if err != nil {

	}
	openapi.Resp(w, http.StatusOK, response)
}

func (s Server) RetrieveSpecifications(w http.ResponseWriter, r *http.Request, params gen.RetrieveSpecificationsParams) {
	return
}
