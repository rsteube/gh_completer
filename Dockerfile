FROM golang

# shells
RUN wget https://packages.microsoft.com/config/debian/10/packages-microsoft-prod.deb \
 && dpkg -i packages-microsoft-prod.deb \
 && apt-get update

RUN apt-get install -y bash-completion \ 
                       fish \
                       elvish \
                       powershell \
                       python3-pip \
                       zsh

ENV GOPATH /go
RUN ln -s /gh/cmd/gh/gh /usr/local/bin/gh

# bash
RUN echo "\n\
source /usr/share/bash-completion/bash_completion \n\
source <(gh completion)" \
       > /root/.bashrc

# fish
RUN mkdir -p /root/.config/fish \
 && echo "gh completion | source" \
       > /root/.config/fish/config.fish

# elvish
RUN curl https://dl.elv.sh/linux-amd64/elvish-HEAD.tar.gz | tar -xvz \
 && mv elvish-* /usr/local/bin/elvish

RUN mkdir -p /root/.elvish/lib \
 && echo "eval (gh completion|slurp)" \
  > /root/.elvish/rc.elv

# powershell
RUN mkdir -p /root/.config/powershell \
 && echo "\n\
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete\n\
gh completion | Out-String | Invoke-Expression" \
       > /root/.config/powershell/Microsoft.PowerShell_profile.ps1

# xonsh
RUN pip3 install --no-cache-dir --disable-pip-version-check xonsh \
 && ln -s $(which xonsh) /usr/bin/xonsh

RUN mkdir -p ~/.config/xonsh \
 && echo "\n\
\$COMPLETIONS_CONFIRM=True\n\
exec(\$(gh completion))"\
  > ~/.config/xonsh/rc.xsh

# zsh
RUN echo "\n\
zstyle ':completion:*' menu select \n\
zstyle ':completion:*' matcher-list 'm:{a-zA-Z}={A-Za-z}' 'r:|=*' 'l:|=* r:|=*' \n\
\n\
autoload -U compinit && compinit \n\
source <(gh completion)"\
  > /root/.zshrc

WORKDIR /gh
