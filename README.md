# Query Builder

![CI](https://github.com/Qyroxen/Query-Builder/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Query-Builder/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Query-Builder?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Query-Builder)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Query-Builder)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Query-Builder?style=social)](https://github.com/Qyroxen/Query-Builder/stargazers)

## What is it?

Query Builder is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Query-Builder.git
cd Query-Builder
go build -o querybuilder .

# Run
./querybuilder --help
```

## CLI Usage

```bash
# Basic usage
./querybuilder

# With flags
./querybuilder --verbose --output json

# Get help
./querybuilder --help
```

## Examples

```bash
# Example 1
./querybuilder example1

# Example 2
./querybuilder example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o querybuilder .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Query-Builder/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Query-Builder?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Query-Builder/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Query-Builder?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Query-Builder/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Query-Builder" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Query-Builder/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Query-Builder" alt="Pull Requests">
  </a>
</p>
