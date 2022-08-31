package main

import (
	"fmt"
	"testing"

	. "github.com/bjartek/overflow"
	"github.com/zeebo/assert"
)

func TestNBAMint(t *testing.T) {
	o, err := OverflowTesting()
	assert.NoError(t, err)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		t.Error(err)
	}

	metaData := map[string]string{
		"type": "SlamDunk",
	}

	o.Tx("admin/create_play", WithSignerServiceAccount(),
		WithArgs("metadata", metaData),
	).AssertSuccess(t)

	o.Tx("admin/create_set", WithSignerServiceAccount(),
		WithArgs("setName", "Overflow"),
	).AssertSuccess(t)

	playId, _ := o.Script("plays/get_nextPlayId").GetAsInterface()
	assert.NotNil(t, playId)

	setId, _ := o.Script("sets/get_nextSetId").GetAsInterface()
	assert.NotNil(t, setId)

	//convert type of playId to uint64
	playIdUint := playId.(uint32)
	playIdUint -= 1

	setIdUint := setId.(uint32)
	setIdUint -= 1

	fmt.Printf("PlayId: %v, SetId: %v \n", playIdUint, setIdUint)

	o.Tx("admin/add_play_to_set", WithSignerServiceAccount(),
		WithArgs("setID", setIdUint),
		WithArgs("playID", playIdUint),
	).AssertSuccess(t).Print()

	o.Tx("admin/mint_moment", WithSignerServiceAccount(),
		WithArgs("setID", setIdUint),
		WithArgs("playID", playIdUint),
		WithArgs("recipientAddr", "account"),
	).AssertSuccess(t).Print()

}
