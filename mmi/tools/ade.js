 import CryptoJS from "./crypto.js"
 
 var key1 = "linkbook1qaz*WSX";
 // var str = {
 //       name: "菜鸟教程11",
 //       site: "http://www.runoob.com"
 //     }
//var plaintText = JSON.stringify(str)
 
 function encodeAes(plaintText) {
var key = CryptoJS.enc.Utf8.parse(key1);
      var encryptedData = CryptoJS.AES.encrypt(plaintText, key, {
        iv: key,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7
      });
      encryptedData = encryptedData.ciphertext.toString();
      // console.log("加密后-no-hex：" + encryptedData);
      return encryptedData
    }
	
 function decodeAes(encryptedDataStr) {
      var key = CryptoJS.enc.Utf8.parse(key1);
      var encryptedHexStr = CryptoJS.enc.Hex.parse(encryptedDataStr);
      // console.log("解密前hex：" + encryptedHexStr);
      var encryptedBase64Str = CryptoJS.enc.Base64.stringify(encryptedHexStr);
      // console.log("解密前：" + encryptedBase64Str);

      var decryptedData = CryptoJS.AES.decrypt(encryptedBase64Str, key, {
        iv: key,
        mode: CryptoJS.mode.CBC,
        padding: CryptoJS.pad.Pkcs7
      });

      var decryptedStr = decryptedData.toString(CryptoJS.enc.Utf8);
      // console.log("解密后:" + decryptedStr);
       return decryptedStr
    }
	
	function GenderImg(u){
		if(u==1){
			return "/static/icon/man.png"
		}
		    return "/static/icon/woman.png"
	}
	
	function dataAlls(str,num){
	
		let date = new Date();
		
		let y = date.getFullYear();
		let MM = date.getMonth() + 1;
		MM = MM < 10 ? ('0' + MM) : MM;//月补0
		let d = date.getDate();
		d = d < 10 ? ('0' + d) : d;//天补0
		let h = date.getHours();
		h = h < 10 ? ('0' + h) : h;//小时补0
		let m = date.getMinutes();
		m = m < 10 ? ('0' + m) : m;//分钟补0
		let s = date.getSeconds();
		s = s < 10 ? ('0' + s) : s;//秒补0
		switch(str){
			case "*":
			return y + '-' + MM + '-' + d;
			break;
			case "y":
			return y;
			break;
			case "m":
			return MM;
			break;
			case "d":
			return d;
		}
	}
	

function formatDate(timeStamp,str) { 
    var date = new Date();
    date.setTime(timeStamp * 1000);
    var y = date.getFullYear();    
    var m = date.getMonth() + 1;    
    m = m < 10 ? ('0' + m) : m;    
    var d = date.getDate();    
    d = d < 10 ? ('0' + d) : d;    
    var h = date.getHours();  
    h = h < 10 ? ('0' + h) : h;  
    var minute = date.getMinutes();  
    var second = date.getSeconds();  
    minute = minute < 10 ? ('0' + minute) : minute;    
    second = second < 10 ? ('0' + second) : second; 
	  switch(str){
		  case "*":
			return y + '-' + m + '-' + d+' '+h+':'+minute+':'+second;  
		  break;
		  case "y":
		   return y;
		  break;
		  case "m":
		   return m;
		  break;
		  case "d":
		   return d;
	  }
    // return y + '-' + m + '-' + d+' '+h+':'+minute+':'+second;    
}; 



// 	function public look 
module.exports = {
	en_bert:encodeAes,
	de_bert:decodeAes,
	gender_img:GenderImg,
	timetoday:dataAlls,
	formatDates:formatDate,
	api:'http://127.0.0.1:8080'
}