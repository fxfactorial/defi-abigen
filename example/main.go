package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fxfactorial/defi-abigen/contracts/aave/lending_pool"
	"github.com/fxfactorial/defi-abigen/contracts/aave/lending_pool_addresses_provider"
	"github.com/pkg/errors"
)

var (
	Client_dial = flag.String(
		"client_dial", "ws://192.168.1.2:9551", "could be websocket or IPC",
	)
	AAVE_ADDR_PROVIDER = common.HexToAddress("0x24a42fD28C976A61Df5D00D0599C34c4f90748c8")
)

const (
	ETHERSCAN_TXNHASH_LINK = "https://etherscan.io/tx/"
	ETHERSCAN_ADDR_LINK    = "https://etherscan.io/address/"
)

func etherscan_of_hash(txn_hash common.Hash) string {
	return ETHERSCAN_TXNHASH_LINK + strings.ToLower(txn_hash.Hex())
}

func program() error {
	client, err := ethclient.Dial(*Client_dial)
	defer client.Close()

	if err != nil {
		return errors.Wrapf(err, "can't connect")
	}

	addr_provider, err := lending_pool_addresses_provider.NewLendingPoolAddressesProvider(
		AAVE_ADDR_PROVIDER, client,
	)

	pool_addr, err := addr_provider.GetLendingPool(nil)

	if err != nil {
		return err
	}

	lendingPool, err := lending_pool.NewLendingPool(pool_addr, client)

	if err != nil {
		return errors.Wrapf(err, "cant load lending pool contract")
	}

	incoming := make(chan *lending_pool.LendingPoolFlashLoan)
	sub, err := lendingPool.WatchFlashLoan(nil, incoming, nil, nil)

	for {
		select {
		case oops := <-sub.Err():
			return oops
		case payload := <-incoming:
			s, _ := json.MarshalIndent(payload, " ", "  ")
			fmt.Printf(
				"aave flashloan %s \n %s\n",
				string(s),
				etherscan_of_hash(payload.Raw.TxHash),
			)

		}
	}
}

func main() {
	if err := program(); err != nil {
		fmt.Println("some problem", err)
		os.Exit(-1)
	}
}
