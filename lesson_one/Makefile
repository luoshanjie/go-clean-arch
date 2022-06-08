.DEFAULT_GOAL=run
BINARY=app
PROJECT="lesson_one"

build: clean
	@echo "[${PROJECT}] build"
	@go build -o bin/${BINARY} main.go

run: build
	@echo "[${PROJECT}] run"
	bin/${BINARY}

clean:
	@echo "[${PROJECT}] clean"
	@if [ -f bin/${BINARY} ]; then rm bin/${BINARY}; fi

test:
	@echo "[${PROJECT}] test"
	@go test ./...

.PHONY: clean build
