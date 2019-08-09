<%: func GenerateMainFile(m *project.Metadata, buffer *bytes.Buffer) %>
package main

import (
	<%= m.RepoName %> "<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/config"
	"github.com/uauteam/ecot"
)

func main() {

	e := ecot.New()

	if err := e.Register(<%= m.RepoName %>.Config); err != nil {
		e.Logger.Fatal(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}