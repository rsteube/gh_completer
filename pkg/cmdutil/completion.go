package cmdutil

import (
	"github.com/cli/cli/pkg/cmdutil/action"
	"github.com/spf13/cobra"
)

var completions = make([]func(), 0)

// DeferCompletion defers completion configuration until the command stucture is finalized (carapace requirement)
func DeferCompletion(completion func()) {
	completions = append(completions, completion)
}

// InitCompletions finalizes completion configuration
func InitCompletions(cmd *cobra.Command) {
	for _, completion := range completions {
		completion()
	}
	addAliasCompletion(cmd)
}

func addAliasCompletion(cmd *cobra.Command) {
	if c, _, err := cmd.Find([]string{"_carapace"}); err == nil {
		c.Annotations = map[string]string{
			"skipAuthCheck": "true",
		}
		c.PreRun = func(cmd *cobra.Command, args []string) {
			if aliases, err := action.Aliases(); err == nil {
				for key, value := range aliases {
					cmd.Root().AddCommand(&cobra.Command{Use: key, Short: value, Run: func(cmd *cobra.Command, args []string) {}})
				}
			}
		}
	}
}
