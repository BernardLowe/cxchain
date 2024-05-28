package account

import (
	"cxchain223/utils/hash"
)


type Account struct {
	Balance uint64
	Nonce uint64
	RootHash hash.Hash
	CodeHash hash.Hash
}