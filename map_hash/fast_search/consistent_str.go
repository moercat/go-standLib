package fast_search

import (
	"errors"
	"hash/crc32"
	"hash/fnv"
	"sort"
	"strconv"
)

type uints []uint32

// Len returns the length of the uints array.
func (x uints) Len() int { return len(x) }

// Less returns true if element i is less than element j.
func (x uints) Less(i, j int) bool { return x[i] < x[j] }

// Swap exchanges elements i and j.
func (x uints) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// ErrEmptyCircle is the error returned when trying to get an element when nothing has been added to hash.
var ErrEmptyCircle = errors.New("empty circle")

// ConsistentStr holds the information about the members of the consistent hash circle.
type ConsistentStr struct {
	oneElt           string
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
	c.circle = make(map[uint32]string, 10000)
	c.strMsg = make([]StrMsg, 0, 100)
	c.sortedHashes = make(uints, 0, 10000)
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

// AddAll inserts a string element in the consistent hash.
func (c *ConsistentStr) AddAll(args []StrMsg) {
	c.strMsg = args
	c.Finally()
}

func (c *ConsistentStr) Finally() {
	c.add(c.strMsg)
	c.strMsg = make([]StrMsg, 0)
}

func (c *ConsistentStr) add(nodes []StrMsg) {
	if len(nodes) == 1 {
		c.oneElt = nodes[0].Key
		return
	}
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
func (c *ConsistentStr) Get(name string) (string, int, error) {
	if len(c.oneElt) > 0 {
		return c.oneElt, 0, nil
	}
	if len(c.circle) == 0 {
		return "", 0, ErrEmptyCircle
	}

	i := c.search(c.hashKey(name))
	return c.circle[c.sortedHashes[i]], i, nil
}

func (c *ConsistentStr) GetWithIdx(name string, idx int) (string, int, error) {
	if len(c.oneElt) > 0 {
		return c.oneElt, 0, nil
	}
	if len(c.circle) == 0 {
		return "", 0, ErrEmptyCircle
	}
	key := c.hashKey(name)
	idxValue := c.sortedHashes[idx]
	if idxValue == key {
		return c.circle[idxValue], idx, nil
	} else if idxValue > key {
		lidx := idx - 250
		if lidx < 0 {
			lidx = 0
		}
		if c.sortedHashes[lidx] > key {
			i := c.search(c.hashKey(name))
			return c.circle[c.sortedHashes[i]], i, nil
		}

		i := c.searchWithIdx(key, lidx, idx)
		return c.circle[c.sortedHashes[i]], i, nil
	} else {
		uidx := idx + 250
		if uidx >= len(c.sortedHashes) {
			uidx = len(c.sortedHashes) - 1
		}

		if c.sortedHashes[uidx] < key {
			i := c.search(c.hashKey(name))
			return c.circle[c.sortedHashes[i]], i, nil
		}

		i := c.searchWithIdx(key, idx, uidx)
		return c.circle[c.sortedHashes[i]], i, nil
	}

	i := c.search(c.hashKey(name))
	return c.circle[c.sortedHashes[i]], i, nil
}

func (c *ConsistentStr) searchWithIdx(key uint32, lidx int, uidx int) (i int) {

	list := c.sortedHashes[lidx:uidx]
	f := func(x int) bool {
		return list[x] > key
	}
	i = sort.Search(len(list), f)
	if i >= len(list) {
		i = 0
	}

	i += lidx
	return
}

// GetNotNeedHash returns an element close to where name hashes to in the circle.
func (c *ConsistentStr) GetNotNeedHash(name uint32) (string, error) {
	if len(c.oneElt) > 0 {
		return c.oneElt, nil
	}
	if len(c.circle) == 0 {
		return "", ErrEmptyCircle
	}
	return c.circle[c.sortedHashes[c.search(name)]], nil
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
