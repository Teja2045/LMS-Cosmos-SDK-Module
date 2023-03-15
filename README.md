


# IBC Setup Documentation

## Local-Setup

### Chain-1

#### choose first binary (eg lmsd)
- git clone <repo> (eg : lmsd)

#### Initialize chain
 - lmsd init --chain-id lmsdnet lmsdmoniker

#### Add Key
- lmsd keys add lmsdkey 

#### Store The Nemonic
- copy the nemonic and paste it somewhere as we need it later during hermes setup

#### Add Genesis Account
- lmsd add-genesis-account lmsdkey 1000000000000stake

#### Gentx To Create a Validator
- lmsd  gentx lmsdkey 1000000000stake --chain-id lmsdnet

#### Collect
- lmsd collect-gentxs

#### Start the node
- lmsd start



### Chain2

NOTE : Use --home flag when using same binary

#### choose second binary (eg gaiad)
- git clone <repo> (eg : gaia)

#### Initialize chain
 - gaiad init --chain-id gaianet gaiamoniker

#### Add Key
- gaiad keys add gaiakey 

#### Store The Nemonic
- copy the nemonic and paste it somewhere as we need it later during hermes setup

#### Add Genesis Account
- gaiad add-genesis-account gaiakey 1000000000000stake

#### Gentx To Create a Validator
- gaiad  gentx gaiakey 1000000000stake --chain-id lmsdnet

### Changes in configuration files

- go to .gaia -> config -> config.toml

- change rpc laddr = "tcp://127.0.0.1:26657" ->  "tcp://127.0.0.1:16657"

- change above p2p, pprof_laddr = "localhost:6060" -> "localhost:6061"

- change p2p laddr = "tcp://0.0.0.0:26656" -> laddr = "tcp://0.0.0.0:16656"

- go to app.toml

- change grpc address = "0.0.0.0:9090" -> address = "0.0.0.0:9092"

- change grpc-web address = "0.0.0.0:9091" -> address = "0.0.0.0:9095"

#### Collect
- gaiad collect-gentxs

#### Start the node
- gaiad start

### Hermes-Configuration

#### Change to .hermes directory
- cd .hermes

#### Create lmsd.json
- touch lmsd.json
- gedit lmsd.json
- paste the nemonic of first key (lmsdkey)

#### Create gaia.json
- touch gaia.json
- gedit gaia.json
- paste the nemonic of second key (gaiakey)

#### config.toml
- follow the below example

```
global]
log_level = 'info'
[mode]
[mode.clients]
enabled = true
refresh = true
misbehaviour = true
[mode.connections]
enabled = true
[mode.channels]
enabled = true
[mode.packets]
enabled = true
clear_interval = 100
clear_on_start = true
tx_confirmation = true
[telemetry]
enabled = true
host = '127.0.0.1'
port = 3001
[rest]
enabled = true
host    = '127.0.0.1'
port    = 3000
[[chains]]
id = 'gaianet'
rpc_addr = 'http://localhost:16657'
grpc_addr = 'http://localhost:9092'
websocket_addr = 'ws:/localhost:16657/websocket'
rpc_timeout = '15s'
account_prefix = 'regen'
key_name = 'gaiakey'
store_prefix = 'ibc'
gas_price = { price = 0.01, denom = 'stake' }
max_gas = 10000000
clock_drift = '5s'
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }
[[chains]]
id = 'lmsdnet'
rpc_addr = 'http://localhost:26657'
grpc_addr = 'http://localhost:9090'
websocket_addr = 'ws:/localhost:26657/websocket'
rpc_timeout = '15s'
account_prefix = 'cosmos'
key_name = 'lmsdkey'
store_prefix = 'ibc'
gas_price = { price = 0.01, denom = 'stake' }
max_gas = 10000000
clock_drift = '5s'
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }
```
 
NOTE : Lookout for id, keys, address

### Start Hermes

#### Directory to .hermes
- cd .hermes

#### Remove keys before restoring (if not the first time/ doing a restart)
- rm -rf keys

#### Restore Keys
- hermes keys add --mnemonic-file lmsdkey.json --chain lmsdnet 
- hermes keys add --mnemonic-file gaiakey.json --chain gaianet

#### Create a channel between the chains
- hermes create channel --a-chain lmsdnet --b-chain gaianet --a-port transfer --b-port transfer --new-client-connection

#### Start
- hermes start


## Setup accross different Systems (should connect to same wifi)

