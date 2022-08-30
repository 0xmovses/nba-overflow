package main

import (
	"fmt"
	"testing"

	. "github.com/bjartek/overflow"
	"github.com/zeebo/assert"
)

func TestNBA(t *testing.T) {
	fmt.Println("Hello, overflow!")
	o, err := OverflowTesting()
	fmt.Println("called overflow!")
	assert.NoError(t, err)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		t.Error(err)
	}

	o.TxFileNameFN("create_set", WithSignerServiceAccount(),
		WithArgs("setName", "test-set"),
	)

	t.Run("set ids should not be nil after create set", func(t *testing.T) {
		ids := o.Script("sets/get_setIDs_ny_name", WithArg("setName", "test-set"))
		fmt.Printf("%v\n", ids)

		assert.NotNil(t, ids)
	})
}
