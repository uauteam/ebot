package generator

import (
	"bufio"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/inflection"
	"github.com/uauteam/ago/str"
	"github.com/uauteam/ebot/project"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	repoHost   string
	repoUser   string
	repoName   string
	moduleName string

	workDir string
)

// init project metadata and work dir
func InitProjectMetadata(m *project.Metadata) {
	// go path
	//m.GoPath = os.Getenv("GOPATH")
	//if m.Workspace == "" {
	//	m.Workspace = os.Getenv("GOPATH")
	//}
	if strings.TrimSpace(m.Workspace) == "" {
		fmt.Println("Input project workspace (e.g. /Users/kingphang/Workspaces/uau)")
		fmt.Scanln(&(m.Workspace))
	}

	// repo name
	fmt.Println("Input repo name (e.g. account):")
	fmt.Scanln(&repoName)
	if strings.TrimSpace(repoName) == "" {
		log.Printf("repo name must not be empty")
		return
	}
	m.RepoName = repoName

	// module name
	fmt.Println("Input module name (e.g. user):")
	if strings.TrimSpace(repoName) == "" {
		log.Printf("module name must not be empty")
		return
	}
	fmt.Scanln(&moduleName)
	m.ModuleName = moduleName
	m.ModuleNamePlural = inflection.Plural(moduleName)
	m.ModuleTitleName = strings.Title(moduleName)
	m.ModuleTitleNamePlural = inflection.Plural(m.ModuleTitleName)

	// repo host
	fmt.Println("Input src repo address (defalt: github.com):")
	fmt.Scanln(&repoHost)
	if repoHost == "" {
		repoHost = "github.com"
	}
	m.RepoHost = repoHost

	// repo user
	fmt.Println("Input src repo user (defalt: uauteam):")
	fmt.Scanln(&repoUser)
	if repoUser == "" {
		repoUser = "uauteam"
	}
	m.RepoUser = repoUser

	// work dir
	workDir = m.Workspace + "/src/" + repoHost + "/" + repoUser + "/" + repoName + "/"
	if m.Workspace != os.Getenv("GOPATH") {
		workDir =  m.Workspace + "/" + repoName + "/"
	}
	fmt.Printf("work dir: %s\n", workDir)

	if _, err := os.Stat(workDir); os.IsNotExist(err) {
		os.MkdirAll(workDir, 0755)
	}

	// entity struct fields
	if !m.Designed {
		return
	}

	fmt.Println("Input fields of the module entity struct")
	fmt.Println("One line for one filed (e.g. Username string `gorm:\"type:varchar(100);unique_index\"`)")
	fmt.Println("Press 「Ctrl + D」 for completing designing entity fields")
	fmt.Println("First field starts:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fieldInfo := scanner.Text()
		fieldInfo = strings.TrimSpace(fieldInfo)
		if fieldInfo == "" {
			continue
		}

		fieldInfoArr := strings.Split(fieldInfo, " ")
		if len(fieldInfoArr) < 2 {
			log.Printf("field type should be specified, input again")
			continue
		}

		m.EntityFields = append(m.EntityFields, fieldInfo)
	}

	fmt.Println("\nThe entity fields list: ")
	for i, f := range m.EntityFields {
		fmt.Printf("%d. %s\n", i, f)
	}

	if len(m.EntityFields) == 0 {
		return
	}

	// request struct fields
	fmt.Println()
	fmt.Println("Select the request fields from entity")
	fmt.Println("Input index of the fields to select the request field(e.g. 1 for the first field of entity)")
	fmt.Printf("Press 「Ctrl + D」 for completing designing request fields")
	fmt.Println("\nThe entity fields list: ")
	for i, f := range m.EntityFields {
		fmt.Printf("%d. %s\n", i+1, f)
	}
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		index := scanner.Text()
		if strings.TrimSpace(index) == "" {
			continue
		}

		i, err := strconv.Atoi(index)
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		if i > len(m.EntityFields) {
			log.Printf("Out of the index of the total entity fields")
			continue
		}

		fieldInfo := strings.TrimSpace(m.EntityFields[i-1])
		fieldInfoArr := strings.Split(fieldInfo, " ")
		fieldName := fieldInfoArr[0]
		fieldTag := fmt.Sprintf("`" + fmt.Sprintf(`json:"%s"`, str.LowerFirstLetter(fieldName)) + "`")
		reqField := fmt.Sprintf("%s %s %s", fieldName, fieldInfoArr[1], fieldTag)
		fmt.Println(reqField)
		m.RequestFields = append(m.RequestFields, reqField)
	}

	// response struct fields
	fmt.Println()
	fmt.Println("Select the response fields from entity")
	fmt.Println("Input index of the fields to select the request field(e.g. 1 for the first field of entity)")
	fmt.Printf("Press 「Ctrl + D」 for completing designing response fields")
	fmt.Println("\nThe entity fields list: ")
	for i, f := range m.EntityFields {
		fmt.Printf("%d. %s\n", i+1, f)
	}
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		index := scanner.Text()
		if strings.TrimSpace(index) == "" {
			continue
		}

		i, err := strconv.Atoi(index)
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		if i > len(m.EntityFields) {
			log.Printf("Out of the index of the total entity fields")
			continue
		}

		fieldInfo := strings.TrimSpace(m.EntityFields[i-1])
		fieldInfoArr := strings.Split(fieldInfo, " ")
		fieldName := fieldInfoArr[0]
		fieldTag := fmt.Sprintf("`" + fmt.Sprintf(`json:"%s"`, str.LowerFirstLetter(fieldName)) + "`")
		respField := fmt.Sprintf("%s %s %s", fieldName, fieldInfoArr[1], fieldTag)
		fmt.Println(respField)
		m.ResponseFields = append(m.ResponseFields, respField)
	}

	// query struct fields
	fmt.Println()
	fmt.Println("Select the query fields from entity")
	fmt.Println("Input index of the fields to select the query field(e.g. 1 for the first field of entity)")
	fmt.Printf("Press 「Ctrl + D」 for completing designing query fields")
	fmt.Println("\nThe entity fields list: ")
	for i, f := range m.EntityFields {
		fmt.Printf("%d. %s\n", i+1, f)
	}
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		index := scanner.Text()
		if strings.TrimSpace(index) == "" {
			continue
		}

		i, err := strconv.Atoi(index)
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		if i > len(m.EntityFields) {
			log.Printf("Out of the index of the total entity fields")
			continue
		}

		fieldInfo := strings.TrimSpace(m.EntityFields[i-1])
		fieldInfoArr := strings.Split(fieldInfo, " ")
		fieldName := fieldInfoArr[0]
		fieldTag := fmt.Sprintf("`" + fmt.Sprintf(`query:"%s"`, str.LowerFirstLetter(fieldName)) + "`")
		queryField := fmt.Sprintf("%s %s %s", fieldName, fieldInfoArr[1], fieldTag)
		fmt.Println(queryField)
		m.QueryFields = append(m.QueryFields, queryField)
	}

	// protected fields struct
	fmt.Println()
	fmt.Println("Select the protected fields from entity")
	fmt.Println("Input index of the fields to select the protected field(e.g. 1 for the first field of entity)")
	fmt.Printf("Press 「Ctrl + D」 for completing designing protected fields")
	fmt.Println("\nThe entity fields list: ")
	for i, f := range m.EntityFields {
		fmt.Printf("%d. %s\n", i+1, f)
	}
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		index := scanner.Text()
		if strings.TrimSpace(index) == "" {
			continue
		}

		i, err := strconv.Atoi(index)
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		if i > len(m.EntityFields) {
			log.Printf("Out of the index of the total entity fields")
			continue
		}

		fieldInfo := strings.TrimSpace(m.EntityFields[i-1])
		fieldInfoArr := strings.Split(fieldInfo, " ")
		fieldName := fieldInfoArr[0]
		protectedField := gorm.ToColumnName(fieldName)
		fmt.Println(protectedField)
		m.ProtectedFields = append(m.ProtectedFields, protectedField)
	}
	m.ProtectedFieldsStr = "nil"
	if m.ProtectedFields != nil && len(m.ProtectedFields) > 0 {
		m.ProtectedFieldsStr = "[]string{"
		for _, field := range m.ProtectedFields {
			m.ProtectedFieldsStr = m.ProtectedFieldsStr + fmt.Sprintf(`"%s", `, field)
		}
		m.ProtectedFieldsStr = strings.TrimRight(m.ProtectedFieldsStr, ", ")
		m.ProtectedFieldsStr = m.ProtectedFieldsStr + "}"
	}

	fmt.Println("\n\n--------------------")
	fmt.Printf("repo name: %s\n\n", m.RepoName)
	fmt.Printf("module name: %s\n\n", m.ModuleName)
	fmt.Printf("repo host: %s\n\n", m.RepoHost)
	fmt.Printf("repo user: %s\n\n", m.RepoUser)
	fmt.Printf("work dir: %s\n\n", workDir)

	fmt.Printf("entity fields:\n")
	for _, field := range m.EntityFields {
		fmt.Println(field)
	}

	fmt.Printf("\n\nrequest fields:\n")
	for _, field := range m.RequestFields {
		fmt.Println(field)
	}

	fmt.Printf("\n\nresponse fields:\n")
	for _, field := range m.ResponseFields {
		fmt.Println(field)
	}

	fmt.Printf("\n\nquery fields:\n")
	for _, field := range m.QueryFields {
		fmt.Println(field)
	}

	fmt.Printf("\n\nprotected fields:\n")
	for _, field := range m.ProtectedFields {
		fmt.Println(field)
	}

	// continue or cancel
	fmt.Println("Continue or not? (y/n)")
	scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.ToUpper(strings.TrimSpace(scanner.Text())) != "Y" {
			fmt.Println("canceled")
			return
		}
		break
	}

}
