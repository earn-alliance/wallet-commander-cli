package vault

import (
	"encoding/json"
	"errors"
	"github.com/earn-alliance/wallet-commander-cli/internal/log"
	"io/ioutil"
	"os"
)

type Vault interface {
	GetPrivateKey(roninWallet string) (string, error)
}

type WalletCommanderVault struct {
	secretsMap map[string]string
}

func New(secretsFile string) (Vault, error) {
	secretsJsonFile, err := os.Open(secretsFile)
	defer secretsJsonFile.Close()

	if err != nil {
		return nil, err
	}

	var secretsMap map[string]string
	secretesJsonBytes, _ := ioutil.ReadAll(secretsJsonFile)
	if err := json.Unmarshal(secretesJsonBytes, &secretsMap); err != nil {
		return nil, err
	}

	log.Logger().Infof("[VAULT] Successfully loaded %d wallets and their secrets", len(secretsMap))

	return &WalletCommanderVault{
		secretsMap: secretsMap,
	}, nil

}

func (w WalletCommanderVault) GetPrivateKey(roninWallet string) (string, error) {
	if key, ok := w.secretsMap[roninWallet]; ok {
		return key, nil
	} else {
		return "", errors.New("roninWallet does not exist in secrets map")
	}
}
