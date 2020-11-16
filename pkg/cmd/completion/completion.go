package completion

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/cli/cli/pkg/cmdutil"
	"github.com/cli/cli/pkg/iostreams"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func NewCmdCompletion(io *iostreams.IOStreams) *cobra.Command {
	var shellType string

	cmd := &cobra.Command{
		Use:   "completion",
		Short: "Generate shell completion scripts",
		Long: heredoc.Doc(`
			Generate shell completion scripts for GitHub CLI commands.

			The output of this command will be computer code and is meant to be saved to a
			file or immediately evaluated by an interactive shell.

			For example, for bash you could add this to your '~/.bash_profile':

				eval "$(gh completion -s bash)"

			When installing GitHub CLI through a package manager, however, it's possible that
			no additional shell configuration is necessary to gain completion support. For
			Homebrew, see https://docs.brew.sh/Shell-Completion
		`),
		RunE: func(cmd *cobra.Command, args []string) error {
			if snippet, err := carapace.Gen(cmd).Snippet(shellType, true); err != nil {
				return err
			} else {
				fmt.Fprintln(io.Out, snippet)
			}
			return nil
		},
	}

	cmdutil.DisableAuthCheck(cmd)

	cmd.Flags().StringVarP(&shellType, "shell", "s", "", "Shell type: {bash|elvish|fish|powershell|xonsh|zsh}")

	cmdutil.DeferCompletion(func() {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"shell": carapace.ActionValues("bash", "elvish", "fish", "powershell", "xonsh", "zsh"),
		})
	})

	return cmd
}
