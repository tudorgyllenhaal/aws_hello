package hello_backend

import (
	"learning/hello_backend/service"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func register(r *server.Hertz) {
	customizeRegister(r)
	r.GET("/ping", service.Ping)
}

func customizeRegister(r *server.Hertz) {
	r.GET("/hello", service.HelloWithTime)
}
