# Local ('l') project scope
l/test:
	go mod tidy
	go test ./... -cover

# Docker ('d') project scope
d/test:
	@docker run --rm -w /multiplyLogic -v $(PWD)/.:/multiplyLogic golang:1.17-alpine go test ./... -cover
