package server

import (
	"github.com/mholt/binding"
)

type ListForm struct {
}

type AppForm struct {
	Index  int
	Host   string
	Desc   string
	Method HttpMethod
	Param  []struct {
		Key   string
		Value string
	}
}

func (l *AppForm) FieldMap() binding.FieldMap {
	return binding.FieldMap{
		&l.Host:   "host",
		&l.Method: "method",
		&l.Desc:   "desc",
		&l.Param:  "param",
		&l.Index:  "index",
	}
}

type HttpMethod int

const (
	GET  HttpMethod = 0
	POST HttpMethod = 1
)
