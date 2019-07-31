<%: func GenerateApiRespFile(m *project.Metadata, buffer *bytes.Buffer) %>
package resp


type <%= m.ModuleTitleName %>Resp struct {
    <% for _, field := range m.ResponseFields { %>
    <%== "\n\t" %>
    <%==s field %>
    <% } %>
}