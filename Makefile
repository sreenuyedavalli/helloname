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


