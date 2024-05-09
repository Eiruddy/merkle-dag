package merkledag

// Hash2File 根据 hash 和 path 返回对应的文件内容
func Hash2File(store KVStore, hash []byte, path string, hp HashPool) []byte {
	// 从 KVStore 中获取 hash 对应的数据
	data, err := store.Get(hash)
	if err != nil {
		return nil
	}
	return data
}
