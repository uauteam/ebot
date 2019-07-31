<%: func GenerateRepoDatabaseModuleFile(m *project.Metadata, buffer *bytes.Buffer) %>
	entities = append(entities, &entity.<%= m.ModuleTitleName %>{})

