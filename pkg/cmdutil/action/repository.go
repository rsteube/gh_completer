package action

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type repository struct {
	NameWithOwner string
	Description   string
}

type repositoryQuery struct {
	Data struct {
		RepositoryOwner struct {
			Repositories struct {
				Nodes []repository
			}
		}
	}
}

func ActionRepositories(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		owner := "$owner"
		if carapace.CallbackValue != "" {
			owner = fmt.Sprintf(`"%v"`, strings.Split(carapace.CallbackValue, "/")[0])
		}

		var queryResult repositoryQuery
		return GraphQlAction(cmd, fmt.Sprintf(`repositoryOwner(login: %v){ repositories(first: 100) { nodes { nameWithOwner, description }  }  }`, owner), &queryResult, func() carapace.Action {
			repositories := queryResult.Data.RepositoryOwner.Repositories.Nodes
			vals := make([]string, len(repositories)*2)
			for index, repo := range repositories {
				vals[index*2] = repo.NameWithOwner
				vals[index*2+1] = repo.Description
			}
			return carapace.ActionValuesDescribed(vals...)
		})

	})
}
