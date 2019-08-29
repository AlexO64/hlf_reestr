# hlf_reestr
everything is happened in fabric-samples directory

1. create directory and copy reestr.go in it
# inside fabric-samples/chaincode
$ mkdir reestr && cd reestr
$ <put the reestr.go file here> from chaincode/reestr

2. create application folder
# inside fabric-samples
$ mkdir reestr && cd reestr

3. init all dependency for application
# inside fabric-samples/reestr
$ npm init -y
$ npm install fabric-ca-client fabric-network -S
$ <put the five files enrollAdmin.js, registerUser.js, saveNewDoc.js and checkDoc.js and checkHash.js here> from reestr

4. Start fabrich basic network
# inside fabric-samples/basic-network
$ ./start.sh 
$ docker ps
$ docker-compose -f docker-compose.yml up -d cli
$ docker ps

Check all five containers are up and running.

5. Start chaincode
# any directory
$ docker exec cli peer chaincode install -n reestr -v 1.0 -p "github.com/reestr"
$ docker exec cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n reestr -v 1.0 -c '{"Args":[]}' -P "OR ('Org1MSP.member')"

6. create admin account
# inside fabric-samples/reestr
$ node enrollAdmin.js
$ ls wallet

7. create user account
# inside fabric-samples/reestr
$ node registerUser.js alex
$ ls wallet

8. Save document
# inside fabric-sample/reestr
$ node saveNewDoc.js alex "Hello, World."
Doc record created with TrxUid 3f8bad31cfb379ebb079e7db54bebe649b869753a776e999b4c303774e18943c
$

9. Check document
# inside fabric-sample/reestr
$ node checkDoc.js alex 3f8bad31cfb379ebb079e7db54bebe649b869753a776e999b4c303774e18943c
Transaction has been evaluated
Doc record found, result is [object Object]


