package hello_backend

import (
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	h := server.Default()

	h.Spin()
}
