package root

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/cli/cli/internal/config"
	"github.com/cli/cli/pkg/cmdutil"
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command

func InitCompletions(f *cmdutil.Factory, c *cobra.Command) {
	rootCmd = c.Root()

	// this rather belongs in the init function of the corresponding command,
	// but since command registration is inverted in this project simply add them separately
	initAlias(f, c)
	initApi(f, c)
	initAuth(f, c)
	initCompletion(f, c)
	initConfig(f, c)
	initGist(f, c)
	initHelp(f, c)
	initIssue(f, c)
	initPr(f, c)
	initRepo(f, c)
}

func initAlias(f *cmdutil.Factory, c *cobra.Command) {
	if cmd, _, err := rootCmd.Find([]string{"alias", "delete"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionAliases(f),
		)
	}
}

func initApi(f *cmdutil.Factory, c *cobra.Command) {}

func initAuth(f *cmdutil.Factory, c *cobra.Command) {
	if cmd, _, err := rootCmd.Find([]string{"auth", "login"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"hostname": carapace.ActionHosts(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"auth", "logout"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"hostname": ActionConfigHosts(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"auth", "refresh"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"hostname": ActionConfigHosts(),
			"scopes":   ActionAuthScopes(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"auth", "status"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"hostname": ActionConfigHosts(),
		})
	}

}

func initCompletion(f *cmdutil.Factory, c *cobra.Command) {
	if cmd, _, err := rootCmd.Find([]string{"completion"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"shell": carapace.ActionValues("bash", "elvish", "fish", "powershell", "zsh"),
		})
	}
}

func initConfig(f *cmdutil.Factory, c *cobra.Command) {
	if cmd, _, err := rootCmd.Find([]string{"config", "get"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"host": carapace.ActionHosts(),
		})

		carapace.Gen(cmd).PositionalCompletion(
			carapace.ActionValuesDescribed(
				"git_protocol", "What protocol to use when performing git operations.",
				"editor", "What editor gh should run when creating issues, pull requests, etc.",
				"aliases", "Aliases allow you to create nicknames for gh commands",
			),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"config", "set"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"host": carapace.ActionHosts(),
		})

		carapace.Gen(cmd).PositionalCompletion(
			carapace.ActionValuesDescribed(
				"git_protocol", "What protocol to use when performing git operations.",
				"editor", "What editor gh should run when creating issues, pull requests, etc.",
				"aliases", "Aliases allow you to create nicknames for gh commands",
			),
			carapace.ActionCallback(func(args []string) carapace.Action {
				switch args[0] {
				case "git_protocol":
					return carapace.ActionValues("ssh", "https")
				case "editor":
					return carapace.ActionValues("vim", "emacs") // TODO path binaries
				default:
					return carapace.ActionValues()
				}
			}),
		)
	}
}

func initGist(f *cmdutil.Factory, c *cobra.Command) {
	if cmd, _, err := rootCmd.Find([]string{"gist", "create"}); err == nil {
		carapace.Gen(cmd).PositionalAnyCompletion(carapace.ActionFiles(""))
	}
}

func initHelp(f *cmdutil.Factory, c *cobra.Command) {}

