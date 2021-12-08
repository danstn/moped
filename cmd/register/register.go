package register

import (
	"log"

	"github.com/danstn/moped/services/pipeline"
	"github.com/spf13/cobra"
)

const (
	flagFile = "file"
)

type RegisterCmdConfig interface {
	GetPipelineService() pipeline.API
}

type registerCLI struct {
	pipelineService pipeline.API
}

func NewCommand(config RegisterCmdConfig) *cobra.Command {
	c := &registerCLI{
		pipelineService: config.GetPipelineService(),
	}
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register new pipeline.",
		Run:   c.RunCmd,
	}
	cmd.Flags().StringP(flagFile, "f", "", "Pipeline file")
	_ = cmd.MarkFlagRequired(flagFile)
	return cmd
}

func (c *registerCLI) RunCmd(cmd *cobra.Command, args []string) {
	// get flags
	file, err := cmd.Flags().GetString(flagFile)
	if err != nil {
		log.Fatalln("missing flag: ", flagFile)
	}
	// create pipeline
	pipelineID, err := c.pipelineService.CreatePipelineFromFile(cmd.Context(), file)
	if err != nil {
		log.Fatalf("failed registering pipeline: %v", err)
	}
	// report
	log.Printf("created a new pipeline: %v", pipelineID)
}
