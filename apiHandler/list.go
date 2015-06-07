package api

import (
	"encoding/json"
	. "github.com/beyondblog/go-api-test/server"
	"io/ioutil"
	"net/http"
	"strings"
)

type ApiRequestInfo struct {
	HostName string
	Count    int
	Requests []ApiRequest
}

func List(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir(CONFIG_PATH)
	hosts := []ApiRequestInfo{}

	for _, f := range files {
		if i := strings.LastIndex(f.Name(), "_config.json"); i > 0 {
			//parser
			hostRequests := []ApiRequest{}
			file, err := ioutil.ReadFile(CONFIG_PATH + f.Name())
			if err != nil {
				continue
			}

			if err := json.Unmarshal(file, &hostRequests); err != nil {
				continue
			}
			hosts = append(hosts, ApiRequestInfo{HostName: f.Name()[:i], Count: len(hostRequests), Requests: hostRequests})
		}
	}

	w.Header().Set("Content-Type:", "application/json")
	var jsonRes JsonResponse
	jsonRes.Code = 200
	jsonRes.Data = hosts
	result, _ := json.Marshal(jsonRes)
	w.Write(result)
}
