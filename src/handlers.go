package speka_space

import (
	"fmt"
	gen "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/openapi"
	guuid "github.com/google/uuid"
	"net/http"
)

func (s Server) CreateSpecification(w http.ResponseWriter, r *http.Request) { return }

func (s Server) ReadSpecification(w http.ResponseWriter, r *http.Request, uuid string) {
	s.logger.Debug().Msg("Handle ReadSpecification request")
	uuidObject, err := guuid.Parse(uuid)
	if err != nil {
		s.logger.Error().Msg("Could not parse uuid string")
		return
	}

	response, err := s.repo.Client.Specification.Get(r.Context(), uuidObject)
	if err != nil {
		s.logger.Error().Msg(fmt.Sprintf("Unable to load specification from DB: %s", err))
		return
	}
	openapi.Resp(w, http.StatusOK, response)
}

func (s Server) RetrieveSpecifications(w http.ResponseWriter, r *http.Request, params gen.RetrieveSpecificationsParams) {
	return
}
