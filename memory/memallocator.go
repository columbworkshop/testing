package main

import (
	"log/slog"
	"os"
	"runtime"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Start program")
	//
	var m runtime.MemStats
	var memchank_gb = 1
	//
	for i := range 2 {
		logger.Info("Memory allocation", "iteration", i+1, "memory_chunk", memchank_gb)
		//
		memobj := make([]byte, memchank_gb*1024*1024*1024)
		memobj[0] = 1
		//
		runtime.ReadMemStats(&m)
		logger.Info("Runtime memory allocated", "size_bytes", m.Alloc)
		time.Sleep(30 * time.Second)
		logger.Info("Releasing memory", "iteration", i+1, "memory_chunk", memchank_gb)
		memobj = nil
		time.Sleep(120 * time.Second)
	}
}
