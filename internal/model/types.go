package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MinedTxStatus uint64

const (
	Failed MinedTxStatus = iota
	Success
)

func (ms MinedTxStatus) String() string {
	switch {
	case ms == Failed:
		return "failed"
	case ms == Success:
		return "success"
	default:
		return "unknown"
	}
}

// TransactionType represents the type of Ethereum transaction
type TransactionType int

const (
	LegacyTxType TransactionType = iota
	AccessListTxType
	DynamicFeeTxType
	BlobTxType
	SetCodeTxType
)

func (t TransactionType) String() string {
	switch t {
	case LegacyTxType:
		return "LegacyTxType"
	case AccessListTxType:
		return "AccessListTxType"
	case DynamicFeeTxType:
		return "DynamicFeeTxType"
	case BlobTxType:
		return "BlobTxType"
	case SetCodeTxType:
		return "SetCodeTxType"
	default:
		return "unknown"
	}
}

func (t Tx) String() string {
	switch t.Type {
	case 0:
		return "LegacyTxType"
	case 1:
		return "AccessListTxType"
	case 2:
		return "DynamicFeeTxType"
	case 3:
		return "BlobTxType"
	case 4:
		return "SetCodeTxType"
	default:
		return "unknown"
	}
}

// TransactionStatus represents the current state of a transaction
type TransactionStatus string

const (
	StatusReceived TransactionStatus = "received"
	StatusPending  TransactionStatus = "pending"
	StatusQueued   TransactionStatus = "queued"
	StatusMined    TransactionStatus = "mined"
	StatusDropped  TransactionStatus = "dropped"
)

// TransactionMetadata contains additional metadata for filtering and grouping
type TransactionMetadata struct {
	Status       TransactionStatus `json:"status"`        // Current status of the tx
	TimeReceived int64             `json:"time_received"` // When seen in mempool
	TimePending  *int64            `json:"time_pending"`
	TimeQueued   int64             `json:"time_queued"`
	TimeMined    *int64            `json:"time_mined"`
	TimeDropped  int64             `json:"time_dropped"`
	BlockNumber  uint64            `json:"block_number"`
	BlockHash    string            `json:"block_hash"`
	MineStatus   string            `json:"mine_status"`
	GasUsed      uint64            `json:"gasUsed"`
}

type Tx struct {
	ChainID            string   `json:"chain_id"`
	From               string   `json:"from"`
	To                 string   `json:"to,omitempty"`
	IsContractCreation bool     `json:"isContractCreation"`
	Nonce              uint64   `json:"nonce"`
	Value              string   `json:"value"`
	Gas                uint64   `json:"gas"`
	GasPrice           *big.Int `json:"gas_price,omitempty"`
	MaxFeePerGas       string   `json:"max_fee_per_gas,omitempty"`
	MaxPriorityFee     string   `json:"max_priority_fee,omitempty"`
	MaxFeePerBlobGas   string   `json:"max_fee_per_blob_gas,omitempty"`
	Data               string   `json:"data,omitempty"`
	Type               uint8    `json:"type"`
}

// StoredTransaction represents a transaction with its metadata
type StoredTransaction struct {
	Hash     string              `json:"hash"`
	Tx       Tx                  `json:"tx"`
	Metadata TransactionMetadata `json:"metadata"`
}
type TxDiff struct {
	Tx       map[string]map[string]interface{} `json:"tx"`
	Metadata map[string]map[string]interface{} `json:"metadata"`
}
type TxBlock struct {
	Tx       map[string]interface{} `json:"tx"`
	Metadata map[string]interface{} `json:"metadata"`
}

type ApiTxResponse struct {
	Hash    string   `json:"hash"`
	Clients []string `json:"clients"`
	Diff    TxDiff   `json:"diff"`
	Common  TxBlock  `json:"common"`
}

type TxSummary struct {
	Hash    string  `json:"hash"`
	From    string  `json:"from"`
	Value   string  `json:"value"`
	GasUsed float64 `json:"gasUsed"`
	Nonce   uint64  `json:"nonce"`
	Type    string  `json:"type"`
}

type RPCRequest struct {
	Method  string   `json:"method"`
	Params  []string `json:"params"`
	Id      int      `json:"id"`
	Jsonrpc string   `json:"jsonrpc"`
}

type RPCResponse struct {
	JSONRPC string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  Result `json:"result"`
}

type SubscriptionResponse struct {
	Jsonrpc string             `json:"jsonrpc"`
	Method  string             `json:"method"`
	Params  SubscriptionParams `json:"params"`
}

type SubscriptionParams struct {
	Subscription string `json:"subscription"`
	TxHash       string `json:"result"`
}

type Result struct {
	Pending map[string]map[string]*types.Transaction `json:"pending"`
	Queued  map[string]map[string]*types.Transaction `json:"queued"`
}

type SSEMessage struct {
	Slot                       string   `json:"slot"`
	ValidatorIndex             string   `json:"validator_index"`
	InclusionListCommitteeRoot string   `json:"inclusion_list_committee_root"`
	Transactions               []string `json:"transactions"`
}

type Data struct {
	Message   SSEMessage `json:"message"`
	Signature string     `json:"signature"`
}

type MempoolMessage struct {
	Version string `json:"version"`
	Data    Data   `json:"data"`
}

type InclusionReport struct {
	Included []common.Hash    `json:"included"`
	Missing  []common.Hash    `json:"missing"`
	Summary  InclusionSummary `json:"summary"`
}

type InclusionSummary struct {
	Total    int `json:"total"`
	Included int `json:"included"`
	Missing  int `json:"missing"`
}

type InclusionListWithSlot struct {
	Slot   int             `json:"slot"`
	Report InclusionReport `json:"report"`
}
