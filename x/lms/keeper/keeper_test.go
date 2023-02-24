package keeper_test

import (
	"fmt"
	"sync"
	"testing"
	"time"

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
	fmt.Println("hi ruth")
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

type AddLeaveTest struct {
	arg1     types.MsgApplyLeaveRequest
	student  types.Student
	expected error
}

var (
	exampleAdress string = sdk.AccAddress("address1").String()
	format        string = "2006-Jan-06"
	from1, _             = time.Parse(format, "2023-Feb-23")
	to1, _               = time.Parse(format, "2023-Feb-27")
	from2, _             = time.Parse(format, "2023-Feb-23")
	to2, _               = time.Parse(format, "2023-Feb-27")
	// from3, _             = time.Parse(format, "2023-Feb-23")
	// to3, _               = time.Parse(format, "2023-Feb-27")
)
var leaveTests = []AddLeaveTest{
	{
		arg1: types.MsgApplyLeaveRequest{
			Address: exampleAdress,
			Reason:  "very sick",
			From:    &from1,
			To:      &to1,
		},
		expected: types.ErrStudentDoesNotExist,
	},
	{
		arg1: types.MsgApplyLeaveRequest{
			Address: exampleAdress,
			Reason:  "sick",
			From:    &from2,
			To:      &to2,
		},
		student: types.Student{
			Address: exampleAdress,
			Name:    "Saiteja",
			Id:      "B162045",
		},
		expected: nil,
	},
}

func (suite *TestSuite) TestKeeper_Leaves() {

	require := suite.Require()
	for _, test := range leaveTests {

		//adding this line so that previous test wouldnt effect current test
		//require := suite.require
		if test.student.Address != "" {
			suite.stdntKeeper.AddStudent(suite.ctx, &test.student)
		}
		err := suite.stdntKeeper.AddLeave(suite.ctx, &test.arg1)
		require.Equal(err, test.expected)
	}

}

// demo test
func (suite *TestSuite) TestKeeper_ApplyLeave() {

	require := suite.Require()
	exampleStudentAdress1 := sdk.AccAddress("address1").String()
	//format := "2006-Jan-06"
	from, _ := time.Parse(format, "2023-Feb-23")
	to, _ := time.Parse(format, "2023-Feb-27")
	// from := time.Now()
	// to := time.Now()
	leave := &types.MsgApplyLeaveRequest{
		Address: exampleStudentAdress1,
		Reason:  "very sick",
		From:    &from,
		To:      &to,
	}
	fmt.Println("check1")
	suite.stdntKeeper.AddLeave(suite.ctx, leave)
	fmt.Println("check2")

	//cdc := suite.stdntKeeper.GetCodec()

	// retrieve the student from the keeper using the GetStudent method
	retrievedLeave, _ := suite.stdntKeeper.GetLeave(suite.ctx, exampleStudentAdress1)

	require.Equal(leave, retrievedLeave)

	//err := suite.stdntKeeper.AddStudent(suite.ctx, student)
	//require.Equal(err, types.ErrStudentAlreadyExists)
}

var exampleAdress2 string = sdk.AccAddress("address2").String()

type LeaveAcceptTest struct {
	student         types.Student
	admin           types.MsgRegisterAdminRequest
	leave           types.MsgApplyLeaveRequest
	studentaddress  string
	adminaddress    string
	registerStudent bool
	registerAdmin   bool
	applyleave      bool
	expected        error
}

var leaveAcceptTests = []LeaveAcceptTest{
	{
		studentaddress: exampleAdress,
		adminaddress:   exampleAdress2,
		expected:       types.ErrAdminDoesNotExist,
	},
	{
		studentaddress: exampleAdress,
		adminaddress:   exampleAdress2,
		admin: types.MsgRegisterAdminRequest{
			Name:    "Saiteja",
			Address: exampleAdress2,
		},
		registerAdmin: true,
		expected:      types.ErrStudentDoesNotExist,
	},
	{
		studentaddress: exampleAdress,
		adminaddress:   exampleAdress2,
		admin: types.MsgRegisterAdminRequest{
			Name:    "Saiteja",
			Address: exampleAdress2,
		},
		student: types.Student{
			Address: exampleAdress,
			Name:    "sai",
			Id:      "B162045",
		},
		leave: types.MsgApplyLeaveRequest{
			Address: exampleAdress,
			Reason:  "fever",
			From:    &from1,
			To:      &to1,
		},
		registerAdmin:   true,
		registerStudent: true,
		applyleave:      true,
		expected:        nil,
	},
}

func (suite *TestSuite) TestKeeper_AcceptLeave() {

	for _, test := range leaveAcceptTests {

		//adding this line so that previous test wouldnt effect current test
		require := suite.require
		if test.registerStudent {
			suite.stdntKeeper.AddStudent(suite.ctx, &test.student)
		}
		if test.registerAdmin {
			suite.stdntKeeper.AdminRegister(suite.ctx, &test.admin)
		}
		if test.applyleave {
			suite.stdntKeeper.AddLeave(suite.ctx, &test.leave)
		}

		err := suite.stdntKeeper.Accept(suite.ctx, test.studentaddress, test.adminaddress)

		require.Equal(err, test.expected)
	}

}
