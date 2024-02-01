package dids_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/tbd54566975/web5-go/dids"
	"github.com/tbd54566975/web5-go/jwk"
	"github.com/tbd54566975/web5-go/jws"
)

func Test_ToKeys(t *testing.T) {
	did, err := dids.NewDIDJWK()
	assert.NoError(t, err)

	portableDID, err := did.ToKeys()
	assert.NoError(t, err)

	assert.Equal[string](t, did.URI, portableDID.URI)
	assert.True(t, len(portableDID.VerificationMethod) == 1, "expected 1 key")

	vm := portableDID.VerificationMethod[0]

	assert.NotEqual[jwk.JWK](t, vm.PublicKeyJWK, jwk.JWK{}, "expected publicKeyJwk to not be empty")
	assert.NotEqual[jwk.JWK](t, vm.PrivateKeyJWK, jwk.JWK{}, "expected privateKeyJWK to not be empty")
}

func TestBearerDIDFromKeys(t *testing.T) {
	did, err := dids.NewDIDJWK()
	assert.NoError(t, err)

	portableDID, err := did.ToKeys()
	assert.NoError(t, err)

	importedDID, err := dids.BearerDIDFromKeys(portableDID)
	assert.NoError(t, err)

	compactJWS, err := jws.Sign("hi", did)
	assert.NoError(t, err)

	compactJWSAgane, err := jws.Sign("hi", importedDID)
	assert.NoError(t, err)

	assert.Equal[string](t, compactJWS, compactJWSAgane, "failed to produce same signature with imported did")
}
