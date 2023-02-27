package types

import "strconv"

const (
	ModuleName   = "lms"
	StoreKey     = ModuleName
	RouterKey    = ModuleName
	QuerierRoute = ModuleName
)

var (
	//0x01+adminaddress -> admin details, we can use this to check if an address is registered as admin
	AdminKey = []byte{0x01}

	//0x02+studentaddress -> student details, we can use this to check if a student is registered
	StudentKey = []byte{0x02}

	//0x03+studentaddess+leaveid(value from leave counter) -> leaverequest details, we can use this to keep track of leaves of students
	// Note : contains pending leaves + handled leaves
	LeaveKey = []byte{0x03}

	//0x04+studentaddress -> leaveId, an integer which starts from 1, used to track number of
	//leaves and most recently applied leave for each user
	LeaveCounterKey = []byte{0x04}

	//0x05+studenadress+leaveid(value from leave counter) -> leave status
	LeaveStatusKey = []byte{0x05}

	//key to access pending leaves... It is a global key
	PendingLeaveStudentsKey = []byte{0x06}
)

func AdminstoreKey(admin string) []byte {
	key := make([]byte, len(AdminKey)+len(admin))
	copy(key, AdminKey)
	copy(key[len(AdminKey):], []byte(admin))
	return key
}

func StudentStoreKey(studentAddress string) []byte {
	key := make([]byte, len(StudentKey)+len(studentAddress))
	copy(key, StudentKey)
	copy(key[len(StudentKey):], studentAddress)
	return key
}

func LeaveStoreKey(studentid string, leaveid int) []byte {
	leaveId := strconv.Itoa(leaveid)
	key := make([]byte, len(LeaveKey)+len(studentid)+len(leaveId))
	copy(key, LeaveKey)
	copy(key[len(LeaveKey):], []byte(studentid))
	copy(key[len(LeaveKey)+len(studentid):], []byte(leaveId))
	return key
}

func LeaveCounterStoreKey(studentid string) []byte {
	key := make([]byte, len(LeaveCounterKey)+len(studentid))
	copy(key, LeaveCounterKey)
	copy(key[len(LeaveCounterKey):], studentid)
	return key
}

func LeaveStatusStoreKey(studentAddress string, leaveid int) []byte {
	leaveId := strconv.Itoa(leaveid)
	key := make([]byte, len(LeaveStatusKey)+len(studentAddress)+len(leaveId))
	copy(key, LeaveStatusKey)
	copy(key[len(LeaveStatusKey):], []byte(studentAddress))
	copy(key[len(LeaveStatusKey)+len(studentAddress):], []byte(leaveId))
	return key
}

func PendingLeaveStudentsStoreKey() []byte {
	return PendingLeaveStudentsKey
}
