ifndef LOCATION
$(error LOCATION is not set)
endif

.PHONY: all plan apply destroy

all: plan

plan:
	cd $(LOCATION) && \
	terraform plan -var-file terraform.tfvars -out terraform.tfplan

create:
	cd $(LOCATION) && \
	terraform apply -var-file terraform.tfvars

destroy:
	cd $(LOCATION) && \
	terraform plan -destroy -var-file terraform.tfvars -out terraform.tfplan
	terraform apply terraform.tfplan
