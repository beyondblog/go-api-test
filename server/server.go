package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const CONFIG_PATH = "config/"

type JsonResponse struct {
	Code    int
	Message string
	Data    interface{}
}

type ApiRequest struct {
	Desc   string
	Host   string
	Param  map[string]string
	Method HttpMethod
}
type ApiResponse struct {
	Code   int
	Result string
}

func RunTest(request ApiRequest) (reqp *ApiResponse, err error) {
	var response *ApiResponse
	var resp *http.Response
	var e error

	if request.Method == GET {
		resp, e = http.Get(request.Host)
	} else {
		formData := make(url.Values)
		for key, value := range request.Param {
			if len(key) > 0 {
				formData.Add(key, value)
			}
		}
		resp, e = http.PostForm(request.Host, formData)

	}

	if e != nil {
		err = e
		return
	}

	if body, e := ioutil.ReadAll(resp.Body); e != nil {
		err = e
		return
	} else {
		response = &ApiResponse{Code: resp.StatusCode, Result: string(body)}
	}
	return response, err
}

func PrintResponse(response ApiResponse) {
	fmt.Println(response.Code)
	fmt.Println(response.Result)
}

func SaveToConfig(configFile string, apiRequest ApiRequest, index int) error {
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

	hostRequests[index].Desc = apiRequest.Desc
	hostRequests[index].Host = apiRequest.Host
	hostRequests[index].Method = apiRequest.Method
	hostRequests[index].Param = apiRequest.Param

	fout, _ := os.Create(configFile)
	defer fout.Close()
	b, _ := json.Marshal(hostRequests)
	fout.Write(b)
	return nil

}

func WriteToConfig(configFile string, apiRequest ApiRequest) error {
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

//func main() {
//
//	//	requestArray := make([]ApiRequest, 0)
//	//	getInfoApi := ApiRequest{Name: "getInfoApi", Desc: "get User Info", Host: "xxxx", Param: map[string]string{"phone": "xxxx", "passwd": "xxxx"}, Method: POST}
//	//
//	//	requestArray = append(requestArray, getInfoApi)
//	//	for _, request := range requestArray {
//	//		PrintResponse(RunTest(request))
//	//	}
//
//	http.HandleFunc("/", handler)
//
//	http.HandleFunc("/api/", apiHandler)
//
//	http.HandleFunc("/views/", func(w http.ResponseWriter, r *http.Request) {
//		fmt.Println("GET " + r.URL.Path[1:])
//		http.ServeFile(w, r, r.URL.Path[1:])
//
//	})
//
//	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
//		w.Header().Add("server", "go server")
//		start := time.Now()
//		defer func(start time.Time) {
//			fmt.Printf(" %dms \n", ((time.Now().Sub(start).Nanoseconds() * 1.0) / 10000))
//		}(start)
//		fmt.Printf("GET " + r.URL.Path[1:])
//		http.ServeFile(w, r, r.URL.Path[1:])
//	})
//
//	fmt.Println("Server start at 127.0.0.1:8080")
//	http.ListenAndServe(":8080", nil)
//}
