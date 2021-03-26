var express= require('express');
var app = express();
var path = require('path');
app.use(express.static(path.join(__dirname,'public')));
app.get('/',function(req,res){
	res.sendFile(path.join(__dirname + '/index-dealer.html'));
});
app.get('/watches',function(req,res){
	res.sendFile(path.join(__dirname + '/table-dealer.html'));
});
app.get('/transfer',function(req,res){
	res.sendFile(path.join(__dirname + '/transfertoCustomer.html'));
});
app.get('/watch-history',function(req,res){
	res.sendFile(path.join(__dirname + '/watch-history.html'));
});
app.get('/watch-customer',function(req,res){
	res.sendFile(path.join(__dirname + '/watch-customer.html'));
});
app.get('/insurer',function(req,res){
	res.sendFile(path.join(__dirname + '../insurer/insurer.html'));
});
app.listen(8081);