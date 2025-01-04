<div align="center">

# Go Faucet

![Go version](https://img.shields.io/badge/Go-1.23%2B-00ADD8?style=for-the-badge&logo=go)
![Go Ethereum version](https://img.shields.io/badge/Go--Ethereum-1.14.12-363636?style=for-the-badge&logo=ethereum)
![DiscordGo](https://img.shields.io/badge/DiscordGo-0.23.0-7289DA?style=for-the-badge&logo=discord)

<br/>

Faucet core engine written in go with discord bot interface

</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#%EF%B8%8F-quick-start">⚡️ Quick Start</a>
    <li><a href="#prerequisites">Prerequisites</a></li>
    <li><a href="#executing-program">Executing Program</a></li>
  </ol>
</details>


## ⚡️ Quick start

### Prerequisites

The project utilizes the following tools:
- **Go** `1.23+`: A statically typed, compiled programming language designed for simplicity, concurrency, and performance. Go is used for building the bot and handling interactions with the blockchain.
- **Go-Ethereum** `1.14.12`: A Go implementation of the Ethereum protocol, enabling the bot to interact with the Ethereum blockchain for cryptocurrency distribution.
- **Blockchain Network**: A local or test network like Ethereum’s Rinkeby, Goerli, or a private Ethereum network for testing and development of the faucet functionality.
- **Make** `4.3+` **(optional)**: A build automation tool used for managing tasks. The Makefile provides predefined commands for setting up, building, and running the project, simplifying the development workflow.

### Executing Program

Clone the repository

```bash
git clone git@github.com:alitdarmaputra/go-faucet.git
cd go-faucet
```

Install module

```bash
go mod tidy
```

Configure env according to your own value

```bash
mv .env.example .env
```

Run the program

```bash
go run src/main.go

# or run using make

make run
```

Build the project (optional)

```bash
make build

# execute the binary

make exec
```
