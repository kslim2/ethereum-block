package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("wss://rinkeby.infura.io/ws/v3/7d14d40090fb43268ce98e468a098c03")
    if err != nil {
        log.Fatal(err)
    }

    headers := make(chan *types.Header)
    sub, err := client.SubscribeNewHead(context.Background(), headers)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case header := <-headers:
            fmt.Println(header.Hash().Hex()) // 0xbc10defa8dda384c96a17640d84de5578804945d347072e091b4e5f390ddea7f

            block, err := client.BlockByHash(context.Background(), header.Hash())
            if err != nil {
                log.Fatal(err)
            }

			fmt.Println("received at: ", block.ReceivedAt)
			fmt.Println("received from: ", block.ReceivedFrom)
			fmt.Println("base fee: ", block.BaseFee())
			fmt.Println("bloom:", block.Bloom())
			fmt.Println("body: ", block.Body())
			fmt.Println("coin base: ", block.Coinbase())
			fmt.Println("extra: ", block.Extra())
			fmt.Println("gas limit: ", block.GasLimit())
			fmt.Println("gas used: ", block.GasUsed())
			fmt.Println("block hash: ", block.Hash())
			fmt.Println("block header", block.Header())
			fmt.Println("transaction hash: ", block.Header())
			fmt.Println("block mixed digest: ", block.MixDigest())
			fmt.Println("block nonce: ", block.Nonce())
			fmt.Println("block number: ", block.Number())
			fmt.Println("parent hash: ", block.ParentHash())
			fmt.Println("receipt hash: ", block.ReceiptHash())
			fmt.Println("root: ", block.Root())
			fmt.Println("size: ", block.Size())
			fmt.Println("time: ", block.Time())
			fmt.Println("transactions: ", block.Transactions())
			fmt.Println("txhash: ", block.TxHash())
			fmt.Println("uncle hash: ", block.UncleHash())
			fmt.Println("uncles: ", block.Uncles())
			fmt.Println("==============================================================================")
			
        }
    }
}