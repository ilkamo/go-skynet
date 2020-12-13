package skynet

import (
	"encoding/binary"
	"golang.org/x/crypto/blake2b"
)

func encodeNumber(number int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(number))
	return b
}

func encodeString(toEncode string) []byte {
	encodedNumber := encodeNumber(int64(len(toEncode)))
	return append(encodedNumber, []byte(toEncode)...)
}

func hashDataKey(dataKey string) ([]byte, error) {
	encodedDataKey := encodeString(dataKey)
	hash := blake2b.Sum256(encodedDataKey)
	return hash[:], nil
}

func hashAll(args ...[]byte) []byte {
	var bytes []byte
	for _, arg := range args {
		bytes = append(bytes, arg...)
	}
	hash := blake2b.Sum256(bytes)
	return hash[:]
}

func hashRegistryEntry(s SignedEntry) ([]byte, error) {
	dataKeyHash, err := hashDataKey(s.Entry.DataKey)
	if err != nil {
		return nil, err
	}

	return hashAll(
		dataKeyHash,
		encodeString(s.Entry.Data),
		encodeNumber(s.Entry.Revision),
	), nil
}