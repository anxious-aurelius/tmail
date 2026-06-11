# tmail

A terminal-based email client built in Go. Designed to be lightweight, scriptable, and fully open source.

## Features

- Send emails from the command line via SMTP
- List emails via IMAP
- Full interactive TUI powered by [Bubble Tea](https://github.com/charmbracelet/bubbletea) *(coming soon)*
- Works with any standard IMAP/SMTP provider
- Simple TOML config file

## Installation

```bash
git clone https://github.com/anxious-aurelius/tmail
cd tmail
go build -o tmail .
```

## Configuration

Create a `config.toml` file in the project root:

```toml
[smtp]
host     = "smtp.example.com"
port     = 587
username = "you@example.com"
password = "your-password"

[imap]
host     = "imap.example.com"
port     = 993
username = "you@example.com"
password = "your-password"
```

> **Note:** `config.toml` is already in `.gitignore` — keep your credentials out of version control.

## Usage

```bash
# Send an email
tmail send

# List emails from your inbox
tmail list

# View or debug loaded config
tmail config
```

## Architecture

See [ARCHITECTURE.md](ARCHITECTURE.md) for the package map, dependency rules, and contributor conventions for keeping CLI code separate from mail-domain logic.

## Project Structure

```
tmail/
├── cmd/                  # Cobra CLI commands (thin launchers)
│   ├── root.go
│   ├── send.go
│   ├── list.go
│   └── config.go
├── internal/
│   ├── config/           # Config loading and parsing (TOML)
│   │   └── config.go
│   ├── imap/             # IMAP logic
│   │   └── imap.go
│   ├── mail/             # Shared domain types (Message, Envelope, Address)
│   │   └── mail.go
│   └── smtp/             # SMTP send logic
│       └── smtp.go
├── main.go
└── config.toml           # Your local credentials (gitignored)
```

## Tech Stack

- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework ( Future Implementation)
- [go-imap](https://github.com/emersion/go-imap) — IMAP client
- [BurntSushi/toml](https://github.com/BurntSushi/toml) — Config parsing

## Contributing

Contributions are welcome! Feel free to open issues or pull requests. This project follows standard Go conventions and values clean separation of concerns — email logic lives in `internal/`, completely decoupled from the UI layer.

## License

MIT
