module score_addr::PrimitiveTypes {
    use std::string::{Self, String};

    #[view]
    public fun foo(_a: u8, _b: u16, _c: u32, _d: u64, _e: u128, _f: u256, _g: bool,
        _h: address, _i: String, _j: vector<u8>
    ): (u8, u16, u32, u64, u128, u256, bool, address, String, vector<u8>) {
        (
            12,
            12345,
            12345678,
            1234567890123,
            12345678901234567890,
            123456789012345678901234567890,
            true,
            @0x1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef,
            string::utf8(b"Hello, Aptos!"),
            b"byte_data"
        )
    }
}