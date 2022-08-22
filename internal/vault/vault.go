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
	AddRoninWallet(roninWallet string, privateKey string) bool
	GetWalletCounts() int
	Save() error
}

type WalletCommanderVault struct {
	secretsFile string
	secretsMap  map[string]string
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
		secretsFile: secretsFile,
		secretsMap:  secretsMap,
	}, nil

}

func (w WalletCommanderVault) AddRoninWallet(roninWallet string, privateKey string) bool {
	if _, ok := w.secretsMap[roninWallet]; ok {
		return false // wallet already exists, don't update
	}

	w.secretsMap[roninWallet] = privateKey

	return true
}

func (w WalletCommanderVault) Save() error {

	secretsJsonBytes, err := json.MarshalIndent(w.secretsMap, "", "    ")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(w.secretsFile, secretsJsonBytes, 0660)
}

func (w WalletCommanderVault) GetPrivateKey(roninWallet string) (string, error) {
	if key, ok := w.secretsMap[roninWallet]; ok {
		return key, nil
	} else {
		return "", errors.New("roninWallet does not exist in secrets map")
	}
}

func (w WalletCommanderVault) GetWalletCounts() int {
	return len(w.secretsMap)
}
