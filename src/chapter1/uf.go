// +build ignore

package main

import "fmt"

type UF struct {
	id []int // 分量id
	n  int   // 连通分量的数量
}

func NewUF(n int) *UF {
	uf := UF{
		id: make([]int, n),
	}
	// 初始化每个分量
	for i := range uf.id {
		uf.id[i] = i
	}
	return &uf
}

// O(1)
func (uf UF) count() int {
	return uf.n
}

// O(1)
func (uf UF) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

// O(1)
func (uf UF) find(p int) int {
	return uf.id[p]
}

// O(n)
func (uf UF) union(p, q int) {
	pID, qID := uf.find(p), uf.find(q)
	if pID == qID {
		return
	}
	for i := range uf.id {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
	uf.n--
}

func main() {
	n := 10
	uf := NewUF(n)
	fmt.Println(uf.find(1))
	fmt.Println(uf.find(2))
	fmt.Println(uf.find(3))
	fmt.Println(uf.connected(1, 2))
	uf.union(1, 2)
	fmt.Println(uf.connected(1, 2))
}
