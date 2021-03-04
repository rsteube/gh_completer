package action

import (
	"github.com/rsteube/carapace"
)

// https://docs.github.com/en/rest/overview/api-previews
func ActionApiPreviews() carapace.Action {
	return carapace.ActionValuesDescribed(
		"wyandotte-preview", "Migrations",
		"ant-man-preview", "Enhanced deployments",
		"squirrel-girl-preview", "Reactions",
		"mockingbird-preview", "Timeline",
		"inertia-preview", "Projects",
		"cloak-preview", "Commit search",
		"mercy-preview", "Repository topics",
		"scarlet-witch-preview", "Codes of conduct",
		"zzzax-preview", "Require signed commits",
		"luke-cage-preview", "Require multiple approving reviews",
		"starfox-preview", "Project card details",
		"fury-preview", "GitHub App Manifests",
		"flash-preview", "Deployment statuses",
		"surtur-preview", "Repository creation permissions",
		"corsair-preview", "Content attachments",
		"switcheroo-preview", "Enable and disable Pages",
		"groot-preview", "List branches or pull requests for a commit",
		"dorian-preview", "Enable or disable vulnerability alerts for a repository",
		"lydian-preview", "Update a pull request branch",
		"london-preview", "Enable or disable automated security fixes",
		"baptiste-preview", "Create and use repository templates",
		"nebula-preview", "New visibility parameter for the Repositories API",
	)
}
