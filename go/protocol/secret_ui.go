// Auto-generated by avdl-compiler v1.1.1 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/secret_ui.avdl

package keybase1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type SecretEntryArg struct {
	Desc           string `codec:"desc" json:"desc"`
	Prompt         string `codec:"prompt" json:"prompt"`
	Err            string `codec:"err" json:"err"`
	Cancel         string `codec:"cancel" json:"cancel"`
	Ok             string `codec:"ok" json:"ok"`
	Reason         string `codec:"reason" json:"reason"`
	UseSecretStore bool   `codec:"useSecretStore" json:"useSecretStore"`
}

type SecretEntryRes struct {
	Text        string `codec:"text" json:"text"`
	Canceled    bool   `codec:"canceled" json:"canceled"`
	StoreSecret bool   `codec:"storeSecret" json:"storeSecret"`
}

type GetPassphraseArg struct {
	SessionID int             `codec:"sessionID" json:"sessionID"`
	Pinentry  GUIEntryArg     `codec:"pinentry" json:"pinentry"`
	Terminal  *SecretEntryArg `codec:"terminal,omitempty" json:"terminal,omitempty"`
}

type SecretUiInterface interface {
	GetPassphrase(context.Context, GetPassphraseArg) (GetPassphraseRes, error)
}

func SecretUiProtocol(i SecretUiInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.secretUi",
		Methods: map[string]rpc.ServeHandlerDescription{
			"getPassphrase": {
				MakeArg: func() interface{} {
					ret := make([]GetPassphraseArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]GetPassphraseArg)
					if !ok {
						err = rpc.NewTypeError((*[]GetPassphraseArg)(nil), args)
						return
					}
					ret, err = i.GetPassphrase(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type SecretUiClient struct {
	Cli rpc.GenericClient
}

func (c SecretUiClient) GetPassphrase(ctx context.Context, __arg GetPassphraseArg) (res GetPassphraseRes, err error) {
	err = c.Cli.Call(ctx, "keybase.1.secretUi.getPassphrase", []interface{}{__arg}, &res)
	return
}
