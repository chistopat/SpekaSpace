package speka_space

import (
	gen "git.redmadrobot.com/internship/backend/lim-ext/src/generated"
	"git.redmadrobot.com/internship/backend/lim-ext/src/pkg/openapi"
	"git.redmadrobot.com/internship/backend/lim-ext/src/repository/ent/specification"
	"net/http"
)

func (s Server) RetrieveSpecifications(w http.ResponseWriter, r *http.Request, params gen.RetrieveSpecificationsParams) {
	s.logger.Debug().Msg("Handle RetrieveSpecifications request")
	author := params.Author
	if author != nil {
		specs, err := s.repo.Client.Specification.Query().
			Where(specification.AuthorEQ(*author)).All(r.Context())
		if err != nil {
		}
		openapi.Resp(w, http.StatusOK, specs)
		return
	}
	//status := params.Status
	//if status != nil {
	//	specs, err := s.repo.Client.Specification.Query().
	//		Where(specification.StatusEQ(*status)).All(r.Context())
	//	if err != nil {
	//	}
	//	openapi.Resp(w, http.StatusOK, specs)
	//	return
	//}

	specs, err := s.repo.Client.Specification.Query().All(r.Context())
	if err != nil {
	}
	openapi.Resp(w, http.StatusOK, specs)
}
