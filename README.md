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

- address: cosmos15etl0x6q53zextm0jq2jfp5rcn54lp6ts0v0eu
  name: a1
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A3HqwAUNowHNiTtO+DU8k1iIl5mOZobX5tgi7fQtw75E"}'
  type: local

- address: cosmos1skv6g33ddt9x5teeqz5k4nxnw72zlr5z284wlw
  name: a2
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A4qdCtMThlta2rsKfrlRuTW9FhEHAUkv5sq2nhJe5zPZ"}'
  type: local

- address: cosmos1a0xz5ufeat7cj5cwgnav5vzz0kvp8v9cnu4jkg
  name: a3
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A65IhpR+1zLh8bS2Hdv+icYMQUtJWDQ3VJ07W3EiArin"}'
  type: local

- address: cosmos12985k8zaxjqdtwvghqxg5z5hfqefcz9565zz88
  name: a4
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A41hZCuLXZejqXaFvc6XkhYe7QO3wlqUf9L1TukqFx99"}'
  type: local

- address: cosmos1flg656awzar09mhpayt5lmd4lzfwkcu9qzmr5u
  name: s1
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A4AyXEL4S3CPg4W+a/LzmA3B5CF3Y0t5RNuOpu3t5VM5"}'
  type: local

- address: cosmos14z670pth37n8wyjhx7sfw87hl5kc0aywgz5nxa
  name: s2
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A3yYf7FJrirsMHkR9wIprjMyzZv0iQ4v9U0i9omEFaEt"}'
  type: local

- address: cosmos186auduv6c5n5qats26pktwg4kazldv650qeqhw
  name: s3
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A8EYIY7oozHMUgoddtwVDcunmvHimIl4LswBuEG2wI1u"}'
  type: local

- address: cosmos1d9xrzq2a4kusg60u4gd6wgsq76qmmj66ly06ad
  name: s4
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AnWN2V1RR5co0hIfvaEVjscBpwEamb1HPJsocn0NQxg3"}'
  type: local

- address: cosmos1zr0c72rn26x5vndky3zyvmzf2mjswvpz3pm35n
  name: validator-key
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A1wxcdib/3exc0rpAjBhR10/uxpF7TKVXKLBBleM0YyI"}'
  type: local


### Local Testnet commands 

#### to add keys
- ./simd keys add s1 

#### to initialise a testnet with a chain id "testnet"
- ./simd init --chain-id testnet myvalidator

#### to some money to accounts (keys above)
- ./simd add-genesis-account validator-key 100000000000stake


#### to add a validator (do a transaction)
- ./simd gentx validator-key 1000000000stake --chain-id testnet

#### collect the transactions
- ./simd collect-gentxs

#### start the chain
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