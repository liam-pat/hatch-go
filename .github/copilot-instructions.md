# Copilot Instructions for hatch-go

## Project Overview
- This repository is a comprehensive Go language study resource, organized by topic and feature.
- Code is grouped by concept: data structures (`a_data_structure/`), functions (`a_function/`), syntax (`a_syntax/`), types (`a_type/`), concurrency (`b_channel/`, `b_goroutine/`), error handling (`c_error/`), serialization (`c_json_xml/`), encryption (`d_encryption/`), reflection (`d_reflection/`), RPC (`e_rpc/`), web (`e_web/`), standard packages (`f_package/`), networking (`f_tcp_udp/`), algorithms (`g_leetcode/`, `g_user_fun/`), and tests (`gotest/`).
- Each subfolder contains focused, idiomatic Go examples. Use these as reference for style and structure.

## Key Workflows
- **Module Management:**
  - Use `go mod tidy` to update dependencies, `go mod vendor` to vendor them, and `go mod graph` to inspect relationships.
  - Replace modules with `go mod edit -replace ...` for local or alternate sources.
- **Build & Format:**
  - Build with `go build` (main package) or `go install` (library).
  - Format code using `go fmt` or `go format`.
- **Testing:**
  - Run all tests: `go test .`
  - Coverage: `go test -coverprofile=c.out` then `go tool cover -html=c.out -o coverage.html`
  - Benchmark: `go test -bench . -cpuprofile cpu.out` then `go tool pprof cpu.out`

## Patterns & Conventions
- **File Naming:**
  - Folders and files are named by topic and function, e.g., `chain.go` for linked lists, `goroutine.go` for concurrency.
- **Example-Driven:**
  - Each file is a self-contained example. Prefer adding new concepts as new files in the relevant topic folder.
- **Minimal External Dependencies:**
  - Most code is standard library; exceptions are documented in comments or `go.mod`.
- **Testing:**
  - Tests are in `gotest/` and use Go's standard `testing` and `benchmark` patterns.
- **RPC/Web:**
  - Client/server code is split by protocol in `e_rpc/` and `e_web/`.

## Integration Points
- **Standard Library:**
  - Heavy use of Go's built-in packages for IO, concurrency, networking, and encoding.
- **External Modules:**
  - Managed via `go.mod`; update and tidy as needed.

## Examples
- To add a new data structure, create a file in `a_data_structure/` and follow the style of `list.go` or `queue.go`.
- For a new algorithm, use `g_user_fun/` or `g_leetcode/` as a template.
- For package demos, see `f_package/` subfolders (e.g., `bufio/main.go`).

## References
- See `README.md` for build, test, and module commands.
- Use `gotest/` for test and benchmark patterns.
- For RPC/web, see `e_rpc/client/` and `e_rpc/server/`.

---

If any section is unclear or missing, please provide feedback to improve these instructions.
