package rest

import (
	"fmt"
	"github.com/hot3246624/TCPNetwork/x/tcp"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"

	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/gorilla/mux"
)


// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *codec.Codec, storeName string) {
	r.HandleFunc(fmt.Sprintf("custom/%s", storeName), transferHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("custom/%s", storeName), deployContractHandler(cdc, cliCtx)).Methods("POST")
	r.HandleFunc(fmt.Sprintf("custom/%s", storeName), execContractHandler(cdc, cliCtx)).Methods("POST")
}


type transferReq struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Name    string       `json:"name"`
	Amount  string       `json:"amount"`
}

// transferHandler
func transferHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO

		var req transferReq
		var msg tcp.MsgTransfer

		baseReq := req.BaseReq.Sanitize()
		clientrest.WriteGenerateStdTxResponse(w, cdc, cliCtx, baseReq, []sdk.Msg{msg})
	}
}


// deployContractHandler
func deployContractHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO
	}
}

// execContractHandler
func execContractHandler(cdc *codec.Codec, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO

	}
}
