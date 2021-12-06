package register

import (
	"log"

	"github.com/danstn/moped/internal/pipeline"
	"github.com/spf13/cobra"
)

const (
	flagFile = "file"
)

type RegisterCmdConfig interface {
	GetPipeline() pipeline.API
}

type registerCLI struct {
	pipeline pipeline.API
}

func NewCommand(config RegisterCmdConfig) *cobra.Command {
	c := &registerCLI{
		pipeline: config.GetPipeline(),
	}

	// command
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register new pipeline.",
		Run:   c.RunCmd,
	}

	// flags
	cmd.Flags().StringP(flagFile, "f", "", "Pipeline file")
	_ = cmd.MarkFlagRequired(flagFile)

	return cmd
}

func (c *registerCLI) RunCmd(cmd *cobra.Command, args []string) {
	file, err := cmd.Flags().GetString(flagFile)
	if err != nil {
		log.Fatalln("missing flag: ", flagFile)
	}

	log.Println("reading pipeline from file:", file)
	definition, err := pipeline.FromYAMLFile(file)
	if err != nil {
		log.Fatalf("failed reading pipeline fie: %v", err)
	}

	_, err = c.pipeline.CreatePipeline(cmd.Context(), definition)
	if err != nil {
		log.Fatalf("failed registering pipeline: %v", err)
	}
}
