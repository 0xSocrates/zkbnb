// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	account "github.com/zecrey-labs/zecrey-legend/service/api/app/internal/handler/account"
	info "github.com/zecrey-labs/zecrey-legend/service/api/app/internal/handler/info"
	pair "github.com/zecrey-labs/zecrey-legend/service/api/app/internal/handler/pair"
	root "github.com/zecrey-labs/zecrey-legend/service/api/app/internal/handler/root"
	transaction "github.com/zecrey-labs/zecrey-legend/service/api/app/internal/handler/transaction"
	"github.com/zecrey-labs/zecrey-legend/service/api/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/",
				Handler: root.GetStatusHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/account/getAccountStatusByPubKey",
				Handler: account.GetAccountStatusByPubKeyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/account/getAccountStatusByAccountName",
				Handler: account.GetAccountStatusByAccountNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/account/getAccountInfoByAccountName",
				Handler: account.GetAccountInfoByAccountNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/account/getBalanceByAssetIdAndAccountName",
				Handler: account.GetBalanceByAssetIdAndAccountNameHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/account/getAssetsByAccountName",
				Handler: account.GetAssetsByAccountNameHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/account/getAccountLiquidityPairs",
				Handler: account.GetAccountLiquidityPairsHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getAssetsList",
				Handler: info.GetAssetsListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getCurrencyPriceBySymbol",
				Handler: info.GetCurrencyPriceBySymbolHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getCurrencyPrices",
				Handler: info.GetCurrencyPricesHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getGasFee",
				Handler: info.GetGasFeeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getWithdrawGasFee",
				Handler: info.GetWithdrawGasFeeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getL1AmountByAssetid",
				Handler: info.GetL1AmountByAssetidHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getL1AmountList",
				Handler: info.GetL1AmountListHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/pair/getSwapAmount",
				Handler: pair.GetSwapAmountHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/pair/getAvailablePairs",
				Handler: pair.GetAvailablePairsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/pair/getLPValue",
				Handler: pair.GetLPValueHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/pair/getPairInfo",
				Handler: pair.GetPairInfoHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getTxsByAccountIndexAndTxType",
				Handler: transaction.GetTxsByAccountIndexAndTxTypeHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getTxsByAccountName",
				Handler: transaction.GetTxsByAccountNameHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/getTxsByPubKey",
				Handler: transaction.GetTxsByPubKeyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getTxByHash",
				Handler: transaction.GetTxByHashHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendTx",
				Handler: transaction.SendTxHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getMempoolTxs",
				Handler: transaction.GetMempoolTxsHandler(serverCtx),
			},
		},
	)
}
