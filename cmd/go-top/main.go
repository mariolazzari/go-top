package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/mariolazzari/go-top/internal/config"
	"github.com/mariolazzari/go-top/internal/terminal"
)

func main() {
	// Parse and validate command-line arguments
	cfg, err := config.ParseFlags()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
		os.Exit(1)
	}

	// Global context with graceful shutdown on SIGINT (Ctrl+C) / SIGTERM
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// TODO
	// Initialize and start the centralized metrics Hub in the background
	// monitor := monitor,New()
	//go monitor.Start(ctx, cfg.Interval)

	// Dispatch based on the selected operating mode
	switch cfg.Mode {
	case "text":
		terminal.ClearScreen()
		fmt.Printf("Starting go-top in text mode (interval: %v)...", cfg.Interval)
		//client.RunTextClient(ctx, hub)
		<-ctx.Done()

	case "web":
		fmt.Printf("Web mode selected (interval: %v). REST/SSE server initialization pending...\n", cfg.Interval)
		<-ctx.Done()

	case "tui":
		fmt.Printf("TUI mode selected (interval: %v). Interactive terminal UI pending...\n", cfg.Interval)
		<-ctx.Done()
	}

	fmt.Println("\nExiting go-top gracefully. Goodbye!")
}
