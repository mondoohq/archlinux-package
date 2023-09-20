update: update-mondoo update-cnquery update-cnspec

update-mondoo:
	go run ./generator/main.go mondoo ./mondoo

update-cnquery:
	go run ./generator/main.go cnquery ./cnquery

update-cnspec:
	go run ./generator/main.go cnspec ./cnspec

# Copywrite Check Tool: https://github.com/hashicorp/copywrite
license: license/headers/check

license/headers/check:
	copywrite headers --plan

license/headers/apply:
	copywrite headers