package blockchain

import (
	"bytes"
	"crypto/sha256"
	"math"
	"math/big"
)

const targetBits = 24

type Proof_of_Work struct {
	block  *Block
	target *big.Int
}

func NewProof_Work(b *Block) *Proof_of_Work {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))
	pow := &Proof_of_Work{b, target}
	return pow
}

func (pow *Proof_of_Work) prepData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToHex(pow.block.Timestamp),
			IntToHex(int64(targetBits)),
			IntToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *Proof_of_Work) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0
	for nonce < math.MaxInt64 {
		data := pow.prepData(nonce)
		hash = sha256.Sum256(data)
		hashInt.SetBytes(hash[:])
	}
}
