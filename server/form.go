package server

type ListForm struct {
}

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
