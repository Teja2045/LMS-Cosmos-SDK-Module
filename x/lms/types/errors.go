package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrStudentDoesNotExist = sdkerrors.Register(ModuleName, 1, "Student Does not exist")
	ErrAdminDoesNotExist   = sdkerrors.Register(ModuleName, 2, "Admin Does not Exist")
	ErrStudentNameNil      = sdkerrors.Register(ModuleName, 5, "Student Name should not be nil")

	ErrStudentDetailsNil    = sdkerrors.Register(ModuleName, 6, "Student Details should not be nil")
	ErrEmptyReason          = sdkerrors.Register(ModuleName, 7, "Reason should not be empty")
	ErrStudentAlreadyExists = sdkerrors.Register(ModuleName, 8, "Student is already there in store")
	ErrLeaveNeverApplied    = sdkerrors.Register(ModuleName, 9, "Student never applied leave")
	ErrLeaveAlreadyHandled  = sdkerrors.Register(ModuleName, 10, "Leave is already handled")
	ErrAdminNameNil         = sdkerrors.Register(ModuleName, 11, "admin name should not be nil")
	ErrStudentIdNil         = sdkerrors.Register(ModuleName, 12, "student id should not be nil")
)

// cosmos1pawsu36n7m5ta02gdvd3zuhrgyww3v7hl6g8ca
// validator-key
// '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AiG+995s1qpSArYClbVjGlvX/xtCiQ+0d8t6+WSl2mT3"}'

//address: cosmos1mmk3jg09e6l9w0y5czvk2eesesky2u4u4enghv
//name: "2045"
//pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AxCAOLRyuQuxHXvOZrQEWXL6pBR+stpQ75NnyKPKxCQu"}'
//type: local

// address: cosmos1dtg0z3k30fy5tqght7w49kfc6pezlffus5jl0v
//name: abcd
//pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AifS8Il5JeEYpqV4IaS4JY1YL28YuN3gLV6AoLdYkSk8"}'
//type: local

/*
 address: cosmos1sanc3p05nrtahr7fgskulhmzqzcuvhr5gts2ys
  name: validator-key
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A1BJTogflABBlSw9rT4Uz51bnH+F9c+cVLYG2/U+hRLR"}'
  type: local


- address: cosmos1lawnlykgkkl85xk5u0xmmcsf6l48k7h7085avy
  name: mykey1
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2vicYBDS0WK5InqGonpe87jYKavFcLbB+7AEWiCEne6"}'
  type: local

  address: cosmos1yevrejew5twcuwg4eaxc60hehy37cxqg23rshu
  name: mykey2
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AgCi0xyuebm0n4PPbjurZ49xv8jxqWPRg6P6JyaGTvzg"}'
  type: local

*/
