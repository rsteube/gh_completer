package cmdutil

import (
	"github.com/cli/cli/pkg/cmdutil/action"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var completions = make([]func(), 0)

func DeferCompletion(completion func()) {
	completions = append(completions, completion)
}

func InitCompletions() {
	var current []func()
	current, completions = completions, make([]func(), 0)
	for _, completion := range current {
		completion()
	}
}

func AddAliasCompletion(cmd *cobra.Command) {
	carapace.Gen(cmd) // ensure _carapace subcommand exists
	if c, _, err := cmd.Find([]string{"_carapace"}); err == nil {
		// TODO move this to separate func
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
