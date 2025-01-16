package tx

import (
	"errors"

	"github.com/ethereum/go-ethereum/crypto"
)

type VoteCode int64

const (
	VoteIgnoreProposal  VoteCode = 200
	VoteProcessProposal VoteCode = 201
	VoteAcceptProposal  VoteCode = 202
	VoteRejectProposal  VoteCode = 203
	VoteGrantNewMember  VoteCode = 204
	VoteRejectNewMember VoteCode = 205
)

type HACTxType uint8
type HACTxCompressType uint8
type HACTxEncodingType uint8

const (
	HACTxTypeUnknown        HACTxType = 0
	HACTxTypeProposal       HACTxType = 1
	HACTxTypeDiscussion     HACTxType = 2
	HACTxTypeGrant          HACTxType = 3
	HACTxTypeRetract        HACTxType = 4
	HACTxTypeSettleProposal HACTxType = 5

	HACTxTypeGeneric HACTxType = 255
)

const (
	HACTxPrefixLen = 4
	HACTxSuffixLen = 8 + 8 + crypto.SignatureLength
	HACTxCheckLen  = HACTxPrefixLen + HACTxSuffixLen
)

const (
	HACTxVersion0 uint8 = 0
	HACTxVersion1 uint8 = 1
)

var (
	ErrInvalidTx         = errors.New("invalid tx")
	ErrUnsupportedTxType = errors.New("unsupported tx type")
	ErrUnmatchedTxType   = errors.New("unmatched tx type")

	ErrUnsupportedTxVersion  = errors.New("unsupported tx version")
	ErrUnsupportedTxCompress = errors.New("unsupported tx compress")
	ErrUnsupportedTxEncoding = errors.New("unsupported tx encoding")
	ErrUnsupportedTxData     = errors.New("unsupported tx data")
)

type HACTxHeader struct {
	Version uint8
	Type    HACTxType
}

type HACExtTxHeader struct {
	HACTxHeader
	Encoding HACTxEncodingType
	Compress HACTxCompressType
}
