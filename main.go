package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math/big"
)

func main() {
	coin := os.Args[1]

	keysPerPage := 128

	switch coin {
	case "btc":
		printBitcoinKeys(os.Args[2], keysPerPage)
	case "btc-search":
		printBtcWifSearch(os.Args[2], keysPerPage)
	case "eth":
		printEthereumKeys(os.Args[2], keysPerPage)
	case "eth-search":
		printEthPrivateKeySearch(os.Args[2], keysPerPage)
	case "eth-range":
		printEthKeyForRange(os.Args[2],os.Args[3],keysPerPage)
	case "eth-addr-search":
		printSearchedEthAddr(os.Args[2],keysPerPage)
	default:
		log.Fatal("Invalid coin type")
	}
}

func printSearchedEthAddr(ethereumAddress string,keysPerPage int) {
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)
	var one = big.NewInt(1)
	for i:=big.NewInt(1);i.Cmp(&limit) < 0 ; i.Add(i,one) {
		fmt.Print("%v",i.String())
		ethereumKeys := generateEthereumKeys(i.String(), keysPerPage)
		length := len(ethereumKeys)
		for j, key := range ethereumKeys {
			eth := key.public
			if(strings.Contains(eth,ethereumAddress)){
				fmt.Printf("%v", key)
				break
			}
			if j != length-1 {
				fmt.Print("\n")
			}
		}
	}
}

func printEthKeyForRange(rangeStart string,rangeEnd string,keysPerPage int) {
	start,_ := strconv.Atoi(rangeStart)
	end,_ := strconv.Atoi(rangeEnd)
	for i :=start; i<=end;i++ {
		printEthereumKeys(strconv.Itoa(i),keysPerPage)
	}
}

func printBitcoinKeys(pageNumber string, keysPerPage int) {
	bitcoinKeys := generateBitcoinKeys(pageNumber, keysPerPage)

	length := len(bitcoinKeys)

	for i, key := range bitcoinKeys {
		fmt.Printf("%v", key)

		if i != length-1 {
			fmt.Print("\n")
		}
	}
}

func printBtcWifSearch(wif string, keysPerPage int) {
	pageNumber := findBtcWifPage(wif, keysPerPage)

	fmt.Printf("%v", pageNumber)
}

func printEthereumKeys(pageNumber string, keysPerPage int) {
	ethereumKeys := generateEthereumKeys(pageNumber, keysPerPage)

	length := len(ethereumKeys)

	for i, key := range ethereumKeys {
		fmt.Printf("%v", key)

		if i != length-1 {
			fmt.Print("\n")
		}
	}
}

func printEthPrivateKeySearch(privateKey string, keysPerPage int) {
	pageNumber := findEthPrivateKeyPage(privateKey, keysPerPage)

	fmt.Printf("%v", pageNumber)
}
