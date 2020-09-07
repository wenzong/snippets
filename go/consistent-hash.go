// Copy Paste From https://github.com/golang/groupcache/blob/master/consistenthash/consistenthash.go
// Modified to support customized replicaName function
package main

import (
	"crypto/md5"
	"fmt"
	"hash/crc32"
	"sort"
	"strconv"
)

type Hash func(data []byte) uint32
type ReplicaName func(index int, key string) []byte

type Map struct {
	hash        Hash
	replicaName ReplicaName
	replicas    int
	keys        []int // Sorted
	hashMap     map[int]string
}

func New(replicas int, hash Hash, replicaName ReplicaName) *Map {
	m := &Map{
		replicas:    replicas,
		hash:        hash,
		hashMap:     make(map[int]string),
		replicaName: replicaName,
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	if m.replicaName == nil {
		m.replicaName = func(index int, key string) []byte {
			return []byte(strconv.Itoa(index) + key)
		}

	}
	return m
}

// IsEmpty returns true if there are no items available.
func (m *Map) IsEmpty() bool {
	return len(m.keys) == 0
}

// Add adds some keys to the hash.
func (m *Map) Add(keys ...string) {
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash(m.replicaName(i, key)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

// Get gets the closest item in the hash to the provided key.
func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	// Binary search for appropriate replica.
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })

	// Means we have cycled back to the first replica.
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}

func main() {
	m := New(128, nil, func(index int, key string) []byte {
		array := md5.Sum([]byte(key + "-" + strconv.Itoa(index)))
		return array[:]
	})
	m.Add("redis://1", "redis://2")
	fmt.Println(m.Get("test"))
}
