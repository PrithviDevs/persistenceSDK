package base

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/constants/errors"
	"github.com/persistenceOne/persistenceSDK/utilities/meta"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_HeightData(t *testing.T) {

	heightValue := NewHeight(123)
	testHeightData := NewHeightData(heightValue)
	testHeightData2 := NewHeightData(NewHeight(-1))

	require.Equal(t, "123", testHeightData.String())
	require.Equal(t, meta.Hash("123"), testHeightData.GenerateHash())
	require.Equal(t, "", testHeightData2.GenerateHash())

	dataAsString, Error := testHeightData.AsString()
	require.Equal(t, "", dataAsString)
	require.Equal(t, errors.EntityNotFound, Error)

	dataAsHeight, Error := testHeightData.AsHeight()
	require.Equal(t, heightValue, dataAsHeight)
	require.Equal(t, nil, Error)

	dataAsDec, Error := testHeightData.AsDec()
	require.Equal(t, sdkTypes.Dec{}, dataAsDec)
	require.Equal(t, errors.EntityNotFound, Error)

	dataAsID, Error := testHeightData.AsID()
	require.Equal(t, id{}, dataAsID)
	require.Equal(t, errors.EntityNotFound, Error)

	require.Equal(t, heightValue, testHeightData.Get())

	data, Error := ReadHeightData("")
	require.Equal(t, heightData{Value: height{Height: -1}}, data)
	require.Nil(t, Error)

	data, Error = ReadHeightData("testString")
	require.Equal(t, nil, data)
	require.NotNil(t, Error)

	data, Error = ReadHeightData("123")
	require.Equal(t, heightData{Value: height{Height: 123}}, data)
	require.Nil(t, Error)

	require.Equal(t, false, testHeightData.Equal(NewStringData("")))
	require.Equal(t, true, testHeightData.Equal(NewHeightData(NewHeight(123))))

}