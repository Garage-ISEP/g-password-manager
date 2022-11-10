
write-conf:
	cd terraform; cp backend-default.conf backend.conf
	cd terraform; sed -i 's/ENV/$(env)/g' backend.conf
	cd terraform; sed -i 's/STACK/$(stack)/g' backend.conf

dev:
	cd terraform; cp backend-default.conf backend.conf
	cd terraform; sed -i 's/ENV/dev/g' backend.conf
	cd terraform; sed -i 's/STACK/g-password-manager/g' backend.conf
	terraform -chdir=terraform init -backend-config=backend.conf
	terraform -chdir=terraform validate
	terraform -chdir=terraform apply -var env=dev
