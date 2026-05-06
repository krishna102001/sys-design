package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash/fnv"
	"io"
)

const TotalShardDB uint32 = 5
const SECRET_KEY = "ILOVEYOU"

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

func GetAESShardIndex(shard_id string) (string, error) {
	block, _ := aes.NewCipher([]byte(SECRET_KEY))

	aesGCM, _ := cipher.NewGCM(block)

	nonce := make([]byte, aesGCM.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipherText := aesGCM.Seal(nonce, nonce, []byte(shard_id), nil)
	return hex.EncodeToString(cipherText), nil
}

func GetShardIdFromAES(aes_idx string) (string, error) {
	data, _ := hex.DecodeString(aes_idx)

	block, _ := aes.NewCipher([]byte(SECRET_KEY))

	aesGCM, _ := cipher.NewGCM(block)

	nonceSize := aesGCM.NonceSize()
	nonce, cipherText := data[:nonceSize], data[nonceSize:]

	shard_id, _ := aesGCM.Open(nil, nonce, cipherText, nil)

	return string(shard_id), nil
}

func main() {
	objectId := "65f1a2b3c4d5e6f7a8b9c0d4"
	shard := GetShardIndex(objectId)
	fmt.Printf("The ID %s belongs in Shard: %d\n", objectId, shard)
	secureShard := GetSecureShardIndex(objectId)
	fmt.Printf("The ID %s belongs in secure Shard: %s\n", objectId, secureShard)
	AesShard, _ := GetAESShardIndex(objectId)
	fmt.Printf("The ID %s belongs in aes Shard: %s\n", objectId, AesShard)
	shard_id, _ := GetShardIdFromAES(AesShard)
	fmt.Printf("The ID %s belongs in aes Shard decryptions: %s\n", objectId, shard_id)
}
