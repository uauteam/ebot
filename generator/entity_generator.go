package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateEntityFile(m *project.Metadata) {
	log.Printf("generating entity file")

	// entity dir
	entityDir := workDir + "entity/"

	if _, err := os.Stat(entityDir); os.IsNotExist(err) {
		os.MkdirAll(entityDir, 0755)
	}

	// generate entity/{moduleName}.go
	entityFile := entityDir + moduleName + ".go"
	if _, err := os.Stat(entityFile); !os.IsNotExist(err) {
		log.Printf("entity file existed: %s", entityFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateEntityFile(m, buf)

	f, err := os.Create(entityFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
