package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBlockingNode(t *testing.T) {
	//r, err := miniredis.Run()
	//assert.Nil(t, err)
	node, err := CreateBlockingNode(New())
	assert.Nil(t, err)
	node.Close()
	node, err = CreateBlockingNode(New(Cluster()))
	assert.Nil(t, err)
	node.Close()
}
