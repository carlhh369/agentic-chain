package tx

import (
	"bytes"
	"encoding/json"
)

type HACTx struct {
	Version   uint8
	Type      HACTxType
	Nonce     uint64
	Validator uint64
	Tx        any
	Sig       [][]byte
}

type GrantTx struct {
	Grants []GrantSt
}

type GrantSt struct {
	Statement string
	Amount    uint64
	Pubkey    []byte
}

func (d *GrantSt) Equal(grant GrantSt) bool {
	if d.Amount != grant.Amount || !bytes.Equal(d.Pubkey, grant.Pubkey) {
		return false
	}
	return true
}

type DiscussionTx struct {
	Proposal uint64
	Data     []byte
}

type ProposalTx struct {
	Proposer  uint64
	EndHeight uint64
	Data      []byte
}

type SettleProposalTx struct {
	Proposal uint64
}

type RetractTx struct {
	Amount uint64
}

type hacTxTmpl[Tx any] struct {
	Version   uint8
	Type      HACTxType
	Nonce     uint64
	Validator uint64
	Tx        Tx
	Sig       [][]byte
}

func (tx *HACTx) SigData(ext []byte) (dat []byte, err error) {
	ntx := *tx
	ntx.Sig = [][]byte{ext}
	dat, err = json.Marshal(ntx)
	return
}

func parseHACTxType(dat []byte) HACTxType {
	n := len(dat)
	if n < 4 {
		return HACTxTypeUnknown
	}
	v := dat[0]
	if v >= 0xc2 && v <= 0xf7 {
		return HACTxType(dat[2])
	} else if v > 0xf7 {
		return HACTxType(dat[1+v-0xf7+1])
	}
	return HACTxTypeUnknown
}

func unmarshalHACTx[Tx any](dat []byte) (btx *HACTx, err error) {
	var txt hacTxTmpl[Tx]
	err = json.Unmarshal(dat, &txt)
	if err != nil {
		return
	}
	btx = new(HACTx)
	btx.Version = txt.Version
	btx.Type = txt.Type
	btx.Nonce = txt.Nonce
	btx.Validator = txt.Validator
	btx.Tx = &txt.Tx
	btx.Sig = txt.Sig
	return
}

func UnmarshalHACTx(dat []byte) (btx *HACTx, err error) {
	tp := parseHACTxType(dat)
	switch tp {
	case HACTxTypeProposal:
		return unmarshalHACTx[ProposalTx](dat)
	case HACTxTypeDiscussion:
		return unmarshalHACTx[DiscussionTx](dat)
	case HACTxTypeGrant:
		return unmarshalHACTx[GrantTx](dat)
	case HACTxTypeRetract:
		return unmarshalHACTx[RetractTx](dat)
	case HACTxTypeSettleProposal:
		return unmarshalHACTx[SettleProposalTx](dat)
	default:
		err = ErrUnsupportedTxType
	}
	return
}

func MarshalHACTx(btx *HACTx) (dat []byte, err error) {
	return json.Marshal(btx)
}
