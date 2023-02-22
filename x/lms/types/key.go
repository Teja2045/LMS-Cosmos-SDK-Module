package types

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

func AdminstoreKey(admin string) []byte {
	key := make([]byte, len(AdminKey)+len(admin))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], []byte(admin))
	return key
}

func StudentStoreKey(studentid string) []byte {
	key := make([]byte, len(StudentKey)+len(studentid))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], studentid)
	return key
}

func LeaveStoreKey(studentid string) []byte {
	key := make([]byte, len(LeaveKey)+len(studentid))
	copy(key, LeaveKey)
	copy(key[len(LeaveKey):], []byte(studentid))
	return key
}

func LeaveStoreCounterKey(studentid string) []byte {
	key := make([]byte, len(LeaveCounterKey)+len(studentid))
	copy(key, LeaveCounterKey)
	copy(key[len(LeaveCounterKey):], studentid)
	return key
}
