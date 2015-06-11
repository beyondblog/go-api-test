package api

import (
	"encoding/json"
	. "github.com/beyondblog/go-api-test/server"
	"github.com/mholt/binding"
	"log"
	"net/http"
)

func Save(w http.ResponseWriter, r *http.Request) {
	log.Println("save request")
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
	var index = appForm.Index
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

		//check host config exist
		configFile := CONFIG_PATH + apiRequest.Host + "_config.json"

		jsonRes.Code = 200
		if err := SaveToConfig(configFile, apiRequest, index); err != nil {
			jsonRes.Code = 400
			jsonRes.Message = err.Error()
		}

	}
	w.Header().Set("Content-Type:", "application/json")
	result, _ := json.Marshal(jsonRes)
	w.Write(result)
}
