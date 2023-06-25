package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"io/ioutil"
)

// Config struct represents the structure of the JSON config file
type Config struct {
	Protocol       string `json:"protocol"`
	Host           string `json:"host"`
	Port           int    `json:"port"`
	MessageHandler string `json:"message_handler"`
}

func main() {
	// Parse command-line flags
	configFile := flag.String("config", "", "path to the JSON config file")
	helpFlag := flag.Bool("help", false, "print help message")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	if *helpFlag {
		flag.Usage()
		return
	}

	if *configFile == "" {
		log.Fatal("Please provide a JSON config file using the --config flag")
	}

	// Read the config file
	configData, err := ReadConfigFile(*configFile)
	if err != nil {
		log.Fatal("Error reading config file:", err)
	}

	// Boot server based on config
	switch configData.Protocol {
	case "tcp":
		fmt.Println("Booting TCP server...")
		// Implement your TCP server logic here using the host, port, and message handler
		// Example: StartTCPServer(configData.Host, configData.Port, configData.MessageHandler)
	case "udp":
		fmt.Println("Booting UDP server...")
		// Implement your UDP server logic here using the host, port, and message handler
		// Example: StartUDPServer(configData.Host, configData.Port, configData.MessageHandler)
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
