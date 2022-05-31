package hex_to_int

import (
	"log"
	"math/big"
	"strconv"
	"strings"
)

type IntObj struct {
	hex string
	bytesNumber int
	littleEndian *big.Int
	bigEndian *big.Int
}

func FromHexBigEndianToBigInt(hex string) *big.Int {
	return fromHexToBigInt(handleHex(hex))
}

func FromHexLittleEndianToBigInt(hex string) *big.Int {
	return fromHexToBigInt(reverseHex(handleHex(hex)))
}

func FromStringToHex(s string) string {
	n := fromStringToBigInt(s)

	hex := ""
	for n.Cmp(big.NewInt(16)) == 1 {
		base := big.NewInt(16)

		nBase := new(big.Int)
		nBase.Set(n)

		n.Div(n, base)

		remainder := big.NewInt(0).Sub(nBase, big.NewInt(0).Mul(n, base))
		hex = strconv.FormatInt(remainder.Int64(), 16) + hex

		if n.Cmp(big.NewInt(16)) == -1 {
			hex = strconv.FormatInt(n.Int64(), 16) + hex
		}
	}

	hex = "0x" + hex

	return hex
}

func fromStringToBigInt(s string) *big.Int {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
	if !ok {
		log.Fatal("Invalid number in string")
	}

	return n
}

func fromHexToBigInt(hex string,) *big.Int {
	result := big.NewInt(0)
	for i := 0; i < len(hex); i++ {
		decimal, err := strconv.ParseInt(string(hex[i]), 16, 64)
		if err != nil {
			log.Fatal(err)
		}

		result.Add(result, big.NewInt(0).Mul( big.NewInt(decimal), big.NewInt(0).Exp(big.NewInt(16), big.NewInt(int64(len(hex) - i - 1)), nil)))
	}

	return result
}

func handleHex(hex string) string {
	if len(hex) % 2 != 0 {
		log.Fatal("Invalid HEX length")
	}

	hex = strings.Replace(hex, "0x", "", -1)
	hex = strings.Replace(hex, "0X", "", -1)

	return hex
}

func reverseHex(hex string) string {
	reversed := ""
	for i := len(hex); i > 0; i-=2 {
		reversed = reversed + hex[i-2:i]
	}

	return reversed
}