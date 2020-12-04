package deputize

import (
	sdkTypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/common"
	"github.com/persistenceOne/persistenceSDK/modules/maintainers/internal/module"
	"github.com/persistenceOne/persistenceSDK/schema/types/base"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Deputize_Message(t *testing.T) {

	testFromID := base.NewID("fromID")
	testToID := base.NewID("toID")
	testClassificationID := base.NewID("classificationID")

	fromAddress := "cosmos1pkkayn066msg6kn33wnl5srhdt3tnu2vzasz9c"
	fromAccAddress, Error := sdkTypes.AccAddressFromBech32(fromAddress)
	require.Nil(t, Error)

	maintainedTrait := "maintainedTraits:S|maintainedTraits"
	maintainedProperties, Error := base.ReadProperties(maintainedTrait)
	require.Equal(t, nil, Error)

	testMessage := newMessage(fromAccAddress, testFromID, testToID, testClassificationID, maintainedProperties, false, false, false)
	require.Equal(t, message{From: fromAccAddress, FromID: testFromID, ToID: testToID, ClassificationID: testClassificationID, MaintainedTraits: maintainedProperties, AddMaintainer: false, RemoveMaintainer: false, MutateMaintainer: false}, testMessage)
	require.Equal(t, module.Name, testMessage.Route())
	require.Equal(t, Transaction.GetName(), testMessage.Type())
	require.Equal(t, nil, testMessage.ValidateBasic())
	require.NotNil(t, message{}.ValidateBasic())
	require.Equal(t, sdkTypes.MustSortJSON(common.Codec.MustMarshalJSON(testMessage)), testMessage.GetSignBytes())
	require.Equal(t, []sdkTypes.AccAddress{fromAccAddress}, testMessage.GetSigners())
	require.Equal(t, testMessage, messageFromInterface(testMessage))
	require.Equal(t, message{}, messageFromInterface(nil))
	require.Equal(t, message{}, messagePrototype())

}