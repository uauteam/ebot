<%: func GenerateApiRegisterModuleFile(m *project.Metadata, buffer *bytes.Buffer) %>

	<%= m.ModuleName %>RouteGroup := arre.RouteGroup{
	    Routes:[]arre.Route{
            {Method: echo.POST, Path:"", Handler:Create<%= m.ModuleTitleName %>},
            {Method: echo.GET, Path:"/:id", Handler:Get<%= m.ModuleTitleName %>},
            {Method: echo.GET, Path:"", Handler:Find<%= m.ModuleTitleNamePlural %>},
            {Method: echo.PUT, Path:"/:id", Handler:Update<%= m.ModuleTitleName %>},
		},
	}
	m["/<%= m.ModuleNamePlural %>"] = <%= m.ModuleName %>RouteGroup

