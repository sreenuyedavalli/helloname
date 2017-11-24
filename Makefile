ifndef LOCATION
$(error Please export LOCATION to TFVAR file path)
endif

.PHONY: all plan apply destroy

all: plan

plan:
	@echo "Terraform Plan:"
	cd $(LOCATION) && \
	terraform plan -var-file terraform.tfvars -out terraform.tfplan

create:
	@echo "Creating Helloname Infra:"
	cd $(LOCATION) && \
	terraform apply -var-file terraform.tfvars

destroy:
	@echo "Destroying Helloname Infra:"
	cd $(LOCATION) && \
	terraform plan -destroy -var-file terraform.tfvars -out terraform.tfplan && \
	terraform apply terraform.tfplan

deploy:
	@echo "Deploying helloname"
	helm upgrade --namespace prd hellonconnect ./helloname-chart

install: 
	@echo "Installing Initial Helm Chart"
	helm install -namespace prd helloconnect ./helloname-chart

run:	
	@echo "Go running app"
	go run main.go model.go app.go

build:  
	@echo "Building docker image"
	docker build -t sreenuyedavalli/helloname:latest .

push:   
	@echo "Pushing docker image to sreenuyedavalli/helloname"
	docker push sreenuyedavalli/helloname:latest
