package cmd

import (
	"github.com/spf13/cobra"

	"github.com/ZupIT/ritchie-cli/pkg/metric"
	"github.com/ZupIT/ritchie-cli/pkg/prompt"
	"github.com/ZupIT/ritchie-cli/pkg/stream"
)

type metricsCmd struct {
	file  stream.FileWriteReadExister
	input prompt.InputList
}

func NewMetricsCmd(file stream.FileWriteReadExister, inList prompt.InputList) *cobra.Command {
	m := &metricsCmd{
		file:  file,
		input: inList,
	}

	cmd := &cobra.Command{
		Use:   "metrics",
		Short: "Turn metrics on and off",
		Long:  "Stop or start to send anonymous metrics to ritchie team.",
		RunE:  m.run(),
	}

	return cmd

}

func (m metricsCmd) run() CommandRunnerFunc {
	return func(cmd *cobra.Command, args []string) error {
		options := []string{"yes", "no"}
		choose, err := m.input.List(metric.AcceptQuestion, options)
		if err != nil {
			return err
		}

		if err := m.file.Write(metric.FilePath, []byte(choose)); err != nil {
			return err
		}

		message := "You are now sending anonymous metrics. Thank you!"
		if choose == "no" {
			message = "You are no longer sending anonymous metrics."
		}
		prompt.Info(message)
		return nil
	}
}
