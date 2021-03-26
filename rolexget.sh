echo "GET query chaincode on peer1 of Org1"
echo
curl -s -X GET \
  "http://localhost:4000/channels/mychannel/chaincodes/mycc?peer=peer0.org1.rolex.com&fcn=getAllWatches&args=%5B%22Rolex%22%5D" \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NDg0NTUzNTEsInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1NDg0MTkzNTF9.JvBxuO7fsTK9Vpyj5c0ef01OIKWi7IeF_NCoSawkGBA" \
  -H "content-type: application/json"
echo
echo
