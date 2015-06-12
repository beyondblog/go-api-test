package api

import (
	"encoding/json"
	. "github.com/beyondblog/go-api-test/server"
	"github.com/mholt/binding"
	"log"
	"net/http"
)

func Run(w http.ResponseWriter, r *http.Request) {
	log.Println("run request")

	appForm := new(AppForm)
	err := binding.Bind(r, appForm)
	if err.Handle(w) {
		return
	}

	var apiRequest ApiRequest
	apiRequest.Desc = appForm.Desc
	apiRequest.Host = appForm.Host
	apiRequest.Param = make(map[string]string)
	apiRequest.Method = appForm.Method
	var jsonRes JsonResponse

	if len(apiRequest.Host) == 0 {
		jsonRes.Code = 400
		jsonRes.Message = "host error"
	} else {
		for _, param := range appForm.Param {
			if len(param.Key) > 0 {
				apiRequest.Param[param.Key] = param.Value
			}
		}
		if response, err := RunTest(apiRequest); err != nil {
			jsonRes.Code = 500
			jsonRes.Message = err.Error()
		} else {
			jsonRes.Code = response.Code
			jsonRes.Message = response.Result
		}

	}
	w.Header().Set("Content-Type:", "application/json")
	result, _ := json.Marshal(jsonRes)
	w.Write(result)

}
