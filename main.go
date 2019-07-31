package main

import (
	"github.com/uauteam/ebot/cmd"
	"github.com/uauteam/ebot/project"
)

var (
	projectMetadata project.Metadata
)

func main()  {
	cmd.Execute()

	//// init project metadata and work dir
	//generator.InitProjectMetadata(&projectMetadata)
	//
	//// generate entity file
	//generator.GenerateEntityFile(&projectMetadata)
	//
	//// generate constant app file
	//generator.GenerateConstantAppFile(&projectMetadata)
	//
	//// generate svc file
	//generator.GenerateSvcFile(&projectMetadata)
	//
	//// generate api query file
	//generator.GenerateApiQueryFile(&projectMetadata)
	//
	//// generate api req file
	//generator.GenerateApiReqFile(&projectMetadata)
	//
	//// generate api resp file
	//generator.GenerateApiRespFile(&projectMetadata)
	//
	//// generate api module file
	//generator.GenerateApiModuleFile(&projectMetadata)
	//
	//// generate api register file
	//generator.GenerateApiRegisterFile(&projectMetadata)
	//
	//// generate repo database file
	//generator.GenerateRepoDatabaseFile(&projectMetadata)
	//
	//// generate app config file
	//generator.GenerateAppConfigFile(&projectMetadata)
	//
	//// generate main file
	//generator.GenerateMainFile(&projectMetadata)
}