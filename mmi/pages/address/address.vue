<template>
	<view>
		<view class="context">
			<view class="address-buy">
				<view class="address-buy-top">
					<view class="left">你好</view>
					<view class="right">13975515540</view>
				</view>	
				<view class="address-buy-detail">
						我说的日方提供用户就会给发短信的日方提供一份多少
					
				</view>
				<view class="address-buy-bottom">
						<view class="bottom-left">
							<view class="bottom-left-buy">
								<image src="/static/icon/defaults.png"></image>
                                 设为默认
							</view>
						</view>
						<view class="bottom-right">
							<view class="rights-left" @click="putaddress(1)" style="color:#3B4144;">
								编辑
							</view>
							<view class="rights-right" style="color:#008000;">
								删除
							</view>
						</view>
					</text>
				</view>
			</view>
		</view>
	     <view class="context">
	     	<view class="address-buy">
	     		<view class="address-buy-top">
	     			<view class="left">你好</view>
	     			<view class="right">13975515540</view>
	     		</view>	
	     		<view class="address-buy-detail">
	     				我说的日方提供用户就会给发短信的日方提供一份多少
	     			
	     		</view>
	     		<view class="address-buy-bottom">
	     				<view class="bottom-left">
	     					<view class="bottom-left-buy">
	     						<image src="/static/icon/default.png"></image>
	                              设为默认
	     					</view>
	     				</view>
	     				<view class="bottom-right">
	     					<view class="rights-left" @click="putaddress(2)" style="color:#3B4144;">
	     						编辑
	     					</view>
	     					<view class="rights-right" style="color:#008000;">
	     						删除
	     					</view>
	     				</view>
	     			</text>
	     		</view>
	     	</view>
	     </view>
	     	
	
	</view>
</template>

<script>
	export default {
		data() {
			return {
				
			}
		},
		methods: {
			putaddress:(e)=>{
			  uni.navigateTo({
			  	url:'putaddress?id='+e
			  })
			}
		}
	}
</script>

<style>
 .context{
	 width: 100%;
	 padding-top:30upx;
	 /* margin: auto; */
	 height: 400upx;
	 /* border: 1upx solid red; */
 }
 .address-buy{
	 width:94%;
	 margin: auto;
	 height: 370upx;
	 border:1rpx solid #00CD98 ;
	 border-radius: 10upx;
 }
 
 .address-buy-top{
	 /* padding: 30rpx; */
	 
	width: 100%;
	height: 80upx;
 }
 
 .address-buy-top .left,.right{
   width:45%;
   height: 80upx;
   float: left;
    /* border: 1rpx solid red; */
   line-height: 80upx;
   font-size:35upx;
   padding-left: 20upx;
   /* text-align: center; */
   color:#141821;
 }
 
 .address-buy-top .right{
	 padding-right:10upx ;
	 text-align: right;
	  /* border: ; */
 }
 
 .address-buy-detail{
	 width:94%;
	 margin: auto;
	 font-size: 30upx;
	 height: 150upx;
	 font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
	 border-bottom: 1rpx solid #CCCCCC;
 }
 
 .address-buy-detail text{
	 font-size: 17upx;
	 color:#333333;
 }
 
 .address-buy-bottom{
	 width: 100%;
	 padding-top: 20upx;
	 height: 80upx;
	 /* border-top:; */
	 /* border: 3rpx solid #000000; */
 }
 .address-buy-bottom .bottom-left,.bottom-right{
	 width: 48%;
	 height: 100%;
	 /* background-color: 10upx solid #007AFF; */
	 float: left;
	 /* border: 5upx solid #007AFF; */
 }
 
 .bottom-left-buy{
	 width: 300upx;
	 margin: auto;
	 height: 80upx;
	 line-height: 80upx; 
	 font-size: 30upx;
	 /* border: 1rpx solid red; */
 }
 
 .bottom-left-buy image{
	 width:70upx;
	 vertical-align: middle;
	 height: 70upx;
 }
 .bottom-right .rights-right,.rights-left{
	 width:40%;
	 float: left;
	 height: 80upx;
	 font-size: 30upx;
	 line-height: 80upx;
	 text-align:right;
 }
</style>
