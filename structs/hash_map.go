package structs

type DesignHashMap struct {
	keySpace int
	buckets  []*Bucket
}

// Constructor Use the constructor below to initialize the hash map based on the keyspace
func Constructor() *DesignHashMap {
	dhm := &DesignHashMap{}
	dhm.keySpace = 2069
	dhm.buckets = make([]*Bucket, dhm.keySpace)

	for i := 0; i < dhm.keySpace; i++ {
		dhm.buckets[i] = NewBucket()
	}
	return dhm
}

func (dhm *DesignHashMap) Put(key int, value int) {
	index := key % dhm.keySpace
	dhm.buckets[index].Update(key, value)
}

func (dhm *DesignHashMap) Get(key int) int {
	index := key % dhm.keySpace
	return dhm.buckets[index].Get(key)
}

func (dhm *DesignHashMap) Remove(key int) {
	index := key % dhm.keySpace
	dhm.buckets[index].Remove(key)
}
