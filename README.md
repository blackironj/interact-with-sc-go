# Interact with smart contract using golang
 It is a sample project for interacting with a smart contract using golang

## Prerequisite
- node v16
- Burnable wallet address & private key for testing
  > :warning: I would recommand that do not use a real wallet for testing!
- go >= 1.17
- solidity >= 0.8.10
  - https://docs.soliditylang.org/en/v0.8.0/installing-solidity.html
- ethereum tool
  - https://geth.ethereum.org/docs/install-and-build/installing-geth
  - we need to `abigen` tool
- Install yarn
  ```sh
  npm install -g yarn
  ```

## Quick start
- installation
    ```sh
    make all
    ```
- How to use it
    ```sh
    ./sc-interactor --help
    ```
    ```sh
    Usage:
        sc-interactor [command]
    
    Available Commands:
        help        help about any command

        balance     account token balance
        deploy      deploy a contract to network
        mint        mint tokens
        burn        burn tokens
        transfer    transfer tokens to others
    ```
