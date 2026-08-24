# 📈 Query Builder

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v4.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Analytics tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`analytics` `reporting` `cli` `golang`

---

## What is Query-Builder?

**Query-Builder** is a CLI tool built with Go for fast, offline-capable operations.

## Features

- ✅ Formatted output
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Query-Builder.git
cd Query-Builder

# Build
go build -o query-builder .

# Run
./query-builder <table> [cols] [--where k=v] [--order col] [--limit N]
```

### Or directly with `go run`:
```bash
go run main.go <table> [cols] [--where k=v] [--order col] [--limit N]
```

## Usage

```bash
# Basic usage
./query-builder <table> [cols] [--where k=v] [--order col] [--limit N]

# With flags
./query-builder <table> [cols] [--where k=v] [--order col] [--limit N] value <table> [cols] [--where k=v] [--order col] [--limit N]
```

### Example Output

```
$ ./query-builder <table> [cols] [--where k=v] [--order col] [--limit N]
<table> [cols] [--where k=v] [--order col] [--limit N]
SELECT %s FROM %s%s%s%s;\n
```

## Project Structure

```
Query-Builder/
  main.go          # Entry point (44 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
