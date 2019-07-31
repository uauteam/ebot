package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateApiRespFile(m *project.Metadata) {
	log.Printf("generating api resp file")

	// api/req dir
	apiRespDir := workDir + "api/resp/"

	if _, err := os.Stat(apiRespDir); os.IsNotExist(err) {
		os.MkdirAll(apiRespDir, 0755)
	}

	// generate api/resp/{moduleName}.go
	apiRespFile := apiRespDir + moduleName + ".go"

	if _, err := os.Stat(apiRespFile); !os.IsNotExist(err) {
		log.Printf("api resp file existed: %s", apiRespFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateApiRespFile(m, buf)

	f, err := os.Create(apiRespFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
