<%: func GenerateApiReqFile(m *project.Metadata, buffer *bytes.Buffer) %>
package req


type <%= m.ModuleTitleName %>Req struct {
    <% for _, field := range m.RequestFields { %>
    <%== "\n\t" %>
    <%==s field %>
    <% } %>
}