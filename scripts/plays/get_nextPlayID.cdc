import TopShot from 0xf8d6e0586b0a20c7

// This script reads the public nextPlayID from the TopShot contract and 
// returns that number to the caller

// Returns: UInt32
// the nextPlayID field in TopShot contract

pub fun main(): UInt32 {

    log(TopShot.nextPlayID)

    return TopShot.nextPlayID
}
