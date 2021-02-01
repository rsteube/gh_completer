package action

import "github.com/rsteube/carapace"

// https://raw.githubusercontent.com/github/rest-api-description/main/descriptions/api.github.com/api.github.com.json
//    cat api.github.com.json | jq '.paths | keys[]' | sed -e 's_/__' -e 's/$/,/' -e 's#gitignore/templates/{name}#gitignore/templates/{gitignore_name}#'
var v3Paths []string = []string{
	"app",
	"app-manifests/{code}/conversions",
	"app/hook/config",
	"app/installations",
	"app/installations/{installation_id}",
	"app/installations/{installation_id}/access_tokens",
	"app/installations/{installation_id}/suspended",
	"applications/grants",
	"applications/grants/{grant_id}",
	"applications/{client_id}/grant",
	"applications/{client_id}/grants/{access_token}",
	"applications/{client_id}/token",
	"applications/{client_id}/token/scoped",
	"applications/{client_id}/tokens/{access_token}",
	"apps/{app_slug}",
	"authorizations",
	"authorizations/clients/{client_id}",
	"authorizations/clients/{client_id}/{fingerprint}",
	"authorizations/{authorization_id}",
	"codes_of_conduct",
	"codes_of_conduct/{key}",
	"content_references/{content_reference_id}/attachments",
	"emojis",
	"enterprises/{enterprise}/actions/permissions",
	"enterprises/{enterprise}/actions/permissions/organizations",
	"enterprises/{enterprise}/actions/permissions/organizations/{org_id}",
	"enterprises/{enterprise}/actions/permissions/selected-actions",
	"enterprises/{enterprise}/actions/runner-groups",
	"enterprises/{enterprise}/actions/runner-groups/{runner_group_id}",
	"enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations",
	"enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/organizations/{org_id}",
	"enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners",
	"enterprises/{enterprise}/actions/runner-groups/{runner_group_id}/runners/{runner_id}",
	"enterprises/{enterprise}/actions/runners",
	"enterprises/{enterprise}/actions/runners/downloads",
	"enterprises/{enterprise}/actions/runners/registration-token",
	"enterprises/{enterprise}/actions/runners/remove-token",
	"enterprises/{enterprise}/actions/runners/{runner_id}",
	"enterprises/{enterprise}/audit-log",
	"enterprises/{enterprise}/settings/billing/actions",
	"enterprises/{enterprise}/settings/billing/packages",
	"enterprises/{enterprise}/settings/billing/shared-storage",
	"events",
	"feeds",
	"gists",
	"gists/public",
	"gists/starred",
	"gists/{gist_id}",
	"gists/{gist_id}/comments",
	"gists/{gist_id}/comments/{comment_id}",
	"gists/{gist_id}/commits",
	"gists/{gist_id}/forks",
	"gists/{gist_id}/star",
	"gists/{gist_id}/{sha}",
	"gitignore/templates",
	"gitignore/templates/{gitignore_name}",
	"installation/repositories",
	"installation/token",
	"issues",
	"licenses",
	"licenses/{license}",
	"markdown",
	"markdown/raw",
	"marketplace_listing/accounts/{account_id}",
	"marketplace_listing/plans",
	"marketplace_listing/plans/{plan_id}/accounts",
	"marketplace_listing/stubbed/accounts/{account_id}",
	"marketplace_listing/stubbed/plans",
	"marketplace_listing/stubbed/plans/{plan_id}/accounts",
	"meta",
	"networks/{owner}/{repo}/events",
	"notifications",
	"notifications/threads/{thread_id}",
	"notifications/threads/{thread_id}/subscription",
	"octocat",
	"organizations",
	"orgs/{org}",
	"orgs/{org}/actions/permissions",
	"orgs/{org}/actions/permissions/repositories",
	"orgs/{org}/actions/permissions/repositories/{repository_id}",
	"orgs/{org}/actions/permissions/selected-actions",
	"orgs/{org}/actions/runner-groups",
	"orgs/{org}/actions/runner-groups/{runner_group_id}",
	"orgs/{org}/actions/runner-groups/{runner_group_id}/repositories",
	"orgs/{org}/actions/runner-groups/{runner_group_id}/repositories/{repository_id}",
	"orgs/{org}/actions/runner-groups/{runner_group_id}/runners",
	"orgs/{org}/actions/runner-groups/{runner_group_id}/runners/{runner_id}",
	"orgs/{org}/actions/runners",
	"orgs/{org}/actions/runners/downloads",
	"orgs/{org}/actions/runners/registration-token",
	"orgs/{org}/actions/runners/remove-token",
	"orgs/{org}/actions/runners/{runner_id}",
	"orgs/{org}/actions/secrets",
	"orgs/{org}/actions/secrets/public-key",
	"orgs/{org}/actions/secrets/{secret_name}",
	"orgs/{org}/actions/secrets/{secret_name}/repositories",
	"orgs/{org}/actions/secrets/{secret_name}/repositories/{repository_id}",
	"orgs/{org}/audit-log",
	"orgs/{org}/blocks",
	"orgs/{org}/blocks/{username}",
	"orgs/{org}/credential-authorizations",
	"orgs/{org}/credential-authorizations/{credential_id}",
	"orgs/{org}/events",
	"orgs/{org}/failed_invitations",
	"orgs/{org}/hooks",
	"orgs/{org}/hooks/{hook_id}",
	"orgs/{org}/hooks/{hook_id}/config",
	"orgs/{org}/hooks/{hook_id}/pings",
	"orgs/{org}/installation",
	"orgs/{org}/installations",
	"orgs/{org}/interaction-limits",
	"orgs/{org}/invitations",
	"orgs/{org}/invitations/{invitation_id}",
	"orgs/{org}/invitations/{invitation_id}/teams",
	"orgs/{org}/issues",
	"orgs/{org}/members",
	"orgs/{org}/members/{username}",
	"orgs/{org}/memberships/{username}",
	"orgs/{org}/migrations",
	"orgs/{org}/migrations/{migration_id}",
	"orgs/{org}/migrations/{migration_id}/archive",
	"orgs/{org}/migrations/{migration_id}/repos/{repo_name}/lock",
	"orgs/{org}/migrations/{migration_id}/repositories",
	"orgs/{org}/outside_collaborators",
	"orgs/{org}/outside_collaborators/{username}",
	"orgs/{org}/packages/{package_type}/{package_name}",
	"orgs/{org}/packages/{package_type}/{package_name}/restore",
	"orgs/{org}/packages/{package_type}/{package_name}/versions",
	"orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id}",
	"orgs/{org}/packages/{package_type}/{package_name}/versions/{package_version_id}/restore",
	"orgs/{org}/projects",
	"orgs/{org}/public_members",
	"orgs/{org}/public_members/{username}",
	"orgs/{org}/repos",
	"orgs/{org}/settings/billing/actions",
	"orgs/{org}/settings/billing/packages",
	"orgs/{org}/settings/billing/shared-storage",
	"orgs/{org}/team-sync/groups",
	"orgs/{org}/teams",
	"orgs/{org}/teams/{team_slug}",
	"orgs/{org}/teams/{team_slug}/discussions",
	"orgs/{org}/teams/{team_slug}/discussions/{discussion_number}",
	"orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments",
	"orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}",
	"orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions",
	"orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/comments/{comment_number}/reactions/{reaction_id}",
	"orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions",
	"orgs/{org}/teams/{team_slug}/discussions/{discussion_number}/reactions/{reaction_id}",
	"orgs/{org}/teams/{team_slug}/invitations",
	"orgs/{org}/teams/{team_slug}/members",
	"orgs/{org}/teams/{team_slug}/memberships/{username}",
	"orgs/{org}/teams/{team_slug}/projects",
	"orgs/{org}/teams/{team_slug}/projects/{project_id}",
	"orgs/{org}/teams/{team_slug}/repos",
	"orgs/{org}/teams/{team_slug}/repos/{owner}/{repo}",
	"orgs/{org}/teams/{team_slug}/team-sync/group-mappings",
	"orgs/{org}/teams/{team_slug}/teams",
	"projects/columns/cards/{card_id}",
	"projects/columns/cards/{card_id}/moves",
	"projects/columns/{column_id}",
	"projects/columns/{column_id}/cards",
	"projects/columns/{column_id}/moves",
	"projects/{project_id}",
	"projects/{project_id}/collaborators",
	"projects/{project_id}/collaborators/{username}",
	"projects/{project_id}/collaborators/{username}/permission",
	"projects/{project_id}/columns",
	"rate_limit",
	"reactions/{reaction_id}",
	"repos/{owner}/{repo}",
	"repos/{owner}/{repo}/actions/artifacts",
	"repos/{owner}/{repo}/actions/artifacts/{artifact_id}",
	"repos/{owner}/{repo}/actions/artifacts/{artifact_id}/{archive_format}",
	"repos/{owner}/{repo}/actions/jobs/{job_id}",
	"repos/{owner}/{repo}/actions/jobs/{job_id}/logs",
	"repos/{owner}/{repo}/actions/permissions",
	"repos/{owner}/{repo}/actions/permissions/selected-actions",
	"repos/{owner}/{repo}/actions/runners",
	"repos/{owner}/{repo}/actions/runners/downloads",
	"repos/{owner}/{repo}/actions/runners/registration-token",
	"repos/{owner}/{repo}/actions/runners/remove-token",
	"repos/{owner}/{repo}/actions/runners/{runner_id}",
	"repos/{owner}/{repo}/actions/runs",
	"repos/{owner}/{repo}/actions/runs/{run_id}",
	"repos/{owner}/{repo}/actions/runs/{run_id}/approvals",
	"repos/{owner}/{repo}/actions/runs/{run_id}/artifacts",
	"repos/{owner}/{repo}/actions/runs/{run_id}/cancel",
	"repos/{owner}/{repo}/actions/runs/{run_id}/jobs",
	"repos/{owner}/{repo}/actions/runs/{run_id}/logs",
	"repos/{owner}/{repo}/actions/runs/{run_id}/pending_deployments",
	"repos/{owner}/{repo}/actions/runs/{run_id}/rerun",
	"repos/{owner}/{repo}/actions/runs/{run_id}/timing",
	"repos/{owner}/{repo}/actions/secrets",
	"repos/{owner}/{repo}/actions/secrets/public-key",
	"repos/{owner}/{repo}/actions/secrets/{secret_name}",
	"repos/{owner}/{repo}/actions/workflows",
	"repos/{owner}/{repo}/actions/workflows/{workflow_id}",
	"repos/{owner}/{repo}/actions/workflows/{workflow_id}/disable",
	"repos/{owner}/{repo}/actions/workflows/{workflow_id}/dispatches",
	"repos/{owner}/{repo}/actions/workflows/{workflow_id}/enable",
	"repos/{owner}/{repo}/actions/workflows/{workflow_id}/runs",
	"repos/{owner}/{repo}/actions/workflows/{workflow_id}/timing",
	"repos/{owner}/{repo}/assignees",
	"repos/{owner}/{repo}/assignees/{assignee}",
	"repos/{owner}/{repo}/automated-security-fixes",
	"repos/{owner}/{repo}/branches",
	"repos/{owner}/{repo}/branches/{branch}",
	"repos/{owner}/{repo}/branches/{branch}/protection",
	"repos/{owner}/{repo}/branches/{branch}/protection/enforce_admins",
	"repos/{owner}/{repo}/branches/{branch}/protection/required_pull_request_reviews",
	"repos/{owner}/{repo}/branches/{branch}/protection/required_signatures",
	"repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks",
	"repos/{owner}/{repo}/branches/{branch}/protection/required_status_checks/contexts",
	"repos/{owner}/{repo}/branches/{branch}/protection/restrictions",
	"repos/{owner}/{repo}/branches/{branch}/protection/restrictions/apps",
	"repos/{owner}/{repo}/branches/{branch}/protection/restrictions/teams",
	"repos/{owner}/{repo}/branches/{branch}/protection/restrictions/users",
	"repos/{owner}/{repo}/branches/{branch}/rename",
	"repos/{owner}/{repo}/check-runs",
	"repos/{owner}/{repo}/check-runs/{check_run_id}",
	"repos/{owner}/{repo}/check-runs/{check_run_id}/annotations",
	"repos/{owner}/{repo}/check-suites",
	"repos/{owner}/{repo}/check-suites/preferences",
	"repos/{owner}/{repo}/check-suites/{check_suite_id}",
	"repos/{owner}/{repo}/check-suites/{check_suite_id}/check-runs",
	"repos/{owner}/{repo}/check-suites/{check_suite_id}/rerequest",
	"repos/{owner}/{repo}/code-scanning/alerts",
	"repos/{owner}/{repo}/code-scanning/alerts/{alert_number}",
	"repos/{owner}/{repo}/code-scanning/alerts/{alert_number}/instances",
	"repos/{owner}/{repo}/code-scanning/analyses",
	"repos/{owner}/{repo}/code-scanning/analyses/{analysis_id}",
	"repos/{owner}/{repo}/code-scanning/sarifs",
	"repos/{owner}/{repo}/code-scanning/sarifs/{sarif_id}",
	"repos/{owner}/{repo}/collaborators",
	"repos/{owner}/{repo}/collaborators/{username}",
	"repos/{owner}/{repo}/collaborators/{username}/permission",
	"repos/{owner}/{repo}/comments",
	"repos/{owner}/{repo}/comments/{comment_id}",
	"repos/{owner}/{repo}/comments/{comment_id}/reactions",
	"repos/{owner}/{repo}/comments/{comment_id}/reactions/{reaction_id}",
	"repos/{owner}/{repo}/commits",
	"repos/{owner}/{repo}/commits/{commit_sha}/branches-where-head",
	"repos/{owner}/{repo}/commits/{commit_sha}/comments",
	"repos/{owner}/{repo}/commits/{commit_sha}/pulls",
	"repos/{owner}/{repo}/commits/{ref}",
	"repos/{owner}/{repo}/commits/{ref}/check-runs",
	"repos/{owner}/{repo}/commits/{ref}/check-suites",
	"repos/{owner}/{repo}/commits/{ref}/status",
	"repos/{owner}/{repo}/commits/{ref}/statuses",
	"repos/{owner}/{repo}/community/code_of_conduct",
	"repos/{owner}/{repo}/community/profile",
	"repos/{owner}/{repo}/compare/{base}...{head}",
	"repos/{owner}/{repo}/contents/{path}",
	"repos/{owner}/{repo}/contributors",
	"repos/{owner}/{repo}/deployments",
	"repos/{owner}/{repo}/deployments/{deployment_id}",
	"repos/{owner}/{repo}/deployments/{deployment_id}/statuses",
	"repos/{owner}/{repo}/deployments/{deployment_id}/statuses/{status_id}",
	"repos/{owner}/{repo}/dispatches",
	"repos/{owner}/{repo}/environments",
	"repos/{owner}/{repo}/environments/{environment_name}",
	"repos/{owner}/{repo}/events",
	"repos/{owner}/{repo}/forks",
	"repos/{owner}/{repo}/git/blobs",
	"repos/{owner}/{repo}/git/blobs/{file_sha}",
	"repos/{owner}/{repo}/git/commits",
	"repos/{owner}/{repo}/git/commits/{commit_sha}",
	"repos/{owner}/{repo}/git/matching-refs/{ref}",
	"repos/{owner}/{repo}/git/ref/{ref}",
	"repos/{owner}/{repo}/git/refs",
	"repos/{owner}/{repo}/git/refs/{ref}",
	"repos/{owner}/{repo}/git/tags",
	"repos/{owner}/{repo}/git/tags/{tag_sha}",
	"repos/{owner}/{repo}/git/trees",
	"repos/{owner}/{repo}/git/trees/{tree_sha}",
	"repos/{owner}/{repo}/hooks",
	"repos/{owner}/{repo}/hooks/{hook_id}",
	"repos/{owner}/{repo}/hooks/{hook_id}/config",
	"repos/{owner}/{repo}/hooks/{hook_id}/pings",
	"repos/{owner}/{repo}/hooks/{hook_id}/tests",
	"repos/{owner}/{repo}/import",
	"repos/{owner}/{repo}/import/authors",
	"repos/{owner}/{repo}/import/authors/{author_id}",
	"repos/{owner}/{repo}/import/large_files",
	"repos/{owner}/{repo}/import/lfs",
	"repos/{owner}/{repo}/installation",
	"repos/{owner}/{repo}/interaction-limits",
	"repos/{owner}/{repo}/invitations",
	"repos/{owner}/{repo}/invitations/{invitation_id}",
	"repos/{owner}/{repo}/issues",
	"repos/{owner}/{repo}/issues/comments",
	"repos/{owner}/{repo}/issues/comments/{comment_id}",
	"repos/{owner}/{repo}/issues/comments/{comment_id}/reactions",
	"repos/{owner}/{repo}/issues/comments/{comment_id}/reactions/{reaction_id}",
	"repos/{owner}/{repo}/issues/events",
	"repos/{owner}/{repo}/issues/events/{event_id}",
	"repos/{owner}/{repo}/issues/{issue_number}",
	"repos/{owner}/{repo}/issues/{issue_number}/assignees",
	"repos/{owner}/{repo}/issues/{issue_number}/comments",
	"repos/{owner}/{repo}/issues/{issue_number}/events",
	"repos/{owner}/{repo}/issues/{issue_number}/labels",
	"repos/{owner}/{repo}/issues/{issue_number}/labels/{name}",
	"repos/{owner}/{repo}/issues/{issue_number}/lock",
	"repos/{owner}/{repo}/issues/{issue_number}/reactions",
	"repos/{owner}/{repo}/issues/{issue_number}/reactions/{reaction_id}",
	"repos/{owner}/{repo}/issues/{issue_number}/timeline",
	"repos/{owner}/{repo}/keys",
	"repos/{owner}/{repo}/keys/{key_id}",
	"repos/{owner}/{repo}/labels",
	"repos/{owner}/{repo}/labels/{name}",
	"repos/{owner}/{repo}/languages",
	"repos/{owner}/{repo}/license",
	"repos/{owner}/{repo}/merges",
	"repos/{owner}/{repo}/milestones",
	"repos/{owner}/{repo}/milestones/{milestone_number}",
	"repos/{owner}/{repo}/milestones/{milestone_number}/labels",
	"repos/{owner}/{repo}/notifications",
	"repos/{owner}/{repo}/pages",
	"repos/{owner}/{repo}/pages/builds",
	"repos/{owner}/{repo}/pages/builds/latest",
	"repos/{owner}/{repo}/pages/builds/{build_id}",
	"repos/{owner}/{repo}/projects",
	"repos/{owner}/{repo}/pulls",
	"repos/{owner}/{repo}/pulls/comments",
	"repos/{owner}/{repo}/pulls/comments/{comment_id}",
	"repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions",
	"repos/{owner}/{repo}/pulls/comments/{comment_id}/reactions/{reaction_id}",
	"repos/{owner}/{repo}/pulls/{pull_number}",
	"repos/{owner}/{repo}/pulls/{pull_number}/comments",
	"repos/{owner}/{repo}/pulls/{pull_number}/comments/{comment_id}/replies",
	"repos/{owner}/{repo}/pulls/{pull_number}/commits",
	"repos/{owner}/{repo}/pulls/{pull_number}/files",
	"repos/{owner}/{repo}/pulls/{pull_number}/merge",
	"repos/{owner}/{repo}/pulls/{pull_number}/requested_reviewers",
	"repos/{owner}/{repo}/pulls/{pull_number}/reviews",
	"repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}",
	"repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/comments",
	"repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/dismissals",
	"repos/{owner}/{repo}/pulls/{pull_number}/reviews/{review_id}/events",
	"repos/{owner}/{repo}/pulls/{pull_number}/update-branch",
	"repos/{owner}/{repo}/readme",
	"repos/{owner}/{repo}/readme/{dir}",
	"repos/{owner}/{repo}/releases",
	"repos/{owner}/{repo}/releases/assets/{asset_id}",
	"repos/{owner}/{repo}/releases/latest",
	"repos/{owner}/{repo}/releases/tags/{tag}",
	"repos/{owner}/{repo}/releases/{release_id}",
	"repos/{owner}/{repo}/releases/{release_id}/assets",
	"repos/{owner}/{repo}/secret-scanning/alerts",
	"repos/{owner}/{repo}/secret-scanning/alerts/{alert_number}",
	"repos/{owner}/{repo}/stargazers",
	"repos/{owner}/{repo}/stats/code_frequency",
	"repos/{owner}/{repo}/stats/commit_activity",
	"repos/{owner}/{repo}/stats/contributors",
	"repos/{owner}/{repo}/stats/participation",
	"repos/{owner}/{repo}/stats/punch_card",
	"repos/{owner}/{repo}/statuses/{sha}",
	"repos/{owner}/{repo}/subscribers",
	"repos/{owner}/{repo}/subscription",
	"repos/{owner}/{repo}/tags",
	"repos/{owner}/{repo}/tarball/{ref}",
	"repos/{owner}/{repo}/teams",
	"repos/{owner}/{repo}/topics",
	"repos/{owner}/{repo}/traffic/clones",
	"repos/{owner}/{repo}/traffic/popular/paths",
	"repos/{owner}/{repo}/traffic/popular/referrers",
	"repos/{owner}/{repo}/traffic/views",
	"repos/{owner}/{repo}/transfer",
	"repos/{owner}/{repo}/vulnerability-alerts",
	"repos/{owner}/{repo}/zipball/{ref}",
	"repos/{template_owner}/{template_repo}/generate",
	"repositories",
	"repositories/{repository_id}/environments/{environment_name}/secrets",
	"repositories/{repository_id}/environments/{environment_name}/secrets/public-key",
	"repositories/{repository_id}/environments/{environment_name}/secrets/{secret_name}",
	"scim/v2/enterprises/{enterprise}/Groups",
	"scim/v2/enterprises/{enterprise}/Groups/{scim_group_id}",
	"scim/v2/enterprises/{enterprise}/Users",
	"scim/v2/enterprises/{enterprise}/Users/{scim_user_id}",
	"scim/v2/organizations/{org}/Users",
	"scim/v2/organizations/{org}/Users/{scim_user_id}",
	"search/code",
	"search/commits",
	"search/issues",
	"search/labels",
	"search/repositories",
	"search/topics",
	"search/users",
	"teams/{team_id}",
	"teams/{team_id}/discussions",
	"teams/{team_id}/discussions/{discussion_number}",
	"teams/{team_id}/discussions/{discussion_number}/comments",
	"teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}",
	"teams/{team_id}/discussions/{discussion_number}/comments/{comment_number}/reactions",
	"teams/{team_id}/discussions/{discussion_number}/reactions",
	"teams/{team_id}/invitations",
	"teams/{team_id}/members",
	"teams/{team_id}/members/{username}",
	"teams/{team_id}/memberships/{username}",
	"teams/{team_id}/projects",
	"teams/{team_id}/projects/{project_id}",
	"teams/{team_id}/repos",
	"teams/{team_id}/repos/{owner}/{repo}",
	"teams/{team_id}/team-sync/group-mappings",
	"teams/{team_id}/teams",
	"user",
	"user/blocks",
	"user/blocks/{username}",
	"user/email/visibility",
	"user/emails",
	"user/followers",
	"user/following",
	"user/following/{username}",
	"user/gpg_keys",
	"user/gpg_keys/{gpg_key_id}",
	"user/installations",
	"user/installations/{installation_id}/repositories",
	"user/installations/{installation_id}/repositories/{repository_id}",
	"user/interaction-limits",
	"user/issues",
	"user/keys",
	"user/keys/{key_id}",
	"user/marketplace_purchases",
	"user/marketplace_purchases/stubbed",
	"user/memberships/orgs",
	"user/memberships/orgs/{org}",
	"user/migrations",
	"user/migrations/{migration_id}",
	"user/migrations/{migration_id}/archive",
	"user/migrations/{migration_id}/repos/{repo_name}/lock",
	"user/migrations/{migration_id}/repositories",
	"user/orgs",
	"user/packages/{package_type}/{package_name}",
	"user/packages/{package_type}/{package_name}/restore",
	"user/packages/{package_type}/{package_name}/versions",
	"user/packages/{package_type}/{package_name}/versions/{package_version_id}",
	"user/packages/{package_type}/{package_name}/versions/{package_version_id}/restore",
	"user/projects",
	"user/public_emails",
	"user/repos",
	"user/repository_invitations",
	"user/repository_invitations/{invitation_id}",
	"user/starred",
	"user/starred/{owner}/{repo}",
	"user/subscriptions",
	"user/teams",
	"users",
	"users/{username}",
	"users/{username}/events",
	"users/{username}/events/orgs/{org}",
	"users/{username}/events/public",
	"users/{username}/followers",
	"users/{username}/following",
	"users/{username}/following/{target_user}",
	"users/{username}/gists",
	"users/{username}/gpg_keys",
	"users/{username}/hovercard",
	"users/{username}/installation",
	"users/{username}/keys",
	"users/{username}/orgs",
	"users/{username}/packages/{package_type}/{package_name}",
	"users/{username}/packages/{package_type}/{package_name}/versions",
	"users/{username}/packages/{package_type}/{package_name}/versions/{package_version_id}",
	"users/{username}/projects",
	"users/{username}/received_events",
	"users/{username}/received_events/public",
	"users/{username}/repos",
	"users/{username}/settings/billing/actions",
	"users/{username}/settings/billing/packages",
	"users/{username}/settings/billing/shared-storage",
	"users/{username}/starred",
	"users/{username}/subscriptions",
	"zen",
}

func ActionPackageTypes() carapace.Action {
	return carapace.ActionValues("npm", "maven", "rubygems", "nuget", "docker", "container")
}
