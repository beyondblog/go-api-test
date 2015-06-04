package server

import (
	"fmt"
	"github.com/mholt/binding"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AppForm struct {
	Host   string
	Desc   string
	Method HttpMethod
	Param  []struct {
		Key   string
		Value string
	}
}

type HttpMethod int

const (
	GET  HttpMethod = 0
	POST HttpMethod = 1
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

func RunTest(request ApiRequest) ApiResponse {
	var response ApiResponse
	if request.Method == GET {
		http.Get(request.Host)
	} else {
		formData := make(url.Values)
		for key, value := range request.Param {
			if len(key) > 0 {
				formData.Add(key, value)
			}
		}
		resp, _ := http.PostForm(request.Host, formData)
		body, _ := ioutil.ReadAll(resp.Body)
		response.Code = resp.StatusCode
		response.Result = string(body)
	}
	return response
}

func PrintResponse(response ApiResponse) {
	fmt.Println(response.Code)
	fmt.Println(response.Result)
}

func (l *AppForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&l.Host:   "host",
		&l.Method: "method",
		&l.Desc:   "desc",
		&l.Param:  "param",
	}
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
