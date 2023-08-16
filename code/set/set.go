package set

// UF 并查集
type UF struct {
	// 记录连通分量
	count int
	// 节点的x的父节点是 parent[x]
	parent []int
}

// NewUF 构造函数
func NewUF(n int) *UF {
	parent := make([]int, n)
	// 父节点指针最初指向自己
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UF{
		count:  n,
		parent: parent,
	}
}

// Find 返回某个节点的根节点
// 递归
func (uf *UF) Find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.Find(uf.parent[x])
	}
	return uf.parent[x]
}

// Find2 压缩路径，查找根节点的同时将树进行扁平压缩
// 迭代
func (uf *UF) Find2(x int) int {
	root := x
	// 一直向上遍历，直到找到根
	for uf.parent[root] != root {
		root = uf.parent[root]
	}
	// 将所有非叶子节点都直接连接根节点，变成叶子
	oldParent := uf.parent[x]
	for x != root {
		uf.parent[x] = root
		x = oldParent
		oldParent = uf.parent[x]
	}
	return root
}

// Union 连通
func (uf *UF) Union(p int, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}
	// 两棵树合并
	uf.parent[rootP] = rootQ
	uf.count--
}

// Union2 简易版
func (uf *UF) Union2(p, q int) {
	uf.parent[uf.Find(p)] = uf.Find(q)
	uf.count--
}

// count 返回连通分量个数
func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) IsConnected(p, q int) bool {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	return rootP == rootQ
}
