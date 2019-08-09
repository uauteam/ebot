<%: func GenerateApiQueryFile(m *project.Metadata, buffer *bytes.Buffer) %>
package query

import "github.com/uauteam/ecot/dto/qry"

type <%= m.ModuleTitleName %>Query struct {
    qry.PageQuery

    <% for _, field := range m.QueryFields { %>
    <%== "\n\t" %>
    <%==s field %>
    <% } %>
}