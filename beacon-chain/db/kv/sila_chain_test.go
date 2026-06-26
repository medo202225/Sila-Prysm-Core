package kv

import (
	"testing"

	v2 "github.com/sila-chain/Sila-Consensus-Core/v7/proto/sila/v1alpha1"
)

func TestStore_SavePowchainData(t *testing.T) {
	type args struct {
		data *v2.SilaChainData
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "nil data",
			args: args{
				data: nil,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := setupDB(t)
			if err := store.SaveSilaChainData(t.Context(), tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("SaveSilaChainData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
