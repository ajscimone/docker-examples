package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	// Create a Zap logger configuration
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Error creating logger: %v\n", err)
		return
	}
	defer logger.Sync()

	// Log "Hello, World!"
	logger.Info("Hello, World!")
}