package modules

import (
	"context"
	"log"
	"math/big"

	Models "github.com/LuisAcerv/goeth-api/models"
	"github.com/ethereum/go-ethereum/ethclient"
)

// GetLatestBlock from blockchain
func GetLatestBlock(client ethclient.Client) *Models.Block {
	header, _ := client.HeaderByNumber(context.Background(), nil)
	blockNumber := big.NewInt(header.Number.Int64())
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	_block := &Models.Block{
		BlockNumber:       block.Number().Int64(),
		Timestamp:         block.Time(),
		Difficulty:        block.Difficulty().Uint64(),
		Hash:              block.Hash().String(),
		TransactionsCount: len(block.Transactions()),
	}

	return _block
}
