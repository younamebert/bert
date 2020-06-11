<template>
	<!-- header("Access-Control-Allow-Origin: *"); -->
	<!-- <meta http-equiv="Access-Control-Allow-Origin" content="*"> -->
  <view class="context"> 
  
  <!-- <view></view> -->
    <view class="context-top">login</view>
	
    <view class="context-form">
            <view>
                <form @submit="formSubmit" @reset="formReset">
                   
                    <view class="uni-form-item uni-column uni-padding" style="border-bottom: 1px solid #3B4144;">
                        <input class="uni-input" type="text" name="name"  placeholder="邮箱/电话/用户名~" />
                    </view>
					
					
					<view class="uni-form-item uni-column uni-padding" style="border-bottom: 1px solid #3B4144;">
					    <input class="uni-input"  type="password" name="password" placeholder="请输入密码~" />
					</view>
					
                    <view class="uni-btn-v uni-padding">
                        <button form-type="submit">登录</button>
						<br/>
						
                        <uni-popup ref="popup" type="top">123456</uni-popup>
						<p class="rest-btnfunc">
							<uni-link :href="item.qq" :text="item.text" v-for="(item,index) in bottomHref" :key="index"></uni-link>
						</p>
                    </view>
                </form>
            </view>
        </view>
  </view>
  
</template>

<script>
import tools from '@/tools/ade.js'	
import uniPort from '@/components/uni-popup/uni-popup.vue'
export default {
	components: {
		uniPort,
	    tools,
	},
		data() {
			return {
				bottomHref:[
					{
						qq:"/qq",
						text:"QQ登录"
					},
					{
						wx:"/wx",
						text:"微信登录"
					}
				]
			}
		},
		methods: {
           formSubmit:(e)=>{
			   
			
			 let obj = e.detail.value;
			 try{
			 	// 字段是否存在
			 	if(!('name' in obj) && !('password' in obj)){
			 		throw "字段不存在"
			 	}
				
				let common = /[A-Z0-9a-z]{5,}/
				var f = Object.values(obj).forEach((val,index)=>{
					if(val.length==0){
						throw "不能为空"
						return
					}
					if(common.test(val)==false){
                       	throw "密码或者用户长度不够"
						return false
					}
				 })

				 let urls = tools.api + "/login";
				 let tokens = tools.en_bert(JSON.stringify({'Name':obj.name,'Password':obj.password}))
				 console.log(tokens)
				 
				
			  uni.request({
			  	url:urls,
				data:{
					token:tokens,
					query:"login",
				},
				dataType:'multipart/form-data',
				headers: {
				      'Accept':'multipart/form-data',
				      'contentType':'Application/json'
				  },
				success:(res)=>{
				 let data = JSON.parse(res.data);
				 
				    console.log(data)
				     if(data.status==200){
						// console.log(tools.de_bert(data.data))
						 uni.setStorage({
						     key: 'userdata',
						     data: JSON.stringify(data),
						 })
						 
						uni.navigateTo({
						    url: '/pages/home'
						});
						 
					 }
				}
			  })
				 
			 }catch(e){
				uni.showToast({
					title: e,
					position:'center',
					duration: 2000,
					image:"/static/icon/out.png"
				});
			 } 
		   }
		}
	}
</script>

<style >
	.context{
		width: 90%;
		padding-top:50upx;
		/* height; */
		/* border: 1px solid red; */
		margin: auto;
		/* border: 1px solid red; */
	}
	
	.context .context-top{
		 height: 120upx;
		text-align: center;
		/* border: ; */
		line-height: 120upx;
		/* font-size: ; */
		border-bottom: 3upx solid blueviolet;
	}
	.context-text{
		/* height: ; */
		width: 100%;
		
		/* border: 1px solid red; */
	}
	
	.context-text-form{
		width: 100%;
		/* border: 1px solid red; */
	}
	
	#inputs-title{
	/* font-family: ; */
	/* color:#007AFF; */
	/* font-family: "楷体"; */
	color:#333333;
	}
	
	.context-form{
		margin-top: 10upx;
		
		/* border: 1upx solid red; */
		/* width:98%; */
		/* margin: auto; */
	}
	
	.uni-padding{
		
		padding-top: 200upx;
		height: 40upx;
		/* border-bottom:1px solid red; */
	}
	
	.rest-btnfunc{
	  /* margin:; */
	  padding-top: 40upx;
	}
</style>
