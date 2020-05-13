package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fmt"
)

const subsidy = 10

// Transaction 由交易ID，输入和输出构成，一笔交易可能有多个输出
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// 交易输入 TXInput
// Txid: 一个交易输入引用了之前一笔交易的一个输出，ID表明是之前哪一笔交易
// Vout: Vout为输出的索引
// ScriptSig: 提供解锁输出 Txid:Vout 的数据
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

// 交易输出 TXOutput
// Value: 有多少币，就是存储在 Value 里面
// ScriptPubKey: 对输出进行锁定
// 当前实现中，ScriptPubKey 将仅用一个字符串来代替
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

// IsCoinbase 判断是否是 coinbase 交易
func (tx Transaction) IsCoinbase() bool {
	return len(tx.Vin) == 1 && len(tx.Vin[0].Txid) == 0 && tx.Vin[0].Vout == -1
}

func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// 这里的 unlockingData 可以理解为地址
func (in *TXInput) CanUnlockOutputWith(unlockingData string) bool {
	return in.ScriptSig == unlockingData
}

func (out *TXOutput) CanBeUnlockedWith(unlockingData string) bool {
	return out.ScriptPubKey == unlockingData
}

// 创建一个 coinbase 交易，该没有输入，只有一个输出。
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}

// 创建一笔新的交易
func NewUTXOTransaction(from, to string, amount int, bc *Blockchain) *Transaction {
	var inputs []TXInput
	var outputs []TXOutput

	// 找到足够的未花费输出
	acc, validOutputs := bc.FindSpendableOutputs(from, amount)
	if acc < amount {
		panic("Error: Not enough funds")
	}
	for txid, outs := range validOutputs {
		txID, err := hex.DecodeString(txid)
		if err != nil {
			panic(err)
		}
		for _, out := range outs {
			input := TXInput{Txid: txID, Vout: out, ScriptSig: from}
			inputs = append(inputs, input)
		}
	}
	outputs = append(outputs, TXOutput{Value: amount, ScriptPubKey: to})
	// 如果 UTXO 总数超过所需，则产生找零
	if acc > amount {
		outputs = append(outputs, TXOutput{Value: acc - amount, ScriptPubKey: from})
	}
	tx := Transaction{nil, inputs, outputs}
	tx.SetID()
	return &tx
}
