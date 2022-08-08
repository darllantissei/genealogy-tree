generate-mocks:
	@rm -rf $$(pwd)/mocks/
	@clear
	@set -e ;\
	folderName="" ;\
	fileName="" ;\
	destinyOfMock="" ;\
	for file_interface in $$(find $$(pwd)/application -name "interface.go"); do \
		echo "" ;\
		folderName=$$( echo $$file_interface | rev | cut -d "/" -f -2 | rev | cut -d '/' -f 1) ;\
		fileName=$$(echo $$folderName.go); \
		echo gerando mock para $$fileName ;\
		destinyOfMock=$$(pwd)/mocks/$$folderName/$$fileName ;\
		GOPATH=$$(go env GOPATH) ~/go/bin/mockgen -source=$$file_interface -destination=$$destinyOfMock ;\
		echo mock gerado ;\
		echo $$destinyOfMock ;\
	done;

deploy-app:
	@docker network inspect app-backend >/dev/null 2>&1 || docker network create app-backend
	@docker-compose build --no-cache && docker-compose down -v --remove-orphans && docker-compose up -d --force-recreate
	@docker-compose logs

reset-database:
	@rm -rf database/genealogy-tree.sqlite && sqlite3 database/genealogy-tree.sqlite < sql/schema.sql