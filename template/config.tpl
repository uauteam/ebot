<%: func GenerateConfigFile(m *project.Metadata, buffer *bytes.Buffer) %>
package config

import (
	"bitbucket.org/squarre/core/arre"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/api"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/app"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/repo"
	"fmt"
)

var Default = arre.Config{
	Name:app.Name,
	DBDialect:"sqlite3",
	DBArgs:[]interface{}{fmt.Sprintf("./repo/%s.db", app.Name)},

	AutoMigrateEntityRegister:repo.AutoMigrateEntityRegister,

	ApiRegister: api.Register,
}

func Config(c arre.Config)func()arre.Config {
	return func()arre.Config {
		if c.Name == "" {
			c.Name = Default.Name
		}
		if c.DBDialect == "" {
			c.DBDialect = Default.DBDialect
		}
		if len(c.DBArgs) == 0 {
			c.DBArgs = Default.DBArgs
			if c.DBDialect == "sqlite3" {
                c.DBArgs = []interface{}{fmt.Sprintf("./repo/%s.db", c.Name)}
            }
		}
		if c.AutoMigrateEntityRegister == nil {
			c.AutoMigrateEntityRegister = Default.AutoMigrateEntityRegister
		}
		if c.ApiRegister == nil {
			c.ApiRegister = Default.ApiRegister
		}

		app.Name = c.Name

		return c
	}
}