func initIssue(f *cmdutil.Factory, c *cobra.Command) {
	if cmd, _, err := rootCmd.Find([]string{"issue"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"repo": ActionRepositories(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"issue", "close"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionIssues("open"),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"issue", "create"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"assignee":  ActionAssignableUsers(),
			"label":     ActionLabels(), // TODO ActionMultiParts
			"milestone": ActionMilestones(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"issue", "list"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"assignee":  ActionAssignableUsers(),
			"author":    ActionMentionableUsers(),
			"label":     ActionLabels(), // TODO ActionMultiParts
			"mention":   ActionAssignableUsers(),
			"milestone": ActionMilestones(),
			"state":     carapace.ActionValues("open", "closed", "all"),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"issue", "reopen"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionIssues("closed"),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"issue", "view"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionIssues("open"),
		)
	}
}

func initPr(f *cmdutil.Factory, c *cobra.Command) {
	if cmd, _, err := rootCmd.Find([]string{"pr"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"repo": ActionRepositories(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "checkout"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionPullRequests("open"),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "close"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionPullRequests("open"),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "create"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"assignee":  ActionAssignableUsers(),
			"base":      ActionBranches(),
			"label":     ActionLabels(),
			"milestone": ActionMilestones(),
			// TODO "project": ActionProjects(),
			"reviewer": ActionAssignableUsers(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "diff"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionPullRequests("open"),
		)
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"color": carapace.ActionValues("always", "never", "auto"),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "list"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"assignee": ActionAssignableUsers(),
			"base":     ActionBranches(),
			"label":    ActionLabels(),
			"state":    carapace.ActionValues("open", "closed", "merged", "all"),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "merge"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionPullRequests("open"),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "ready"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionPullRequests("open"), // TODO is this for draft pr? needs to be filtered
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "reopen"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionPullRequests("closed"),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "review"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionPullRequests("open"),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"pr", "view"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionPullRequests("open"),
		)
	}
}

func initRepo(f *cmdutil.Factory, c *cobra.Command) {
	if cmd, _, err := rootCmd.Find([]string{"repo", "clone"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionRepositories(),
			carapace.ActionDirectories(),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"repo", "create"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			// TODO team
			"template": ActionRepositories(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"repo", "fork"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionRepositories(),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"repo", "view"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionRepositories(),
		)
	}
}

func ActionConfigHosts() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if config, err := config.ParseDefaultConfig(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			if hosts, err := config.Hosts(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				return carapace.ActionValues(hosts...)
			}
		}
	})
}

func ActionAliases(f *cmdutil.Factory) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if config, err := config.ParseDefaultConfig(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			if aliasCfg, err := config.Aliases(); err != nil {
				return carapace.ActionMessage(err.Error())
			} else {
				values := make([]string, 0)
				for key, value := range aliasCfg.All() {
					values = append(values, key, value)
				}
				return carapace.ActionValuesDescribed(values...)
			}
		}
	})
}

type label struct {
	Name        string
	Description string
}

type labelsQuery struct {
	Data struct {
		Repository struct {
			Labels struct {
				Nodes []label
			}
		}
	}
}

