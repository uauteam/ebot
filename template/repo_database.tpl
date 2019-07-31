<%: func GenerateRepoDatabaseFile(m *project.Metadata, buffer *bytes.Buffer) %>
package repo

import "<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/entity"

func AutoMigrateEntityRegister()(entities []interface{}) {

	entities = append(entities, &entity.<%= m.ModuleTitleName %>{})

	return
}