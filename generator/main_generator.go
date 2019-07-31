package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateMainFile(m *project.Metadata) {
	log.Printf("generating the main file")

	// generate main.go
	mainFile := workDir + "main.go"

	if _, err := os.Stat(mainFile); !os.IsNotExist(err) {
		log.Printf("main file existed: %s", mainFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateMainFile(m, buf)

	f, err := os.Create(mainFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
