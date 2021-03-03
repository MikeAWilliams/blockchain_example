package network_test

import (
	"blockchain_example/network"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_BlockIsValid_Happy(t *testing.T) {
	testObject := &network.Network{}
	f1Called := 0
	var f1Arg string
	testObject.SubscribeForNewData(func(data string) {
		f1Called++
		f1Arg = data
	})
	f2Called := 0
	var f2Arg string
	testObject.SubscribeForNewData(func(data string) {
		f2Called++
		f2Arg = data
	})

	testObject.PostNewData("TACO")

	require.Equal(t, 1, f1Called)
	require.Equal(t, 1, f2Called)
	require.Equal(t, "TACO", f1Arg)
	require.Equal(t, "TACO", f2Arg)
}
