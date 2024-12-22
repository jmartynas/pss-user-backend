PROGRAM_NAME = pss-user-backend

run: build
	env $$(cat .env) ./$(PROGRAM_NAME)

build:
	go build -o $(PROGRAM_NAME)

.PHONY: all build
