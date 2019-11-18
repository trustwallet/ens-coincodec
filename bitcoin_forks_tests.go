package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestLitecoinDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "LaMT348PWRnrqeeWArpwQPbuanpXDZGEUz",
			output: "76a914a5f4d12ce3685781b227c1f39548ddef429e978388ac",
		},
		{
			name:   "P2SH",
			input:  "MQMcJhpWHYVeQArcZR3sBgyPZxxRtnH441",
			output: "a914b48297bff5dadecc5f36145cec6a5f20d57c8f9b87",
		},
		{
			name:   "P2WPKH",
			input:  "ltc1qdp7p2rpx4a2f80h7a4crvppczgg4egmv5c78w8",
			output: "0014687c150c26af5493befeed7036043812115ca36c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinDecodeToBytes(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BitcoinDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("BitcoinDecodeToBytes() = %v, want %v, err: %v", hex.EncodeToString(got), tt.output, tt.err)
				}
			}
		})
	}
}

func TestLitecoinEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a914a5f4d12ce3685781b227c1f39548ddef429e978388ac")
	script2, _ := hex.DecodeString("a914b48297bff5dadecc5f36145cec6a5f20d57c8f9b87")
	script3, _ := hex.DecodeString("0014687c150c26af5493befeed7036043812115ca36c")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  script1,
			output: "LaMT348PWRnrqeeWArpwQPbuanpXDZGEUz",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "MQMcJhpWHYVeQArcZR3sBgyPZxxRtnH441",
		},
		{
			name:   "P2WPKH",
			input:  script3,
			output: "ltc1qdp7p2rpx4a2f80h7a4crvppczgg4egmv5c78w8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BitcoinEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("BitcoinEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}

func TestDogeDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "DBXu2kgc3xtvCUWFcxFE3r9hEYgmuaaCyD",
			output: "76a9144620b70031f0e9437e374a2100934fba4911046088ac",
		},
		{
			name:   "P2SH",
			input:  "AF8ekvSf6eiSBRspJjnfzK6d1EM6pnPq3G",
			output: "a914f8f5d99a9fc21aa676e74d15e7b8134557615bda87",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinDecodeToBytes(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BitcoinDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("BitcoinDecodeToBytes() = %v, want %v, err: %v", hex.EncodeToString(got), tt.output, tt.err)
				}
			}
		})
	}
}

func TestDogeEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a9144620b70031f0e9437e374a2100934fba4911046088ac")
	script2, _ := hex.DecodeString("a914f8f5d99a9fc21aa676e74d15e7b8134557615bda87")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  script1,
			output: "DBXu2kgc3xtvCUWFcxFE3r9hEYgmuaaCyD",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "AF8ekvSf6eiSBRspJjnfzK6d1EM6pnPq3G",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BitcoinEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("BitcoinEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}

func TestDashDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "XtAG1982HcYJVibHxRZrBmdzL5YTzj4cA1",
			output: "76a914bfa98bb8a919330c432e4ff16563c5ab449604ad88ac",
		},
		{
			name:   "P2SH",
			input:  "7gks9gWVmGeir7m4MhsSxMzXC2eXXAuuRD",
			output: "a9149d646d71f0815c0cfd8cd08aa9d391cd127f378687",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinDecodeToBytes(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BitcoinDecodeToBytes() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if !reflect.DeepEqual(hex.EncodeToString(got), tt.output) {
					t.Errorf("BitcoinDecodeToBytes() = %v, want %v, err: %v", hex.EncodeToString(got), tt.output, tt.err)
				}
			}
		})
	}
}

func TestDashEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a914bfa98bb8a919330c432e4ff16563c5ab449604ad88ac")
	script2, _ := hex.DecodeString("a9149d646d71f0815c0cfd8cd08aa9d391cd127f378687")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  script1,
			output: "XtAG1982HcYJVibHxRZrBmdzL5YTzj4cA1",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "7gks9gWVmGeir7m4MhsSxMzXC2eXXAuuRD",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := BitcoinEncodeToString(tt.input)
			if tt.err != nil {
				if err.Error() != tt.err.Error() {
					t.Errorf("BitcoinEncodeToString() error = %v, wantErr %v", err, tt.err)
					return
				}
			} else {
				if got != tt.output {
					t.Errorf("BitcoinEncodeToString() = %v, want %v", got, tt.output)
				}
			}
		})
	}
}
