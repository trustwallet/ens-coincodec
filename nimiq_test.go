package coincodec

import (
	"bytes"
	"encoding/hex"
	"math/rand"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func TestNimiqEncodeToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		{
			name:   "Normal",
			input:  "NQ19 46LK 9YHV D9LB TDJ4 8Y2P 3J4C 37HR YDL5",
			output: "21a934fe3d6a68bdb64447c571c88c19e39fb685",
		},
		{
			name:   "No Spaces",
			input:  "NQ1946LK9YHVD9LBTDJ48Y2P3J4C37HRYDL5",
			output: "21a934fe3d6a68bdb64447c571c88c19e39fb685",
		},
		{
			name:  "Prefix",
			input: "LD20 46LK 9YHV D9LB TDJ4 8Y2P 3J4C 37HR YDL5",
			err:   errors.New("not a NQ address"),
		},
		{
			name:  "Checksum",
			input: "NQ20 46LK 9YHV D9LB TDJ4 8Y2P 3J4C 37HR YDL5",
			err:   errors.New("wrong checksum"),
		},
		{
			name:  "Long",
			input: "NQ19 46LK 9YHV D9LB TDJ4 8Y2P 3J4C 37HR YDL5 A",
			err:   errors.New("invalid length"),
		},
		{
			name:  "Unicode",
			input: "NQ19 😳 9YHV D9LB TDJ4 8Y2P 3J4C 37HR YDL5",
			err:   errors.New("unexpected character"),
		},
		{
			name:  "Invalid Base32",
			input: "NQ85 I999 9YHV D9LB TDJ4 8Y2P 3J4C 37HR YDL5",
			err:   errors.New("decoding base32 failed: illegal base32 data at input byte 0"),
		},
	}

	RunTestsEncode(t, slip44.NIMIQ, tests)

	// Since the runtime of NimiqEncodeToString is
	// dependent on the input, we should fuzz it.
	t.Run("Random", func(t *testing.T) {
		rand.Seed(time.Now().UnixNano())
		const iterations = 1024
		for i := 0; i < iterations; i++ {
			// Generate random address
			var raw [20]byte
			_, _ = rand.Read(raw[:])
			// Encode
			ufa, err := NimiqEncodeToString(raw[:])
			if err != nil {
				t.Fatalf("EncodeToString(%s).err => %s",
					hex.EncodeToString(raw[:]), err)
			}
			// Decode
			raw2, err := NimiqDecodeToBytes(ufa)
			if err != nil {
				t.Fatalf("DecodeToBytes(%s).err => %s", ufa, err)
			}
			if !bytes.Equal(raw[:], raw2) {
				t.Fatalf("Encoding of %s corrupt",
					hex.EncodeToString(raw[:]))
			}
		}
	})
}

func TestNimiqDecodeToString(t *testing.T) {
	tests := []TestcaseDecode{
		{
			name:   "Good",
			input:  "21a934fe3d6a68bdb64447c571c88c19e39fb685",
			output: "NQ19 46LK 9YHV D9LB TDJ4 8Y2P 3J4C 37HR YDL5",
		},
		{
			name:  "Length",
			input: "21a934fe3d6a68bdb64447c571c88c19e39fb68599",
			err:   errors.New("invalid length"),
		},
	}

	RunTestsDecode(t, slip44.NIMIQ, tests)
}
