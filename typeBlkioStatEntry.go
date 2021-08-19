package iotmakerdockerbuilder

// BlkioStatEntry is one small entity to store a piece of Blkio stats
// Not used on Windows.
type BlkioStatEntry struct {
	Major uint64 `json:"major"`
	Minor uint64 `json:"minor"`
	Op    string `json:"op"`
	Value uint64 `json:"value"`
}
