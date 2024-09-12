package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"log/slog"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithDeadline(ctx, time.Now().Add(1000 * time.Millisecond))

	defer cancel()
	slog.

	
}

