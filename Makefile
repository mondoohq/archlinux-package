update: update-mondoo update-cnquery

update-mondoo:
	go run ./generator/main.go mondoo ./mondoo

update-cnquery:
	go run ./generator/main.go cnquery ./cnquery

