package speka_space

import (
	gen "git.redmadrobot.com/internship/backend/lim-ext/service/speka_space/generated"
	"net/http"
)

func (s Server) CreateSpecification(w http.ResponseWriter, r *http.Request)          { return }
func (s Server) ReadSpecification(w http.ResponseWriter, r *http.Request, id string) { return }
func (s Server) RetrieveSpecifications(w http.ResponseWriter, r *http.Request, params gen.RetrieveSpecificationsParams) {
	return
}
