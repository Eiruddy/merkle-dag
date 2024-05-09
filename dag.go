package merkledag

import (
	"bytes"
	"hash"
	"hash/fnv"
)

// Add 将 Node 中的数据保存在 KVStore 中，然后计算出 Merkle Root
func Add(store KVStore, node Node, hp HashPool) []byte {
	rootHash := computeMerkleRoot(node, store, hp)
	return rootHash
}

// computeMerkleRoot 计算 Merkle Root
func computeMerkleRoot(node Node, store KVStore, hp HashPool) []byte {
	switch n := node.(type) {
	case *File:
		// 对文件进行哈希计算
		h := hp.Get()
		h.Write(n.Bytes())
		return h.Sum(nil)
	case *Dir:
		// 对文件夹中的每个文件或文件夹进行哈希计算
		h := hp.Get()
		for iter := n.It(); iter.Next(); {
			childNode := iter.Node()
			childHash := computeMerkleRoot(childNode, store, hp)
			h.Write(childHash)
		}
		return h.Sum(nil)
	default:
		return nil
	}
}
