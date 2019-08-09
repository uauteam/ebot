<%: func GenerateConfigFile(m *project.Metadata, buffer *bytes.Buffer) %>
package config

import (
	"fmt"
	//_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/api"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/app"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/repo"
	"github.com/uauteam/ecot"
)

var Default = ecot.Config{
	Name:app.Name,
	//DBDialect:"sqlite3",
	//DBArgs:[]interface{}{fmt.Sprintf("./repo/%s.db", app.Name)},
	DBDialect:"mysql",
    DBArgs:[]interface{}{fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root",
        "localhost:3306", "db_" + app.Name)},

	AutoMigrateEntityRegister:repo.AutoMigrateEntityRegister,

	ApiRegister: api.Register,
}

func Config(c ecot.Config)func()ecot.Config {
	return func()ecot.Config {
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