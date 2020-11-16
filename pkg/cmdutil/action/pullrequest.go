package action

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type pullrequest struct {
	Number int
	Title  string
}

type pullRequestsQuery struct {
	Data struct {
		Repository struct {
			PullRequests struct {
				Nodes []pullrequest
			}
		}
	}
}

type PullRequestOpts struct {
	Closed bool
	Merged bool
	Open   bool
}

func (p *PullRequestOpts) states() string {
	states := make([]string, 0)
	for index, include := range []bool{p.Closed, p.Merged, p.Open} {
		if include {
			states = append(states, []string{"CLOSED", "MERGED", "OPEN"}[index])
		}
	}
	return fmt.Sprintf("[%v]", strings.Join(states, ","))
}

func ActionPullRequests(cmd *cobra.Command, opts PullRequestOpts) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult pullRequestsQuery
		return GraphQlAction(cmd, fmt.Sprintf(`repository(owner: $owner, name: $repo){ pullRequests(first: 100, states: %v) { nodes { number, title } } }`, opts.states()), &queryResult, func() carapace.Action {
			pullrequests := queryResult.Data.Repository.PullRequests.Nodes
			vals := make([]string, len(pullrequests)*2)
			for index, pullrequest := range pullrequests {
				vals[index*2] = strconv.Itoa(pullrequest.Number)
				vals[index*2+1] = pullrequest.Title // TODO shorten title
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
