package main

import (
	"context"
	"net/http"
)

func handle(req *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	if req.Method != http.MethodGet {

		return
	}
	defer cancel()
	doHandle(ctx)
}

func doHandle(ctx context.Context) {
  handle(nil)
	println(ctx)
}
