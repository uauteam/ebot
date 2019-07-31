package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateApiQueryFile(m *project.Metadata) {
	log.Printf("generating api query file")

	// api/query dir
	apiQueryDir := workDir + "api/query/"

	if _, err := os.Stat(apiQueryDir); os.IsNotExist(err) {
		os.MkdirAll(apiQueryDir, 0755)
	}

	// generate api/query/{moduleName}.go
	apiQueryFile := apiQueryDir + moduleName + ".go"

	if _, err := os.Stat(apiQueryFile); !os.IsNotExist(err) {
		log.Printf("api query file existed: %s", apiQueryFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateApiQueryFile(m, buf)

	f, err := os.Create(apiQueryFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
