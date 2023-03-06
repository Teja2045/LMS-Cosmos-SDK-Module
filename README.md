# LMS-Cosmos-SDK-Module

### Store Details:


#### admin
 0x01|adminaddress -> admin

#### student
 0x02|studentaddress -> student


#### student-leavecounter
0x04|studentaddress -> number
    
    - number above starts from 1 for each student

    - the number above can be used in different ways, leaveId when combined with student address, to find the number of leaves the student applied, to track all the leaves of that student

    - this is the connecting part of students and their leave, can be used for scaling and implementing other 

#### all leaves
0x03|studentaddress|leaveCounter -> leave
    
    - each time the student applies for a leave, his leave counter is updated, and combining studentaddress+counterNumber forms a unique key for leaves


#### handled leaves
0x05|student address -> signed/handled leaves

    - after admin handles(accept/reject) student's most recently applied leave, it will be stored here. These handled leaves also contains the info about the Admin who signed/ handled that leave

#### pending leaaves' students
0x06 -> a list of students whose leaves are pending

    - Each time a student applies for a leave, his address will be stored here.

    -Each time an admin handled a student's leave, his address will removed from here 


address: cosmos1mmk3jg09e6l9w0y5czvk2eesesky2u4u4enghv
  name: "2045"
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AxCAOLRyuQuxHXvOZrQEWXL6pBR+stpQ75NnyKPKxCQu"}'
  type: local


### example Addresses

- address: cosmos1dtg0z3k30fy5tqght7w49kfc6pezlffus5jl0v
  name: abcd
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AifS8Il5JeEYpqV4IaS4JY1YL28YuN3gLV6AoLdYkSk8"}'
  type: local


- address: cosmos1lawnlykgkkl85xk5u0xmmcsf6l48k7h7085avy
  name: mykey1
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2vicYBDS0WK5InqGonpe87jYKavFcLbB+7AEWiCEne6"}'
  type: local


- address: cosmos1yevrejew5twcuwg4eaxc60hehy37cxqg23rshu
  name: mykey2
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AgCi0xyuebm0n4PPbjurZ49xv8jxqWPRg6P6JyaGTvzg"}'
  type: local


- address: cosmos1en6xphyvhpq07zfht2xlnsna3pjd9mh0c803q9
  name: student1
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A6MKVpygNZhnIJ3TaQFY0kmnPvjyELwvHVZNM1z2Lcmh"}'
  type: local


- address: cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n
  name: validator-key
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A1wxcdib/3exc0rpAjBhR10/uxpF7TKVXKLBBleM0YyI"}'
  type: local

address: cosmos1smt2maj353hrcerc23r3am6dqjc0v83snxmkr7
  name: admina
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A5WZ5x2Z8ENR9QJKUJUc8ULkAHqVAI1tR//RAhcTIpau"}'
  type: local

  address: cosmos17h0gpjkdgf565r3tcwlxc77x07k0hug40m3wnk
  name: studenta
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A2nExIUzgyDFVCixrYRtCIOSghcVbUF8K5y4qbF1nKW7"}'
  type: local

  address: cosmos1vznegugh3pfkut6rmt0zhfx0es3gnl0cn5akvj
  name: studentb
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A4swMVXQvHzhDIeQgEAV+e4P2Nb3iefRibumAP2V4r4l"}'
  type: local



### Example Commands

#### Transactions

./simd tx lms RegisterAdmin cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n admin cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n --from validator-key --chain-id testnet


./simd tx lms AddStudents cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n saiteja b162045 cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n --from validator-key testnet

cosmos1yevrejew5twcuwg4eaxc60hehy37cxqg23rshu

./simd tx lms ApplyLeave cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n sick 2006-Jan-06 2006-Jan-06 cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n --from validator-key --chain-id testnet

./simd tx lms AcceptLeave cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n  cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n --from validator-key --chain-id testnet

./simd tx lms ApplyLeave cosmos1yevrejew5twcuwg4eaxc60hehy37cxqg23rshu sick 2006-Jan-06 2006-Jan-06 cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n --from validator-key --chain-id testnet

./simd tx lms AcceptLeave cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n cosmos1yevrejew5twcuwg4eaxc60hehy37cxqg23rshu  cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n --from validator-key --chain-id testnet --broadcast-mode block




#### Queries

./simd query lms LeaveStatus cosmos1en6xphyvhpq07zfht2xlnsna3pjd9mh0c803q9 student


 ./simd query lms LeaveStatus cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n student

 ./simd query lms LeaveStatus cosmos1mmk3jg09e6l9w0y5czvk2eesesky2u4u4enghv student

 ./simd query lms ListLeaves cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n admin