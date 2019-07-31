package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateAppGlobalFile(m *project.Metadata) {
	log.Printf("generating app global file")

	// app dir
	appDir := workDir + "app/"

	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		os.MkdirAll(appDir, 0755)
	}

	// generate app/global.go
	appGlobalFile := appDir + "global.go"

	if _, err := os.Stat(appGlobalFile); !os.IsNotExist(err) {
		log.Printf("app global file existed: %s", appGlobalFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateAppGlobalFile(m, buf)

	f, err := os.Create(appGlobalFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
