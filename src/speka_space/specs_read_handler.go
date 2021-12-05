package speka_space

import (
	"encoding/json"
	gen "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/openapi"
	"net/http"
)

func (s Server) RetrieveSpecifications(w http.ResponseWriter, r *http.Request, params gen.RetrieveSpecificationsParams) {
	s.logger.Debug().Msg("Handle RetrieveSpecifications request")
	var req gen.RetrieveSpecificationsParams
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	specs, err := s.repo.Client.Specification.Query().All(r.Context())
	openapi.Resp(w, http.StatusOK, specs)
}
