package BLC
// 交易输出
type TxOutput struct {
	// 有多少钱（金额）
	Value int64
	// 钱是谁的（用户）
	ScriptPubkey string
}
