package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(1000 * time.Millisecond))

	defer cancel()

	doRequest(ctx, "http://ya.ru")

}


func doRequest(ctx context.Context, requestStr string) {
	request, _ := http.NewRequest(http.MethodGet, requestStr, nil)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Printf("response completed. status code = %d\n", response.StatusCode)
	case <-ctx.Done():
		fmt.Println("time to long")
	}
}
