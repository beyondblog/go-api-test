package api

import (
	"encoding/json"
	"errors"
	. "github.com/beyondblog/go-api-test/server"
	"github.com/mholt/binding"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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
		if err := writeToConfig(configFile, apiRequest); err != nil {
			jsonRes.Code = 400
			jsonRes.Message = err.Error()
		}

	}
	w.Header().Set("Content-Type:", "application/json")
	result, _ := json.Marshal(jsonRes)
	w.Write(result)
}

func writeToConfig(configFile string, apiRequest ApiRequest) error {
	hostRequests := []ApiRequest{}
	if _, err := os.Stat(configFile); err == nil {
		//file is exist append
		file, err := ioutil.ReadFile(configFile)
		if err != nil {
			return errors.New("config error")
		}

		if err := json.Unmarshal(file, &hostRequests); err != nil {
			return errors.New("config error")
		}
	}

	hostRequests = append(hostRequests, apiRequest)
	fout, _ := os.Create(configFile)
	defer fout.Close()
	b, _ := json.Marshal(hostRequests)
	fout.Write(b)
	return nil
}
