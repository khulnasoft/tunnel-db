# tunnel-db

![Build DB](https://go.khulnasoft.com/tunnel-db/workflows/Tunnel%20DB/badge.svg)
[![GitHub Release][release-img]][release]
![Downloads][download]
[![Go Report Card][report-card-img]][report-card]
[![Go Doc][go-doc-img]][go-doc]
[![License][license-img]][license]

[download]: https://img.shields.io/github/downloads/khulnasoft/tunnel-db/total?logo=github
[release-img]: https://img.shields.io/github/release/khulnasoft/tunnel-db.svg?logo=github
[release]: https://go.khulnasoft.com/tunnel-db/releases
[report-card-img]: https://goreportcard.com/badge/go.khulnasoft.com/tunnel-db
[report-card]: https://goreportcard.com/report/go.khulnasoft.com/tunnel-db
[go-doc-img]: https://godoc.org/go.khulnasoft.com/tunnel-db?status.svg
[go-doc]: https://godoc.org/go.khulnasoft.com/tunnel-db
[code-cov]: https://codecov.io/gh/khulnasoft/tunnel-db/branch/main/graph/badge.svg
[license-img]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[license]: https://go.khulnasoft.com/tunnel-db/blob/main/LICENSE

## Overview

`tunnel-db` is a CLI tool and a library to manipulate Tunnel DB.

### Library

Tunnel uses `tunnel-db` internally to manipulate vulnerability DB. This DB has vulnerability information from NVD, Red Hat, Debian, etc.

### CLI

`tunnel-db` builds vulnerability DBs on GitHub Actions and uploads them to GitHub Release periodically.

```
NAME:
   tunnel-db - Tunnel DB builder

USAGE:
   main [global options] command [command options] image_name

VERSION:
   0.0.1

COMMANDS:
     build    build a database file
     upload   upload database files to GitHub Release
     help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### Building the DB

You can utilize `make db-all` to build the database, the DB artifact is outputted to the assets folder.

If you want to build the light DB, please set your environment to contain `DB_TYPE=tunnel-light`.

Alternatively Docker is supported, you can run `docker build . -t tunnel-db`.

If you want to build the light DB, please run `docker build --build-arg DB_TYPE=tunnel-light . -t tunnel-db-light`

If you want to build a tunnel integration test DB, please run `make create-test-db`

## Update interval

Every 6 hours
