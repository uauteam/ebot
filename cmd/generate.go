package cmd

import (
	"github.com/spf13/cobra"
	"github.com/uauteam/ebot/generator"
	"github.com/uauteam/ebot/project"
	"log"
)

var projectMetadata project.Metadata

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate basic code",
	Run: func(cmd *cobra.Command, args []string) {
		designed, err := cmd.Flags().GetBool("designed")
		if err != nil {
			log.Printf(err.Error())
		}
		projectMetadata.Designed = designed

		// init project metadata and work dir
		generator.InitProjectMetadata(&projectMetadata)

		// init git ignore file
		generator.GenerateGitIgnoreFile(&projectMetadata)

		// generate entity file
		generator.GenerateEntityFile(&projectMetadata)

		// generate svc file
		generator.GenerateSvcFile(&projectMetadata)

		// generate api query file
		generator.GenerateApiQueryFile(&projectMetadata)

		// generate api req file
		generator.GenerateApiReqFile(&projectMetadata)

		// generate api resp file
		generator.GenerateApiRespFile(&projectMetadata)

		// generate api module file
		generator.GenerateApiModuleFile(&projectMetadata)

		// generate api register file
		generator.GenerateApiRegisterFile(&projectMetadata)

		// generate repo database file
		generator.GenerateRepoDatabaseFile(&projectMetadata)

		// generate app global file
		generator.GenerateAppGlobalFile(&projectMetadata)

		// generate config file
		generator.GenerateConfigFile(&projectMetadata)

		// generate main file
		generator.GenerateMainFile(&projectMetadata)
	},
}