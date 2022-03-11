package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"time"
)

const (
	buyOperation      = "BUY"
	sellOperation     = "SELL"
	convertOperation  = "CONVERT"
	withdrawOperation = "WITHDRAW"
	dateFormat        = "01/02/2006 15:04:05 -0700"
	dateFileFormat    = "01-02-2006-15:04:05"
)

type config struct {
	rotationInterval         time.Duration
	transactionInterval      time.Duration
	transactionTotalLifetime time.Duration
}

func main() {
	directoryFlag := flag.String("dir", "testdata", "the directory to store the transaction files in")
	intervalFlag := flag.Duration("interval", time.Minute, "interval between each transaction")
	rotationFlag := flag.Duration("rotation", time.Hour, "rotation interval between each transaction file")
	totalFlag := flag.Duration("total", time.Hour*10, "total lifetime of all transactions")

	flag.Parse()

	err := os.MkdirAll(*directoryFlag, 0777)
	if err != nil {
		log.Fatalf("could not craete directory: %v", err)
	}

	then := time.Now().UTC()
	cfg := config{
		rotationInterval:         *rotationFlag,
		transactionInterval:      *intervalFlag,
		transactionTotalLifetime: *totalFlag,
	}

	now := time.Now().UTC()
	for {
		if now.Sub(then) > cfg.transactionTotalLifetime {
			return
		}

		bs := generateTransactions(cfg, now)
		filename := fmt.Sprintf("transaction-%s.txt", now.Format(dateFileFormat))
		file, err := os.Create(path.Join(*directoryFlag, filename))
		if err != nil {
			log.Fatalf("could not create file: %v", err)
		}
		_, err = file.Write(bs)
		if err != nil {
			log.Fatalf("could not write to file: %v", err)
		}

		now = now.Add(cfg.rotationInterval+cfg.transactionInterval)
	}
}

func generateTransactions(cfg config, now time.Time) []byte {
	addresses := []string{
		"0xa42c9E5B5d936309D6B4Ca323B0dD5739643D2Dd",
		"0x7F1C681EF8aD3E695b8dd18C9aD99Ad3A1469CEb",
		"0xD534d113C3CdDFB34bC9D78d85caE4433E6B6326",
		"0x3ddda9438c70f06ce31Bb364788b47EF113e06F9",
		"0x1312395388f9f8F0AF11bfc50Bae8284962732b1",
		"0x980Bc04e435C5E948B1f70a69cD377783500757b",
		"0x120aE479935B4dB6e8bAea92Ac82Efed60165777",
		"0xFfEC835E4fEF2038F8CBC1170fD5d3bf3122bCd5",
		"0x72C3996FC71f485D95C705aE8A167380e4a891af",
		"0x2e23acC09912b6327766179E5F861679D50b5a9b",
		"0x07bb6FBE0e76492FeA01f740D01Ec796e5468968",
		"0x1C28aA9E5Bd21c62153Dae1AD19F6cc9305C15c1",
		"0xf56167Fa1CD74FD6d761E015758a3CE6BE4466F5",
		"0xd1ABA973674601DD10FEF7Abb239E4e975E26a44",
		"0x4bA6b63527B81B82d6b5eDf75E960e071FA21937",
		"0xc68c701B5904fB27Ec72Cc8ff062530a0ffd2015",
		"0xeeaFf5e4B8B488303A9F1db36edbB9d73b38dFcf",
		"0x3a623858c4e9E8649D9Fbb01e7aE3248d12D2b3E",
		"0x00B2cf90D4aDD5023A0e2CF29516fE72E3A02e2c",
		"0xf9Fb58eB4871590764987ac1b1244b3AE4135626",
	}
	cryptoCoins := []string{"BTC", "ETH", "USDT", "BUSD", "SOL", "DOT", "LUNA"}
	fiatCoins := []string{"USD", "EUR", "GBP"}
	operations := []string{buyOperation, sellOperation, convertOperation, withdrawOperation}
	maxAmounts := map[string]float64{
		"BTC":  2,
		"ETH":  20,
		"USDT": 5000,
		"BUSD": 5000,
		"SOL":  50,
		"DOT":  100,
		"LUNA": 80,
	}
	prices := map[string]float64{
		"BTC":  41000,
		"ETH":  2700,
		"USDT": 0.9999,
		"BUSD": 0.9999,
		"SOL":  90,
		"DOT":  18,
		"LUNA": 90,
	}
	buyFee := 2.0
	sellFee := 3.0
	withdrawFee := 15

	buf := &bytes.Buffer{}
	then := now
	for {
		if now.Sub(then) > cfg.rotationInterval {
			break
		}
		rand.Seed(time.Now().UnixNano())
		addressIndex := rand.Intn(len(addresses))
		address := addresses[addressIndex]
		cryptoCoinIndex := rand.Intn(len(cryptoCoins))
		cryptoCoinIndexAlt := rand.Intn(len(cryptoCoins))
		cryptoCoin := cryptoCoins[cryptoCoinIndex]
		cryptoCoinAlt := cryptoCoins[cryptoCoinIndexAlt]
		fiatCoinIndex := rand.Intn(len(fiatCoins))
		fiatCoin := fiatCoins[fiatCoinIndex]
		operationIndex := rand.Intn(len(operations))
		operation := operations[operationIndex]
		amount := rand.Float64() * maxAmounts[cryptoCoin]
		price := (prices[cryptoCoin]*rand.Float64() + prices[cryptoCoin]) / 2
		date := now.Format(dateFormat)

		line := ""
		switch operation {
		case buyOperation:
			line = fmt.Sprintf("%s %s %s/%s:%.2f %s:%.2f %v%%(%.2f %s) %s", address, operation, cryptoCoin, fiatCoin, price, fiatCoin, amount, buyFee, amount*buyFee/100, fiatCoin, date)
		case sellOperation:
			line = fmt.Sprintf("%s %s %s/%s:%.2f %s:%.2f %v%%(%.2f %s) %s", address, operation, cryptoCoin, fiatCoin, price, fiatCoin, amount, sellFee, amount*sellFee/100, fiatCoin, date)
		case convertOperation:
			if cryptoCoin == cryptoCoinAlt {
				continue
			}
			line = fmt.Sprintf("%s %s %s/%s:%.2f %s:%.2f %v%% %s", address, operation, cryptoCoin, cryptoCoinAlt, price, cryptoCoinAlt, price*amount, 0, date)
		case withdrawOperation:
			line = fmt.Sprintf("%s %s %s/%s:%.2f %s:%.2f %v%s %s", address, operation, cryptoCoin, fiatCoin, price, fiatCoin, amount*price, withdrawFee, fiatCoin, date)
		}

		buf.WriteString(fmt.Sprintf("%s\n", line))
		now = now.Add(cfg.transactionInterval)
	}

	return buf.Bytes()
}
