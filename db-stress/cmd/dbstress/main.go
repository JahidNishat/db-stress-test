package main

import (
	"db-stress/internal/runner"
	"db-stress/internal/ui"
	"db-stress/internal/workload"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	workers  int
	duration string
	dsn      string
)

var rootCmd = &cobra.Command{
	Use:   "dbstress",
	Short: "A database stress testing tool",
	Run: func(cmd *cobra.Command, args []string) {
		dur, err := time.ParseDuration(duration)
		if err != nil {
			fmt.Printf("Invalid duration: %v\n", err)
			os.Exit(1)
		}

		cfg := runner.Config{
			Workers:  workers,
			Duration: dur,
			DSN:      dsn,
		}
		wl := &workload.HTTPWorkload{}

		statsCh := make(chan runner.Stats, 10000)

		go func() {
			_ = runner.Run(cfg, wl, statsCh)
		}()

		p := tea.NewProgram(ui.InitialModel(statsCh))
		if _, err := p.Run(); err != nil {
			fmt.Println("Error running UI: ", err)
			os.Exit(1)
		}

		// if err := runner.Run(cfg, wl, statsCh); err != nil {
		// 	fmt.Println("Error: ", err)
		// 	os.Exit(1)
		// }
	},
}

func init() {
	rootCmd.Flags().IntVarP(&workers, "workers", "w", 10, "Number of concurrent workers")
	rootCmd.Flags().StringVarP(&duration, "duration", "d", "10s", "Duration of the test")
	rootCmd.Flags().StringVar(&dsn, "dsn", "", "Postgres Connection String")
}

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
