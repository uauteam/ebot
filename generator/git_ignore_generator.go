package generator

import (
	"bytes"
	"github.com/uauteam/ebot/project"
	"github.com/uauteam/ebot/template"
	"log"
	"os"
)

func GenerateGitIgnoreFile(m *project.Metadata) {
	log.Printf("generating git ignore file")

	// generate .gitignore
	gitIgnoreFile := workDir + ".gitignore"

	if _, err := os.Stat(gitIgnoreFile); !os.IsNotExist(err) {
		log.Printf("git ignore file existed: %s", gitIgnoreFile)
		return
	}

	buf := new(bytes.Buffer)
	template.GenerateGitIgnoreFile(m, buf)

	f, err := os.Create(gitIgnoreFile)
	if err != nil {
		log.Printf(err.Error())
	}

	if _, err := buf.WriteTo(f); err != nil {
		log.Printf(err.Error())
	}
	f.Close()
}
