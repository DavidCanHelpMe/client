// Auto-generated by avdl-compiler v1.1.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/prove_common.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
)

type ProofState int

const (
	ProofState_NONE         ProofState = 0
	ProofState_OK           ProofState = 1
	ProofState_TEMP_FAILURE ProofState = 2
	ProofState_PERM_FAILURE ProofState = 3
	ProofState_LOOKING      ProofState = 4
	ProofState_SUPERSEDED   ProofState = 5
	ProofState_POSTED       ProofState = 6
	ProofState_REVOKED      ProofState = 7
)

type ProofStatus int

const (
	ProofStatus_NONE              ProofStatus = 0
	ProofStatus_OK                ProofStatus = 1
	ProofStatus_LOCAL             ProofStatus = 2
	ProofStatus_FOUND             ProofStatus = 3
	ProofStatus_BASE_ERROR        ProofStatus = 100
	ProofStatus_HOST_UNREACHABLE  ProofStatus = 101
	ProofStatus_PERMISSION_DENIED ProofStatus = 103
	ProofStatus_FAILED_PARSE      ProofStatus = 106
	ProofStatus_DNS_ERROR         ProofStatus = 107
	ProofStatus_AUTH_FAILED       ProofStatus = 108
	ProofStatus_HTTP_429          ProofStatus = 129
	ProofStatus_HTTP_500          ProofStatus = 150
	ProofStatus_TIMEOUT           ProofStatus = 160
	ProofStatus_INTERNAL_ERROR    ProofStatus = 170
	ProofStatus_BASE_HARD_ERROR   ProofStatus = 200
	ProofStatus_NOT_FOUND         ProofStatus = 201
	ProofStatus_CONTENT_FAILURE   ProofStatus = 202
	ProofStatus_BAD_USERNAME      ProofStatus = 203
	ProofStatus_BAD_REMOTE_ID     ProofStatus = 204
	ProofStatus_TEXT_NOT_FOUND    ProofStatus = 205
	ProofStatus_BAD_ARGS          ProofStatus = 206
	ProofStatus_CONTENT_MISSING   ProofStatus = 207
	ProofStatus_TITLE_NOT_FOUND   ProofStatus = 208
	ProofStatus_SERVICE_ERROR     ProofStatus = 209
	ProofStatus_TOR_SKIPPED       ProofStatus = 210
	ProofStatus_TOR_INCOMPATIBLE  ProofStatus = 211
	ProofStatus_HTTP_300          ProofStatus = 230
	ProofStatus_HTTP_400          ProofStatus = 240
	ProofStatus_HTTP_OTHER        ProofStatus = 260
	ProofStatus_EMPTY_JSON        ProofStatus = 270
	ProofStatus_DELETED           ProofStatus = 301
	ProofStatus_SERVICE_DEAD      ProofStatus = 302
	ProofStatus_BAD_SIGNATURE     ProofStatus = 303
	ProofStatus_BAD_API_URL       ProofStatus = 304
	ProofStatus_UNKNOWN_TYPE      ProofStatus = 305
	ProofStatus_NO_HINT           ProofStatus = 306
	ProofStatus_BAD_HINT_TEXT     ProofStatus = 307
)

type ProofType int

const (
	ProofType_NONE             ProofType = 0
	ProofType_KEYBASE          ProofType = 1
	ProofType_TWITTER          ProofType = 2
	ProofType_GITHUB           ProofType = 3
	ProofType_REDDIT           ProofType = 4
	ProofType_COINBASE         ProofType = 5
	ProofType_HACKERNEWS       ProofType = 6
	ProofType_GENERIC_WEB_SITE ProofType = 1000
	ProofType_DNS              ProofType = 1001
	ProofType_ROOTER           ProofType = 100001
)

type Prove_commonInterface interface {
}

func Prove_commonProtocol(i Prove_commonInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "keybase.1.prove_common",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type Prove_commonClient struct {
	Cli rpc.GenericClient
}
