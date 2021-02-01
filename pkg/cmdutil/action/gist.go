package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

type gist struct {
	Name        string
	Description string
	Files       []struct {
		Name string
		Size int
	}
}

type gistQuery struct {
	Data struct {
		Viewer struct {
			Gists struct {
				Edges []struct {
					Node gist
				}
			}
		}
	}
}

func ActionGists(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult gistQuery
		return GraphQlAction(cmd, `viewer { gists(first:100, privacy:ALL) { edges { node { name, description } } } }`, &queryResult, func() carapace.Action {
			gists := queryResult.Data.Viewer.Gists.Edges
			vals := make([]string, len(gists)*2)
			for index, gist := range gists {
				vals[index*2] = gist.Node.Name
				vals[index*2+1] = gist.Node.Description
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionGistFiles(cmd *cobra.Command, name string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		// TODO query/filter specific gist by name
		var queryResult gistQuery
		return GraphQlAction(cmd, `viewer { gists(first:100, privacy:ALL) { edges { node { name, description, files { name, size } } } } }`, &queryResult, func() carapace.Action {
			gists := queryResult.Data.Viewer.Gists.Edges
			vals := make([]string, 0)
			for _, gist := range gists {
				if gist.Node.Name == name {
					for _, file := range gist.Node.Files {
						vals = append(vals, file.Name, byteCountSI(file.Size))
					}

				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
