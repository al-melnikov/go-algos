package unionfind_test

import (
	"testing"

	"github.com/al-melnikov/algos/unionfind"

	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	assert := assert.New(t)
	uf := unionfind.NewUnionFind(5)
	for i := 0; i < 5; i++ {
		assert.Equal(i, uf.Find(i))
	}

	assert.False(uf.Connected(0, 4))

	uf.UnionSet(1, 2)
	assert.Equal(uf.Find(1), uf.Find(2))

	uf.UnionSet(3, 4)
	assert.Equal(uf.Find(3), uf.Find(4))

	uf.UnionSet(0, 3)
	assert.Equal(uf.Find(0), uf.Find(3))
	assert.Equal(uf.Find(0), uf.Find(4))
	assert.True(uf.Connected(0, 3))

	uf.UnionSet(4, 1)
	for i := 1; i < 5; i++ {
		assert.Equal(uf.Find(0), uf.Find(i))
	}
}
