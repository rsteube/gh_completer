package action

import (
	"bytes"
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionSecrets(cmd *cobra.Command, org string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		repo := ""
		if cmd.Flag("repo") != nil {
			repo = cmd.Flag("repo").Value.String()
		}
		var stderr bytes.Buffer
		cmd := exec.Command("gh", "secret", "--repo", repo, "list", "--org", org)
		cmd.Stderr = &stderr
		if output, err := cmd.Output(); err != nil {
			return carapace.ActionMessage(stderr.String())
		} else {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0, len(lines)-1)
			for _, line := range lines[:len(lines)-1] {
				vals = append(vals, strings.SplitN(line, "\t", 2)...)
			}
			return carapace.ActionValuesDescribed(vals...)
		}
	})
}
