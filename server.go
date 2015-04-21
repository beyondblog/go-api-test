package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

type httpMethod int

const (
	GET  httpMethod = 0
	POST httpMethod = 1
)

type ApiRequest struct {
	Name   string
	Desc   string
	Host   string
	Param  map[string]string
	Method httpMethod
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
			formData.Add(key, value)
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

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/index.html")
	t.Execute(w, nil)
}

func main() {

	//	requestArray := make([]ApiRequest, 0)
	//	getInfoApi := ApiRequest{Name: "getInfoApi", Desc: "get User Info", Host: "xxxx", Param: map[string]string{"phone": "xxxx", "passwd": "xxxx"}, Method: POST}
	//
	//	requestArray = append(requestArray, getInfoApi)
	//	for _, request := range requestArray {
	//		PrintResponse(RunTest(request))
	//	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/views/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET " + r.URL.Path[1:])
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("server", "go server")
		fmt.Println("GET " + r.URL.Path[1:])
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	fmt.Println("Server start at 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
