var express= require('express');
var app = express();
var path = require('path');
app.use(express.static(path.join(__dirname,'public')));
app.get('/',function(req,res){
	res.sendFile(path.join(__dirname + '/Home-page.html'));
});
app.get('/login',function(req,res){
	res.sendFile(path.join(__dirname + '/index-manufacturer.html'));
});
app.get('/home',function(req,res){
	res.sendFile(path.join(__dirname + '/manufacturer.html'));
});
app.get('/watches',function(req,res){
	res.sendFile(path.join(__dirname + '/table-mfg.html'));
});
app.get('/sign-up',function(req,res){
	res.sendFile(path.join(__dirname + '/sign-up.html'));
});
app.get('/watch-dealer',function(req,res){
	res.sendFile(path.join(__dirname + '/watch-dealer.html'));
});
app.get('/watch-history',function(req,res){
	res.sendFile(path.join(__dirname + '/watch-history.html'));
});
app.get('/transfer',function(req,res){
	res.sendFile(path.join(__dirname + '/transfertoDealer.html'));
});
app.get('/insurer',function(req,res){
	res.sendFile(path.join(__dirname + '../insurer/insurer.html'));
});
app.listen(8080);