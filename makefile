
write-conf:
	cd terraform; cp backend-default.conf backend.conf
	cd terraform; sed -i 's/ENV/$(env)/g' backend.conf
	cd terraform; sed -i 's/STACK/$(stack)/g' backend.conf

dev: build
	cd terraform; cp backend-default.conf backend.conf
	cd terraform; sed -i 's/ENV/dev/g' backend.conf
	cd terraform; sed -i 's/STACK/g-password-manager/g' backend.conf
	terraform -chdir=terraform init -backend-config=backend.conf
	terraform -chdir=terraform validate
	terraform -chdir=terraform apply -var env=dev

prod: build
	cd terraform; cp backend-default.conf backend.conf
	cd terraform; sed -i 's/ENV/prod/g' backend.conf
	cd terraform; sed -i 's/STACK/g-password-manager/g' backend.conf
	terraform -chdir=terraform init -backend-config=backend.conf
	terraform -chdir=terraform validate
	terraform -chdir=terraform apply -var env=prod

init:
	terraform -chdir=terraform init -backend-config=backend.conf

deploy:
	terraform -chdir=terraform apply -var env=$(env)

destroy:
	terraform -chdir=terraform destroy -var env=$(env)

output:
	terraform -chdir=terraform output > outputs

build:
	cd api; make build
	cd app; make build

clean:
	rm -rf dist/*