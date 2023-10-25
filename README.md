### Golang Hexagonal Architecture

Boilerplate for [hexagonal architecture](https://www.happycoders.eu/software-craftsmanship/hexagonal-architecture/) in Golang.
There are few endpoints like Auth and Furniture for demonstration purpose. Feel free to remove them (or anything else).

### Dev Setup

1. The following are required for development purpose:
    - system packages: [just](https://github.com/casey/just), [pre-commit](https://pre-commit.com), [go 1.21+](https://go.dev/doc/install), [protobuf](https://grpc.io/docs/protoc-installation/), [`atlas`](https://github.com/ariga/atlas)
    - go binaries: `just install-bins`

2. Make sure to add pre-commit hooks in local git repo: `pre-commit install`

3. Install runtime dependencies: `just tidy`. This updates go.mod and installs those packages.

4. Verify that things are working
    - copy `.env.default` as `.env` at project root
    - apply migration: `just runc migrate`
    - run server: `just runc server` (in container) or `just run server` (natively with hot-reloading)
    - run tests: `just test`

Run `just list` to view shorthands for other project-specific commands. It is highly recommended to go through them. These will make life a bit easier.

### Code Guidelines

**Readability**
- Keep the most important things in a file at top so that a reader knows at the first glance what concerns the file handles. This typically includes exported functions, structs, etc.
- Use the standard settings for formatting long lines: `just golines path/to/source`.

**Nomenclature**
- Package names:
  - should tell what it contains
  - must be [short, lowercase, and singular (no plurals)](https://go.dev/blog/package-names)
  - prefer unserscore over mixed-caps where we have to resort to multi-word package name
- File names:
  - should tell what it contains
  - must be short and lowercase, but not necessarily singular
  - use period as separator in case multi-word names cannot be avoided
- Struct, function, and variable names:
  - must be lowercase or mixedCaps for unexported ones
  - must be CamelCase for exported ones
  - must NOT contain underscores
  - must NOT be same as some package name
    - e.g. if `repo` is a package, then `repos := repo.InitRepo()` is okay, but `repo = repo.InitRepo()` is not
- Test:
  - functions must follow the nomenclature `TestXxx_Yyy`, where `Xxx` identifies the function, and `Yyy`  describes the test

**Error handling**
- Do not ignore error values returned from function calls. Handle them.
- Usually, functions should return `(ReturnType, error)`.
- Use `logger.Panic(err)` wherever possible instead of built-in `panic()`.
- Log errors at the deepest (terminal) point in the call chain. Usually, this point is where we return the control back to the previous layer; or call a server, a datastore, or a library. This way, the entire call stack is logged.
  - We do this because there is no built-in exception handling in Golang. [Errors are just values](https://go.dev/blog/errors-are-values). A stack trace of errors is not available to the upper layers in a call chain.

**Other recommended practices**
- Write [idiomatic](https://go.dev/doc/effective_go) Golang code.
- [Accept Interface, Return Struct](https://bryanftan.medium.com/accept-interfaces-return-structs-in-go-d4cab29a301b)
- `ctx context.Context` should be first parameter of a function wherever necessary e.g. long-running functions.
- Dependencies like logger, config, etc. should be at the beggining of the function parameter list.
- Avoid silencing the linter using `//nolint`. It must be followed most of the times.
