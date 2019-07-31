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

func GenerateRepoDatabaseFile(m *project.Metadata) {
	log.Printf("generating repo database file")

	// repo dir
	repoDir := workDir + "repo/"

	if _, err := os.Stat(repoDir); os.IsNotExist(err) {
		os.MkdirAll(repoDir, 0755)
	}

	// generate repo/database.go
	repoDatabaseFile := repoDir + "database.go"

	// repo database file not generated
	if _, err := os.Stat(repoDatabaseFile); os.IsNotExist(err) {
		log.Printf("repo database file is not generated: %s", repoDatabaseFile)

		buf := new(bytes.Buffer)
		template.GenerateRepoDatabaseFile(m, buf)

		f, err := os.Create(repoDatabaseFile)
		if err != nil {
			log.Printf(err.Error())
		}

		if _, err := buf.WriteTo(f); err != nil {
			log.Printf(err.Error())
		}
		f.Close()

		return
	}

	// repo database file has already generated
	log.Printf("repo database file has already generated: %s", repoDatabaseFile)
	lines, err := file.ContentLines(repoDatabaseFile)
	if err != nil {
		log.Printf("error getting file content lines: %s", repoDatabaseFile)
	}

	buf := new(bytes.Buffer)
	template.GenerateRepoDatabaseModuleFile(m, buf)

	// insert appending code of module entity before the last `return` statement in repo/database.gp
	lines = append(lines[:len(lines)-2], append([]string{buf.String()}, lines[len(lines)-2:]...)...)

	s := strings.Join(lines, "\n")

	err = ioutil.WriteFile(repoDatabaseFile, []byte(s), 0644)
	if err != nil {
		log.Printf("error writing content lines to file: %s", repoDatabaseFile)
	}
}
