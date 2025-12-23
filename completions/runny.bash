#!/usr/bin/env bash

_runny_comp() {
	local cword cur prev
	cword="$COMP_CWORD"
	cur="${COMP_WORDS[cword]}"
	prev="${COMP_WORDS[cword - 1]}"

	local flags modes
	flags="-v -h -m"
	modes="apps path"

	if [ $cword -eq 1 ]; then
		COMPREPLY=($(compgen -W "${flags}" -- "${cur}"))
		return 0
	fi

	if [ "$prev" = "-m" ] && [ $cword -eq 2 ]; then
		COMPREPLY=($(compgen -W "$modes" -- "$cur"))
		return 0
	fi

	COMPREPLY=()
	return 0
}

complete -F _runny_comp runny
