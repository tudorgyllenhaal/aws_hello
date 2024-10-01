package service

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
)

func HelloWithTime(ctx context.Context, c *app.RequestContext) {
	c.JSON(200, map[string]string{
		"message": "hello back",
		"time":    time.Now().Format(time.RFC3339),
	})

}
