package coincodec

import (
	"testing"

	"github.com/pkg/errors"
)

func TestFilecoinToBytes(t *testing.T) {
	tests := []TestcaseEncode{
		// ID addresses
		{
			name:   "f00",
			input:  "f00",
			output: "0000",
		},
		{
			name:   "f01",
			input:  "f01",
			output: "0001",
		},
		{
			name:   "f010",
			input:  "f010",
			output: "000a",
		},
		{
			name:   "f0150",
			input:  "f0150",
			output: "009601",
		},
		{
			name:   "f0499",
			input:  "f0499",
			output: "00f303",
		},
		{
			name:   "f01024",
			input:  "f01024",
			output: "008008",
		},
		{
			name:   "f01729",
			input:  "f01729",
			output: "00c10d",
		},
		{
			name:   "f0999999",
			input:  "f0999999",
			output: "00bf843d",
		},
		{
			name:   "f018446744073709551615",
			input:  "f018446744073709551615",
			output: "00ffffffffffffffffff01",
		},
		// secp256k1 addresses
		{
			name:   "secp1",
			input:  "f15ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7pdrq",
			output: "01ea0f0ea039b291a0f08fd179e0556a8c3277c0d3",
		},
		{
			name:   "secp2",
			input:  "f12fiakbhe2gwd5cnmrenekasyn6v5tnaxaqizq6a",
			output: "01d1500504e4d1ac3e89ac891a4502586fabd9b417",
		},
		{
			name:   "secp3",
			input:  "f1wbxhu3ypkuo6eyp6hjx6davuelxaxrvwb2kuwva",
			output: "01b06e7a6f0f551de261fe3a6fe182b422ee0bc6b6",
		},
		{
			name:   "secp4",
			input:  "f1xtwapqc6nh4si2hcwpr3656iotzmlwumogqbuaa",
			output: "01bcec07c05e69f92468e2b3e3bf77c874f2c5da8c",
		},
		{
			name:   "secp5",
			input:  "f1xcbgdhkgkwht3hrrnui3jdopeejsoatkzmoltqy",
			output: "01b882619d46558f3d9e316d11b48dcf211327026a",
		},
		{
			name:   "secp6",
			input:  "f17uoq6tp427uzv7fztkbsnn64iwotfrristwpryy",
			output: "01fd1d0f4dfcd7e99afcb99a8326b7dc459d32c628",
		},
		// Actor addresses
		{
			name:   "act1",
			input:  "f24vg6ut43yw2h2jqydgbg2xq7x6f4kub3bg6as6i",
			output: "02e54dea4f9bc5b47d261819826d5e1fbf8bc5503b",
		},
		{
			name:   "act2",
			input:  "f25nml2cfbljvn4goqtclhifepvfnicv6g7mfmmvq",
			output: "02eb58bd08a15a6ade19d0989674148fa95a8157c6",
		},
		{
			name:   "act3",
			input:  "f2nuqrg7vuysaue2pistjjnt3fadsdzvyuatqtfei",
			output: "026d21137eb4c4814269e894d296cf6500e43cd714",
		},
		{
			name:   "act4",
			input:  "f24dd4ox4c2vpf5vk5wkadgyyn6qtuvgcpxxon64a",
			output: "02e0c7c75f82d55e5ed55db28033630df4274a984f",
		},
		{
			name:   "act5",
			input:  "f2gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr23y",
			output: "02316b4c1ff5d4afb7826ceab5bb0f2c3e0f364053",
		},
		// BLS addresses
		{
			name:   "bls1",
			input:  "f3vvmn62lofvhjd2ugzca6sof2j2ubwok6cj4xxbfzz4yuxfkgobpihhd2thlanmsh3w2ptld2gqkn2jvlss4a",
			output: "03ad58df696e2d4e91ea86c881e938ba4ea81b395e12797b84b9cf314b9546705e839c7a99d606b247ddb4f9ac7a3414dd",
		},
		{
			name:   "bls2",
			input:  "f3wmuu6crofhqmm3v4enos73okk2l366ck6yc4owxwbdtkmpk42ohkqxfitcpa57pjdcftql4tojda2poeruwa",
			output: "03b3294f0a2e29e0c66ebc235d2fedca5697bf784af605c75af608e6a63d5cd38ea85ca8989e0efde9188b382f9372460d",
		},
		{
			name:   "bls3",
			input:  "f3s2q2hzhkpiknjgmf4zq3ejab2rh62qbndueslmsdzervrhapxr7dftie4kpnpdiv2n6tvkr743ndhrsw6d3a",
			output: "0396a1a3e4ea7a14d49985e661b22401d44fed402d1d0925b243c923589c0fbc7e32cd04e29ed78d15d37d3aaa3fe6da33",
		},
		{
			name:   "bls4",
			input:  "f3q22fijmmlckhl56rn5nkyamkph3mcfu5ed6dheq53c244hfmnq2i7efdma3cj5voxenwiummf2ajlsbxc65a",
			output: "0386b454258c589475f7d16f5aac018a79f6c1169d20fc33921dd8b5ce1cac6c348f90a3603624f6aeb91b64518c2e8095",
		},
		{
			name:   "bls5",
			input:  "f3u5zgwa4ael3vuocgc5mfgygo4yuqocrntuuhcklf4xzg5tcaqwbyfabxetwtj4tsam3pbhnwghyhijr5mixa",
			output: "03a7726b038022f75a384617585360cee629070a2d9d28712965e5f26ecc40858382803724ed34f2720336f09db631f074",
		},
		// invalid
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Too short"),
		},
		{
			name:  "Negative",
			input: "f0-1",
			err:   errors.New("Invalid ID"),
		},
		{
			name:  "Too  large",
			input: "f018446744073709551616",
			err:   errors.New("Invalid ID"),
		},
		{
			name:  "Too  long",
			input: "f0111111111111111111111111111111111111111111111111111111111111",
			err:   errors.New("Invalid ID"),
		},
		{
			name:  "Invalid number",
			input: "f0C12",
			err:   errors.New("Invalid ID"),
		},
		{
			name:  "Embedded NUL",
			input: "f15ihq5ibzwki2b4ep2f46avlkr\000zhpqgtga7pdrq",
			err:   errors.New("decoding base32 failed"),
		},
		{
			name:  "Testnet",
			input: "t15ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7pdrq",
			err:   errors.New("Invalid network"),
		},
		{
			name:  "Unknown net",
			input: "a15ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7pdrq",
			err:   errors.New("Invalid network"),
		},
		{
			name:  "Unknown address type",
			input: "f95ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7pdrq",
			err:   errors.New("Invalid type"),
		},
		{
			name:  "Invalid checksum case1",
			input: "f15ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7rdrr",
			err:   errors.New("Wrong checksum"),
		},
		{
			name:  "Invalid checksum case2",
			input: "f24vg6ut43yw2h2jqydgbg2xq7x6f4kub3bg6as66",
			err:   errors.New("Wrong checksum"),
		},
		{
			name:  "Invalid checksum case3",
			input: "f3vvmn62lofvhjd2ugzca6sof2j2ubwok6cj4xxbfzz4yuxfkgobpihhd2thlanmsh3w2ptld2gqkn2jvlss44",
			err:   errors.New("Wrong checksum"),
		},
	}

	// 461 slip44.FILECOIN
	RunTestsEncode(t, 461, tests)
}

