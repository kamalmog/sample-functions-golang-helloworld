package main

import (
//	"context"
)

type Event struct {
	Name string `json:"name"`
}

type Response struct {
	Body string `json:"body"`
}

/*func Main(ctx context.Context, event Event) Response {
	if event.Name == "" {
		event.Name = "stranger"
	}
	version := ctx.Value("function_version").(string)
	return Response {
		Body: "Hello " + event.Name + "! This is function version " + version,
	}
}*/

type Args struct {
	Path string `json:"path"`
	All map[string]interface{}
}

func Main(params Args) map[string]interface{}{
	msg := make(map[string]interface{})
	msg["body"] = "Hello Kamal!"
	return msg
}