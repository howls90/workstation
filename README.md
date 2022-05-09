## Monorepo

Framework to handler and speed up the creation of new projects.

- [Documentation](https://howls90.github.io/monorepo/)

```bash
.
├── docs
├── backend
│   ├── apps
│   └── libs
├── frontend
│   ├── apps
│   └── libs
└── cli
```

## Requirements

- [Earthly](https://docs.earthly.dev/)
- [Golang](https://golang.org)

## Setup process

Install the necessary packages:

```shell
earthly +install
```

Add the following `github secrets` to your repository:

```shell
GH_TOKEN
```

## Usage

Create new project

```shell
./cli
```
