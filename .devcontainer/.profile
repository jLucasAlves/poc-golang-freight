# in this file you can customize the sh behavior updating PS1, setting envs and so many more...
# by thefault PS1 will display user@ folder <git branch>
# when you customize this file, you only need to close the current terminal instance
# clicking in the garbage icon and then, open again and the new config will be applied

export PS1="\[\033[36m\]\u\[\033[m\]@\[\033[32m\] \[\033[33;1m\]\w\[\033[m\] \[\033[31m\]\$(git branch --show-current 2>/dev/null)\[\033[m\] \$ "

alias ll="ls -la"
