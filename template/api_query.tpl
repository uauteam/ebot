<%: func GenerateApiQueryFile(m *project.Metadata, buffer *bytes.Buffer) %>
package query


type <%= m.ModuleTitleName %>Query struct {
    <% for _, field := range m.QueryFields { %>
    <%== "\n\t" %>
    <%==s field %>
    <% } %>
}