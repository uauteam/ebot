<%: func GenerateServiceFile(m *project.Metadata, buffer *bytes.Buffer) %>
package svc

import (
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/entity"
	"github.com/uauteam/ecot"
	"<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/api/req"
    "github.com/jinzhu/copier"
    "<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/api/resp"
    "<%= m.RepoHost %>/<%= m.RepoUser %>/<%= m.RepoName %>/api/query"
)

func Create<%= m.ModuleTitleName %>(<%= m.ModuleName %>Req *req.<%= m.ModuleTitleName %>Req)(<%= m.ModuleName %>Resp resp.<%= m.ModuleTitleName %>Resp, err error) {
	<%= m.ModuleName %> := new(entity.<%= m.ModuleTitleName %>)
    if err = copier.Copy(<%= m.ModuleName %>, <%= m.ModuleName %>Req); err != nil {
        return
    }

    err = repo.Create(<%= m.ModuleName %>)
    if err != nil {
        return
    }

    if err = copier.Copy(&<%= m.ModuleName %>Resp, <%= m.ModuleName %>); err != nil {
        return
    }

    return
}

func Get<%= m.ModuleTitleName %>(id uint)(<%= m.ModuleName %>Resp resp.<%= m.ModuleTitleName %>Resp, err error) {
	<%= m.ModuleName %> := new(entity.<%= m.ModuleTitleName %>)
    err = repo.Get(id, <%= m.ModuleName %>)
    if err != nil {
        return
    }

    if err = copier.Copy(&<%= m.ModuleName %>Resp, <%= m.ModuleName %>); err != nil {
        return
    }

	return
}

func Find<%= m.ModuleTitleNamePlural %>(<%= m.ModuleName %>Query *query.<%= m.ModuleTitleName %>Query)(<%= m.ModuleName %>RespList []resp.<%= m.ModuleTitleName %>Resp, err error) {
	<%= m.ModuleName %> := new(entity.<%= m.ModuleTitleName %>)
    if err = copier.Copy(<%= m.ModuleName %>, <%= m.ModuleName %>Query); err != nil {
        return
    }

    var <%= m.ModuleNamePlural %> []entity.<%= m.ModuleTitleName %>
    err = repo.Find(<%= m.ModuleName %>, &<%= m.ModuleNamePlural %>)
    if err != nil {
        return
    }

    if err = copier.Copy(&<%= m.ModuleName %>RespList, &<%= m.ModuleNamePlural %>); err != nil {
        return
    }

	return
}

func Update<%= m.ModuleTitleName %>(id uint, <%= m.ModuleName %>Req *req.<%= m.ModuleTitleName %>Req)(<%= m.ModuleName %>Resp resp.<%= m.ModuleTitleName %>Resp, err error) {
	<%= m.ModuleName %> := new(entity.<%= m.ModuleTitleName %>)
    if err = copier.Copy(<%= m.ModuleName %>, <%= m.ModuleName %>Req); err != nil {
        return
    }

    err = repo.Update(id, <%= m.ModuleName %>)

    if err != nil {
        return
    }

    return Get<%= m.ModuleTitleName %>(id)
}