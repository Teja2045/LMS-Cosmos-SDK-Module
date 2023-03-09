# LMS-Cosmos-SDK-Module

## Store Details:


### Admin
 - ##### 0x01 | adminaddress --> admin

### Student
 - ##### 0x02 | studentaddress --> student


### Leavecounter
 - ##### 0x04 | studentaddress --> number
    
    - number above starts from 1 for each student

    - the number above can be used in different ways, leaveId when combined with student address, to find the number of leaves the student applied, to track all the leaves of that student

    - this is the connecting part of students and their leave, can be used for scaling and implementing other 

### All Leaves
 - ##### 0x03 | studentaddress | leaveCounter --> leave
    
    - each time the student applies for a leave, his leave counter is updated, and combining studentaddress+counterNumber forms a unique key for leaves


### Handled Leaves
 - ##### 0x05 | student address --> signed/handled leaves

    - after admin handles(accept/reject) student's most recently applied leave, it will be stored here. These handled leaves also contains the info about the Admin who signed/ handled that leave

### Pending Leaves' Students
 - ##### 0x06 -> a list of students whose leaves are pending

    - Each time a student applies for a leave, his address will be stored here.

    -Each time an admin handled a student's leave, his address will removed from here 



### Example Addresses

- address: cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu
  name: a1
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A3HqwAUNowHNiTtO+DU8k1iIl5mOZobX5tgi7fQtw75E"}'
  type: local


- address: cosmos1flg656awzar09mhpayt5lmd4lzfwkcu9qzmr5u
  name: s1
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A4AyXEL4S3CPg4W+a/LzmA3B5CF3Y0t5RNuOpu3t5VM5"}'
  type: local



### Local Testnet commands 

#### To add keys

- ./simd keys add s1 

#### To initialise a testnet with a chain id "testnet"
- ./simd init --chain-id testnet myvalidator

#### To some money to accounts (keys above)
- ./simd add-genesis-account validator-key 100000000000stake


#### To add a validator (do a transaction)
- ./simd gentx validator-key 1000000000stake --chain-id testnet

#### Collect the transactions
- ./simd collect-gentxs

#### Start the chain
- ./simd start


### Example Commands

#### Transactions

- ./simd tx lms add-students cosmos1flg656awzar09mhpayt5lmd4lzfwkcu9qzmr5u saiteja 2045 --from adminaddress --chain-id testnet

- ./simd tx lms apply-leave sick 12-Jan-2022 13-Jan-2022 --from studentaddress --chain-id testnet

- ./simd tx lms register-admin adminname --from adminaddress --chain-id testnet

- ./simd tx lms accept-leave cosmos1flg656awzar09mhpayt5lmd4lzfwkcu9qzmr5u --from adminaddress --chain-id testnet


#### Queries

- ./simd query lms leave-status cosmos1flg656awzar09mhpayt5lmd4lzfwkcu9qzmr5u studentname

- ./simd query lms list-pending-leaves cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu adminname

 - ./simd query lms list-students cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu