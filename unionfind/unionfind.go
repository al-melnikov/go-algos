package unionfind

type UnionFind struct {
	root []int
	rank []int
}

// return new union find structure
func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{
		root: make([]int, n),
		rank: make([]int, n),
	}

	for i := 0; i < n; i++ {
		uf.root[i] = i
		uf.rank[i] = 1
	}

	return &uf
}

func (uf *UnionFind) Find(x int) int {
	if x == uf.root[x] {
		return x
	}
	uf.root[x] = uf.Find(uf.root[x])
	return uf.root[x]
}

func (uf *UnionFind) UnionSet(x int, y int) {
	rootX := uf.Find(x)
	rootY := uf.Find(y)
	if rootX != rootY {
		if uf.rank[rootX] > uf.rank[rootY] {
			uf.root[rootY] = rootX
		} else if uf.rank[rootX] < uf.rank[rootY] {
			uf.root[rootX] = rootY
		} else {
			uf.root[rootY] = rootX
			uf.rank[rootX] += 1
		}
	}
}

func (uf *UnionFind) Connected(x int, y int) bool {
	return uf.Find(x) == uf.Find(y)
}
