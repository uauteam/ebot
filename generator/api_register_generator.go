package generator

import (
	"bytes"
	"github.com/uauteam/ago/file"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func GenerateApiRegisterFile(m *project.Metadata) {
	log.Printf("generating api register file")

	// api dir
	apiDir := workDir + "api/"

	if _, err := os.Stat(apiDir); os.IsNotExist(err) {
		os.MkdirAll(apiDir, 0755)
	}

	// generate api/api.go
	apiRegisterFile := apiDir + "api.go"

	// api register file not generated
	if _, err := os.Stat(apiRegisterFile); os.IsNotExist(err) {
		log.Printf("api register file is not generated: %s", apiRegisterFile)

		buf := new(bytes.Buffer)
		template.GenerateApiRegisterFile(m, buf)

		f, err := os.Create(apiRegisterFile)
		if err != nil {
			log.Printf(err.Error())
		}

		if _, err := buf.WriteTo(f); err != nil {
			log.Printf(err.Error())
		}
		f.Close()

		return
	}

	// api register file has already generated
	log.Printf("api register file has already generated: %s", apiRegisterFile)
	lines, err := file.ContentLines(apiRegisterFile)
	if err != nil {
		log.Printf("error getting file content lines: %s", apiRegisterFile)
	}

	buf := new(bytes.Buffer)
	template.GenerateApiRegisterModuleFile(m, buf)

	// insert registering code of module api before the last `return` statement in api/api.gp
	lines = append(lines[:len(lines)-2], append([]string{"\n\t" + buf.String()}, lines[len(lines)-2:]...)...)

	s := strings.Join(lines, "\n")

	err = ioutil.WriteFile(apiRegisterFile, []byte(s), 0644)
	if err != nil {
		log.Printf("error writing content lines to file: %s", apiRegisterFile)
	}
}