func TestFilecoinToString(t *testing.T) {
	tests := []TestcaseDecode{
		// ID addresses
		{
			name:   "f00",
			input:  "0000",
			output: "f00",
		},
		{
			name:   "f01",
			input:  "0001",
			output: "f01",
		},
		{
			name:   "f010",
			input:  "000a",
			output: "f010",
		},
		{
			name:   "f0150",
			input:  "009601",
			output: "f0150",
		},
		{
			name:   "f0499",
			input:  "00f303",
			output: "f0499",
		},
		{
			name:   "f01024",
			input:  "008008",
			output: "f01024",
		},
		{
			name:   "f01729",
			input:  "00c10d",
			output: "f01729",
		},
		{
			name:   "f0999999",
			input:  "00bf843d",
			output: "f0999999",
		},
		{
			name:   "f018446744073709551615",
			input:  "00ffffffffffffffffff01",
			output: "f018446744073709551615",
		},
		// secp256k1 addresses
		{
			name:   "secp1",
			input:  "01ea0f0ea039b291a0f08fd179e0556a8c3277c0d3",
			output: "f15ihq5ibzwki2b4ep2f46avlkrqzhpqgtga7pdrq",
		},
		{
			name:   "secp2",
			input:  "01d1500504e4d1ac3e89ac891a4502586fabd9b417",
			output: "f12fiakbhe2gwd5cnmrenekasyn6v5tnaxaqizq6a",
		},
		{
			name:   "secp3",
			input:  "01b06e7a6f0f551de261fe3a6fe182b422ee0bc6b6",
			output: "f1wbxhu3ypkuo6eyp6hjx6davuelxaxrvwb2kuwva",
		},
		{
			name:   "secp4",
			input:  "01bcec07c05e69f92468e2b3e3bf77c874f2c5da8c",
			output: "f1xtwapqc6nh4si2hcwpr3656iotzmlwumogqbuaa",
		},
		{
			name:   "secp5",
			input:  "01b882619d46558f3d9e316d11b48dcf211327026a",
			output: "f1xcbgdhkgkwht3hrrnui3jdopeejsoatkzmoltqy",
		},
		{
			name:   "secp6",
			input:  "01fd1d0f4dfcd7e99afcb99a8326b7dc459d32c628",
			output: "f17uoq6tp427uzv7fztkbsnn64iwotfrristwpryy",
		},
		// Actor addresses
		{
			name:   "act1",
			input:  "02e54dea4f9bc5b47d261819826d5e1fbf8bc5503b",
			output: "f24vg6ut43yw2h2jqydgbg2xq7x6f4kub3bg6as6i",
		},
		{
			name:   "act2",
			input:  "02eb58bd08a15a6ade19d0989674148fa95a8157c6",
			output: "f25nml2cfbljvn4goqtclhifepvfnicv6g7mfmmvq",
		},
		{
			name:   "act3",
			input:  "026d21137eb4c4814269e894d296cf6500e43cd714",
			output: "f2nuqrg7vuysaue2pistjjnt3fadsdzvyuatqtfei",
		},
		{
			name:   "act4",
			input:  "02e0c7c75f82d55e5ed55db28033630df4274a984f",
			output: "f24dd4ox4c2vpf5vk5wkadgyyn6qtuvgcpxxon64a",
		},
		{
			name:   "act5",
			input:  "02316b4c1ff5d4afb7826ceab5bb0f2c3e0f364053",
			output: "f2gfvuyh7v2sx3patm5k23wdzmhyhtmqctasbr23y",
		},
		// BLS addresses
		{
			name:   "bls1",
			input:  "03ad58df696e2d4e91ea86c881e938ba4ea81b395e12797b84b9cf314b9546705e839c7a99d606b247ddb4f9ac7a3414dd",
			output: "f3vvmn62lofvhjd2ugzca6sof2j2ubwok6cj4xxbfzz4yuxfkgobpihhd2thlanmsh3w2ptld2gqkn2jvlss4a",
		},
		{
			name:   "bls2",
			input:  "03b3294f0a2e29e0c66ebc235d2fedca5697bf784af605c75af608e6a63d5cd38ea85ca8989e0efde9188b382f9372460d",
			output: "f3wmuu6crofhqmm3v4enos73okk2l366ck6yc4owxwbdtkmpk42ohkqxfitcpa57pjdcftql4tojda2poeruwa",
		},
		{
			name:   "bls3",
			input:  "0396a1a3e4ea7a14d49985e661b22401d44fed402d1d0925b243c923589c0fbc7e32cd04e29ed78d15d37d3aaa3fe6da33",
			output: "f3s2q2hzhkpiknjgmf4zq3ejab2rh62qbndueslmsdzervrhapxr7dftie4kpnpdiv2n6tvkr743ndhrsw6d3a",
		},
		{
			name:   "bls4",
			input:  "0386b454258c589475f7d16f5aac018a79f6c1169d20fc33921dd8b5ce1cac6c348f90a3603624f6aeb91b64518c2e8095",
			output: "f3q22fijmmlckhl56rn5nkyamkph3mcfu5ed6dheq53c244hfmnq2i7efdma3cj5voxenwiummf2ajlsbxc65a",
		},
		{
			name:   "bls5",
			input:  "03a7726b038022f75a384617585360cee629070a2d9d28712965e5f26ecc40858382803724ed34f2720336f09db631f074",
			output: "f3u5zgwa4ael3vuocgc5mfgygo4yuqocrntuuhcklf4xzg5tcaqwbyfabxetwtj4tsam3pbhnwghyhijr5mixa",
		},
		// Invalid
		{
			name:  "Empty",
			input: "",
			err:   errors.New("Data too short"),
		},
		{
			name:  "Short",
			input: "00",
			err:   errors.New("Data too short"),
		},
		{
			name:  "Invalid type",
			input: "c200",
			err:   errors.New("Invalid type"),
		},
		{
			name:  "Secp, too long",
			input: "01ea0f0ea039b291a0f08fd179e0556a8c3277c0d30303",
			err:   errors.New("Invalid length"),
		},
	}

	// 461 slip44.FILECOIN
	RunTestsDecode(t, 461, tests)
}
