package coincodec

import (
	"bytes"
	"encoding/base32"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/wealdtech/go-slip44"
)

func init() {
	toBytesMap[slip44.NIMIQ] = NimiqDecodeToBytes
	toStringMap[slip44.NIMIQ] = NimiqEncodeToString
}

var base32NimiqEncoding = base32.NewEncoding("0123456789ABCDEFGHJKLMNPQRSTUVXY")

// NimiqDecodeToBytes converts the "user-friendly Nimiq address"
// to its raw byte representation.
// The input must be uppercase, spaces are ignored.
func NimiqDecodeToBytes(input string) ([]byte, error) {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) != 36 {
		return nil, errors.New("invalid length")
	}
	if input[:2] != "NQ" {
		return nil, errors.New("not a NQ address")
	}
	check, err := nimiqChecksum(input)
	if err != nil {
		return nil, err
	}
	if check != 1 {
		println(check)
		return nil, errors.New("wrong checksum")
	}
	var raw [20]byte
	_, err = base32NimiqEncoding.Decode(raw[:], []byte(input[4:]))
	if err != nil {
		return nil, errors.Wrap(err, "decoding base32 failed")
	}
	return raw[:], nil
}

// NimiqEncodeToString a raw address to its user-friendly representation.
// If len(buf) != 20, encoding fails.
func NimiqEncodeToString(buf []byte) (string, error) {
	if len(buf) != 20 {
		return "", errors.New("invalid length")
	}
	// Encode the user-friendly address without spaces
	// and with a dummy checksum.
	noSpaces := "NQ00" + base32NimiqEncoding.EncodeToString(buf)
	check, _ := nimiqChecksum(noSpaces)
	check = 98 - check
	// Insert correct checksum and put space after every 4th char.
	var result strings.Builder
	result.WriteString("NQ")
	result.Write([]byte{
		'0' + (check%100)/10, // Digit X of XY
		'0' + check%10,       // Digit Y of XY
	})
	for i := 4; i < 36; i += 4 {
		// Insert a space every 4 characters
		result.WriteByte(' ')
		result.WriteString(noSpaces[i : i+4])
	}
	return result.String(), nil
}

// nimiqChecksum calculates a checksum similar to the IBAN modulo 97 method.
// input must be a string with len(s) = 36.
func nimiqChecksum(input string) (uint8, error) {
	bigNumber, err := nimiqToNumber(input[4:] + input[:4])
	if err != nil {
		return 0, err
	}

	// ceil(len(bigNumber) / 6)
	blockCount := (len(bigNumber) + 5) / 6

	var chunk bytes.Buffer
	var check int

	for i := 0; i < blockCount; i++ {
		offset := i * 6
		var stop int
		if len(bigNumber) <= offset+6 {
			stop = len(bigNumber)
		} else {
			stop = offset + 6
		}

		block := bigNumber[offset:stop]

		// Compress using modulo
		chunk.WriteString(block)
		tmp := chunk.String()
		check, _ = strconv.Atoi(tmp)
		check %= 97
		chunk.Reset()
		chunk.WriteString(strconv.Itoa(check))
	}

	return uint8(check), nil
}

// nimiqToNumber converts an alphanumerical string to a numerical string.
// Required to calculate the Nimiq checksum.
// Example: "6789ABcd" => "678910111213".
func nimiqToNumber(input string) (string, error) {
	var buf strings.Builder
	for _, r := range input {
		switch {
		case r >= '0' && r <= '9':
			_ = buf.WriteByte(byte(r))
		case r >= 'A' && r <= 'Z':
			num := 10 + (r - 'A')
			_, _ = buf.WriteString(strconv.Itoa(int(num)))
		default:
			return "", errors.New("unexpected character")
		}
	}
	return buf.String(), nil
}
