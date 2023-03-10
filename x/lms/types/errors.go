package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrStudentDoesNotExist = sdkerrors.Register(ModuleName, 1, "Student Does not exist")
	ErrAdminDoesNotExist   = sdkerrors.Register(ModuleName, 2, "Admin Does not Exist")
	ErrStudentNameNil      = sdkerrors.Register(ModuleName, 5, "Student Name should not be nil")
	ErrStudentDetailsNil   = sdkerrors.Register(ModuleName, 6, "Student Details should not be nil")
	ErrEmptyReason         = sdkerrors.Register(ModuleName, 7, "Reason should not be empty")
)
