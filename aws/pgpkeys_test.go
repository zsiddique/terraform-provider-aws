// Legacy code migrated from hashicorp/vault/helper/pgpkeys

package aws

import (
	"bytes"
	"encoding/base64"
	"fmt"

	"github.com/keybase/go-crypto/openpgp"
	"github.com/keybase/go-crypto/openpgp/packet"
)

// pgpkeysDecryptBytes takes in base64-encoded encrypted bytes and the base64-encoded
// private key and decrypts it. A bytes.Buffer is returned to allow the caller
// to do useful thing with it (get it as a []byte, get it as a string, use it
// as an io.Reader, etc), and also because this function doesn't know if what
// comes out is binary data or a string, so let the caller decide.
func pgpkeysDecryptBytes(encodedCrypt, privKey string) (*bytes.Buffer, error) {
	privKeyBytes, err := base64.StdEncoding.DecodeString(privKey)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 private key: %s", err)
	}

	cryptBytes, err := base64.StdEncoding.DecodeString(encodedCrypt)
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 crypted bytes: %s", err)
	}

	entity, err := openpgp.ReadEntity(packet.NewReader(bytes.NewBuffer(privKeyBytes)))
	if err != nil {
		return nil, fmt.Errorf("error parsing private key: %s", err)
	}

	entityList := &openpgp.EntityList{entity}
	md, err := openpgp.ReadMessage(bytes.NewBuffer(cryptBytes), entityList, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error decrypting the messages: %s", err)
	}

	ptBuf := bytes.NewBuffer(nil)
	ptBuf.ReadFrom(md.UnverifiedBody)

	return ptBuf, nil
}
