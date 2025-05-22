package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"strconv"
)


type FooResult struct {
	U8       uint8
	U16      uint16
	U32      uint32
	U64      uint64
	U128     *big.Int
	U256     *big.Int
	Bool     bool
	Address  [20]byte
	String   string
	VectorU8 []byte
}

func main() {
	accountAddress := "0xc8fbacb88102686835801c46eb5bc15be4308de80f9fc58a4103bfb26ed10871"

	functionID := fmt.Sprintf("%s::PrimitiveTypes::foo", accountAddress)

	url := "https://api.testnet.aptoslabs.com/v1/view"

	pU8 := uint8(42)
	pU16 := uint16(12345)
	pU32 := uint32(12345678)
	pU64 := "1234567890123"
	pU128 := "12345678901234567890"
	pU256 := "123456789012345678901234567890"
	pBool := true
	pAddress := "0xc8fbacb88102686835801c46eb5bc15be4308de80f9fc58a4103bfb26ed10871"
	pString := "Hello, World!"
	pVector := hex.EncodeToString([]byte("hello"))

	payload := []byte(fmt.Sprintf(`{
		"function": "%s",
		"type_arguments": [],
		"arguments": [
			%d,
			%d,
			%d,
			"%s",
			"%s",
			"%s",
			%t,
			"%s",
			"%s",
			"0x%s"
		]
	}`, functionID, pU8, pU16, pU32, pU64, pU128, pU256, pBool, pAddress, pString, pVector))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Printf("Error making HTTP request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}
	fmt.Printf("Raw response: %s\n", body)

	var result []interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("Error parsing JSON response: %v\n", err)
		return
	}

	if len(result) != 10 {
		fmt.Printf("Expected 10 results, got %d\n", len(result))
		return
	}

	var fooResult FooResult

	u8, err := parseNumber(result[0], 8)
	if err != nil {
		fmt.Printf("Error parsing u8: %v\n", err)
		return
	}
	fooResult.U8 = uint8(u8)

	u16, err := parseNumber(result[1], 16)
	if err != nil {
		fmt.Printf("Error parsing u16: %v\n", err)
		return
	}
	fooResult.U16 = uint16(u16)

	u32, err := parseNumber(result[2], 32)
	if err != nil {
		fmt.Printf("Error parsing u32: %v\n", err)
		return
	}
	fooResult.U32 = uint32(u32)

	u64, err := parseStringNumber(result[3], 64)
	if err != nil {
		fmt.Printf("Error parsing u64: %v\n", err)
		return
	}
	fooResult.U64 = uint64(u64)

	var ok bool
	fooResult.U128, ok = new(big.Int).SetString(result[4].(string), 10)
	if !ok {
		fmt.Printf("Error parsing u128: %s\n", result[4])
		return
	}

	fooResult.U256, ok = new(big.Int).SetString(result[5].(string), 10)
	if !ok {
		fmt.Printf("Error parsing u256: %s\n", result[5])
		return
	}

	boolVal, ok := result[6].(bool)
	if !ok {
		fmt.Printf("Error parsing bool: %v\n", result[6])
		return
	}
	fooResult.Bool = boolVal

	addrStr, ok := result[7].(string)
	if !ok {
		fmt.Printf("Error parsing address: %v\n", result[7])
		return
	}
	addrBytes, err := hex.DecodeString(addrStr[2:])
	if err != nil {
		fmt.Printf("Error decoding address hex: %v\n", err)
		return
	}
	if len(addrBytes) != 32 {
		fmt.Printf("Expected 32-byte address, got %d bytes\n", len(addrBytes))
		return
	}
	var addr [20]byte
	copy(addr[:], addrBytes[12:32])
	fooResult.Address = addr

	str, ok := result[8].(string)
	if !ok {
		fmt.Printf("Error parsing string: %v\n", result[8])
		return
	}
	fooResult.String = str

	vec, ok := result[9].(string)
	if !ok {
		fmt.Printf("Error parsing vector<u8>: %v\n", result[9])
		return
	}
	fooResult.VectorU8, err = hex.DecodeString(vec[2:])
	if err != nil {
		fmt.Printf("Error decoding vector<u8>: %v\n", err)
		return
	}

	fmt.Printf("Results from foo:\n")
	fmt.Printf("u8: %d\n", fooResult.U8)
	fmt.Printf("u16: %d\n", fooResult.U16)
	fmt.Printf("u32: %d\n", fooResult.U32)
	fmt.Printf("u64: %d\n", fooResult.U64)
	fmt.Printf("u128: %s\n", fooResult.U128.String())
	fmt.Printf("u256: %s\n", fooResult.U256.String())
	fmt.Printf("bool: %v\n", fooResult.Bool)
	fmt.Printf("address: 0x%x\n", fooResult.Address)
	fmt.Printf("string: %s\n", fooResult.String)
	fmt.Printf("vector<u8): %x\n", fooResult.VectorU8)
}

func parseNumber(val interface{}, bitSize int) (uint64, error) {
	f, ok := val.(float64)
	if !ok {
		return 0, fmt.Errorf("expected number, got %T", val)
	}
	if f < 0 || f != float64(uint64(f)) {
		return 0, fmt.Errorf("value %v is not a valid unsigned integer", f)
	}
	return uint64(f), nil
}

func parseStringNumber(val interface{}, bitSize int) (uint64, error) {
	s, ok := val.(string)
	if !ok {
		return 0, fmt.Errorf("expected string, got %T", val)
	}
	return strconv.ParseUint(s, 10, bitSize)
}
