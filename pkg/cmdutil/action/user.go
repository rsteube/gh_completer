package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type mentionableUsersQuery struct {
	Data struct {
		Repository struct {
			MentionableUsers struct {
				Nodes []user
			}
		}
	}
}

func ActionMentionableUsers(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult mentionableUsersQuery
		return GraphQlAction(cmd, `repository(owner: $owner, name: $repo){ mentionableUsers(first: 100) { nodes { login, name } } }`, &queryResult, func() carapace.Action {
			users := queryResult.Data.Repository.MentionableUsers.Nodes
			vals := make([]string, len(users)*2)
			for index, user := range users {
				vals[index*2] = user.Login
				vals[index*2+1] = user.Name
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

type user struct {
	Login string
	Name  string
}

type assignableUserQuery struct {
	Data struct {
		Repository struct {
			AssignableUsers struct {
				Nodes []user
			}
		}
	}
}

func ActionAssignableUsers(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult assignableUserQuery
		return GraphQlAction(cmd, `repository(owner: $owner, name: $repo){ assignableUsers(first: 100) { nodes { login, name } } }`, &queryResult, func() carapace.Action {
			users := queryResult.Data.Repository.AssignableUsers.Nodes
			vals := make([]string, len(users)*2)
			for index, user := range users {
				vals[index*2] = user.Login
				vals[index*2+1] = user.Name
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
