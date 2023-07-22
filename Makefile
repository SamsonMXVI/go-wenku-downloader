test-scraper:
	go test -v ./scraper

test-downloader:
	go test -v ./downloader

build: 
	rm -rf build/wenku_downloader && go build -o build/wenku_downloader

build-exe-amd64:
	rm -rf build/build/wenku_downloader_amd64.exe && GOOS=windows GOARCH=amd64 go build -o build/wenku_downloader_amd64.exe main.go

build-exe-386:
	rm -rf build/wenku_downloader_x86.exe && GOOS=windows GOARCH=amd64 go build -o build/wenku_downloader_x86.exe main.go

build-all: 
	make build && make build-exe-amd64 && make build-exe-386
 
.PHONY: test-scraper test-downloader build-exe-amd64 build-exe-386 build-all build