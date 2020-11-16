package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type issue struct {
	Number int
	Title  string
}

type issueQuery struct {
	Data struct {
		Repository struct {
			Issues struct {
				Nodes []issue
			}
		}
	}
}

func ActionIssues(cmd *cobra.Command, state string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult issueQuery
		return GraphQlAction(cmd, fmt.Sprintf(`repository(owner: $owner, name: $repo){ issues(first: 100, states: %v) { nodes { number, title } } }`, strings.ToUpper(state)), &queryResult, func() carapace.Action {
			issues := queryResult.Data.Repository.Issues.Nodes
			vals := make([]string, len(issues)*2)
			for index, issue := range issues {
				vals[index*2] = strconv.Itoa(issue.Number)
				vals[index*2+1] = issue.Title // TODO shorten title
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
