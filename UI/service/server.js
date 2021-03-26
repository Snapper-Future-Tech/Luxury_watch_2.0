var express= require('express');
var app = express();
var path = require('path');
app.use(express.static(path.join(__dirname,'public')));
app.get('/',function(req,res){
	res.sendFile(path.join(__dirname + '/index-service.html'));
});
app.get('/home',function(req,res){
	res.sendFile(path.join(__dirname + '/service-home.html'));
});
app.get('/watch-history',function(req,res){
	res.sendFile(path.join(__dirname + '/watch-history.html'));
});
app.get('/update',function(req,res){
	res.sendFile(path.join(__dirname + '/update-service.html'));
});
app.get('/insurer',function(req,res){
	res.sendFile(path.join(__dirname + '../insurer/insurer.html'));
});
app.listen(8082);