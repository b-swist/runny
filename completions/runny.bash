#!/usr/bin/env bash

_runny_comp() {
	local cword cur
	cword="$COMP_CWORD"
	cur="${COMP_WORDS[cword]}"

	local flags
	flags="-v -h"

	if [ $cword -eq 1 ]; then
		COMPREPLY=($(compgen -W "${flags}" -- "${cur}"))
		return 0
	fi

	COMPREPLY=()
	return 0
}

complete -F _runny_comp runny
