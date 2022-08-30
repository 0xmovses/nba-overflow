package main

import (
	"fmt"
	"testing"

	. "github.com/bjartek/overflow"
	"gotest.tools/assert"
)

func testNBA(t *testing.T) {
	o, err := OverflowTesting()
	assert.NilError(t, err)

	fmt.Printf("NBA: %v", o)
	assert.Equal(t, 1, 1)
}
