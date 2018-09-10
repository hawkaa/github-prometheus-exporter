package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hawkaa/github-prometheus-exporter/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/spf13/cobra"

	log "github.com/sirupsen/logrus"
)

const (
	port = 8080
)

var repos = []string{
	"spacemakerai/toadsworth",
	"spacemakerai/toadsworth-sd",
}

func main() {
	var token string
	var rootCmd = &cobra.Command{
		Use: "github-prometheus-exporter",
		Run: func(cmd *cobra.Command, args []string) {
			exporter := exporter.NewExporter(repos, token)
			prometheus.MustRegister(&exporter)

			http.Handle("/metrics/", prometheus.Handler())

			log.Println("Starting github-prometheus-exporter...")
			log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
		},
	}
	rootCmd.Flags().StringVarP(&token, "github-token", "t", os.Getenv("GITHUB_TOKEN"), "Github Token")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
