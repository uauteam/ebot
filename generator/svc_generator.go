package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateSvcFile(m *project.Metadata) {
	log.Printf("generating svc file")

	// svc dir
	svcDir := workDir + "svc/"

	if _, err := os.Stat(svcDir); os.IsNotExist(err) {
		os.MkdirAll(svcDir, 0755)
	}

	// generate svc/{moduleName}.go
	svcFile := svcDir + moduleName + ".go"
	if _, err := os.Stat(svcFile); !os.IsNotExist(err) {
		log.Printf("svc file existed: %s", svcFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateServiceFile(m, buf)

	f, err := os.Create(svcFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
