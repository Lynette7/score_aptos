// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract FooDecoder {
    function decode(bytes calldata data)
        public
        pure
        returns (
            uint8,
            uint16,
            uint32,
            uint64,
            uint128,
            uint256,
            bool,
            address,
            string memory,
            bytes memory
        )
    {
        (
            uint8 _a,
            uint16 _b,
            uint32 _c,
            uint64 _d,
            uint128 _e,
            uint256 _f,
            bool _g,
            address _h,
            string memory _i,
            bytes memory _j
        ) = abi.decode(
            data,
            (uint8, uint16, uint32, uint64, uint128, uint256, bool, address, string, bytes)
        );

        return (_a, _b, _c, _d, _e, _f, _g, _h, _i, _j);
    }
}