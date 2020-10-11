#!/bin/bash 

# source /Users/luis_acerv/.bash_profile

FUNC=$1

function sendTx() {
    # echo "Your private key"
    # read PRIV_KEY
    # echo "Destination address"
    # read TO_ADDR
    # post='{"privKey":"${PRIV_KEY}", "to":"${TO_ADDR}", "amount":1000000000000000000}'
    # echo $post
    # curl -d ${post} -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/eth/send-eth
    curl -d '{"privKey":"6d76e6f38bd3a1e1abbba37f03bb74a65ec7948a36db050ad190aa7f384c45d8", "to":"0x00Ea26936a4b0868aD6BEA4FDb847D457765544B", "amount":1000000000000000000}' -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/eth/send-eth
} 

$FUNC