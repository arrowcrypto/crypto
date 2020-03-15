package bls

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicFunctions(t *testing.T) {
	req := require.New(t)

	blsPrivateKey, err := GenerateKeyPair()
	req.NoError(err)
	req.NotNil(blsPrivateKey)

	blsPublicKey := blsPrivateKey.GetPublicKey()
	req.NotNil(blsPublicKey)

}
