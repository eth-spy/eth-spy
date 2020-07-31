#dir: eth-spy
#bash ./private_net/start_node.sh [nodeID] [rpcport] [wsport] [port]

nodeID=$1
rpcport=$2
wsport=$3
port=$4

dataDir=./private_net/data/node_$nodeID
keyStore=./private_net/keystore/node_$nodeID
genesisDir=./private_net/genesis.json
passwordDir=./private_net/empty_password.txt
gethDir=$GOPATH/src/github.com/ethereum/go-ethereum/build/bin
#gethDir=$GOPATH/src/github.com/lilione/go_ethereum_SGN/build/bin

$gethDir/geth --datadir $dataDir init $genesisDir

$gethDir/geth \
    --datadir $dataDir \
    --keystore $keyStore \
    --mine --unlock 0 --password $passwordDir \
    --rpc --rpcport $rpcport --rpccorsdomain '*' --rpcapi admin,debug,eth,miner,net,personal,txpool,web3,clique \
    --ws --wsport $wsport --wsorigins '*' --wsapi admin,debug,eth,miner,net,personal,txpool,web3 \
    --networkid 438 \
    --port $port \
    --allow-insecure-unlock \
    --syncmode full \
    --ipcdisable \
    console