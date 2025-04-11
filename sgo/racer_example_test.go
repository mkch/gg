package sgo_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/mkch/gg/sgo"
)

func ping(ctx context.Context, url string) string {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return url
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return ""
	}
	return url
}

func ExampleRace_ping() {
	winner := sgo.Race[string](context.Background(), 0, nil,
		func(ctx context.Context) string { return ping(ctx, "https://goproxy.cn") },
		func(ctx context.Context) string { return ping(ctx, "https://goproxy.io") })
	if winner == "" {
		fmt.Println("no winner")
	} else {
		fmt.Printf("The winner is %v", winner)
	}
}
