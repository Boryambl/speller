BIN=speller
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
ifeq ($(OS), Windows_NT)
	BIN=speller.exe
endif

all:
	@go clean
	@echo "Building speller server"
	@echo "Compiling"
	@cd ./cmd/;\
	GOOS=${GOOS} GOARCH=${GOARCH} go version
	@cd ./cmd/;\
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ../bin/${BIN}
	@echo "Done"