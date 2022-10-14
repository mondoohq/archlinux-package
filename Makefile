update: update-mondoo update-cnquery update-cnspec

update-mondoo:
	go run ./generator/main.go mondoo ./mondoo

update-cnquery:
	go run ./generator/main.go cnquery ./cnquery

update-cnspec:
	go run ./generator/main.go cnspec ./cnspec

