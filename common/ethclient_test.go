package common

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestEthClient_GetPartialBeaconBellatrixBlock(t *testing.T) {
	client := NewEthClient(os.Getenv("CHARON_ENDPOINT_NODE"))

	b, err := client.GetPartialBeaconBellatrixBlock(5253886)
	require.NoError(t, err)

	fmt.Println(b.Body.ExecutionPayload.BlockHash)
}
