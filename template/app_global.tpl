<%: func GenerateAppGlobalFile(m *project.Metadata, buffer *bytes.Buffer) %>
package app

var Name = "<%= m.RepoName %>"