package merkledag

import "hash"

type Link struct {
	Name string
	Hash []byte
	Size int
}

type Object struct {
	Links []Link
	Data  []byte
}

// Add 将一个节点添加到 Merkle DAG 中，并返回 Merkle Root。
func Add(store KVStore, node Node, hp HashPool) []byte {
    // 检查节点是否是文件
    if fileNode, ok := node.(File); ok {
        // 如果是文件，则计算其哈希值并将数据存储到 KVStore 中
        data := fileNode.Bytes()
        hash := calculateHash(data)
        store.Put(hash, data)
        return hash
    } else if dirNode, ok := node.(Dir); ok {
        // 如果是文件夹，则递归地添加其子节点
        var childHashes [][]byte
        it := dirNode.It()
        for it.Next() {
            childNode := it.Node()
            childHash := Add(store, childNode, hp)
            childHashes = append(childHashes, childHash)
        }
        // 根据子节点的哈希值计算文件夹节点的哈希值
        dirHash := calculateHash([]byte(fmt.Sprintf("%v", childHashes)))
        return dirHash
    } else {
        // 处理不支持的节点类型
        return nil
    }
}
