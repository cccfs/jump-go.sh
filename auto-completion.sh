
if [[ -n "$ZSH_VERSION" ]]; then
    autoload bashcompinit
    bashcompinit
fi

if [ -z ${ALICLOUD_REGION+x} ]; then
  echo "var is unset"
  exit 1
fi
if [ -z ${ALICLOUD_ACCESS_KEY+x} ]; then
  echo "var is unset"
  exit 1
fi
if [ -z ${ALICLOUD_SECRET_KEY+x} ]; then
  echo "var is unset"
  exit 1
fi

_jumpsh()
{
    local cur opts
    COMPREPLY=()
    cur="${COMP_WORDS[COMP_CWORD]}"
    names="$(~/.jumpgohost/bin/find-hosts.sh | cut -f1)"

    COMPREPLY=( $(compgen -W "${names}" -- ${cur}) )
    return 0
}
complete -F _jumpsh jump-go.sh

export PATH="$HOME/.jumpgohost/bin:$PATH"
