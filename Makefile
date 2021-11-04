.DEFAULT: default
default: help

.PHONY: help
help: ## Shows the help screen
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\-_0-9%:\\/]+:.*?## / {printf "\033[36m%-20s→\033[0m %s\n", $$1, $$2}' Makefile | sed s/:// | sort | fzf | cut -d ' ' -f1 | xargs -r make

