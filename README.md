Business Logic for Multiply two numbers Service.

# 1. Requirements

| Software         | Version | Importance                   |
| ---------------- | ------- | ---------------------------- |
| 🐳 Docker         | 20.10.9 | Required                     |
| 🐙 Docker Compose | 1.29.2  | Required                     |
| 🐹 Go             | 1.17.1  | Optional (if you use docker) |
| 🐃 GNU Make       | 4.2.1   | Optional                     |

# 2. Testing

- I create a `multiply_test.go` to perform tests.
- to learn more see the file `Makefile`

## 2.1 Using your local machine
- Execute make `make l/test` to run the test.

## 2.1 Using Docker compose
- Execute `make d/test` or run:

```shell
docker run --rm -it -w /multiplyLogic -v $PWD/.:/multiplyLogic golang:1.17-alpine go test ./...
```

> ⚠️ Warning ⚠️
>
> This command only works for UNIX-Like systems. On PowerShell make sure to replace the `$PWD` command with `%cd%`

# 3. CI/CD

The `.github` directory contains the pipelines to testing the code in these two cases:

- When someone pushed on `main` branch.
- When someone sent a `Pull Request`.

# 4. How to use into your project
Execute the command:
```shell
go get github.com/kenesparta/multiplyLogic@latest
```
