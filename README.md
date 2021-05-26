# Moved to carapace-bin

Completion moved to [carapace-bin](https://github.com/rsteube/carapace-bin) as external completer:
```sh
carapace gh [bash|elvish|fish|oil|xonsh|zsh]
```


# GitHub CLI

[gh](https://github.com/cli/cli) with dynamic completion

> continued work in the [better-completions](https://github.com/rsteube/gh/tree/better-completions) branch (https://github.com/cli/cli/pull/2728)

[![asciicast](https://asciinema.org/a/358690.svg)](https://asciinema.org/a/358690)

## Usage

```sh
#bash
source <(gh _carapace)

# elvish
eval (gh _carapace|slurp)

# fish
gh _carapace | source

# oil
source <(gh _carapace)

# powershell
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
gh _carapace | Out-String | Invoke-Expression

# xonsh
COMPLETIONS_CONFIRM=True
exec($(gh _carapace))

# zsh
source <(gh _carapace)
```
