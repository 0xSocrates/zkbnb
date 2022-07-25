// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	account "github.com/zecrey-labs/zecrey-legend/service/api/app/internal/handler/account"
	block "github.com/zecrey-labs/zecrey-legend/service/api/app/internal/handler/block"
	info "github.com/zecrey-labs/zecrey-legend/service/api/app/internal/handler/info"
	nft "github.com/zecrey-labs/zecrey-legend/service/api/app/internal/handler/nft"
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
				Method:  http.MethodGet,
				Path:    "/api/v1/account/getAccountStatusByAccountPk",
				Handler: account.GetAccountStatusByAccountPkHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/account/getAccountInfoByPubKey",
				Handler: account.GetAccountInfoByPubKeyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/account/getAccountInfoByAccountIndex",
				Handler: account.GetAccountInfoByAccountIndexHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/account/getAccountStatusByAccountName",
				Handler: account.GetAccountStatusByAccountNameHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/account/getAccountInfoByAccountName",
				Handler: account.GetAccountInfoByAccountNameHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/account/getBalanceByAssetIdAndAccountName",
				Handler: account.GetBalanceByAssetIdAndAccountNameHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/block/getBlocks",
				Handler: block.GetBlocksHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/block/getBlockByCommitment",
				Handler: block.GetBlockByCommitmentHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/block/getBlockByBlockHeight",
				Handler: block.GetBlockByBlockHeightHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/block/getCurrentBlockHeight",
				Handler: block.GetCurrentBlockHeightHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getLayer2BasicInfo",
				Handler: info.GetLayer2BasicInfoHandler(serverCtx),
			},
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
				Path:    "/api/v1/info/getGasFeeAssetList",
				Handler: info.GetGasFeeAssetListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getAccounts",
				Handler: info.GetAccountsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/search",
				Handler: info.SearchHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/info/getGasAccount",
				Handler: info.GetGasAccountHandler(serverCtx),
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
				Path:    "/api/v1/tx/getTxsListByBlockHeight",
				Handler: transaction.GetTxsListByBlockHeightHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getTxsListByAccountIndex",
				Handler: transaction.GetTxsListByAccountIndexHandler(serverCtx),
			},
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
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getTxsByPubKey",
				Handler: transaction.GetTxsByPubKeyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getTxByHash",
				Handler: transaction.GetTxByHashHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getMempoolTxs",
				Handler: transaction.GetMempoolTxsHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getmempoolTxsByAccountName",
				Handler: transaction.GetmempoolTxsByAccountNameHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/tx/getNextNonce",
				Handler: transaction.GetNextNonceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendTx",
				Handler: transaction.SendTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendCreateCollectionTx",
				Handler: transaction.SendCreateCollectionTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendMintNftTx",
				Handler: transaction.SendMintNftTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendAddLiquidityTx",
				Handler: transaction.SendAddLiquidityTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendAtomicMatchTx",
				Handler: transaction.SendAtomicMatchTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendCancelOfferTx",
				Handler: transaction.SendCancelOfferTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendRemoveLiquidityTx",
				Handler: transaction.SendRemoveLiquidityTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendSwapTx",
				Handler: transaction.SendSwapTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendTransferNftTx",
				Handler: transaction.SendTransferNftTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendTransferTx",
				Handler: transaction.SendTransferTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendWithdrawNftTx",
				Handler: transaction.SendWithdrawNftTxHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/v1/tx/sendWithdrawTx",
				Handler: transaction.SendWithdrawTxHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/nft/getMaxOfferId",
				Handler: nft.GetMaxOfferIdHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/v1/nft/getAccountNftList",
				Handler: nft.GetAccountNftListHandler(serverCtx),
			},
		},
	)
}
