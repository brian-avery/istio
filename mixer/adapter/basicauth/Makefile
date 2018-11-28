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

deploy:
	@echo "Deploying to istio-system"
	kubectl delete pod -n istio-system -l app=basicauthadapter
	kubectl apply -f operatorconfig/attributes.yaml -n istio-system
	kubectl apply -f operatorconfig/authorizationtemplate.yaml -n istio-system
	kubectl apply -f operatorconfig/basicauth-k8s.yaml -n istio-system
	kubectl apply -f operatorconfig/basicauthadapter.yaml -n istio-system
	kubectl apply -f operatorconfig/sample_operator_cfg.yaml -n istio-system
