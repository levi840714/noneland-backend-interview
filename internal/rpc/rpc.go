package rpc

import "github.com/google/wire"

var RpcProviderSet = wire.NewSet(
	NewXXExchangeClient,
)
