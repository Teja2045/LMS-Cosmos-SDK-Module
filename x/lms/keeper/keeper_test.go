package keeper_test

import (
	"sync"
	"testing"

	"lmsmodule/x/lms/keeper"
	"lmsmodule/x/lms/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/testutil"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	dbm "github.com/tendermint/tm-db"
)

type TestSuite struct {
	suite.Suite
	ctx         sdk.Context
	stdntKeeper keeper.Keeper
	*assert.Assertions
	mu      sync.RWMutex
	require *require.Assertions
	t       *testing.T
}

func (s *TestSuite) SetupTest() {
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	encCfg := simapp.MakeTestEncodingConfig()
	lmsKey := sdk.NewKVStoreKey(types.StoreKey)
	ctx := testutil.DefaultContext(lmsKey, sdk.NewTransientStoreKey("transient_test"))
	keeper := keeper.NewKeeper(lmsKey, encCfg.Codec)
	cms.MountStoreWithDB(lmsKey, storetypes.StoreTypeIAVL, db)
	s.Require().NoError(cms.LoadLatestVersion())
	s.stdntKeeper = keeper
	s.ctx = ctx
}

func (suite *TestSuite) T() *testing.T {
	suite.mu.RLock()
	defer suite.mu.RUnlock()
	return suite.t
}

// SetT sets the current *testing.T context.
func (suite *TestSuite) SetT(t *testing.T) {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	suite.t = t
	suite.Assertions = assert.New(t)
	suite.require = require.New(t)
}

// Require returns a require context for suite.
func (suite *TestSuite) Require() *require.Assertions {
	suite.mu.Lock()
	defer suite.mu.Unlock()
	if suite.require == nil {
		suite.require = require.New(suite.T())
	}
	return suite.require
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestKeeper_AdminRegister() {

	require := suite.Require()
	//suite := new(TestSuite)

	// create and store an admin in the keeper
	//fmt.Println("fine 1")
	exampleAdress := sdk.AccAddress("address1").String()
	admin := types.NewMsgRegisterAdminRequest("saiteja", exampleAdress)
	//fmt.Println("fine 2", student)
	suite.stdntKeeper.AdminRegister(suite.ctx, admin)

	// get codec and use it to Marshal Admin
	cdc := suite.stdntKeeper.GetCodec()
	//fmt.Println("fine 4")
	expected, _ := cdc.Marshal(admin)
	// fmt.Println("fine 5")

	// retrieve the student from the keeper using the GetAdmin method
	retrievedAdmin := suite.stdntKeeper.GetAdmin(suite.ctx, sdk.AccAddress("address1").String())

	//fmt.Println(retrievedAdmin)
	require.Equal(expected, retrievedAdmin)
}

func (suite *TestSuite) TestKeeper_AddStudent() {

	require := suite.Require()
	exampleStudentAdress1 := sdk.AccAddress("address1").String()

	student := &types.Student{
		Address: exampleStudentAdress1,
		Name:    "saiteja",
		Id:      "2045",
	}
	suite.stdntKeeper.AddStudent(suite.ctx, student)

	cdc := suite.stdntKeeper.GetCodec()
	expected, _ := cdc.Marshal(student)

	// retrieve the student from the keeper using the GetStudent method
	retrievedAdmin := suite.stdntKeeper.GetStudent(suite.ctx, exampleStudentAdress1)

	require.Equal(expected, retrievedAdmin)

	err := suite.stdntKeeper.AddStudent(suite.ctx, student)
	require.Equal(err, types.ErrStudentAlreadyExists)
}
