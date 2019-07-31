<%: func GenerateEntityFile(m *project.Metadata, buffer *bytes.Buffer) %>
package entity

import (
	"github.com/jinzhu/gorm"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/app"
)

// module model struct
type <%= m.ModuleTitleName %> struct {
	gorm.Model
	<% for _, field := range m.EntityFields { %>
	<%== "\n\t" %>
    <%==s field %>
	<% } %>
}

// the database name of the modle struct(table)
func (<%= m.ModuleTitleName %>) DBName()string {
	return "db_" + app.Name
}

// fields that will be ignored when creating or updating the struct
// e.g. []string{"username", "email", "password"}
func (<%= m.ModuleTitleName %>) ProtectedFields()[]string{
	return <%==s m.ProtectedFieldsStr %>
}