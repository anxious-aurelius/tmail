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

After building, move the binary to a directory on your PATH so you can run `tmail` from anywhere:

```bash
mv tmail /usr/local/bin/
```

## Configuration

Run the init command to scaffold a config file at `~/.tmail/config.toml`:

```bash
tmail config init
```

Then open `~/.tmail/config.toml` and fill in your credentials:

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

## Usage

```bash
# Scaffold your config file (first-time setup)
tmail config init

# View or debug loaded config
tmail config

# Send an email
tmail send

# List emails from your inbox
tmail list
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
│   ├── config.go
│   └── configInit.go
├── internal/
│   ├── config/           # Config loading and parsing (TOML)
│   │   └── config.go
│   ├── imap/             # IMAP logic
│   │   └── imap.go
│   ├── mail/             # Shared domain types (Message, Envelope, Address)
│   │   └── mail.go
│   └── smtp/             # SMTP send logic
│       └── smtp.go
└── main.go
```

## Tech Stack

- [Cobra](https://github.com/spf13/cobra): CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea): TUI framework (Future Implementation)
- [go-imap](https://github.com/emersion/go-imap): IMAP client
- [BurntSushi/toml](https://github.com/BurntSushi/toml): Config parsing

## Contributing

Contributions are welcome! Feel free to open issues or pull requests. This project follows standard Go conventions and values clean separation of concerns. Email logic lives in `internal/`, completely decoupled from the UI layer.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
