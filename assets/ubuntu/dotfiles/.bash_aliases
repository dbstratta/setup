# cd aliases
alias cd..='cd ..'
alias c..='c ..'
alias ..='c ..'
alias ...='c ../..'
alias ....='c ../../..'
alias .....='c ../../../..'
alias ......='c ../../../../..'
alias .......='c ../../../../../..'
alias ........='c ../../../../../../..'

# Useful aliases
alias :q='exit'
alias sl='ls'

# ls aliases
alias ls='exa'
alias ll='ls --all --long --header'
alias la='ls --all'

# Color aliases
alias grep='grep --color=auto'
alias fgrep='fgrep --color=auto'
alias egrep='egrep --color=auto'

# Git aliases
alias g='git'
alias s='git status'

# Docker aliases
alias d='docker'
alias dco='docker-compose'

# Kubernetes aliases
alias kb='kubectl'
alias mk='minikube'

# Enables completion for aliases
_completion_loader git
complete -o bashdefault -o default -o nospace -F _git g

_completion_loader docker
complete -o bashdefault -o default -o nospace -F _docker d

_completion_loader docker-compose
complete -o bashdefault -o default -o nospace -F _docker_compose dco

_completion_loader minikube
complete -o bashdefault -o default -o nospace -F __start_minikube mk

# Neovim aliases
alias vim='nvim'
