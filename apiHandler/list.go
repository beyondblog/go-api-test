package api

import (
	"encoding/json"
	"fmt"
	. "github.com/beyondblog/go-api-test/server"
	"io/ioutil"
	"net/http"
	"strings"
)

type ApiRequestInfo struct {
	HostName string
	Count    int
}

func List(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir(CONFIG_PATH)
	hosts := []ApiRequestInfo{}

	for _, f := range files {
		if i := strings.LastIndex(f.Name(), "_config.json"); i > 0 {
			//解析
			hostRequests := []ApiRequest{}
			file, err := ioutil.ReadFile(CONFIG_PATH + f.Name())
			if err != nil {
				fmt.Println("123")
				continue
			}

			if err := json.Unmarshal(file, &hostRequests); err != nil {
				fmt.Println("456")
				continue
			}
			hosts = append(hosts, ApiRequestInfo{HostName: f.Name()[:i], Count: len(hostRequests)})
		}
	}

	w.Header().Set("Content-Type:", "application/json")
	var jsonRes JsonResponse
	jsonRes.Code = 200
	jsonRes.Data = hosts
	result, _ := json.Marshal(jsonRes)
	w.Write(result)
}
