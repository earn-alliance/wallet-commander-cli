#!/bin/sh
abigen --abi "abis/axieinfinity/axie_abi.json" --pkg abi --type Axie --out pkg/abi/axie.go
abigen --abi "abis/axieinfinity/slp_abi.json" --pkg abi --type Slp --out pkg/abi/slp.go
abigen --abi "abis/axieinfinity/marketplace_abi.json" --pkg abi --type Marketplace --out pkg/abi/marketplace.go
abigen --abi "abis/axieinfinity/ronin_balance_abi.json" --pkg abi --type RoninBalance --out pkg/abi/ronin_balance.go
abigen --abi "abis/axieinfinity/portal_abi.json" --pkg abi --type Portal --out pkg/abi/portal.go