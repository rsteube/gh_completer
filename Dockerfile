FROM golang

# shells
RUN wget https://packages.microsoft.com/config/debian/10/packages-microsoft-prod.deb \
 && dpkg -i packages-microsoft-prod.deb

RUN apt-get update \
 && apt-get install -y bash-completion \ 
                       fish \
                       elvish \
                       locales \
                       powershell \
                       python3-pip \
                       zsh

ENV GOPATH /go
RUN ln -s /gh/cmd/gh/gh /usr/local/bin/gh

# bash
RUN echo "\n\
source /usr/share/bash-completion/bash_completion \n\
source <(gh _carapace)" \
       > /root/.bashrc

# fish
# Set the locale
RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen \
 && locale-gen
ENV LANG en_US.UTF-8  
ENV LANGUAGE en_US:en  
ENV LC_ALL en_US.UTF-8     

RUN mkdir -p /root/.config/fish \
 && echo "gh _carapace | source" \
       > /root/.config/fish/config.fish

# elvish
RUN curl https://dl.elv.sh/linux-amd64/elvish-HEAD.tar.gz | tar -xvz \
 && mv elvish-* /usr/local/bin/elvish

RUN mkdir -p /root/.elvish/lib \
 && echo "eval (gh _carapace|slurp)" \
  > /root/.elvish/rc.elv

# powershell
RUN mkdir -p /root/.config/powershell \
 && echo "\n\
Set-PSReadlineKeyHandler -Key Tab -Function MenuComplete\n\
gh _carapace | Out-String | Invoke-Expression" \
       > /root/.config/powershell/Microsoft.PowerShell_profile.ps1

# xonsh
RUN pip3 install --no-cache-dir --disable-pip-version-check xonsh \
 && ln -s $(which xonsh) /usr/bin/xonsh

RUN mkdir -p ~/.config/xonsh \
 && echo "\n\
\$COMPLETIONS_CONFIRM=True\n\
exec(\$(gh _carapace))"\
  > ~/.config/xonsh/rc.xsh

# zsh
RUN echo "\n\
zstyle ':completion:*' menu select \n\
zstyle ':completion:*' matcher-list 'm:{a-zA-Z}={A-Za-z}' 'r:|=*' 'l:|=* r:|=*' \n\
\n\
autoload -U compinit && compinit \n\
source <(gh _carapace)"\
  > /root/.zshrc

WORKDIR /gh
