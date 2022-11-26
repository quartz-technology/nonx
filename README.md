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

## Authors

Made with ❤️ by 🤖 [0xpanoramix](https://github.com/0xpanoramix/) 🤖
