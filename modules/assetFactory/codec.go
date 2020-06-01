package assetFactory

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/mapper"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/burn"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mint"
	"github.com/persistenceOne/persistenceSDK/modules/assetFactory/transactions/mutate"
)

func RegisterCodec(codec *codec.Codec) {
	mapper.RegisterCodec(codec)

	burn.RegisterCodec(codec)
	mint.RegisterCodec(codec)
	mutate.RegisterCodec(codec)
}