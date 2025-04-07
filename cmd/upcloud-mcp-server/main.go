package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	stdlog "log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/UpCloudLtd/upcloud-go-api/v8/upcloud/client"
	iolog "github.com/marinx/upcloud-mcp-server/pkg/log"
	"github.com/marinx/upcloud-mcp-server/pkg/upcloud"
	"github.com/mark3labs/mcp-go/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "version"
	commit  = "commit"
	date    = "date"
)

var (
	rootCmd = &cobra.Command{
		Use:     "server",
		Short:   "Upcloud MCP Server",
		Long:    `A Upcloud MCP server that handles various tools and resources.`,
		Version: fmt.Sprintf("%s (%s) %s", version, commit, date),
	}

	stdioCmd = &cobra.Command{
		Use:   "stdio",
		Short: "Start stdio server",
		Long:  `Start a server that communicates via standard input/output streams using JSON-RPC messages.`,
		Run: func(_ *cobra.Command, _ []string) {
			logFile := viper.GetString("log-file")
			readOnly := viper.GetBool("read-only")
			prettyPrintJSON := viper.GetBool("pretty-print-json")
			logger, err := initLogger(logFile)
			if err != nil {
				stdlog.Fatal("Failed to initialize logger:", err)
			}
			logCommands := viper.GetBool("enable-command-logging")
			cfg := runConfig{
				readOnly:        readOnly,
				logger:          logger,
				logCommands:     logCommands,
				prettyPrintJSON: prettyPrintJSON,
			}
			if err := runStdioServer(cfg); err != nil {
				stdlog.Fatal("failed to run stdio server:", err)
			}
		},
	}
)

func init() {
	cobra.OnInitialize(initConfig)

	// Add global flags that will be shared by all commands
	rootCmd.PersistentFlags().Bool("read-only", false, "Restrict the server to read-only operations")
	rootCmd.PersistentFlags().String("log-file", "", "Path to log file")
	rootCmd.PersistentFlags().Bool("enable-command-logging", false, "When enabled, the server will log all command requests and responses to the log file")
	rootCmd.PersistentFlags().Bool("pretty-print-json", false, "Pretty print JSON output")

	// Bind flag to viper
	_ = viper.BindPFlag("read-only", rootCmd.PersistentFlags().Lookup("read-only"))
	_ = viper.BindPFlag("log-file", rootCmd.PersistentFlags().Lookup("log-file"))
	_ = viper.BindPFlag("enable-command-logging", rootCmd.PersistentFlags().Lookup("enable-command-logging"))
	_ = viper.BindPFlag("pretty-print-json", rootCmd.PersistentFlags().Lookup("pretty-print-json"))

	// Add subcommands
	rootCmd.AddCommand(stdioCmd)
}

func initConfig() {
	// Initialize Viper configuration
	viper.SetEnvPrefix("APP")
	viper.AutomaticEnv()
}

func initLogger(outPath string) (*log.Logger, error) {
	if outPath == "" {
		return log.New(), nil
	}

	file, err := os.OpenFile(outPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	logger := log.New()
	logger.SetLevel(log.DebugLevel)
	logger.SetOutput(file)

	return logger, nil
}

type runConfig struct {
	readOnly        bool
	logger          *log.Logger
	logCommands     bool
	prettyPrintJSON bool
}

// JSONPrettyPrintWriter is a Writer that pretty prints input to indented JSON
type JSONPrettyPrintWriter struct {
	writer io.Writer
}

func (j JSONPrettyPrintWriter) Write(p []byte) (n int, err error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, p, "", "\t"); err != nil {
		return 0, err
	}
	return j.writer.Write(prettyJSON.Bytes())
}

func runStdioServer(cfg runConfig) error {
	// Create app context
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	username := os.Getenv("UPCLOUD_USERNAME")
	if username == "" {
		cfg.logger.Fatal("UPCLOUD_USERNAME not set")
	}
	password := os.Getenv("UPCLOUD_PASSWORD")
	if password == "" {
		cfg.logger.Fatal("UPCLOUD_PASSWORD not set")
	}

	upClient := client.New(username, password, client.WithTimeout(time.Second*30))
	upClient.UserAgent = fmt.Sprintf("upcloud-mcp-server/%s", version)

	// Create
	upServer := upcloud.NewServer(upClient, version, cfg.readOnly)
	stdioServer := server.NewStdioServer(upServer)

	stdLogger := stdlog.New(cfg.logger.Writer(), "stdioserver", 0)
	stdioServer.SetErrorLogger(stdLogger)

	// Start listening for messages
	errC := make(chan error, 1)
	go func() {
		in, out := io.Reader(os.Stdin), io.Writer(os.Stdout)

		if cfg.logCommands {
			loggedIO := iolog.NewIOLogger(in, out, cfg.logger)
			in, out = loggedIO, loggedIO
		}

		if cfg.prettyPrintJSON {
			out = JSONPrettyPrintWriter{writer: out}
		}
		errC <- stdioServer.Listen(ctx, in, out)
	}()

	// Output upcloud-mcp-server string
	_, _ = fmt.Fprintf(os.Stderr, "Upcloud MCP Server running on stdio\n")

	// Wait for shutdown signal
	select {
	case <-ctx.Done():
		cfg.logger.Infof("shutting down server...")
	case err := <-errC:
		if err != nil {
			return fmt.Errorf("error running server: %w", err)
		}
	}

	return nil
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
