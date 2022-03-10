package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const (
	buyOperation      = "BUY"
	sellOperation     = "SELL"
	convertOperation  = "CONVERT"
	withdrawOperation = "WITHDRAW"
	dateFormat        = "01/02/2006 15:04:05 -0700"
)

func main() {
	transactionIntervalFlag := flag.Duration("transaction-interval", time.Minute, "interval between each transaction")
	rotationIntervalFlag := flag.Duration("rotation-interval", time.Hour, "interval between each transaction file")

	flag.Parse()

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
	fiatCoins := []string{"USD", "EUR", "MDL"}
	operations := []string{buyOperation, sellOperation, convertOperation, withdrawOperation}
	//maxAmounts := map[string]float64{
	//}
	buyFee := 2.0
	withdrawFee := 15

	now, then := time.Now().UTC(), time.Now().UTC()
	for {
		if now.Sub(then) > *rotationIntervalFlag {
			return
		}
		rand.Seed(time.Now().UnixNano())
		addressIndex := rand.Intn(len(addresses))
		address := addresses[addressIndex]
		cryptoCoinIndex := rand.Intn(len(cryptoCoins))
		cryptoCoin := cryptoCoins[cryptoCoinIndex]
		fiatCoinIndex := rand.Intn(len(fiatCoins))
		fiatCoin := fiatCoins[fiatCoinIndex]
		operationIndex := rand.Intn(len(operations))
		operation := operations[operationIndex]
		date := now.Format(dateFormat)

		// make sure in and out coins are different otherwise skip iteration

		line := ""
		switch operation {
		case buyOperation:
			line = fmt.Sprintf("%s %s %s:%v %s:%v %v%% %s", address, operation, cryptoCoin, 1, fiatCoin, 123, buyFee, date)
		case sellOperation:
			line = fmt.Sprintf("%s %s %s:%v %s:%v %v%% %s", address, operation, cryptoCoin, 1, fiatCoin, 123, 0, date)
		case convertOperation:
			line = fmt.Sprintf("%s %s %s:%v %s:%v %v%% %s", address, operation, cryptoCoin, 1, cryptoCoin, 123, 0, date)
		case withdrawOperation:
			line = fmt.Sprintf("%s %s %s:%v %s:%v %v%s %s", address, operation, cryptoCoin, 1, cryptoCoin, 123, withdrawFee, fiatCoin, date)
		}

		fmt.Println(line)
		now = now.Add(*transactionIntervalFlag)
	}
}
