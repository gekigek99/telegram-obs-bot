package config

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"tob/lib/model"

	"github.com/google/shlex"
)

var configFileName string = "tob-config.json"

var Default *Configuration = &Configuration{}
var Runtime *Configuration = &Configuration{}

type Configuration struct {
	model.Configuration
}

func Load() error {
	err := Default.loadDefault()
	if err != nil {
		log.Fatal(err)
	}

	err = Runtime.loadRuntime(Default)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (c *Configuration) loadDefault() error {
	// get working directory
	cwdPath, err := os.Getwd()
	if err != nil {
		return err
	}

	// read config file
	configFilePath := filepath.Join(cwdPath, configFileName)
	fmt.Printf("reading config file: \"%s\"\n", configFilePath)
	configData, err := os.ReadFile(configFilePath)
	if err != nil {
		return err
	}

	// write data to config variable
	err = json.Unmarshal(configData, &c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Configuration) loadRuntime(def *Configuration) error {
	// initialize config to base
	*c = *def

	flag.StringVar(&c.BotToken, "token", c.BotToken, "Specify telegram bot token.")

	// specify the usage when there is an error in the arguments
	flag.Usage = func() {
		// not using errco.NewLogln since log time is not needed
		fmt.Println("Usage of tob:")
		flag.PrintDefaults()
	}

	// join os provided args and split them again with shlex.
	args, err := shlex.Split(strings.Join(os.Args[1:], " "))
	if err != nil {
		return err
	}
	flag.CommandLine.Parse(args)

	if c.BotToken != "" {
		// already set
	} else if os.Getenv("BOT_TOKEN") != "" {
		c.BotToken = os.Getenv("BOT_TOKEN")
	} else {
		return fmt.Errorf("bot token unknown")
	}

	return nil
}
