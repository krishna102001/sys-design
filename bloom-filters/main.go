package main

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	box []int64
}

func NewBloomFilter() *BloomFilter {
	return &BloomFilter{
		box: make([]int64, 100000),
	}
}

func (bf *BloomFilter) Set(username string) {
	idx1, idx2, idx3 := bf.HashFunc1(username), bf.HashFunc2(username), bf.HashFunc3(username)
	bf.box[idx1], bf.box[idx2], bf.box[idx3] = 1, 1, 1
}

func (bf *BloomFilter) CheckUserNameExist(username string) bool {
	idx1, idx2, idx3 := bf.HashFunc1(username), bf.HashFunc2(username), bf.HashFunc3(username)
	val1, val2, val3 := bf.box[idx1], bf.box[idx2], bf.box[idx3]
	if val1 == 1 && val2 == 1 && val3 == 1 {
		return true
	}
	return false
}

func (bf *BloomFilter) HashFunc1(username string) int64 {
	hash := fnv.New64a()
	hash.Write([]byte(username))
	hash.Write([]byte("6388"))
	return int64(hash.Sum64() % 100000)
}

func (bf *BloomFilter) HashFunc2(username string) int64 {
	hash := fnv.New64a()
	hash.Write([]byte(username))
	hash.Write([]byte("51459"))
	return int64(hash.Sum64() % 100000)
}

func (bf *BloomFilter) HashFunc3(username string) int64 {
	hash := fnv.New64a()
	hash.Write([]byte(username))
	hash.Write([]byte("94152"))
	return int64(hash.Sum64() % 100000)
}

func main() {
	newBF := NewBloomFilter()
	newBF.Set("krishna")
	fmt.Println(" does user name exist ", newBF.CheckUserNameExist("krishna"))
}
