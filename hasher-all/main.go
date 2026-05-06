package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/fnv"
)

const TotalShardDB uint32 = 5

// we have used here fnv(Fowler-Noll-Vo) non-crypto graphic hash object because of fast hashing which will be easier to scatter the data
// we don't need here secure hash here so we avoid the "SHA-256"
func GetShardIndex(shard_id string) int {
	hasher := fnv.New32a()
	hasher.Write([]byte(shard_id))

	hashValue := hasher.Sum32()

	return int(hashValue % TotalShardDB)
}

func GetSecureShardIndex(shard_id string) string {
	hasher := sha256.New()

	hasher.Write([]byte(shard_id))

	hashValue := hex.EncodeToString(hasher.Sum([]byte("krishna"))) // here the sum byte "krishna" will be added to the last of shard_id
	return hashValue
}

func main() {
	objectId := "65f1a2b3c4d5e6f7a8b9c0d4"
	shard := GetShardIndex(objectId)
	fmt.Printf("The ID %s belongs in Shard: %d\n", objectId, shard)
	secureShard := GetSecureShardIndex(objectId)
	fmt.Printf("The ID %s belongs in secure Shard: %s\n", objectId, secureShard)
}
