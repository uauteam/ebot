<%: func GenerateApiModuleFile(m *project.Metadata, buffer *bytes.Buffer) %>
package api

import (
	"github.com/labstack/echo"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/api/query"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/api/req"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/svc"
	"net/http"
	"strconv"
)

func Create<%= m.ModuleTitleName %>(c echo.Context) (err error) {
	<%= m.ModuleName %>Req := new(req.<%= m.ModuleTitleName %>Req)
	if err = c.Bind(<%= m.ModuleName %>Req); err != nil {
		return
	}

	<%= m.ModuleName %>Resp, err := svc.Create<%= m.ModuleTitleName %>(<%= m.ModuleName %>Req)
    if err != nil {
        return
    }

	return c.JSON(http.StatusOK, <%= m.ModuleName %>Resp)
}

func Get<%= m.ModuleTitleName %>(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	<%= m.ModuleName %>Resp, err := svc.Get<%= m.ModuleTitleName %>(uint(id))
    if err != nil {
        return
    }

	return c.JSON(http.StatusOK, <%= m.ModuleName %>Resp)
}

func Find<%= m.ModuleTitleNamePlural %>(c echo.Context)(err error) {
	<%= m.ModuleName %>Query := new(query.<%= m.ModuleTitleName %>Query)
	if err = c.Bind(<%= m.ModuleName %>Query); err != nil {
		return
	}

	<%= m.ModuleName %>RespList, err := svc.Find<%= m.ModuleTitleNamePlural %>(<%= m.ModuleName %>Query)
    if err != nil {
        return
    }

	return c.JSON(http.StatusOK, <%= m.ModuleName %>RespList)
}

func Update<%= m.ModuleTitleName %>(c echo.Context) (err error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	<%= m.ModuleName %>Req := new(req.<%= m.ModuleTitleName %>Req)
	if err = c.Bind(<%= m.ModuleName %>Req); err != nil {
		return
	}

    <%= m.ModuleName %>Resp, err := svc.Update<%= m.ModuleTitleName %>(uint(id), <%= m.ModuleName %>Req)
	if err != nil {
		return
	}

	return c.JSON(http.StatusOK, <%= m.ModuleName %>Resp)
}