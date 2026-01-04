# mals-ctl

## Build

```sh
go build -o build/mals-ctl cmd/*.go
```

## Usage

```sh
./build/mals-ctl -h
```

## TODO

Commands:

`log`

- `ls` - list all
- `get [<name>]` - print comprehensive info

`listener`

- `ls`
- `get [<name>]`

`lsp`

- `ls`
- `get [<name>]`

`model`

- `ls`
- `get [<name>]`

`usage`

- `ls`
- `get [<name>]`

`scope`

- `tree`

`config`

- `server`
  - `ls`
  - `get [<name>]`
  - `add <name> <url>`
  - `remove <name>`

- `context`
  - `get`
  - `set server <name>`
