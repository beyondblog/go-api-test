package api

import (
	"encoding/json"
	. "github.com/beyondblog/go-api-test/server"
	"github.com/mholt/binding"
	"log"
	"net/http"
	"regexp"
)

func Add(w http.ResponseWriter, r *http.Request) {
	log.Println("add request")
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

	re := regexp.MustCompile("^(http://|https://)?((?:[A-Za-z0-9]+-[A-Za-z0-9]+|[A-Za-z0-9]+)\\.)+([A-Za-z]+)[/\\?\\:]?.*$")

	if len(apiRequest.Host) == 0 {
		jsonRes.Code = 400
		jsonRes.Message = "host error"
	} else if r := re.FindString(apiRequest.Host); r == "" {
		jsonRes.Code = 400
		jsonRes.Message = "host format error"
	} else {

		//check host

		for _, param := range appForm.Param {
			if len(param.Key) > 0 {
				apiRequest.Param[param.Key] = param.Value
			}
		}

		//check host config exist
		configFile := CONFIG_PATH + apiRequest.Host + "_config.json"

		jsonRes.Code = 200
		if err := WriteToConfig(configFile, apiRequest); err != nil {
			jsonRes.Code = 400
			jsonRes.Message = err.Error()
		}

	}
	w.Header().Set("Content-Type:", "application/json")
	result, _ := json.Marshal(jsonRes)
	w.Write(result)
}
