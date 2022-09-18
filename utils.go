package main

import (
	"runtime"

	"go.uber.org/zap"
)

func handleError(e error) {
	if e != nil {
		logger.Error("Error", zap.Error(e))
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	logger.Info("Alloc", zap.Uint64("Mem", bToMb(m.Alloc)))
	logger.Info("TotalAlloc", zap.Uint64("Mem", bToMb(m.TotalAlloc)))
	logger.Info("Sys", zap.Uint64("Mem", bToMb(m.Sys)))
	logger.Info("NumGC", zap.Uint32("Mem", m.NumGC))
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func Recovery() {
	if r := recover(); r != nil {
		logger.Info("Recovered in f", zap.Any("r", r))
	}
}

func initLogger() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
}
