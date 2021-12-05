package speka_space

import (
	"fmt"
	gen "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/openapi"
	googleUuid "github.com/google/uuid"
	"net/http"
)

func (s Server) ReadSpecification(w http.ResponseWriter, r *http.Request, uuid string) {
	s.logger.Debug().Msg("Handle ReadSpecification request")
	uuidObject, err := googleUuid.Parse(uuid)
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
