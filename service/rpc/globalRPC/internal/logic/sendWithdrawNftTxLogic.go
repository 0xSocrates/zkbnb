/*
 * Copyright © 2021 Zecrey Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/zecrey-labs/zecrey-legend/common/commonAsset"
	"github.com/zecrey-labs/zecrey-legend/common/commonConstant"
	"github.com/zecrey-labs/zecrey-legend/common/commonTx"
	"github.com/zecrey-labs/zecrey-legend/common/model/mempool"
	"github.com/zecrey-labs/zecrey-legend/common/model/tx"
	"github.com/zecrey-labs/zecrey-legend/common/sysconfigName"
	"github.com/zecrey-labs/zecrey-legend/common/util"
	"github.com/zecrey-labs/zecrey-legend/common/util/globalmapHandler"
	"github.com/zecrey-labs/zecrey-legend/common/zcrypto/txVerification"
	"math/big"
	"reflect"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

func (l *SendTxLogic) sendWithdrawNftTx(rawTxInfo string) (txId string, err error) {
	// parse transfer tx info
	txInfo, err := commonTx.ParseWithdrawNftTxInfo(rawTxInfo)
	if err != nil {
		errInfo := fmt.Sprintf("[sendWithdrawNftTx.ParseWithdrawNftTxInfo] %s", err.Error())
		logx.Error(errInfo)
		return "", errors.New(errInfo)
	}

	/*
		Check Params
	*/
	// check param: from account index
	err = util.CheckRequestParam(util.TypeAccountIndex, reflect.ValueOf(txInfo.AccountIndex))
	if err != nil {
		errInfo := fmt.Sprintf("[sendWithdrawNftTx] err: invalid accountIndex %v", txInfo.AccountIndex)
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, errors.New(errInfo))
	}
	// check gas account index
	gasAccountIndexConfig, err := l.svcCtx.SysConfigModel.GetSysconfigByName(sysconfigName.GasAccountIndex)
	if err != nil {
		logx.Errorf("[sendWithdrawNftTx] unable to get sysconfig by name: %s", err.Error())
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, err)
	}
	gasAccountIndex, err := strconv.ParseInt(gasAccountIndexConfig.Value, 10, 64)
	if err != nil {
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, errors.New("[sendWithdrawNftTx] unable to parse big int"))
	}
	if gasAccountIndex != txInfo.GasAccountIndex {
		logx.Errorf("[sendWithdrawNftTx] invalid gas account index")
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, errors.New("[sendWithdrawNftTx] invalid gas account index"))
	}

	var (
		accountInfoMap = make(map[int64]*commonAsset.AccountInfo)
	)
	nftInfo, err := globalmapHandler.GetLatestNftInfoForRead(
		l.svcCtx.NftModel,
		l.svcCtx.MempoolModel,
		l.svcCtx.RedisConnection,
		txInfo.NftIndex,
	)
	if err != nil {
		logx.Errorf("[sendWithdrawNftTx] unable to get nft info")
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, err)
	}
	accountInfoMap[txInfo.AccountIndex], err = globalmapHandler.GetLatestAccountInfo(
		l.svcCtx.AccountModel,
		l.svcCtx.MempoolModel,
		l.svcCtx.RedisConnection,
		txInfo.AccountIndex,
	)
	if err != nil {
		logx.Errorf("[sendWithdrawNftTx] unable to get account info: %s", err.Error())
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, err)
	}
	if accountInfoMap[nftInfo.CreatorAccountIndex] == nil {
		// get account info by gas index
		accountInfoMap[nftInfo.CreatorAccountIndex], err = globalmapHandler.GetBasicAccountInfo(
			l.svcCtx.AccountModel,
			l.svcCtx.RedisConnection,
			nftInfo.CreatorAccountIndex)
		if err != nil {
			logx.Errorf("[sendWithdrawNftTx] unable to get account info: %s", err.Error())
			return "", l.HandleCreateFailWithdrawNftTx(txInfo, err)
		}
	}
	// get account info by gas index
	if accountInfoMap[txInfo.GasAccountIndex] == nil {
		// get account info by gas index
		accountInfoMap[txInfo.GasAccountIndex], err = globalmapHandler.GetBasicAccountInfo(
			l.svcCtx.AccountModel,
			l.svcCtx.RedisConnection,
			txInfo.GasAccountIndex)
		if err != nil {
			logx.Errorf("[sendWithdrawNftTx] unable to get account info: %s", err.Error())
			return "", l.HandleCreateFailWithdrawNftTx(txInfo, err)
		}
	}

	if nftInfo.OwnerAccountIndex != txInfo.AccountIndex {
		logx.Errorf("[sendWithdrawNftTx] you're not owner")
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, errors.New("[sendWithdrawNftTx] you're not owner"))
	}

	txInfo.CreatorAccountIndex = nftInfo.CreatorAccountIndex
	txInfo.CreatorAccountNameHash = common.FromHex(accountInfoMap[nftInfo.CreatorAccountIndex].AccountNameHash)
	txInfo.CreatorTreasuryRate = nftInfo.CreatorTreasuryRate
	txInfo.NftContentHash = common.FromHex(nftInfo.NftContentHash)
	txInfo.NftL1Address = nftInfo.NftL1Address
	txInfo.NftL1TokenId, _ = new(big.Int).SetString(nftInfo.NftL1TokenId, 10)
	txInfo.CollectionId = nftInfo.CollectionId

	// check expired at
	now := time.Now().UnixMilli()
	if txInfo.ExpiredAt < now {
		logx.Errorf("[sendWithdrawNftTx] invalid time stamp")
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, errors.New("[sendWithdrawNftTx] invalid time stamp"))
	}

	var (
		txDetails []*mempool.MempoolTxDetail
	)
	// verify transfer tx
	txDetails, err = txVerification.VerifyWithdrawNftTxInfo(
		accountInfoMap,
		nftInfo,
		txInfo,
	)
	if err != nil {
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, err)
	}

	/*
		Check tx details
	*/

	/*
		Create Mempool Transaction
	*/
	// delete key
	key := util.GetNftKeyForRead(txInfo.NftIndex)
	_, err = l.svcCtx.RedisConnection.Del(key)
	if err != nil {
		logx.Errorf("[sendWithdrawNftTx] unable to delete key from redis: %s", err.Error())
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, err)
	}
	// write into mempool
	txInfoBytes, err := json.Marshal(txInfo)
	if err != nil {
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, err)
	}
	txId, mempoolTx := ConstructMempoolTx(
		commonTx.TxTypeWithdrawNft,
		txInfo.GasFeeAssetId,
		txInfo.GasFeeAssetAmount.String(),
		txInfo.NftIndex,
		commonConstant.NilPairIndex,
		commonConstant.NilAssetId,
		commonConstant.NilAssetAmountStr,
		"",
		string(txInfoBytes),
		"",
		txInfo.AccountIndex,
		txInfo.Nonce,
		txInfo.ExpiredAt,
		txDetails,
	)
	err = CreateMempoolTx(mempoolTx, l.svcCtx.RedisConnection, l.svcCtx.MempoolModel)
	if err != nil {
		return "", l.HandleCreateFailWithdrawNftTx(txInfo, err)
	}
	// update redis
	var formatNftInfo *commonAsset.NftInfo
	for _, txDetail := range mempoolTx.MempoolDetails {
		if txDetail.AssetType == commonAsset.NftAssetType {
			formatNftInfo, err = commonAsset.ParseNftInfo(txDetail.BalanceDelta)
			if err != nil {
				logx.Errorf("[sendWithdrawNftTx] unable to parse nft info: %s", err.Error())
				return txId, nil
			}
		}
	}
	nftInfoBytes, err := json.Marshal(formatNftInfo)
	if err != nil {
		logx.Errorf("[sendWithdrawNftTx] unable to marshal: %s", err.Error())
		return txId, nil
	}
	_ = l.svcCtx.RedisConnection.Setex(key, string(nftInfoBytes), globalmapHandler.NftExpiryTime)

	return txId, nil
}

