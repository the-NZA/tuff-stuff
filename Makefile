APP = tuff
BUILD_FLAGS = -v

# Backend commands
.PHONY: build
build:
	go build $(BUILD_FLAGS) ./backend/cmd/$(APP)

.PHONY: run
run:
	go run $(BUILD_FLAGS) ./backend/cmd/$(APP)

.PHONY: clean
clean:
	rm ./$(APP)

.DEFAULT_GOAL := build

# Frontend commands
.PHONY: start
start:
	cd frontend && npm run dev

.PHONY: buildf
buildf:
	cd frontend && npm run build

# Docker commands
.PHONY: create_image
create_image:
	docker build --no-cache -t $(APP):latest .

.PHONY: compose_up
compose_up:
	docker compose -f docker-compose.dev.yml up