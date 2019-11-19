package coincodec

import (
	"encoding/hex"
	"reflect"
	"testing"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
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
			got, err := toBytesMap[slip44.LITECOIN](tt.input)
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
			got, err := toStringMap[slip44.LITECOIN](tt.input)
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
			got, err := toBytesMap[slip44.DOGECOIN](tt.input)
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
			got, err := toStringMap[slip44.DOGECOIN](tt.input)
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
		{
			name:  "Bitcoin Segwit",
			input: "bc1qw508d6qejxtdg4y5r3zarvary0c5xw7kv8f3t4",
			err:   errors.New("invalid format: version and/or checksum bytes missing"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBytesMap[slip44.DASH](tt.input)
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
			got, err := toStringMap[slip44.DASH](tt.input)
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

func TestMonaDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "MHxgS2XMXjeJ4if2PRRbWYcdwZPWfdwaDT",
			output: "76a9146e5bb7226a337fe8307b4192ae5c3fab9fa9edf588ac",
		},
		{
			name:   "P2SH",
			input:  "PHjTKtgYLTJ9D2Bzw2f6xBB41KBm2HeGfg",
			output: "a9146449f568c9cd2378138f2636e1567112a184a9e887",
		},
		{
			name:   "Segwit",
			input:  "mona1qw508d6qejxtdg4y5r3zarvary0c5xw7kg5lnx5",
			output: "0014751e76e8199196d454941c45d1b3a323f1433bd6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBytesMap[slip44.MONACOIN](tt.input)
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

func TestMonaEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a9146e5bb7226a337fe8307b4192ae5c3fab9fa9edf588ac")
	script2, _ := hex.DecodeString("a9146449f568c9cd2378138f2636e1567112a184a9e887")
	script3, _ := hex.DecodeString("0014751e76e8199196d454941c45d1b3a323f1433bd6")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  script1,
			output: "MHxgS2XMXjeJ4if2PRRbWYcdwZPWfdwaDT",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "PHjTKtgYLTJ9D2Bzw2f6xBB41KBm2HeGfg",
		},
		{
			name:   "Segwit",
			input:  script3,
			output: "mona1qw508d6qejxtdg4y5r3zarvary0c5xw7kg5lnx5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toStringMap[slip44.MONACOIN](tt.input)
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

func TestQtumDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "QYJHEEt8kS8TzUuCy1ia7aYe1cpNg4QYnn",
			output: "76a91480485018e46a9c8176282adf0acb4ff3e0de93ff88ac",
		},
		{
			name:   "P2SH",
			input:  "MHhghmmCTASDnuwpgsPUNJVPTFaj61GzaG",
			output: "a9146b85b3dac9340f36b9d32bbacf2ffcb0851ef17987",
		},
		{
			name:   "Segwit",
			input:  "qc1qxssrzt03ncm0uda02vd8tuvrk0eg9wrz8qm2qe",
			output: "00143420312df19e36fe37af531a75f183b3f282b862",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBytesMap[slip44.QTUM](tt.input)
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

func TestQtumEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a91480485018e46a9c8176282adf0acb4ff3e0de93ff88ac")
	script2, _ := hex.DecodeString("a9146b85b3dac9340f36b9d32bbacf2ffcb0851ef17987")
	script3, _ := hex.DecodeString("00143420312df19e36fe37af531a75f183b3f282b862")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  script1,
			output: "QYJHEEt8kS8TzUuCy1ia7aYe1cpNg4QYnn",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "MHhghmmCTASDnuwpgsPUNJVPTFaj61GzaG",
		},
		{
			name:   "Segwit",
			input:  script3,
			output: "qc1qxssrzt03ncm0uda02vd8tuvrk0eg9wrz8qm2qe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toStringMap[slip44.QTUM](tt.input)
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

func TestVIADecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "Vw6bJFaF5Hyiveko7dpqRjVvcTAsjz7eYa",
			output: "76a914e771c6695c5dd189ccc4ef00cd0f3db3096d79bd88ac",
		},
		{
			name:   "P2SH",
			input:  "ESxRxvhJP6ZKtYaMGjj48As1kgCh6hXa6X",
			output: "a9146b85b3dac9340f36b9d32bbacf2ffcb0851ef17987",
		},
		{
			name:   "Segwit",
			input:  "via1qs32zgdhe2tpzcnz55r7d9jvhce33063s8w4xre",
			output: "001484542436f952c22c4c54a0fcd2c997c66317ea30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBytesMap[slip44.VIACOIN](tt.input)
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

func TestVIAEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a914e771c6695c5dd189ccc4ef00cd0f3db3096d79bd88ac")
	script2, _ := hex.DecodeString("a9146b85b3dac9340f36b9d32bbacf2ffcb0851ef17987")
	script3, _ := hex.DecodeString("001484542436f952c22c4c54a0fcd2c997c66317ea30")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  script1,
			output: "Vw6bJFaF5Hyiveko7dpqRjVvcTAsjz7eYa",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "ESxRxvhJP6ZKtYaMGjj48As1kgCh6hXa6X",
		},
		{
			name:   "Segwit",
			input:  script3,
			output: "via1qs32zgdhe2tpzcnz55r7d9jvhce33063s8w4xre",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toStringMap[slip44.VIACOIN](tt.input)
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

func TestDigiByteDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "DBfCffUdSbhqKZhjuvrJ6AgvJofT4E2kp4",
			output: "76a91447825943ca6a936b177fdc7c9dc05251640169c288ac",
		},
		{
			name:   "P2SH",
			input:  "SUngTA1vaC2E62mbnc81Mdos3TcvZHwsVo",
			output: "a91452356ed3d2d31eb8b263ace5d164e3cf3b37096687",
		},
		{
			name:   "Segwit",
			input:  "dgb1q3p2nf26ac6qtdrv4czh5nmp2eshfj9wyn9vv3d",
			output: "0014885534ab5dc680b68d95c0af49ec2acc2e9915c4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBytesMap[slip44.DIGIBYTE](tt.input)
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

func TestDigiByteEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a91447825943ca6a936b177fdc7c9dc05251640169c288ac")
	script2, _ := hex.DecodeString("a91452356ed3d2d31eb8b263ace5d164e3cf3b37096687")
	script3, _ := hex.DecodeString("0014885534ab5dc680b68d95c0af49ec2acc2e9915c4")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  script1,
			output: "DBfCffUdSbhqKZhjuvrJ6AgvJofT4E2kp4",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "SUngTA1vaC2E62mbnc81Mdos3TcvZHwsVo",
		},
		{
			name:   "Segwit",
			input:  script3,
			output: "dgb1q3p2nf26ac6qtdrv4czh5nmp2eshfj9wyn9vv3d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toStringMap[slip44.DIGIBYTE](tt.input)
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

func TestZcoinDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "a4YtT82mWWxHZhLmdx7e5aroW92dqJoRs3",
			output: "76a9142a10f88e30768d2712665c279922b9621ce58bc788ac",
		},
		{
			name:   "P2SH",
			input:  "4CFa4fnAQvFz4VpikGNzQ9XfCDXMmdk6sh",
			output: "a914f010b17a9189e0f2737d71ae9790359eb5bbc13787",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBytesMap[slip44.ZCOIN](tt.input)
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

func TestZcoinEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a9142a10f88e30768d2712665c279922b9621ce58bc788ac")
	script2, _ := hex.DecodeString("a914f010b17a9189e0f2737d71ae9790359eb5bbc13787")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  script1,
			output: "a4YtT82mWWxHZhLmdx7e5aroW92dqJoRs3",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "4CFa4fnAQvFz4VpikGNzQ9XfCDXMmdk6sh",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toStringMap[slip44.ZCOIN](tt.input)
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

func TestRavenDecodeToBytes(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  "RNoSGCX8SPFscj8epDaJjqEpuZa2B5in88",
			output: "76a9149451f4546e09fc2e49ef9b5303924712ec2b038e88ac",
		},
		{
			name:   "P2SH",
			input:  "rPWwn5h4QFZNaz1XmY39rc73sdYGGDdmq1",
			output: "a914bd92088bb7e82d611a9b94fbb74a0908152b784f87",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toBytesMap[slip44.RAVENCOIN](tt.input)
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

func TestRavenEncodeToString(t *testing.T) {
	script1, _ := hex.DecodeString("76a9149451f4546e09fc2e49ef9b5303924712ec2b038e88ac")
	script2, _ := hex.DecodeString("a914bd92088bb7e82d611a9b94fbb74a0908152b784f87")

	tests := []struct {
		name   string
		input  []byte
		output string
		err    error
	}{
		{
			name:   "P2PKH",
			input:  script1,
			output: "RNoSGCX8SPFscj8epDaJjqEpuZa2B5in88",
		},
		{
			name:   "P2SH",
			input:  script2,
			output: "rPWwn5h4QFZNaz1XmY39rc73sdYGGDdmq1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := toStringMap[slip44.RAVENCOIN](tt.input)
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