func (l *SendTxLogic) HandleCreateFailWithdrawNftTx(txInfo *commonTx.WithdrawNftTxInfo, err error) error {
	errCreate := l.CreateFailWithdrawNftTx(txInfo, err.Error())
	if errCreate != nil {
		logx.Error("[sendtransfertxlogic.HandleCreateFailWithdrawNftTx] %s", errCreate.Error())
		return errCreate
	} else {
		errInfo := fmt.Sprintf("[sendtransfertxlogic.HandleCreateFailWithdrawNftTx] %s", err.Error())
		logx.Error(errInfo)
		return errors.New(errInfo)
	}
}

func (l *SendTxLogic) CreateFailWithdrawNftTx(info *commonTx.WithdrawNftTxInfo, extraInfo string) error {
	txHash := util.RandomUUID()
	nativeAddress := "0x00"
	txInfo, err := json.Marshal(info)
	if err != nil {
		errInfo := fmt.Sprintf("[sendtxlogic.CreateFailWithdrawNftTx] %s", err.Error())
		logx.Error(errInfo)
		return errors.New(errInfo)
	}
	// write into fail tx
	failTx := &tx.FailTx{
		// transaction id, is primary key
		TxHash: txHash,
		// transaction type
		TxType: commonTx.TxTypeWithdrawNft,
		// tx fee
		GasFee: info.GasFeeAssetAmount.String(),
		// tx fee l1asset id
		GasFeeAssetId: info.GasFeeAssetId,
		// tx status, 1 - success(default), 2 - failure
		TxStatus: tx.StatusFail,
		// l1asset id
		AssetAId: commonConstant.NilAssetId,
		// AssetBId
		AssetBId: commonConstant.NilAssetId,
		// tx amount
		TxAmount: commonConstant.NilAssetAmountStr,
		// layer1 address
		NativeAddress: nativeAddress,
		// tx proof
		TxInfo: string(txInfo),
		// extra info, if tx fails, show the error info
		ExtraInfo: extraInfo,
		// native memo info
		Memo: "",
	}

	err = l.svcCtx.FailTxModel.CreateFailTx(failTx)
	if err != nil {
		errInfo := fmt.Sprintf("[sendtxlogic.CreateFailWithdrawNftTx] %s", err.Error())
		logx.Error(errInfo)
		return errors.New(errInfo)
	}
	return nil
}
