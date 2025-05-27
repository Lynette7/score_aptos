package main

import (
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/aptos-labs/aptos-go-sdk"
	"github.com/aptos-labs/aptos-go-sdk/bcs"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
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
	client, err := aptos.NewClient(aptos.TestnetConfig)
	if err != nil {
		fmt.Printf("Failed to create client: %v\n", err)
		return
	}

	accountAddress := "0xc8fbacb88102686835801c46eb5bc15be4308de80f9fc58a4103bfb26ed10871"

	pU8 := uint8(42)
	pU16 := uint16(12345)
	pU32 := uint32(12345678)
	pU64 := uint64(1234567890123)
	pU128 := big.NewInt(0)
	pU128.SetString("12345678901234567890", 10)
	pU256 := big.NewInt(0)
	pU256.SetString("123456789012345678901234567890", 10)
	pBool := true
	pAddress := "0xc8fbacb88102686835801c46eb5bc15be4308de80f9fc58a4103bfb26ed10871"
	pString := "Hello, World!"
	pVector := []byte("hello")

	args := [][]byte{}

	u8Bytes, err := bcs.SerializeU8(pU8)
	if err != nil {
		fmt.Printf("Failed to serialize u8: %v\n", err)
		return
	}
	args = append(args, u8Bytes)

	u16Bytes, err := bcs.SerializeU16(pU16)
	if err != nil {
		fmt.Printf("Failed to serialize u16: %v\n", err)
		return
	}
	args = append(args, u16Bytes)

	u32Bytes, err := bcs.SerializeU32(pU32)
	if err != nil {
		fmt.Printf("Failed to serialize u32: %v\n", err)
		return
	}
	args = append(args, u32Bytes)

	u64Bytes, err := bcs.SerializeU64(pU64)
	if err != nil {
		fmt.Printf("Failed to serialize u64: %v\n", err)
		return
	}
	args = append(args, u64Bytes)

	u128Bytes, err := bcs.SerializeU128(*pU128)
	if err != nil {
		fmt.Printf("Failed to serialize u128: %v\n", err)
		return
	}
	args = append(args, u128Bytes)

	u256Bytes, err := bcs.SerializeU256(*pU256)
	if err != nil {
		fmt.Printf("Failed to serialize u256: %v\n", err)
		return
	}
	args = append(args, u256Bytes)

	boolBytes, err := bcs.SerializeBool(pBool)
	if err != nil {
		fmt.Printf("Failed to serialize bool: %v\n", err)
		return
	}
	args = append(args, boolBytes)

	addrBytes, err := hex.DecodeString(pAddress[2:])
	if err != nil {
		fmt.Printf("Failed to decode address hex: %v\n", err)
		return
	}
	if len(addrBytes) != 32 {
		fmt.Printf("Expected 32-byte address, got %d bytes\n", len(addrBytes))
		return
	}
	var addr aptos.AccountAddress
	copy(addr[:], addrBytes)
	addrArgBytes, err := bcs.SerializeSingle(func(ser *bcs.Serializer) {
		ser.FixedBytes(addr[:])
	})
	if err != nil {
		fmt.Printf("Failed to serialize address: %v\n", err)
		return
	}
	args = append(args, addrArgBytes)

	stringBytes, err := bcs.SerializeSingle(func(ser *bcs.Serializer) {
		ser.WriteString(pString)
	})
	if err != nil {
		fmt.Printf("Failed to serialize string: %v\n", err)
		return
	}
	args = append(args, stringBytes)

	vectorBytes, err := bcs.SerializeBytes(pVector)
	if err != nil {
		fmt.Printf("Failed to serialize vector<u8>: %v\n", err)
		return
	}
	args = append(args, vectorBytes)

	moduleAddrBytes, err := hex.DecodeString(accountAddress[2:])
	if err != nil {
		fmt.Printf("Failed to decode module address hex: %v\n", err)
		return
	}
	if len(moduleAddrBytes) != 32 {
		fmt.Printf("Expected 32-byte module address, got %d bytes\n", len(moduleAddrBytes))
		return
	}
	var moduleAddr aptos.AccountAddress
	copy(moduleAddr[:], moduleAddrBytes)
	moduleId := aptos.ModuleId{
		Address: moduleAddr,
		Name:    "PrimitiveTypes",
	}

	viewPayload := aptos.ViewPayload{
		Module:   moduleId,
		Function: "foo",
		ArgTypes: []aptos.TypeTag{},
		Args:     args,
	}

	result, err := client.View(&viewPayload) //[]any
	if err != nil {
		fmt.Printf("Failed to call view function: %v\n", err)
		return
	}
	fmt.Printf("Raw result: %v\n", result)

	if len(result) != 10 {
		fmt.Printf("Expected 10 results, got %d\n", len(result))
		return
	}

	var fooResult FooResult

	f8, ok := result[0].(float64)
	if !ok {
		fmt.Printf("Error converting u8: expected float64, got %T\n", result[0])
		return
	}
	u8Int, err := float64ToInt(f8, 8)
	if err != nil {
		fmt.Printf("Error converting u8: %v\n", err)
		return
	}
	u8, err := aptos.ConvertToU8(u8Int)
	if err != nil {
		fmt.Printf("Error converting u8: %v\n", err)
		return
	}
	fooResult.U8 = u8

	f16, ok := result[1].(float64)
	if !ok {
		fmt.Printf("Error converting u16: expected float64, got %T\n", result[1])
		return
	}
	u16Int, err := float64ToInt(f16, 16)
	if err != nil {
		fmt.Printf("Error converting u16: %v\n", err)
		return
	}
	u16, err := aptos.ConvertToU16(u16Int)
	if err != nil {
		fmt.Printf("Error converting u16: %v\n", err)
		return
	}
	fooResult.U16 = u16

	f32, ok := result[2].(float64)
	if !ok {
		fmt.Printf("Error converting u32: expected float64, got %T\n", result[2])
		return
	}
	u32Int, err := float64ToInt(f32, 32)
	if err != nil {
		fmt.Printf("Error converting u32: %v\n", err)
		return
	}
	u32, err := aptos.ConvertToU32(u32Int)
	if err != nil {
		fmt.Printf("Error converting u32: %v\n", err)
		return
	}
	fooResult.U32 = u32

	u64, err := aptos.ConvertToU64(result[3])
	if err != nil {
		fmt.Printf("Error converting u64: %v\n", err)
		return
	}
	fooResult.U64 = u64

	u128, err := aptos.ConvertToU128(result[4])
	if err != nil {
		fmt.Printf("Error converting u128: %v\n", err)
		return
	}
	fooResult.U128 = u128

	u256, err := aptos.ConvertToU256(result[5])
	if err != nil {
		fmt.Printf("Error converting u256: %v\n", err)
		return
	}
	fooResult.U256 = u256

	boolVal, err := aptos.ConvertToBool(result[6])
	if err != nil {
		fmt.Printf("Error converting bool: %v\n", err)
		return
	}
	fooResult.Bool = boolVal

	addrPtr, err := aptos.ConvertToAddress(result[7])
	if err != nil {
		fmt.Printf("Error converting address: %v\n", err)
		return
	}
	if addrPtr == nil {
		fmt.Printf("Error: converted address is nil\n")
		return
	}
	addr = *addrPtr
	var evmAddr [20]byte
	copy(evmAddr[:], addr[12:32])
	fooResult.Address = evmAddr

	str, ok := result[8].(string)
	if !ok {
		fmt.Printf("Error converting string: expected string, got %T\n", result[8])
		return
	}
	fooResult.String = str

	vecStr, ok := result[9].(string)
	if !ok {
		fmt.Printf("Error converting vector<u8>: expected string, got %T\n", result[9])
		return
	}
	vecBytes, err := aptos.ParseHex(vecStr)
	if err != nil {
		fmt.Printf("Error decoding vector<u8> hex: %v\n", err)
		return
	}
	fooResult.VectorU8 = vecBytes

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
	fmt.Printf("vector<u8>): %x\n", fooResult.VectorU8)

	argumentMarshalings := []abi.ArgumentMarshaling{
		{Name: "u8", Type: "uint8"},
		{Name: "u16", Type: "uint16"},
		{Name: "u32", Type: "uint32"},
		{Name: "u64", Type: "uint64"},
		{Name: "u128", Type: "uint128"},
		{Name: "u256", Type: "uint256"},
		{Name: "bool", Type: "bool"},
		{Name: "address", Type: "address"},
		{Name: "str", Type: "string"},
		{Name: "vec", Type: "bytes"},
	}

	abiType, err := abi.NewType("tuple", "", argumentMarshalings)
	if err != nil {
		fmt.Printf("Failed to create ABI type: %v\n", err)
		return
	}

	argsABI := abi.Arguments{}
	for i, argMarshaling := range argumentMarshalings {
		argsABI = append(argsABI, abi.Argument{
			Name: argMarshaling.Name,
			Type: *abiType.TupleElems[i],
		})
	}

	encoded, err := argsABI.Pack(
		fooResult.U8,
		fooResult.U16,
		fooResult.U32,
		fooResult.U64,
		fooResult.U128,
		fooResult.U256,
		fooResult.Bool,
		common.BytesToAddress(fooResult.Address[:]),
		fooResult.String,
		fooResult.VectorU8,
	)
	if err != nil {
		fmt.Printf("Failed to ABI encode: %v\n", err)
		return
	}
	fmt.Printf("ABI encoded result: 0x%x\n", encoded)
}

func float64ToInt(f float64, bitSize int) (int, error) {
    if f < 0 {
        return 0, fmt.Errorf("value %v is negative, expected unsigned integer", f)
    }
    if f != float64(int64(f)) {
        return 0, fmt.Errorf("value %v is not an integer", f)
    }
    i := int64(f)
    var max int64
    switch bitSize {
    case 8:
        max = 255
    case 16:
        max = 65535
    case 32:
        max = 4294967295
    default:
        return 0, fmt.Errorf("invalid bit size %d", bitSize)
    }
    if i > max {
        return 0, fmt.Errorf("value %v exceeds max for %d-bit integer", f, bitSize)
    }
    return int(i), nil
}
