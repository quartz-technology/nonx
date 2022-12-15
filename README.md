# <h1 align="center"> nonx </h1>

<p align="center">
    <img src="./.github/assets/dalle_cover.png" width="400" alt="A DALL-E representation of a 
photo of the Grim Reaper on a boat with a purple background and a dark theme in cyberpunk style">
</p>

<p align="center">
🪬 A verifier for mev-boost commitments sent by proposers to relays 🪬
</p>

Cover by [DALL-E](https://openai.com/dall-e-2/).

## Introduction

[mev-boost](https://boost.flashbots.net/) introduced a way for the entire Ethereum network to 
verify if a Proposer used a block sent by a Relay without having modified it - remember that a 
proposer commits to use the most profitable payload sent by the relay and only this one.

This is possible because the Relays which implement the Data Transparency API expose these 
payloads. Once they have been proposed, we can compare those two and check if they match or not.

## Nonx

Nonx is a project focused on exposing if the Proposers using mev-boost break their commitments 
of using the payload they received from the Relays.
This can be the case when a Proposer decides to use the most profitable block it received and 
extract more value for itself.

## Usage

First, you must either set a few environment variable values using the ones provided in the [`.envrc.example`](./.envrc.example)
file or set the corresponding values in the [`config.example.json`](./config.example.json) or 
[`config.example.yaml`](./config.example.yaml) file.

Then, run the following command to build the project:
```shell
make all
```

### Single slot verification

To start a single analysis of a given slot, use the following command (don't provide the 
`--slot` flag if you already set it in the environment / configuration file):
```shell
./nonx verify --slot <SLOT_NUMBER>
```

Example:
```shell
$ - ./nonx verify --slot 5253886
INFO[0000] ✅ commitment has been respected by proposer   committed_payload_hash=0xcffee69df32c924ff2de6a3975b52ce69208c540e7afd8f9cdc9b7efac119cff proposed_payload_hash=0xcffee69df32c924ff2de6a3975b52ce69208c540e7afd8f9cdc9b7efac119cff slot=5253886
```

### Watcher

To run the watcher, which will verify every new payload delivered by the relay to the proposers, 
run the following command:
```
./nonx watch
```

Example:
```shell
$ - ./nonx watch
INFO[0000] ✅ commitment has been respected by proposer   committed_payload_hash=0x255ccbf4e495d768ce453d30f2ae7050ed7123d24bf3d5e1b24940ab5908b499 proposed_payload_hash=0x255ccbf4e495d768ce453d30f2ae7050ed7123d24bf3d5e1b24940ab5908b499 slot=5280458
INFO[0024] ✅ commitment has been respected by proposer   committed_payload_hash=0x7b999b60dc70bedb3238b56a94f533cf446cdb077110d9781140323d203dca8b proposed_payload_hash=0x7b999b60dc70bedb3238b56a94f533cf446cdb077110d9781140323d203dca8b slot=5280461
INFO[0048] ✅ commitment has been respected by proposer   committed_payload_hash=0x7d3c120dce693edeecc55a16699e20ab1b88147a6f2f31bc95dfcb806b53eda2 proposed_payload_hash=0x7d3c120dce693edeecc55a16699e20ab1b88147a6f2f31bc95dfcb806b53eda2 slot=5280463
INFO[0060] ✅ commitment has been respected by proposer   committed_payload_hash=0xf2992e00d4cbbda6a62bec6dedcac5440ccfacc1141bd6242fcb79cf2ea3ec49 proposed_payload_hash=0xf2992e00d4cbbda6a62bec6dedcac5440ccfacc1141bd6242fcb79cf2ea3ec49 slot=5280464
```

## Limitations

The current implementation can produce wrong results because the beacon client you might use can 
take some time to get the latest proposed block, which would lead to the analysis to fail 
(comparing the relay delivered payload to an empty hash).


## Disclaimer

This project is a personal project and might contain issues resulting in false data being reported.
Refer to the [License](./LICENSE) for more.
For now, I recommend not to use any of the resulting data for external projects.
Please contact one of the maintainer if you've found any issue.

## Authors

Made with ❤️ by 🤖 [0xpanoramix](https://github.com/0xpanoramix/) 🤖
