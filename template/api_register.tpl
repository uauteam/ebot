<%: func GenerateApiRegisterFile(m *project.Metadata, buffer *bytes.Buffer) %>
package api

import (
	"github.com/labstack/echo"
	"github.com/uauteam/ecot"
)

func Register()map[string]ecot.RouteGroup {
	m := make(map[string]ecot.RouteGroup)

	<%= m.ModuleName %>RouteGroup := ecot.RouteGroup{
        Routes:[]ecot.Route{
            {Method: echo.POST, Path:"", Handler:Create<%= m.ModuleTitleName %>, Version:"v1"},
            {Method: echo.GET, Path:"/:id", Handler:Get<%= m.ModuleTitleName %>, Version:"v1"},
            {Method: echo.GET, Path:"", Handler:Find<%= m.ModuleTitleNamePlural %>, Version:"v1"},
            {Method: echo.PUT, Path:"/:id", Handler:Update<%= m.ModuleTitleName %>, Version:"v1"},
        },
    }
    m["/<%= m.ModuleNamePlural %>"] = <%= m.ModuleName %>RouteGroup

	return m
}