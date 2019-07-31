package project


type Metadata struct {
	GoPath string

	RepoHost string
	RepoUser string
	RepoName string
	ModuleName string // module name e.g. user
	ModuleNamePlural string // pluralizes of the module name e.g. users
	ModuleTitleName string // titled the module name e.g. User
	ModuleTitleNamePlural string // pluralizes pluralizes of the module name e.g. Users

	// whether the module is designed
	// if designed, the entity/request/response/query/ fields detail
	// and protected fields of entity will be asked
	Designed bool
	EntityFields []string // entity fields
	RequestFields []string // struct fields of HTTP request body for creating or updating
	ResponseFields []string // struct fields of HTTP response body of the module
	QueryFields []string // struct fields of HTTP query param of the module
	ProtectedFields []string // fields that will be ignored when creating or updating the struct
	ProtectedFieldsStr string // fields that will be ignored when creating or updating the struct
}