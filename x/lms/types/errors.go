package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrStudentDoesNotExist = sdkerrors.Register(ModuleName, 1, "Student Does not exist")
	ErrAdminDoesNotExist   = sdkerrors.Register(ModuleName, 2, "Admin Does not Exist")
	ErrStudentNameNil      = sdkerrors.Register(ModuleName, 5, "Student Name should not be nil")

	ErrStudentDetailsNil     = sdkerrors.Register(ModuleName, 6, "Student Details should not be nil")
	ErrEmptyReason           = sdkerrors.Register(ModuleName, 7, "Reason should not be empty")
	ErrStudentAlreadyExists  = sdkerrors.Register(ModuleName, 8, "Student is already there in store")
	ErrLeaveNeverApplied     = sdkerrors.Register(ModuleName, 9, "Student never applied leave")
	ErrLeaveAlreadyHandled   = sdkerrors.Register(ModuleName, 10, "Leave is already handled")
	ErrAdminNameNil          = sdkerrors.Register(ModuleName, 11, "admin name should not be nil")
	ErrStudentIdNil          = sdkerrors.Register(ModuleName, 12, "student id should not be nil")
	ErrAdminDetailsNil       = sdkerrors.Register(ModuleName, 13, "admin details should not be nil")
	ErrLeaveDetailsNil       = sdkerrors.Register(ModuleName, 14, "leave details should not be nil")
	ErrNoNewStudentToBeAdded = sdkerrors.Register(ModuleName, 15, "atleast one student should be added")
)
