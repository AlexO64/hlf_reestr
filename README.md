# hlf_reestr
# everything is happened in fabric-samples directory

# 1. create directory and copy reestr.go in it
inside fabric-samples/chaincode<br/>
$ mkdir reestr && cd reestr <br/>
$ <put the chaincode reestr.go enrollAdmin.js, registerUser.js, saveNewDoc.js and checkDoc.js and checkHash.js here> from chaincode/reestr <br/>

# 2. create application folder
inside fabric-samples<br/>
$ mkdir reestr && cd reestr<br/>

# 3. init all dependency for application
inside fabric-samples/reestr<br/>
$ rm -R wallet
$ npm init -y<br/>
$ npm install fabric-ca-client fabric-network -S <br/>
$ <put the five files enrollAdmin.js, registerUser.js, saveNewDoc.js and checkDoc.js and checkHash.js here> from reestr<br/>

# 4. Start fabrich basic network
inside fabric-samples/basic-network<br/>
$ ./start.sh <br/>
$ docker ps<br/>
$ docker-compose -f docker-compose.yml up -d cli<br/>
$ docker ps<br/>

Check all five containers are up and running.<br/>

# 5. Start chaincode
any directory<br/>
$ docker exec cli peer chaincode install -n reestr -v 1.0 -p "github.com/reestr"<br/>
$ docker exec cli peer chaincode instantiate -o orderer.example.com:7050 -C mychannel -n reestr -v 1.0 -c '{"Args":[]}' -P "OR ('Org1MSP.member')"<br/>

# 6. create admin account
inside fabric-samples/reestr<br/>
$ node enrollAdmin.js<br/>
$ ls wallet<br/>

# 7. create user account
inside fabric-samples/reestr<br/>
$ node registerUser.js alex<br/>
$ ls wallet<br/>

# 8. Save document
inside fabric-sample/reestr<br/>
$ node saveNewDoc.js alex "Hello, World."<br/>
Doc record created with TrxUid 3f8bad31cfb379ebb079e7db54bebe649b869753a776e999b4c303774e18943c<br/>
$

# 9. Check document
inside fabric-sample/reestr<br/>
$ node checkDoc.js alex 3f8bad31cfb379ebb079e7db54bebe649b869753a776e999b4c303774e18943c<br/>
Transaction has been evaluated<br/>
Doc record found, result is [object Object]<br/>


