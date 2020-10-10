# Create an API to interact with Ethereum blockchain using Golang

Hi folks! In this tutorial, we are going to learn how to create a simple API to interact with the Ethereum blockchain using Golang. 

Web3.js is the de-facto library to interact for interacting with Ethereum in JavaScript and Node.js. It takes care of encoding payloads and generating the RPC calls. Web3.js is very popular and heavily documented.

On the other hand, (geth), the most popular Ethereum implementation, is written in Go. It's a complete Ethereum node. If you build a dApp in Go, then you'll be using the go-ethereum libraries directly which means you can do everything the node can do.

So, for this tutorial, I chose to use Go as our weapon.

In simple terms, interacting with the blockchain means that you will be making RPC calls over HTTP. Any language will be able to do that for you but, using a compiled language such as Go will give you a better performance… If performance is important to you or your team.

Testing the send transaction endpoint
```bash
curl -d '{"privKey":"12a770fe34a793800abaab0a48b7a394ae440b9117d000178af81b61cda8ff15", "to":"0xa8Ce5Fb2DAB781B8f743a8096601eB01Ff0a246d", "amount":1000000000000000000}' -H "Content-Type: application/json" -X POST http://localhost:8080/api/v1/eth/send-eth

```