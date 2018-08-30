# If not running interactively, don't do anything.
[[ $- != *i* ]] && return

# Don't put duplicate lines or lines starting with space in the history.
HISTCONTROL=ignoreboth

# Append to the history file, don't overwrite it.
shopt -s histappend

# For setting history length see HISTSIZE and HISTFILESIZE in bash(1).
HISTSIZE=1000
HISTFILESIZE=2000

# Check the window size after each command and, if necessary,
# update the values of LINES and COLUMNS.
shopt -s checkwinsize

# If set, the pattern "**" used in a pathname expansion context will
# match all files and zero or more directories and subdirectories.
shopt -s globstar

# Display matches for ambiguous patterns at first tab press.
bind "set show-all-if-ambiguous on"

# Correct spelling errors during tab-completion.
shopt -s dirspell 2> /dev/null
# Correct spelling errors in arguments supplied to cd.
shopt -s cdspell 2> /dev/null

# Set EDITOR to Neovim.
export EDITOR=nvim
export VISUAL=nvim

# Load the Bash Completion script.
[[ -r /usr/share/bash-completion/bash_completion ]] && source /usr/share/bash-completion/bash_completion

# Alias definitions.
[[ -f ~/.bash_aliases ]] && source ~/.bash_aliases

# Bash functions.
[[ -f ~/.bash_functions ]] && source ~/.bash_functions

PROMPT_USER="\[\e[01;32m\]\u"
PROMPT_HOSTNAME="\h\[\e[m\]"
PROMPT_USER_AND_HOSTNAME="${PROMPT_USER}@${PROMPT_HOSTNAME}"
PROMPT_CWD="\[\e[01;34m\]:: \w\[\e[m\]"
PROMPT_GIT_BRANCH="\[\e[01;33m\]\$(parse_git_branch)\[\e[m\]"

# The command prompt.
export PS1="${PROMPT_USER_AND_HOSTNAME} ${PROMPT_CWD} ${PROMPT_GIT_BRANCH}\n\$ "

source <(kubectl completion bash)
source <(minikube completion bash)
