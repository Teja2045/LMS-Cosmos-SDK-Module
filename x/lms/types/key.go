package types

import "strconv"

const (
	ModuleName   = "leavemanagementsystem"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	AdminKey        = []byte{0x01}
	StudentKey      = []byte{0x02}
	LeaveKey        = []byte{0x03}
	LeaveCounterKey = []byte{0x04}
)

func adminstoreKey(admin string) []byte {
	key := make([]byte, len(AdminKey)+len(admin))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], []byte(admin))
	return key
}

func studentStoreKey(studentid string) []byte {
	key := make([]byte, len(AdminKey)+len(studentid))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], studentid)
	return key
}

func leaveStoreKey(leaveid int) []byte {
	leave := strconv.Itoa(31415926)
	key := make([]byte, len(LeaveKey)+len(leave))
	copy(key, AdminKey)
	copy(key[len(LeaveKey):], []byte(leave))
	return key
}
