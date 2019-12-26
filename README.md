# UniChain

[![Build Status](https://travis-ci.org/unichainplatform/unichain.svg?branch=master)](https://travis-ci.org/unichainplatform/unichain)
[![GoDoc](https://godoc.org/github.com/unichainplatform/unichain?status.svg)](https://godoc.org/github.com/unichainplatform/unichain)
[![Coverage Status](https://coveralls.io/repos/github/unichainplatform/unichain/badge.svg?branch=master)](https://coveralls.io/github/unichainplatform/unichain?branch=master)
[![GitHub](https://img.shields.io/github/license/unichainplatform/unichain.svg)](LICENSE)

Welcome to the UniChain source code repository!

## What is UniChain?

UniChain is a high-level blockchain framework that can implement the issuance, circulation, and dividends of tokens efficiently and reliably. UniChain can also steadily implement various community governance functions with voting as the core and foundation. These functions are the foundation for building the token economy of future.

home page: https://unichainproject.com/

## Executables

The unichain project comes with several wrappers/executables found in the `cmd` directory.

|    Command     | Description                                                                                                                                                                                                                                                                                                                                                                                               |
| :------------: | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|    **`uni`**    | Our main unichain CLI client. It is the entry point into the unichain network (main-, test- or private net), It can be used by other processes as a gateway into the unichain network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `uni -h` and the [Command Line Options](https://github.com/unichainplatform/unichain/wiki/Command-Line-Options) for command line options. |
| **`unifinder`** | unifinder is a unichain node discoverer.`unifinder -h` and the [Command Line Options](https://github.com/unichainplatform/unichain/wiki/Command-Line-Options) for command line options.                                                                                                                                                                                                                        |

## Getting Started

The following instructions overview the process of getting the code, building it, and start node.

### Getting the code

To download all of the code:

`git clone https://github.com/unichainplatform/unichain`

### Setting up a build/development environment

Install latest distribution of [Go](https://golang.org/) if you don't have it already. (go version >= go1.10 )

Currently supports the following operating systems:

- Ubuntu 16.04
- Ubuntu 18.04
- MacOS Darwin 10.12 and higher

### Build UniChain

`make all`

more information see: [Installing UniChain](https://github.com/unichainplatform/unichain/wiki/Build-UniChain)

### Running a node

To run `./uni` , you can run your own UNI instance.

`$ uni`

Join the unichain main network see: [Main Network](https://github.com/unichainplatform/unichain/wiki/Main-Network)

Join the unichain test network see: [Test Network](https://github.com/unichainplatform/unichain/wiki/Test-Network)

Operating a private network see:[Private Network](https://github.com/unichainplatform/unichain/wiki/Private-Network)

## Resources

[UniChain Official Website](https://unichainproject.com/)

[UniChain Blog](https://unichainproject.com/blog.html)

More Documentation see [the UniChain wiki](https://github.com/unichainplatform/unichain/wiki)

## License

UniChain is distributed under the terms of the [GPLv3 License](./License).