### Chain-1

#### choose first binary (eg lmsd)
- git clone <repo> (eg : lmsd)

#### Initialize chain
 - lmsd init --chain-id lmsdtejanet teja

#### Add Key
- lmsd keys add lmsdtejakey 

#### Store The Nemonic
- copy the nemonic and paste it somewhere as we need it later during hermes setup

#### Add Genesis Account
- lmsd add-genesis-account lmsdtejakey 1000000000000stake

#### Gentx To Create a Validator
- lmsd  gentx lmsdtejakey 1000000000stake --chain-id lmsdtejanet

### Changes in configuration files

- go to .simapp -> config -> config.toml

- change rpc laddr = "tcp://127.0.0.1:26657" ->  "tcp://0.0.0.0:26657"

#### Collect
- lmsd collect-gentxs

#### Start the node
- lmsd start



### Chain-2

#### chain-2 setup
- follow the same steps mentioned above

#### Lets assume
- ipaddress of system-2 : 192.168.1.32
- chainid : lmsdritviknet
- key : lmsdritvikkey

### Hermes-Configuration

#### Change to .hermes directory
- cd .hermes

#### Create lmsd.json
- touch lmsdteja.json
- gedit lmsdteja.json
- paste the nemonic of first key (lmsdtejakey)

#### Create gaia.json
- touch lmsdritvik.json
- gedit lmsdritvik.json
- paste the nemonic of second key (lmsdritvikkey)

#### config.toml
- follow the below example

```
[global]           
log_level = 'info'
[mode]
[mode.clients]
enabled = true
refresh = true
misbehaviour = true
[mode.connections]
enabled = true
[mode.channels]
enabled = true
[mode.packets]
enabled = true
clear_interval = 100
clear_on_start = true
tx_confirmation = true
[telemetry]
enabled = true
host = '127.0.0.1'
port = 3001
[rest]
enabled = true
host    = '127.0.0.1'
port    = 3000
[[chains]]
id = 'lmsdritviknet'
rpc_addr = 'http://192.168.1.32:26657'
grpc_addr = 'http://192.168.1.32:9090'
websocket_addr = 'ws:/192.168.1.32:26657/websocket'
rpc_timeout = '15s'
account_prefix = 'cosmos'
key_name = 'lmsdtejakey'
store_prefix = 'ibc'
gas_price = { price = 0.01, denom = 'stake' }
max_gas = 10000000
clock_drift = '5s'
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }
[[chains]]
id = 'lmsdtejanet'
rpc_addr = 'http://localhost:26657'
grpc_addr = 'http://localhost:9090'
websocket_addr = 'ws:/localhost:26657/websocket'
rpc_timeout = '15s'
account_prefix = 'cosmos'
key_name = 'lmsdtejakey'
store_prefix = 'ibc'
gas_price = { price = 0.01, denom = 'stake' }
max_gas = 10000000
clock_drift = '5s'
trusting_period = '14days'
trust_threshold = { numerator = '1', denominator = '3' }
 
```

NOTE : Lookout for id, keys, address, ip address


### Start Hermes

#### Directory to .hermes
- cd .hermes

#### Remove keys before restoring (if not the first time/ doing a restart)
- rm -rf keys

#### Restore Keys
- hermes keys add --mnemonic-file lmsdtejakey.json --chain lmsdtejanet 
- hermes keys add --mnemonic-file lmsdritvikkey.json --chain lmsdritviknet

#### Create a channel between the chains
- hermes create channel --a-chain lmsdtejanet --b-chain lmsdritviknet --a-port transfer --b-port transfer --new-client-connection

#### Start
- hermes start




### Transaction across chains using IBC(Hermes)

- lmsd tx ibc-transfer transfer transfer channel-0 <reciever key> 77stake --from lmskey --chain-id lmsdnet

- gaiad q bank balances <your address> --node http://127.0.0.1:16657

- hermes tx ft-transfer --dst-chain test2 --src-chain test1 --src-port transfer --src-channel transfer --amount 1000 --denom stake




### Notes and Errors

#### ip address command
- ifconfig

#### stake errors
- change .hermes -> config.toml -> denom (eg: uregen -> stake)
- change .gaia -> config -> config.toml -> min_gas_price ( eg: "" -> "0stake")

#### hermes keys not restoring
- check the nemonics properly
- cd .hermes

#### address already in use
- restart the system
- kill the process
- change .simapp -> config -> config.toml | app.toml

