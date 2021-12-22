# Golang Commons
[![CI](https://github.com/vpnbeast/golang-commons/workflows/CI/badge.svg?event=push)](https://github.com/vpnbeast/golang-commons/actions?query=workflow%3ACI)
[![Go Report Card](https://goreportcard.com/badge/github.com/vpnbeast/golang-commons)](https://goreportcard.com/report/github.com/vpnbeast/golang-commons)
[![codecov](https://codecov.io/gh/vpnbeast/golang-commons/branch/master/graph/badge.svg)](https://codecov.io/gh/vpnbeast/golang-commons)
[![Go version](https://img.shields.io/github/go-mod/go-version/vpnbeast/golang-commons)](https://github.com/vpnbeast/golang-commons)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is a common library for [Vpnbeast](https://github.com/vpnbeast) backend services which are written with Golang.

## Development
This project requires below tools while developing:
- [Golang 1.16](https://golang.org/doc/go1.16)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)

After installed [pre-commit](https://pre-commit.com/), make sure that you've completed the below final installation steps:
- Make sure that you've installed [pre-commit](https://pre-commit.com/) for our git repository in root directory of the project:
  ```shell
  $ pre-commit install
  ```
- Add below custom variables to `.git/hooks/pre-commit` in the root of our git repository:
  ```python
  # custom variable definition for local development
  os.environ["CONFIG_PATH"] = "./"
  os.environ["ACTIVE_PROFILE"] = "unit-test"
  ```

## License
Apache License 2.0
