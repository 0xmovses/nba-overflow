import TopShot from 0xf8d6e0586b0a20c7 
// This script reads the next Set ID from the TopShot contract and 
// returns that number to the caller

// Returns: UInt32
// Value of nextSetID field in TopShot contract

pub fun main(): UInt32 {

    log(TopShot.nextSetID)

    return TopShot.nextSetID
}
