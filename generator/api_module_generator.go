package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateApiModuleFile(m *project.Metadata) {
	log.Printf("generating api module file")

	// api dir
	apiModuleDir := workDir + "api/"

	if _, err := os.Stat(apiModuleDir); os.IsNotExist(err) {
		os.MkdirAll(apiModuleDir, 0755)
	}

	// generate api/{moduleName}.go
	apiRespFile := apiModuleDir + moduleName + ".go"

	if _, err := os.Stat(apiRespFile); !os.IsNotExist(err) {
		log.Printf("api module file existed: %s", apiRespFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateApiModuleFile(m, buf)

	f, err := os.Create(apiRespFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
