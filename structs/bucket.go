package structs

// Bucket represents a bucket for storing key-value pairs.
type Bucket struct {
	bucket []Pair // Slice to store key-value pairs
}

// Pair represents a key-value pair.
type Pair struct {
	key   int
	value int
}

// NewBucket creates a new instance of Bucket.
func NewBucket() *Bucket {
	return &Bucket{
		bucket: make([]Pair, 0), // Initialize the slice
	}
}

// Get return the value from the Bucket
func (b *Bucket) Get(hashKey int) int {
	for _, pair := range b.bucket {
		if pair.key == hashKey {
			return pair.value
		}
	}
	return -1
}

// Update function put value in the Bucket
func (b *Bucket) Update(hashKey int, value int) {
	found := false
	for i, pair := range b.bucket {
		if pair.key == hashKey {
			b.bucket[i].value = value
			found = true
			break
		}
	}

	if !found {
		b.bucket = append(b.bucket, Pair{hashKey, value})
	}
}

// Remove function delete value from the Bucket
func (b *Bucket) Remove(hashKey int) {
	for i := range b.bucket {
		if b.bucket[i].key == hashKey {
			b.bucket = append(b.bucket[:i], b.bucket[i+1:]...)
			break
		}
	}
}
