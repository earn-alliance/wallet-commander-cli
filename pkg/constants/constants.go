package constants

import "math/big"

const (
	SLP_CONTRACT            = "0xa8754b9fa15fc18bb59458815510e40a12cd2014"
	WETH_CONTRACT           = "0xc99a6a985ed2cac1ef41640596c5a5f9f4e19ef5"
	AXS_CONTRACT            = "0x97a9107c1793bc407d6f527b77e7fff4d812bece"
	AXIE_CONTRACT           = "0x32950db2a7164ae833121501c797d79e7b79d74c"
	MARKETPLACE_CONTRACT    = "0x213073989821f738a7ba3520c3d31a1f9ad31bbd"
	RONIN_PROVIDER_FREE_URI = "https://proxy.roninchain.com/free-gas-rpc"
	RONIN_PROVIDER_RPC_URI  = "https://api.roninchain.com/rpc"

	DEFAULT_GAS_LIMIT = uint64(500000)
)

var (
	RONIN_CHAIN_ID = big.NewInt(2020)
)
