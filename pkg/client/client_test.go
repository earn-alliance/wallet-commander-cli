package client

import (
	"context"
	"github.com/earn-alliance/wallet-commander-cli/pkg/store"
	"testing"
)

func TestAxieClient_GetClaimableAmount(t *testing.T) {
	c := &AxieClient{
		freeEthClient:     nil,
		ethClient:         nil,
		slpClient:         nil,
		axieClient:        nil,
		marketplaceClient: nil,
		jwtStore:          store.JwtStore{},
	}

	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get claimable amount api respond 200",
			args: args{
				address: "ronin:c21dce25e3a2f1af584aa43841c8f6667e3e389e",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetClaimableAmount(context.Background(), tt.args.address)

			if err != nil {
				t.Errorf("Received error getting claimable amount %v", err)
			}
		})
	}
}

func TestAxieClient_GetClaimPayload(t *testing.T) {
	t.SkipNow()
	c := &AxieClient{
		freeEthClient:     nil,
		ethClient:         nil,
		slpClient:         nil,
		axieClient:        nil,
		marketplaceClient: nil,
		jwtStore:          store.JwtStore{},
	}

	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get claim payload api respond 200",
			args: args{
				address: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//TODO: Enter address and private key for this test. Probably need a test wallet for this!
			var err error
			_, err = c.getClaimPayload(context.Background(), "ronin:b789ce98aa566d1847bd31545d7491461c0af099", "{enter-key-here}")

			if err != nil {
				t.Errorf("Received error getting claimable amount %v", err)
			}
		})
	}
}
