# charon
🪬 A verifier for mev-boost commitments sent by proposers to relays 🪬

## Introduction

[mev-boost](https://boost.flashbots.net/) introduced a way for the entire Ethereum network to 
verify if a Proposer used a block sent by a Relay without having modified it - remember that a 
proposer commits to use the most profitable payload sent by the relay and only this one.

This is possible because the Relays which implement the Data Transparency API expose these 
payloads. Once they have been proposed, we can compare those two and check if they match or not.

## Charon

Charon is a project focused on exposing if the Proposers using mev-boost break their commitments 
of using the payload they received from the Relays.
This can be the case when a Proposer decides to use the most profitable block it received and 
extract more value for itself.

## Usage

First, you must either set a few environment variable values using the ones provided in the [`.
envrc.example`](./.envrc.example) file or set the corresponding values in the
[`config.example.json`](./config.example.json) or [`config.example.yaml`](./config.example.yaml) file.

Then, run the following command to build the project:
```shell
go build main.go -o charon
```

To start a single analysis of a given slot, use the following command (don't provide the 
`--slot` flag if you already set it in the environment / configuration file):
```shell
./charon verify --slot <SLOT_NUMBER>
```

## Disclaimer

This project is a personal project and might contain issues resulting in false data being reported.
Refer to the [License](./LICENSE) for more.
For now, I recommend not to use any of the resulting data for external projects.
Please contact one of the maintainer if you've found any issue.

## Authors

Made with ❤️ by 🤖 [0xpanoramix](https://github.com/0xpanoramix/) 🤖
