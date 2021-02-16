package main

import (
	"flag"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/nonsense/filecoin-deal-verify/bindings/oracle"

	logging "github.com/ipfs/go-log/v2"
)

var log = logging.Logger("verify-proof")

var (
	endpoint string
	contract string

	dealID      int
	dataCID     string
	pieceCID    string
	provider    string
	startEpoch  int
	endEpoch    int
	signedEpoch int
	proof       string
)

func init() {
	flag.IntVar(&dealID, "dealID", 0, "")
	flag.StringVar(&dataCID, "dataCID", "", "")
	flag.StringVar(&pieceCID, "pieceCID", "", "")
	flag.StringVar(&provider, "provider", "", "")
	flag.IntVar(&startEpoch, "startEpoch", 0, "")
	flag.IntVar(&endEpoch, "endEpoch", 0, "")
	flag.IntVar(&signedEpoch, "signedEpoch", 0, "")
	flag.StringVar(&proof, "proof", "", "")

	flag.StringVar(&endpoint, "endpoint", "https://rinkeby.infura.io/v3/xxx", "endpoint to an ethereum node")
	flag.StringVar(&contract, "contract", "0xd4375467f6CfB0493b5e4AF0601B3a0f2e7D2FcA", "contract to query and make `VerifyProof` call to")
}

type Deal struct {
	MerkleRoot  string
	DealID      *big.Int
	DataCID     string
	PieceCID    string
	Provider    string
	StartEpoch  *big.Int
	EndEpoch    *big.Int
	SignedEpoch *big.Int
	Proof       string
}

func main() {
	flag.Parse()

	client, err := rpc.Dial(endpoint)
	if err != nil {
		panic(err)
	}
	ethClient := ethclient.NewClient(client)

	fs, err := oracle.NewFilecoinService(common.HexToAddress(contract), ethClient)
	if err != nil {
		panic(err)
	}

	d := &Deal{
		DealID:      big.NewInt(int64(dealID)),
		DataCID:     dataCID,
		PieceCID:    pieceCID,
		Provider:    provider,
		StartEpoch:  big.NewInt(int64(startEpoch)),
		EndEpoch:    big.NewInt(int64(endEpoch)),
		SignedEpoch: big.NewInt(int64(signedEpoch)),
	}

	spew.Dump(d)

	var merkleProof [][32]byte

	entries := len(proof) / 66
	for i := 0; i < entries; i++ {
		start := i * 66
		end := (i + 1) * 66

		entry := proof[start:end]
		spew.Dump(entry)

		slice := common.HexToHash(entry).Bytes()
		var arr [32]byte

		copy(arr[:], slice)

		merkleProof = append(merkleProof, arr)
	}

	spew.Dump(merkleProof)

	tx, err := fs.VerifyProof(nil, d.DataCID, d.PieceCID, d.DealID, d.Provider, d.StartEpoch, d.EndEpoch, d.SignedEpoch, merkleProof)
	if err != nil {
		panic(err)
	}

	log.Info("verifying a proof:")
	spew.Dump(tx)
}
