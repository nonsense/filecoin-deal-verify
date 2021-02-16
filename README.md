# Filecoin Oracle on Ethereum

This is a tool to verify merkle inclusion proofs produced by the Filecoin oracle trusted web service versus the latest uploaded merkle root hash on Ethereum.

It was built as part of the Protocol Labs remote hack week held from February 1st, 2020 to February 5th, 2020.

## Architecture

The experimental Filecoin Oracle consists of two parts:

1. `smart contracts` - [Solidity smart contracts for Ethereum](https://github.com/nonsense/filecoin-oracle)

2. `web oracle` - [A trusted web service which monitors the state of the Filecoin blockchain](https://github.com/dirkmc/filecoin-deal-proofs-svc)

---

The web oracle continuously monitors the Filecoin blockchain, once an hour processes the state for all deals, and produces a merkle tree root hash of the serialized data. This service is backed by the [Filecoin Sentinel](https://github.com/filecoin-project/sentinel).

Users are able to query data CIDs of interest on the web oracle and get a merkle inclusion proof with all the relevant data for the data CID at that point in time:`dataCid`, `pieceCid`, `dealId`, `provider`, `startEpoch`, `endEpoch`, `signedEpoch`

## Installation

```
go install ./...
```

## Usage

```
Usage of verify-proof:
  -dataCID string

  -dealID int

  -endEpoch int

  -pieceCID string

  -proof string

  -provider string

  -signedEpoch int

  -startEpoch int
```
