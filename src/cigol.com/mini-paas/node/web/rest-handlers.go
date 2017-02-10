package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"cigol.com/mini-paas/common/entity"
	"cigol.com/mini-paas/common/utils"
	"github.com/gorilla/mux"
)

func opHandler(w http.ResponseWriter, r *http.Request) {
	// by default, the params are not to be parsed
	// when manually read from the body, should not call this method
	//	r.ParseForm()
	//	r.Body.Read()
	vars := mux.Vars(r)
	op := vars["op"]
	//	op := r.FormValue("op")
	//	message := r.FormValue("message")
	// manually read the body
	message, _ := utils.Read2String(r.Body)

	ret := exec(op, message)
	io.WriteString(w, ret)
}

func exec(op, message string) string {
	var s string
	switch op {
	case "deploy":
		s = doDeploy(message)
	case "start":
		s = doStart(message)
	case "stop":
		s = doStop(message)
	case "destroy":
		s = doDestroy(message)
	default:
		s = utils.WrapMessage("result", false, "message", "unrecognized op found: "+op)
	}
	return s
}

func doDeploy(message string) string {
	var msg map[string]interface{}
	if err := json.Unmarshal([]byte(message), &msg); err != nil {
		return utils.WrapMessage("result", false, "message", err)
	}
	image := msg["image"]
	replicas := msg["replicas"]
	instanceCode := msg["instanceCode"]
	fmt.Println(image, replicas, instanceCode)
	return utils.WrapMessage("result", true, "message", "deploy successfully")
}

func doStart(message string) string {
	return utils.WrapMessage("result", true, "message", "start successfully")
}

func doStop(message string) string {
	return utils.WrapMessage("result", true, "message", "stop successfully")
}

func doDestroy(message string) string {
	return utils.WrapMessage("result", true, "message", "destroy successfully")
}

// construct rest handler array
func GetRestHandlers() []entity.RestHandler {
	// add all the rest handlers here
	handlers := []entity.RestHandler{
		entity.NewRestHandler("/node/{op}", opHandler, "GET", "POST"),
	}
	return handlers
}
