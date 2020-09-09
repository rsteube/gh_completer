package root

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

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
			"hostname": carapace.ActionHosts(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"auth", "refresh"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"hostname": carapace.ActionHosts(),
			"scopes":   ActionAuthScopes(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"auth", "logout"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"hostname": carapace.ActionHosts(),
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
			"repo": ActionRepos(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"issue", "close"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionIssues("open"),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"issue", "create"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"assignee":  ActionContributors(),
			"label":     ActionLabels(), // TODO ActionMultiParts
			"milestone": ActionMilestones(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"issue", "list"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"assignee":  ActionContributors(),
			"author":    ActionContributors(),
			"label":     ActionLabels(), // TODO ActionMultiParts
			"mention":   ActionContributors(),
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
			"repo": ActionRepos(),
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
			"assignee":  ActionContributors(),
			"base":      ActionBranches(),
			"label":     ActionLabels(),
			"milestone": ActionMilestones(),
			// TODO "project": ActionProjects(),
			"reviewer": ActionContributors(),
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
			"assignee": ActionContributors(),
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
			ActionRepos(),
			carapace.ActionDirectories(),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"repo", "create"}); err == nil {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			// TODO team
			"template": ActionRepos(),
		})
	}

	if cmd, _, err := rootCmd.Find([]string{"repo", "fork"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionRepos(),
		)
	}

	if cmd, _, err := rootCmd.Find([]string{"repo", "view"}); err == nil {
		carapace.Gen(cmd).PositionalCompletion(
			ActionRepos(),
		)
	}
}

func ActionAliases(f *cmdutil.Factory) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		cfg, err := f.Config()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		aliasCfg, err := cfg.Aliases()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		values := make([]string, 0)
		for key, value := range aliasCfg.All() {
			values = append(values, key, value)
		}
		return carapace.ActionValuesDescribed(values...)
	})

}

type contributor struct {
	Login         string
	Contributions int
}

func ActionContributors() carapace.Action {
	var contributors []contributor
	return ApiAction("/repos/:owner/:repo/contributors?per_page=100", &contributors, func() carapace.Action {
		vals := make([]string, len(contributors)*2)
		for index, contributor := range contributors {
			vals[index*2] = contributor.Login
			vals[index*2+1] = fmt.Sprintf("%v contributions", contributor.Contributions)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

type label struct {
	Name        string
	Description string
}

func ActionLabels() carapace.Action {
	var labels []label
	return ApiAction("/repos/:owner/:repo/labels?per_page=100", &labels, func() carapace.Action {
		vals := make([]string, len(labels)*2)
		for index, label := range labels {
			vals[index*2] = label.Name
			vals[index*2+1] = label.Description
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

type milestone struct {
	Title       string
	Description string
}

func ActionMilestones() carapace.Action {
	var milestones []milestone
	return ApiAction("/repos/:owner/:repo/milestones?per_page=100", &milestones, func() carapace.Action {
		vals := make([]string, len(milestones)*2)
		for index, label := range milestones {
			vals[index*2] = label.Title
			vals[index*2+1] = label.Description
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

type repo struct {
	FullName    string `json:"full_name"`
	Description string
}

func ActionRepos() carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		owner := ":owner"
		if carapace.CallbackValue != "" {
			owner = strings.Split(carapace.CallbackValue, "/")[0]
		}

		var repos []repo
		return ApiAction(fmt.Sprintf("/users/%v/repos?per_page=100", owner), &repos, func() carapace.Action {
			vals := make([]string, len(repos)*2)
			for index, repo := range repos {
				vals[index*2] = repo.FullName
				vals[index*2+1] = repo.Description
			}
			return carapace.ActionValuesDescribed(vals...)
		})

	})
}

type issue struct {
	Number      int
	Title       string
	PullRequest struct{ Url string } `json:"pull_request"`
}

func ActionIssues(state string) carapace.Action {
	var issues []issue
	return ApiAction(fmt.Sprintf("/repos/:owner/:repo/issues?state=%v&per_page=100", state), &issues, func() carapace.Action {
		vals := make([]string, 0)
		for _, issue := range issues {
			if issue.PullRequest.Url == "" { // no pullrequest
				vals = append(vals, strconv.Itoa(issue.Number), issue.Title)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

type pullrequest struct {
	Number int
	Title  string
}

func ActionPullRequests(state string) carapace.Action {
	var pullrequests []pullrequest
	return ApiAction(fmt.Sprintf("/repos/:owner/:repo/pulls?state=%v&per_page=100", state), &pullrequests, func() carapace.Action {
		vals := make([]string, len(pullrequests)*2)
		for index, pull := range pullrequests {
			vals[index*2] = strconv.Itoa(pull.Number)
			vals[index*2+1] = pull.Title
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

type branch struct {
	Name   string
	Commit struct{ Sha string }
}

func ActionBranches() carapace.Action {
	var branches []branch
	return ApiAction("/repos/:owner/:repo/branches?protected=true&per_page=100", &branches, func() carapace.Action {
		vals := make([]string, len(branches)*2)
		for index, branch := range branches {
			vals[index*2] = branch.Name
			vals[index*2+1] = branch.Commit.Sha
		}
		return carapace.ActionValuesDescribed(vals...)
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

func ApiAction(endpoint string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(args []string) carapace.Action {
		if output, err := exec.Command("gh", "api", repoOverride(endpoint)).Output(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			if json.Unmarshal(output, &v); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return transform()
		}
	})
}

func repoOverride(endpoint string) string {
	result := endpoint
	// TODO support env GH_REPO
	for _, cmd := range rootCmd.Commands() {
		if flag := cmd.Flag("repo"); flag != nil && flag.Changed {
			parts := strings.Split(flag.Value.String(), "/")
			switch len(parts) {
			case 1:
				result = strings.Replace(result, ":owner", parts[0], -1)
			case 2:
				result = strings.Replace(result, ":owner", parts[0], -1)
				result = strings.Replace(result, ":repo", parts[1], -1)
			case 3:
				result = strings.Replace(result, ":owner", parts[1], -1)
				result = strings.Replace(result, ":repo", parts[2], -1)
			}
			break
		}
	}
	return result
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
