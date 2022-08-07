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