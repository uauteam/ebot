package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateConfigFile(m *project.Metadata) {
	log.Printf("generating app config file")

	// config dir
	configDir := workDir + "config/"

	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		os.MkdirAll(configDir, 0755)
	}

	// generate config/config.go
	configFile := configDir + "config.go"

	if _, err := os.Stat(configFile); !os.IsNotExist(err) {
		log.Printf("config file existed: %s", configFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateConfigFile(m, buf)

	f, err := os.Create(configFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
