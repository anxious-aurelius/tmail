# Architecture

`tmail` is organized around a small set of packages with clear dependency boundaries so new commands and transport features can grow without coupling business logic to the CLI.

## Package Map

- `cmd/`: Cobra commands and all user-facing I/O. This layer parses flags, prints results to stdout, and reports fatal command errors back to the process entrypoint.
- `internal/config`: Configuration loading and parsing for SMTP and IMAP credentials.
- `internal/mail`: Dependency-light domain types shared across transports, such as `Message`, `Envelope`, and `Address`.
- `internal/smtp`: Outbound mail delivery built on top of the shared domain types and loaded config.
- `internal/imap`: Inbox access and envelope fetching built on top of the shared domain types and loaded config.
- `internal/ui`: Planned home for Bubble Tea TUI code once the interactive client lands.

## Dependency Rule

Dependencies should always point inward:

`cmd -> {internal/smtp, internal/imap, internal/config} -> internal/mail`

That rule keeps the domain model reusable and prevents transport or presentation code from bleeding across package boundaries.

- `cmd/` may import `internal/config`, `internal/smtp`, `internal/imap`, and future UI packages.
- `internal/smtp` and `internal/imap` may depend on `internal/config` and `internal/mail`.
- `internal/mail` should stay independent of transport, config, and CLI concerns.
- `internal/` packages must never import `cmd/`.
- `internal/` packages must never print directly to the terminal.

## Conventions

- Wrap errors with useful context and return them upward instead of terminating deep in helper packages.
- Send diagnostics to stderr through `slog` once structured logging is wired in.
- Reserve stdout for user-facing command results in `cmd/` only.
- Keep tests beside the code they verify so package boundaries stay obvious.
- Prefer constructors and injected dependencies when adding stateful collaborators, so transports and UI layers remain testable.
