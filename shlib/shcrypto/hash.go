package shcrypto

import (
	"math/big"

	bn256 "github.com/ethereum/go-ethereum/crypto/bn256/cloudflare"
	"golang.org/x/crypto/sha3"
)

func keccak256(ds ...[]byte) []byte {
	state := sha3.NewLegacyKeccak256()
	for _, d := range ds {
		state.Write(d)
	}
	h := []byte{}
	return state.Sum(h)
}

// HashBytesToBlock hashes the given byte slice and returns the result as a block.
func HashBytesToBlock(ds ...[]byte) Block {
	h := keccak256(ds...)
	var b Block
	copy(b[:], h)
	return b
}

// HashBlockToInt hashes a block and returns the result as an integer in Z_q.
func HashBlockToInt(d Block) *big.Int {
	h := keccak256(d[:])
	i := new(big.Int).SetBytes(h)
	i.Mod(i, bn256.Order) // TODO: check if this is fine
	return i
}

// HashGTToBlock hashes an element of GT and returns the result as a block.
func HashGTToBlock(gt *bn256.GT) Block {
	b := gt.Marshal()
	return HashBytesToBlock(b)
}
