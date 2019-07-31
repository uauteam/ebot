<%: func GenerateApiRegisterFile(m *project.Metadata, buffer *bytes.Buffer) %>
package api

import (
	"github.com/labstack/echo"
	"bitbucket.org/squarre/core/arre"
)

func Register()(map[string]arre.RouteGroup) {
	m := make(map[string]arre.RouteGroup)

	<%= m.ModuleName %>RouteGroup := arre.RouteGroup{
        Routes:[]arre.Route{
            {Method: echo.POST, Path:"", Handler:Create<%= m.ModuleTitleName %>},
            {Method: echo.GET, Path:"/:id", Handler:Get<%= m.ModuleTitleName %>},
            {Method: echo.GET, Path:"", Handler:Find<%= m.ModuleTitleNamePlural %>},
            {Method: echo.PUT, Path:"/:id", Handler:Update<%= m.ModuleTitleName %>},
        },
    }
    m["/<%= m.ModuleNamePlural %>"] = <%= m.ModuleName %>RouteGroup

	return m
}