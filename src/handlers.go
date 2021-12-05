package speka_space

import (
	"encoding/json"
	"fmt"
	gen "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/openapi"
	guuid "github.com/google/uuid"
	"net/http"
)

func (s Server) CreateSpecification(w http.ResponseWriter, r *http.Request) {
	s.logger.Debug().Msg("Handle CreateSpecification request")

	var req gen.CreateSpecificationJSONBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	spec, err := s.repo.Client.Specification.Create().
		SetName(*req.Name).
		SetDescription(*req.Description).
		Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	openapi.Resp(w, http.StatusCreated, spec)
}

func (s Server) ReadSpecification(w http.ResponseWriter, r *http.Request, uuid string) {
	s.logger.Debug().Msg("Handle ReadSpecification request")
	uuidObject, err := guuid.Parse(uuid)
	if err != nil {
		s.logger.Error().Msg("Could not parse uuid string")
		return
	}

	response, err := s.repo.Client.Specification.Get(r.Context(), uuidObject)
	if err != nil {
		s.logger.Warn().Msg(fmt.Sprintf("Unable to load specification from DB: %s", err))
		errorType := gen.ErrorTypeNOTFOUND
		message := fmt.Sprintf("Specification with id: %s, not found", uuid)
		var resp = gen.Error{
			Type:    &errorType,
			Message: &message,
		}
		openapi.Resp(w, http.StatusNotFound, resp)
		return
	}

	openapi.Resp(w, http.StatusOK, response)
}

func (s Server) RetrieveSpecifications(w http.ResponseWriter, r *http.Request, params gen.RetrieveSpecificationsParams) {
	return
}
