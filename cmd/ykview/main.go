package main

import (
	"log/slog"
	"os"

	"github.com/YukiYuigishi/ykview/internal/preview"
)

func main() {
	if err := preview.Run(os.Args[1:], os.Args[0]); err != nil {
		slog.Error("ykview exited", "error", err)
		os.Exit(1)
	}
}
