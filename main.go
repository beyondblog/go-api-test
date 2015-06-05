package main

import (
	"fmt"
	"net/http"
	//"time"
	api "github.com/beyondblog/go-api-test/apiHandler"
	"html/template"
	"log"
	"os"
)

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

	logFile, _ := os.OpenFile("http-test.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	http.HandleFunc("/", handler)

	http.HandleFunc("/api/add", api.Add)

	http.HandleFunc("/api/list", api.List)

	http.HandleFunc("/views/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GET " + r.URL.Path[1:])
		http.ServeFile(w, r, r.URL.Path[1:])

	})

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("server", "go server")
		//start := time.Now()
		//defer func(start time.Time) {
		//	fmt.Printf(" %dms \n", ((time.Now().Sub(start).Nanoseconds() * 1.0) / 10000))
		//}(start)
		fmt.Println("GET " + r.URL.Path[1:])
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	fmt.Println("Server start at 127.0.0.1:8080")
	log.Println("Server start at 127.0.0.1:8080")
	http.ListenAndServe(":8080", nil)
}
