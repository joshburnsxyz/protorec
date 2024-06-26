package main

import (
	"encoding/json"
	"fmt"
	flag "github.com/spf13/pflag"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/joshburnsxyz/protorec/servers"
)

// Config struct represents the structure of the JSON config file
type Config struct {
	Protocol            string `json:"protocol"`
	Host                string `json:"host"`
	Port                int    `json:"port"`
	MessageHandler      string `json:"message_handler"`
	LogFilePath         string `json:"logfile"`
	MessageBufferLength int    `json:"message_buffer_length"`
}

// global variable to store config file path
var configFile string
var helpFlag bool

func init() {
	flag.StringVarP(&configFile, "config", "c", "", "path to the JSON config file")
	flag.BoolVarP(&helpFlag, "help", "h", false, "print help message")

	// Custom usage output
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()

	if helpFlag {
		flag.Usage()
		return
	}

	if configFile == "" {
		log.Fatal("Please provide a JSON config file using the --config flag")
	}

	// Read the config file
	configData, err := ReadConfigFile(configFile)
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	// Validate the config data and apply defaults if necessary
	if err := ValidateConfigData(&configData); err != nil {
		log.Fatal("Invalid config data:", err)
	}

	// Create the log file
	logFile, err := os.Create(configData.LogFilePath)
	if err != nil {
		log.Fatal("Error creating log file:", err)
	}
	defer logFile.Close()

	// Create a multi-writer for logging to both stdout and the log file
	logOutput := io.MultiWriter(os.Stdout, logFile)

	// Set the logger output to the multi-writer
	log.SetOutput(logOutput)

	// Boot server based on config
	switch configData.Protocol {
	case "tcp":
		log.Println("Booting TCP server...")
		// Implement your TCP server logic here using the host, port, and message handler
		servers.StartTCPServer(configData.Host, configData.Port, configData.MessageHandler, configData.MessageBufferLength)
	case "udp":
		log.Println("Booting UDP server...")
		// Implement your UDP server logic here using the host, port, and message handler
		servers.StartUDPServer(configData.Host, configData.Port, configData.MessageHandler, configData.MessageBufferLength)
	default:
		log.Fatal("Invalid protocol specified in the config file")
	}
}

// ReadConfigFile reads the JSON config file and returns the config data
func ReadConfigFile(configFile string) (Config, error) {
	var configData Config
	file, err := ioutil.ReadFile(configFile)
	if err != nil {
		return configData, err
	}
	err = json.Unmarshal(file, &configData)
	if err != nil {
		return configData, err
	}
	return configData, nil
}

// ValidateConfigData validates the config data and applies defaults if necessary
func ValidateConfigData(configData *Config) error {
	// Store any generated errors in a slice
	var errs []string

	// Dummy object to hold default config values
	// apply if value is missing from config
	defaultConfig := Config{
		Protocol:            "tcp",
		Host:                "0.0.0.0",
		Port:                55011,
		LogFilePath:         "server.log",
		MessageBufferLength: 1024,
	}

	if configData.Protocol == "" {
		configData.Protocol = defaultConfig.Protocol
	}

	if configData.Host == "" {
		configData.Host = defaultConfig.Host
	}

	if configData.Port == 0 {
		configData.Port = defaultConfig.Port
	}

	if configData.LogFilePath == "" {
		configData.LogFilePath = defaultConfig.LogFilePath
	}

	if configData.MessageBufferLength == 0 {
		configData.MessageBufferLength = defaultConfig.MessageBufferLength
	}

	// No default value for message handler, just generate an error
	if configData.MessageHandler == "" {
		errs = append(errs, "message_handler config is missing")
	}

	// If there is more than 0 errors in the slice return
	// a non-nil error value.
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	} else {
		return nil
	}

}
