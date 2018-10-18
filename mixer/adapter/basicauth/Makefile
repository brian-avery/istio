version?=latest
image=basicauthadapter

full_image=$(image):$(version)
ifdef hub
	full_image=$(hub)/$(image):$(version)
endif

build:
	@echo "Building"
	go generate ./...
	cp config/basicauthadapter.yaml operatorconfig
	CGO_ENABLED=0 go build -o $(image) -installsuffix cgo -v cmd/main.go
	docker build -t $(full_image) .

push:
	@echo "Pushing $(full_image)"
	docker push $(full_image)
