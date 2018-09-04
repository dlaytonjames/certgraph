package fingerprint

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// Fingerprint sha256 of certificate bytes
type Fingerprint [sha256.Size]byte

// HexString print Fingerprint as hex
func (fp *Fingerprint) HexString() string {
	return fmt.Sprintf("%X", *fp)
}

// FromHashBytes returns a Fingerprint generated by the first len(Fingerprint) bytes
func FromHashBytes(data []byte) Fingerprint {
	var fp Fingerprint
	/*if len(data) != sha256.Size {
		v("Data is not correct SHA256 size", data)
	}*/
	for i := 0; i < len(data) && i < len(fp); i++ {
		fp[i] = data[i]
	}
	return fp
}

// FromBytes returns a Fingerprint generated by the provided bytes
func FromBytes(data []byte) Fingerprint {
	var fp Fingerprint
	fp = sha256.Sum256(data)
	return fp
}

// FromB64 returns a Fingerprint from a base64 encoded hash string
func FromB64(hash string) Fingerprint {
	data, _ := base64.StdEncoding.DecodeString(hash)
	/*if err != nil {
		fmt.Println(err)
	}*/
	return FromHashBytes(data)
}

// B64Encode returns the b64 string of a Fingerprint
func (fp *Fingerprint) B64Encode() string {
	return base64.StdEncoding.EncodeToString(fp[:])
}