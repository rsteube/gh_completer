# GitHub CLI

[gh](https://github.com/cli/cli) with dynamic completion


[![asciicast](https://asciinema.org/a/358690.svg)](https://asciinema.org/a/358690)

## Usage

```sh
#bash
source <(gh completion)

# elvish
gh completion > gh.elv
-source gh.elv

# fish
gh completion | source

# powershell
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete
gh completion | Out-String | Invoke-Expression

# zsh
source <(gh completion)
```
