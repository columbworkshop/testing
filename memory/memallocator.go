package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("Start program")
	//
	NUM_ITER, NUM_ITER_found := os.LookupEnv("NUM_ITER")
	MEM_SIZE, MEM_SIZE_found := os.LookupEnv("MEM_SIZE")
	//
	if !NUM_ITER_found {
		logger.Error("ENV VARIABLE is not set", "NUM_ITER", NUM_ITER)
		os.Exit(1)
	} else {
		logger.Info("ENV VARIABLE", "NUM_ITER", NUM_ITER)
	}
	//
	if !MEM_SIZE_found {
		logger.Error("ENV VARIABLE is not set", "MEM_SIZE", MEM_SIZE)
		os.Exit(1)
	} else {
		logger.Info("ENV VARIABLE", "MEM_SIZE", MEM_SIZE)
	}
	//
	var m runtime.MemStats
	memsize, err := strconv.Atoi(MEM_SIZE)
	numiter, err := strconv.Atoi(NUM_ITER)
	if err != nil {
		fmt.Println("Error converting string:", err)
		return
	}

	/*
		Allocate memory for some time
	*/
	for i := range numiter {
		logger.Info("Memory allocation", "iteration", i+1, "memory_chunk", memsize)
		//
		memobj := make([]byte, memsize*1024*1024*1024)
		memobj[0] = 1
		//
		runtime.ReadMemStats(&m)
		logger.Info("Runtime memory allocated", "size_bytes", m.Alloc)
		time.Sleep(5 * 60 * time.Second)
		logger.Info("Releasing memory", "iteration", i+1, "memory_chunk", memsize)
		memobj = nil
		time.Sleep(120 * time.Second)
	}
}
