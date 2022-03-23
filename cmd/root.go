package cmd

import (
	"fmt"
	cnf "github.com/project-ginza/notifications/config"
	logging "github.com/project-ginza/notifications/log"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var loggingCh = logging.LoggingCh

func Root() *cobra.Command {
	rootCmd := cobra.Command{
		Use: "run",
		Run: run,
	}
	return &rootCmd
}

func run(cmd *cobra.Command, args []string) {
	sigs := make(chan os.Signal, 1)

	go loggingCh.LogWorker()
	logLevel := cnf.GlobalConfig.LogLevel
	setLogConfig(logLevel)

	<-sigs
}

func setLogConfig(logLevel string) {
	f, err := os.OpenFile("/applogs/output.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	log.SetOutput(f)

	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	default:
		msg := fmt.Sprintf("Improper log level is given. %s", logLevel)
		loggingCh.PushErrorMessageToChannel(msg)
	}
}
