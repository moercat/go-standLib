package consistent

import (
	"hash/crc32"
	"hash/fnv"
	"sort"
	"strconv"
)

// ConsistentStr holds the information about the members of the consistent hash circle.
type ConsistentStr struct {
	circle           map[uint32]string
	strMsg           []StrMsg
	sortedHashes     uints
	NumberOfReplicas int
	UseFnv           bool
	scratch          [64]byte
}

type StrMsg struct {
	Key    string
	Weight int
}

func NewStr() *ConsistentStr {
	c := new(ConsistentStr)
	c.NumberOfReplicas = 200
	c.circle = make(map[uint32]string)
	c.UseFnv = true
	return c
}

// eltKey generates a string key for an element with an index.
func (c *ConsistentStr) eltKey(elt string, idx int, weigth int) string {
	return strconv.Itoa(idx*10000+weigth) + elt
}

func (c *ConsistentStr) SetNumberOfReplicas(num int) {
	if num < 1 {
		num = 1
	}
	c.NumberOfReplicas = num
}

// Add inserts a string element in the consistent hash.
func (c *ConsistentStr) Add(elt string, weight int) {
	c.strMsg = append(c.strMsg, StrMsg{
		Key:    elt,
		Weight: weight,
	})
}

func (c *ConsistentStr) Finally() {
	c.add(c.strMsg)
	c.strMsg = make([]StrMsg, 0)
}

func (c *ConsistentStr) add(nodes []StrMsg) {
	for _, node := range nodes {
		for i := 0; i < c.NumberOfReplicas; i++ {
			for j := 0; j < node.Weight; j++ {
				c.circle[c.hashKey(c.eltKey(node.Key, i, j))] = node.Key
			}
		}
	}
	c.updateSortedHashes()
}

// Get returns an element close to where name hashes to in the circle.
func (c *ConsistentStr) Get(name string) (string, error) {
	if len(c.circle) == 0 {
		return "", ErrEmptyCircle
	}
	key := c.hashKey(name)
	i := c.search(key)
	return c.circle[c.sortedHashes[i]], nil
}

func (c *ConsistentStr) search(key uint32) (i int) {
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	i = sort.Search(len(c.sortedHashes), f)
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return
}

func (c *ConsistentStr) hashKey(key string) uint32 {
	if c.UseFnv {
		return c.hashKeyFnv(key)
	}
	return c.hashKeyCRC32(key)
}

func (c *ConsistentStr) hashKeyCRC32(key string) uint32 {
	if len(key) < 64 {
		var scratch [64]byte
		copy(scratch[:], key)
		return crc32.ChecksumIEEE(scratch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

func (c *ConsistentStr) hashKeyFnv(key string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(key))
	return h.Sum32()
}

func (c *ConsistentStr) updateSortedHashes() {
	hashes := c.sortedHashes[:0]
	//reallocate if we're holding on to too much (1/4th)
	if cap(c.sortedHashes)/(c.NumberOfReplicas*4) > len(c.circle) {
		hashes = nil
	}
	for k := range c.circle {
		hashes = append(hashes, k)
	}
	sort.Sort(hashes)
	c.sortedHashes = hashes
}