func ActionLabels() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult labelsQuery
		return GraphQlAction(`repository(owner: $owner, name: $repo){ labels(first: 100) { nodes { name, description } } }`, &queryResult, func() carapace.Action {
			labels := queryResult.Data.Repository.Labels.Nodes
			vals := make([]string, len(labels)*2)
			for index, label := range labels {
				vals[index*2] = label.Name
				vals[index*2+1] = label.Description
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

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

func ActionRepositories() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		owner := "$owner"
		if carapace.CallbackValue != "" {
			owner = fmt.Sprintf(`"%v"`, strings.Split(carapace.CallbackValue, "/")[0])
		}

		var queryResult repositoryQuery
		return GraphQlAction(fmt.Sprintf(`repositoryOwner(login: %v){ repositories(first: 100) { nodes { nameWithOwner, description }  }  }`, owner), &queryResult, func() carapace.Action {
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

type mentionableUsersQuery struct {
	Data struct {
		Repository struct {
			MentionableUsers struct {
				Nodes []user
			}
		}
	}
}

func ActionMentionableUsers() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult mentionableUsersQuery
		return GraphQlAction(`repository(owner: $owner, name: $repo){ mentionableUsers(first: 100) { nodes { login, name } } }`, &queryResult, func() carapace.Action {
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

func ActionAssignableUsers() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult assignableUserQuery
		return GraphQlAction(`repository(owner: $owner, name: $repo){ assignableUsers(first: 100) { nodes { login, name } } }`, &queryResult, func() carapace.Action {
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

type milestone struct {
	Title       string
	Description string
}

type milestonQuery struct {
	Data struct {
		Repository struct {
			Milestones struct {
				Nodes []milestone
			}
		}
	}
}

func ActionMilestones() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult milestonQuery
		return GraphQlAction(`repository(owner: $owner, name: $repo){ milestones(first: 100) { nodes { title, description } } }`, &queryResult, func() carapace.Action {
			milestones := queryResult.Data.Repository.Milestones.Nodes
			vals := make([]string, len(milestones)*2)
			for index, repo := range milestones {
				vals[index*2] = repo.Title
				vals[index*2+1] = repo.Description
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

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

func ActionIssues(state string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult issueQuery
		return GraphQlAction(fmt.Sprintf(`repository(owner: $owner, name: $repo){ issues(first: 100, states: %v) { nodes { number, title } } }`, strings.ToUpper(state)), &queryResult, func() carapace.Action {
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

func ActionPullRequests(state string) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult pullRequestsQuery
		return GraphQlAction(fmt.Sprintf(`repository(owner: $owner, name: $repo){ pullRequests(first: 100, states: %v) { nodes { number, title } } }`, strings.ToUpper(state)), &queryResult, func() carapace.Action {
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

type branch struct {
	Name   string
	Target struct{ AbbreviatedOid string } // TODO last commit message?
}

type branchesQuery struct {
	Data struct {
		Repository struct {
			Refs struct {
				Nodes []branch
			}
		}
	}
}

func ActionBranches() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		var queryResult branchesQuery
		return GraphQlAction(`repository(owner: $owner, name: $repo){ refs(first: 100, refPrefix: "refs/heads/") { nodes { name, target { abbreviatedOid } } } }`, &queryResult, func() carapace.Action {
			branches := queryResult.Data.Repository.Refs.Nodes
			vals := make([]string, len(branches)*2)
			for index, branch := range branches {
				vals[index*2] = branch.Name
				vals[index*2+1] = branch.Target.AbbreviatedOid
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func ActionHttpMethods() carapace.Action {
	return carapace.ActionValuesDescribed(
		"POST", "submit an entity to the specified resource",
		"PATCH", "apply partial modifications to a resourc",
		"PUT", "replaces all current representations of the target resource with the request payload",
		"DELETE", "delete the specified resource",
	)
}

func GraphQlAction(query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		params := make([]string, 0)
		if strings.Contains(query, "$owner") {
			params = append(params, "$owner: String!")
		}
		if strings.Contains(query, "$repo") {
			params = append(params, "$repo: String!")
		}
		queryParams := strings.Join(params, ",")
		if queryParams != "" {
			queryParams = "(" + queryParams + ")"
		}

		owner, repo := repoOverride()
		if output, err := exec.Command("gh", "api", "graphql", "-F", "owner="+owner, "-F", "repo="+repo, "-f", fmt.Sprintf("query=query%v {%v}", queryParams, query)).Output(); err != nil {
			return carapace.ActionMessage(string(output))
		} else {
			if json.Unmarshal(output, &v); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return transform()
		}
	})
}

func repoOverride() (string, string) {
	// TODO support env GH_REPO
	owner := ":owner"
	repo := ":repo"
	for _, cmd := range rootCmd.Commands() {
		if flag := cmd.Flag("repo"); flag != nil && flag.Changed {
			parts := strings.Split(flag.Value.String(), "/")
			switch len(parts) {
			case 1:
				owner = parts[0]
			case 2:
				owner = parts[0]
				repo = parts[1]
			case 3:
				owner = parts[1]
				repo = parts[2]
			}
			break
		}
	}
	return owner, repo
}

// TODO escape `:` in zsh
func ActionAuthScopes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"repo", "Grants full access to private and public repositories.",
		"repo:status", "Grants read/write access to public and private repository commit statuses.",
		"repo_deployment", "Grants access to deployment statuses for public and private repositories.",
		"public_repo", "Limits access to public repositories.",
		"repo:invite", "Grants accept/decline abilities for invitations to collaborate on a repository.",
		"security_events", "Grants read and write access to security events in the code scanning API.",
		"admin:repo_hook", "Grants read, write, ping, and delete access to repository hooks in public and private repositories.",
		"write:repo_hook", "Grants read, write, and ping access to hooks in public or private repositories.",
		"read:repo_hook", "Grants read and ping access to hooks in public or private repositories.",
		"admin:org", "Fully manage the organization and its teams, projects, and memberships.",
		"write:org", "Read and write access to organization membership, organization projects, and team membership.",
		"read:org", "Read-only access to organization membership, organization projects, and team membership.",
		"admin:public_key", "Fully manage public keys.",
		"write:public_key", "Create, list, and view details for public keys.",
		"read:public_key", "List and view details for public keys.",
		"admin:org_hook", "Grants read, write, ping, and delete access to organization hooks.",
		"gist", "Grants write access to gists.",
		"notifications", "Grants read access to a user's notifications",
		"user", "Grants read/write access to profile info only.",
		"read:user", "Grants access to read a user's profile data.",
		"user:email", "Grants read access to a user's email addresses.",
		"user:follow", "Grants access to follow or unfollow other users.",
		"delete_repo", "Grants access to delete adminable repositories.",
		"write:discussion", "Allows read and write access for team discussions.",
		"read:discussion", "Allows read access for team discussions.",
		"write:packages", "Grants access to upload or publish a package in GitHub Packages.",
		"read:packages", "Grants access to download or install packages from GitHub Packages.",
		"delete:packages", "Grants access to delete packages from GitHub Packages.",
		"admin:gpg_key", "Fully manage GPG keys.",
		"write:gpg_key", "Create, list, and view details for GPG keys.",
		"read:gpg_key", "List and view details for GPG keys.",
		"workflow", "Grants the ability to add and update GitHub Actions workflow files.",
	)
}
