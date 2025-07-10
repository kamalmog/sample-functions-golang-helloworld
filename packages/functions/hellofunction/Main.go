package main

import (
	"context"
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

func Main(params map[string]interface{}) map[string]interface{}{
	return Response {
		Body: "Hello  Kamal ",
		Path: "hellofunction"
	}
}