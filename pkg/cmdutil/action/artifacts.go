package action

import (
	"fmt"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type artifact struct {
	Id          int
	Name        string
	SizeInBytes int `json:"size_in_bytes"`
}

type artifactQuery struct {
	Artifacts []artifact
}

func ActionWorkflowArtifacts(cmd *cobra.Command, runId string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var queryResult artifactQuery
		return ApiV3Action(cmd, fmt.Sprintf(`repos/%v/%v/actions/runs/%v/artifacts`, repo.RepoOwner(), repo.RepoName(), runId), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, artifact := range queryResult.Artifacts {
				vals = append(vals, artifact.Name, byteCountSI(artifact.SizeInBytes))
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
