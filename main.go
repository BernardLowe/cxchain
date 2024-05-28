package main

import (
	"cxchain223/kvstore"
	"cxchain223/trie"
	"cxchain223/account"
	"cxchain223/utils/rlp"
	"fmt"
)

func main() {
	db := kvstore.NewLevelDB("./leveldb")
	account1 := &account.Account{
		Balance: 10,
		Nonce: 1,
	}

	accoutData, err:=rlp.EncodeToBytes(account1)

	state := trie.NewState(db, trie.EmptyHash)
	key, _ := hexutils.DecodeBytes("0xB2219123")
	state.Store(key, accoutData)
	value, err := state.Load(key)
	// fmt.Println(string(value), err)

	var acc account.Account
	err = rlp.DecodeBytes(value, &acc)
	fmt.Printf("Account.balance = %d\n", acc.Balance)
	fmt.Printf("account.nonce - %d\n", acc.Nonce)
	fmt.Println(err)
}
