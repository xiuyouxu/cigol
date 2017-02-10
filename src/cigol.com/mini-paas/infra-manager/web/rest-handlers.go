package web

import (
	"encoding/json"
	"net/http"
	"strconv"

	"cigol.com/mini-paas/common/entity"
)

func listHosts(w http.ResponseWriter, r *http.Request) {
	//	vars := mux.Vars(r)
	//	name := vars["name"]
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	// set header before write, or set will not work
	w.WriteHeader(http.StatusOK)

	var hosts []entity.Host = make([]entity.Host, 3, 3)
	for i := 0; i < 3; i++ {
		hosts[i] = entity.NewHost("host"+strconv.Itoa(i), "127.0.0."+strconv.Itoa(i+1), 22, "root", "1234")
	}
	b, e := json.Marshal(hosts)
	if e != nil {
		w.Write([]byte("marshalling hosts failed"))
	} else {
		w.Write(b)
	}
}

// construct rest handler array
func GetRestHandlers() []entity.RestHandler {
	// add all the rest handlers here
	handlers := []entity.RestHandler{
		entity.NewRestHandler("/iaas/hosts", listHosts, "GET"),
	}
	return handlers
}
