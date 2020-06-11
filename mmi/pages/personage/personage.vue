<template>
	<view>
		<view class="context">
			<view class="context-top">
				<view class="context-top-left">
					<view class="context-top-left-image" @click="putImage()">
						<image :src="'/static/usertop/'+userpersonage.topimg"></image>
					</view>
				</view>
				<view class="context-top-right">
					<view class="context-top-right-buy">
						<view class="left">{{userpersonage.name}}</view>
						<view class="right">
						<image :src="usergender"></image>
						 </view>
					</view>
				</view>
			</view>
			<uni-list>
			    <uni-list-item title="资料" @click="u(1)" thumb="/static/icon/userdata.png"></uni-list-item>
			    <uni-list-item title="购物" thumb="/static/icon/usershopcart.png"></uni-list-item>
				<uni-list-item title="订单" thumb="/static/icon/buy.png"></uni-list-item>
			    <uni-list-item title="收藏" thumb="/static/icon/collecting.png"></uni-list-item>
				<uni-list-item title="收货" thumb="/static/icon/address.png"></uni-list-item>
				<uni-list-item title="消息" thumb="/static/icon/message.png"></uni-list-item>
			    <uni-list-item title="客服" thumb="/static/icon/service.png"></uni-list-item>
			</uni-list>
			
			<view class="context-bottom">
				<button type="default">退出登录</button>
			</view>
		</view>
	</view>
</template>

<script>
	import tools from '@/tools/ade.js'
	export default {
		data() {
			return {
				userpersonage:"",
				usergender:"",
			}
		},
		onLoad() {
		// uni.clearStorage();
			var _this = this;
			uni.getStorage({
				key:"userdata",
				success:(res)=>{
					let obj = JSON.parse(res.data);
					let urls = tools.api + "/personage";
					uni.getStorage({
						key:"personage",
						success:(res)=>{
							_this.userpersonage = JSON.parse(res.data);
				        
						   _this.usergender =  tools.gender_img(_this.userpersonage.gender)
						   // console.log()
						},
						fail() {
							uni.request({
								url:urls,
								data:{
									query:"personage",
									usertoken:obj.data,
									openid:{"Id":obj.data},
									pages:"personage",
								},
								success:(res)=>{
								let data = JSON.parse(tools.de_bert(res.data.data))
								  if(res.data.status==200){
								  	 uni.setStorage({
										 key: 'personage',
										 data: JSON.stringify(data),
									 })

									_this.userpersonage = data;
									_this.usergender =  tools.gender_img(data.gender)
								  }
								}
							})
							
						}
					})
					
					
				},
				fail() {
					// 没有openid
					uni.navigateTo({
					    url: '/pages/login/login'
					});
				}
			})
			// console.log(this.userpersonage)
		},
		methods: {
			putImage:function(){
				uni.navigateTo({
				    url: 'userimage'
				});
			},
			u:(e)=>{
				switch(e){
					case 1:
					uni.navigateTo({
					    url: '/pages/personage/userdata'
					});
					break
				}
			}
		}
	}
</script>

<style>
.context{
	width: 94%;
	height: 1200upx;
	margin: auto;
	 /* border: 1rpx solid red; */
}
.context-top{
	width: 100%;
	padding-top: 20upx;
	height: 170upx;
	border-bottom: 3rpx solid #C0C0C0;
}
.context-top .context-top-left{
	width: 40%;
	height: 140upx;
	float: left;
	padding-top:9upx ;
	/* font-size: ; */
	 /* border: 1upx solid #007AFF; */
	/* vertical-align: middle; */
}
.context-top .context-top-left .context-top-left-image{
	width: 55%;
	height: 150upx;
	margin: auto;
}
.context-top .context-top-left .context-top-left-image image{
	width: 90%;
	/* padding:10; */
	border-radius: 50%;
	height:90%;
}

.context-top .context-top-right{
	width: 55%;
	float: right;
	/* border: 1rpx solid red; */
	height:150upx;
}

.context-top .context-top-right .context-top-right-buy{
	width: 100%;
	padding-top:20upx;
	font-size: 40upx;
	height: 130upx;
}
.context-top .context-top-right .context-top-right-buy .left,.right{
	width: 80%;
	float: right;
	/* font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif; */
	font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
	/* border: 1upx solid red; */
	height: 50%;
}
.context-top-right-buy .right image{
	width: 40upx;
	height: 40upx;
}
.context-bottom{
	width: 100%;
	padding-top: 60upx;
	height: 100upx;
	/* border: 1rpx solid border-radius; */
}
</style>
