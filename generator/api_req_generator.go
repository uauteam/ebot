package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateApiReqFile(m *project.Metadata) {
	log.Printf("generating api req file")

	// api/req dir
	apiReqDir := workDir + "api/req/"

	if _, err := os.Stat(apiReqDir); os.IsNotExist(err) {
		os.MkdirAll(apiReqDir, 0755)
	}

	// generate api/req/{moduleName}.go
	apiReqFile := apiReqDir + moduleName + ".go"

	if _, err := os.Stat(apiReqFile); !os.IsNotExist(err) {
		log.Printf("api req file existed: %s", apiReqFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateApiReqFile(m, buf)

	f, err := os.Create(apiReqFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
