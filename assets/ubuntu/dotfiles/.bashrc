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
[ -r /usr/share/bash-completion/bash_completion ] && . /usr/share/bash-completion/bash_completion

# Alias definitions.
[ -f ~/.bash_aliases ] && . ~/.bash_aliases

# Bash functions.
[ -f ~/.bash_functions ] && . ~/.bash_functions

# The command prompt.
export PS1="\[\e[01;32m\]\u@\h\[\e[m\] \[\e[01;34m\] \w\[\e[m\] \[\e[01;33m\]\$(parse_git_branch)\[\e[m\]\n\$ "
