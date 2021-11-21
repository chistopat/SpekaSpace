package simple_gateway

import (
	"encoding/json"
	"git.redmadrobot.com/internship/backend/lim-ext/pkg/openapi"
	"net/http"

	"git.redmadrobot.com/internship/backend/lim-ext/service/simple_gateway/generated"
)

func (s Server) Sort(w http.ResponseWriter, r *http.Request) {
	var req generated.SortJSONBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// TODO: Далее реализация метода
	// ...

	response := generated.ArrResponse{}
	openapi.Resp(w, http.StatusOK, response)
}

func (s Server) Uniq(w http.ResponseWriter, r *http.Request) {
	var req generated.SortJSONBody
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	response := generated.ArrResponse{}
	openapi.Resp(w, http.StatusOK, response)
}
