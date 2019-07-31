<%: func GenerateMainFile(m *project.Metadata, buffer *bytes.Buffer) %>
package main

import (
	"github.com/sirupsen/logrus"
	"github.com/uauteam/ecot"
	<%= m.RepoName %> "<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/config"
)

func main() {

	a := arre.New()

	if err := a.Register(<%= m.RepoName %>.Config); err != nil {
		logrus.Fatal(err)
	}

	a.Logger.Fatal(a.Start(":1323"))
}