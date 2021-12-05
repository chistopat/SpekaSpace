package speka_space

import (
	"encoding/json"
	"fmt"
	gen "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/openapi"
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent/specification"
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
	name := *req.Name
	spec, err := s.repo.Client.Specification.Query().Where(specification.NameEQ(name)).Only(r.Context())

	if spec != nil {
		errorType := gen.ErrorTypeCONFLICT
		message := fmt.Sprintf("Specification name: %s already exists", name)
		var resp = gen.Error{
			Type:    &errorType,
			Message: &message,
		}
		openapi.Resp(w, http.StatusConflict, resp)
		return
	}

	spec, err = s.repo.Client.Specification.Create().
		SetName(*req.Name).
		SetDescription(*req.Description).
		Save(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	openapi.Resp(w, http.StatusCreated, spec)
}
