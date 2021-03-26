echo "POST invoke chaincode on peers of Org1, Org2, Org3 and Org4"
echo
TRX_ID=$(curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/mycc \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDk0NzIyNDYsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1NDk0MzYyNDZ9.AboFJdqOLZwyrRtS2MEOtVJqktIaonX7ebQA11pa4I4" \
  -H "content-type: application/json" \
  -d '{
	"peers": ["peer0.org1.rolex.com","peer0.org2.dealer.com", "peer0.org3.service.com"],
	"fcn":	"transferMtoC",
	"args":	["N3157163", "D67564", "Lotus Digital", "01-07-2019"]
}')
echo "Transaction ID is $TRX_ID"
echo
echo


