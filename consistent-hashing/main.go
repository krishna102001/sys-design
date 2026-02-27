package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
	"sort"
)

type StorageNode struct {
	ID   string
	Host string
}

type ConsistentHash struct {
	nodes      []StorageNode
	keys       []*big.Int
	totalSlots *big.Int // hash space that can be 2^256
}

func NewConsistentHash() *ConsistentHash {
	slots := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	return &ConsistentHash{
		nodes:      []StorageNode{},
		keys:       []*big.Int{},
		totalSlots: slots,
	}
}

func (ch *ConsistentHash) hashFunc(key string) *big.Int {
	h := sha256.New()
	h.Write([]byte(key)) // its give length of keys in byte form
	// log.Printf("consistent hash %v", val)
	hashInt := new(big.Int).SetBytes(h.Sum(nil))
	// log.Printf("hash int %v", hashInt)
	hashKey := hashInt.Mod(hashInt, ch.totalSlots)
	// log.Printf("hash key is %v", hashKey)
	return hashKey
}

func (ch *ConsistentHash) AddServerNode(node StorageNode) {
	position := ch.hashFunc(node.Host)

	//sort.Search uses the binary search
	index := sort.Search(len(ch.keys), func(i int) bool {
		return ch.keys[i].Cmp(position) >= 0
	})

	if index < len(ch.keys) && ch.keys[index].Cmp(position) == 0 {
		log.Printf("Collision occurred at position %v", position)
		return
	}

	ch.keys = append(ch.keys, nil)
	copy(ch.keys[index+1:], ch.keys[index:])
	ch.keys[index] = position

	ch.nodes = append(ch.nodes, StorageNode{})
	copy(ch.nodes[index+1:], ch.nodes[index:])
	ch.nodes[index] = node

}

func (ch *ConsistentHash) AssignItem(item string) StorageNode {
	if len(ch.keys) == 0 {
		return StorageNode{}
	}
	position := ch.hashFunc(item)

	index := sort.Search(len(ch.keys), func(i int) bool {
		return ch.keys[i].Cmp(position) >= 0
	})

	if index == len(ch.keys) {
		index = 0
	}
	return ch.nodes[index]
}

func (ch *ConsistentHash) RemoveServerNode(node StorageNode) {
	if len(ch.keys) == 0 {
		return
	}

	position := ch.hashFunc(node.Host)

	index := sort.Search(len(ch.keys), func(i int) bool {
		return ch.keys[i].Cmp(position) >= 0
	})

	if index < len(ch.keys) && ch.keys[index].Cmp(position) == 0 {
		ch.keys = append(ch.keys[:index], ch.keys[index+1:]...)
		ch.nodes = append(ch.nodes[:index], ch.nodes[index+1:]...)
		fmt.Printf("remove the node %v\n", node)
	} else {
		fmt.Printf("no node found %v\n", node)
	}
}

func main() {
	ch := NewConsistentHash()

	// Adding nodes
	ch.AddServerNode(StorageNode{"A", "10.0.0.1"})
	ch.AddServerNode(StorageNode{"B", "10.0.0.2"})
	ch.AddServerNode(StorageNode{"C", "10.0.0.3"})
	ch.AddServerNode(StorageNode{"D", "10.0.0.4"})
	// Assigning items
	items := []string{"video_file_1.mp4", "image_77.png", "user_data_99"}

	for _, item := range items {
		node := ch.AssignItem(item)
		fmt.Printf("Item '%s' is assigned to Node %s (%s)\n", item, node.ID, node.Host)
	}
	ch.RemoveServerNode(StorageNode{"B", "10.0.0.2"})
	for _, item := range items {
		node := ch.AssignItem(item)
		fmt.Printf("Item '%s' is assigned to Node %s (%s)\n", item, node.ID, node.Host)
	}

}
