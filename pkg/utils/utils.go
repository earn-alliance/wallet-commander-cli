package utils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

func GetPublicKeyFromPrivateKeyStr(privateKeyStr string) string {
	privateKey := crypto.ToECDSAUnsafe(common.FromHex(privateKeyStr))
	return crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
}

func RoninAddrToEthAddr(addr string) string {
	return strings.Replace(addr, "ronin:", "0x", -1)
}

func EthAddrToRoninAddr(addr string) string {
	return strings.Replace(addr, "0x", "ronin:", -1)
}

func NodejsHashData(data []byte) []byte {
	// https://stackoverflow.com/questions/54287013/different-signing-results-from-go-ethereum-and-ethereumjs-utils
	return crypto.Keccak256([]byte(fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(data), data)))
}
