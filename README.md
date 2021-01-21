# GitHub CLI

[gh](https://github.com/cli/cli) with dynamic completion

[![asciicast](https://asciinema.org/a/358690.svg)](https://asciinema.org/a/358690)

## Usage

```sh
#bash
source <(gh _carapace)

# elvish
eval (gh _carapace|slurp)

# fish
gh _carapace | source

# powershell
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
gh _carapace | Out-String | Invoke-Expression

# xonsh
COMPLETIONS_CONFIRM=True
exec($(gh _carapace))

# zsh
source <(gh _carapace)
```
