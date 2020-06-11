<template>
	<view>
		<view class="context">
			<view class="details-show-img">
				<view class="uni-padding-wrap">
				            <view class="page-section swiper">
				                <view class="page-section-spacing">
				                    <swiper class="swiper" :indicator-dots="indicatorDots" :autoplay="autoplay" :interval="interval" :duration="duration">
				                        <swiper-item style="height: 400upx;" v-for="(item,index) in product.details_img" :key="index">
				                           <view class="images-box">
											<image :src=" '/static/image/goodsimage/'+item"></image>
										   </view>
				                        </swiper-item>
				                    </swiper>
				                </view>
				            </view>
				</view>
			</view>
			
			
			<view class="title-box">
				<view class='title'>{{product.name}}</view>
				<view class='price'>¥{{product.price}}</view>
			</view>
			
			
			<view class="top-box">
				<view class="top-boxs">
					<view class="top-icon"><image src="/static/icon/fh.png" mode=""></image></view>
				    <view class="top-font">闪电发货</view>
				</view>
				<view class="top-boxs">
					<view class="top-icon"><image src="/static/icon/zp2.png" mode=""></image></view>
				    <view class="top-font">保障正品</view>
				</view>
				<view class="top-boxs">
					<view class="top-icon"><image src="/static/icon/id.png" mode=""></image></view>
				    <view class="top-font">唯一ID</view>
				</view>
				<view class="top-boxs">
					<view class="top-icon"><image src="/static/icon/sh.png" mode=""></image></view>
				    <view class="top-font">售后服务</view>
				</view>
				
			</view>
			
			<view class="looks-box" :style="{height:buyHeight +'upx'}">
				<view class="lately-box">
					<view class="lately-title">最近购买<text>(232)</text></view>
					<view class="lately-box-data" :style="buyHeight" v-for="(item,index) in buydata" :key="index">
						<view class="top">
							<view class="top-image">
								<view style="width: 70%;margin: auto;">
								 <image class="top-img" :src="item.usertop"></image><text style="padding-left:5upx;">{{item.username}}</text>
							    </view>
							</view>
							<view class="top-price">¥{{item.price}}</view>
							<view class="top-size">{{item.version}}</view>
						</view>				
					</view>
				</view>
			</view>
			
			
			<view class="looks-boxs" :style="{height:looksBoxsHeight + 'upx'}">
				<view class="lately-boxs">
					<view class="lately-title">潮流搭配</view>
					<view class="looks-box-data" v-for="(item,index) in laos" :key="index">
						<view class="context-box">
							<view class="left">
								<view class="left-image">
									<image class="left-images" :src="item.shopimg"></image>
								</view>
							</view>
							<view class="right">
								<view class="right-top">
									<view class="right-top-images">
										<view class="right-images-title">{{item.title}}</view>
										<view class="right-images-price">¥{{item.price}}</view>
									</view>
								</view>
								<view class="right-bottom">
									<view class="right-bottom-images">
									  <view class="images-buy" @click="shopcart(item.goodsid)">
										<image class="right-images" src="/static/icon/collect.png"></image>
									  </view>
									</view>
									<view class="right-bottom-images">
									  <view class="images-buy">
										<image class="right-images" src="/static/icon/shopcart.png"></image>
									  </view>
									</view>
								</view>
							</view>
						</view>
					</view>
				</view>
			</view>
			
			<view class="parameter-box" :style="{height:parameterHeight +'upx'}">
				<view class="parameter-box-data" :style="{height:parameterHeight + 'upx'}">
					<view class="parameter-box-title">参数</view>
					<view class="parameter-box-datas" v-for="(item,index) in chanel" :key="index">
						<view class="data-alls">
							<view class="data-alls-left">
								{{index}}
							</view>
							<view class="data-alls-right">
								{{item}}
							</view>
						</view>
						
					</view>
				</view>
			</view>
			
			
			<view class="referral-box">
				<view class="referral-box-title">单品详情</view>
				<view class="referral-box-images">
					<image :src="'/static/image/goodsimage/'+product.details_bottomimg"></image>
				</view>
				<view class="referral-box-font">
					{{product.details_bottomdata}}
				</view>
			</view>
			
			<view class="bottom-buy">
				<uni-popup ref="popupShare" type="share" @change="change">
					<uni-popup-share title="分享到" @select="select" :reciveUserInfo="info"></uni-popup-share>
				</uni-popup>
				<view class="example-body share">
					<button class="button" type="primary" @click="confirmShare"><text class="button-text">立即购买</text></button>
				</view>
			</view>
		
			<view class="ade"></view>
		</view>
	</view>
</template>

<script>
	
	import uniPopupShare from './shares.vue'
	import tools from '@/tools/ade.js'

	export default {
		components: {
			tools,
			uniPopupShare
		},
		data() {
			return {
				info:"",
				indicatorDots: true,
				autoplay: true,
				interval: 2000,
				duration: 500,
				openid:0,
				buyHeight:190,
				windHeight:0,
				looksBoxsHeight:400,
				parameterHeight:150,
				version:"",
				product:"",
				buydata:[	
					{
					  userid:1,
					  usertop:'/static/usertop/1.jpeg',
					  username:'阿德的ni',
					  price:'599',
					  version:'43码',
					},
					{
					  userid:2,
					  usertop:'/static/usertop/2.jpeg',
					  username:'bertni',
					  price:'399',
					  version:'38码',
					},
					{
					  userid:3,
					  usertop:'/static/usertop/3.jpg',
					  username:'阿阿德1213',
					  price:'1599',
					  version:'26码',
					},
					{
					  userid:4,
					  usertop:'/static/usertop/4.jpg',
					  username:'南苑bert',
					  price:'1399',
					  version:'37码',
					},
					{
					  userid:5,
					  usertop:'/static/usertop/5.jpg',
					  username:'太难了没不',
					  price:'1599',
					  version:'41码',
					}
				],
				laos:[
					{   
						goodsid:1,
						shopimg:'/static/image/goodsimage/kui2.jpg',
						title:'supreme破洞牛仔裤',
						price:'1599'
					},
					{
						goodsid:2,
						shopimg:'/static/image/goodsimage/wek2.jpg',
						title:'wekller白色T恤',
						price:'399'
					},
					{
						goodsid:3,
						shopimg:'/static/image/goodsimage/yufu.jpg',
						title:'supreme渔夫帽',
						price:'199'
					}
				],
				chanel:"",
			}
		},
		methods: {
			changeIndicatorDots(e) {
			            this.indicatorDots = !this.indicatorDots
			        },
			        changeAutoplay(e) {
			            this.autoplay = !this.autoplay
			        },
			        intervalChange(e) {
			            this.interval = e.target.value
			        },
			        durationChange(e) {
			            this.duration = e.target.value
			        },
					confirmShare() {
						this.$refs.popupShare.open()
					},
					change(e) {
				    console.log('popup ' + e.type + ' 状态', e.show)
			        },
					select(e, done) {
						uni.showToast({
							title: `您选择了第${e.index+1}项：${e.item.text}`,
							icon: 'none'
						})
						done()
					},
		},
		onLoad(){
			  let _this = this;
			uni.getStorage({
				key:"userdata",
				success(res) {
					_this.openid = JSON.parse(res.data).data
				},
				fail() {
					uni.navigateTo({
					    url: '/pages/login/login'
					});
				}
			})

		
		  let urls = tools.api + "/details";
		  let tokens = tools.en_bert(JSON.stringify({"query":"details","productid":1}));
		  console.log(JSON.stringify({"query":"details","productid":1}))
		  uni.request({
			url:urls,
			data:{"token":tokens,"usertoken":_this.openid,"pages":"details"},
			success:(res)=>{
				let version = res.data.version
				_this.product = res.data.data
				_this.product.details_img = JSON.parse(_this.product.details_img)

				 let r = [];
				 let rsone = [];

				 for(var i=0;i<version.length;i++){
					let e = i%3;
				    if(e<=2){
					rsone.push(version[i])
					 if(e==2){
						 r.push(rsone)
						 rsone = [];
					 }
					 if(i==version.length-1){
						 r.push(rsone)
					 }
					}
				 }

				_this.info =  r;
				console.log(_this.info);
				_this.chanel = JSON.parse(_this.product.details_jsondata)
				 _this.parameterHeight = _this.parameterHeight + (Object.keys(this.chanel).length * 60);

			}
		  })
		},
		mounted() {
			var _this = this;
			uni.getSystemInfo({
				success:function(res){
					_this.windHeight = res.windowHeight;
				}
			})
			
			if (this.buydata.length>1){
				this.buyHeight =  this.buyHeight + (this.buydata.length * 85) - 85;
			}
			    this.looksBoxsHeight = this.looksBoxsHeight + (this.laos.length * 230); 
				// console.log(Object.keys(this.chanel))
		
		}
	}
</script>

<style>
	.context{
		width: 95%;
		height: 500upx;
		margin: auto;
	}
	
	.details-show-img{
		height: 500upx;
	}
	.images-box{
		width: 90%;
		margin: auto;
		height:500upx;
	}
	
	.images-box image{
		width: 100%;
		height: 100%;
	}
	swiper{
		height: 500upx;
	}
	swiper-item{
		height: 500upx;
	}
	
	.title-box{
		width: 100%;
		height: 200upx;
	}
	
	.title-box .title,.price{
		width: 100%;
		height: 100upx;
		text-align: center;
	}
	
	.title-box .title{
		font-size: 40upx;
		line-height: 100upx;
		font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
	}
	
	.title-box .price{
		line-height: 100upx;
		font-size: 46upx;
	}
	
	.top-box{
		width: 100%;
		height:150upx;
		padding-top: 30upx;
		border-top: 1upx solid #333;
		border-bottom: 6upx solid #ccc;
	}
	.top-box .top-boxs{
		width:25%;
		float: left;
		height:150upx;
	}
	.top-icon{
		margin: auto;
		width: 50%;
		height: 50%;
	}
	
	.top-icon image{
		margin: auto;
		width: 90%;
		height: 90%;
	}
	
	.top-font{
		width:96%;
        height: 50%;
		line-height: 60upx;
		font-size: 25upx;
		font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
		text-align: center;
	}
	
	.looks-box{
		width: 100%;
		padding-top:20upx;
	}
	
	.looks-box .lately-box,.lately-boxs{
	   width: 94%;
	   margin: auto;
	   height: 85upx;
	}
	
	.lately-title,.lately-box-data{
		width: 100%;
		height: 100%;
		line-height: 85upx;
		font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
		
	}
	
	.lately-title{
		font-size: 30upx;
	}
	
	.lately-title text{
		font-size: 25upx;
		color:#2C405A;
	}
	
	.lately-box-data .top{
	   width:100%;
	   border-bottom: 1upx solid #CCCCCC;
	}
	
	.top-image,.top-price,.top-size{
		width:33%;
		float: left;
		height: 85upx;
		font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
		font-size: 20upx;
		color:#141821;
	}
	
	.top-image .top-img{
		width: 50upx;
		height: 50upx;
		border-radius: 50upx;
		vertical-align: middle;
	}
	
	.top-price,.top-size{
		color:#333333;
		font-size: 26upx;
		text-align: center;
		line-height: 85upx;
	}
	
	.looks-boxs{
		width: 100%;
	}
	
	.looks-boxs .looks-box-data{
		width:100%;
		height: 330upx;
	}
	
	.looks-boxs .looks-box-data .context-box{
		width: 100%;
		margin: auto;
		height: 280upx;
		border-radius:20upx;
		border: 2upx solid #C0C0C0;
	}
	
	.context-box .left,.right{
		float: left;
		height: 280upx;
	}
	
	.context-box .left{
		width:60%;
		/* border: 1upx solid red; */
	}
	
	.context-box .left .left-image{
		width: 80%;
		height:260upx;
		margin: auto;
		padding-top: 20upx;
	}
	
	.context-box .left .left-image .left-images{
		width: 83%;
		height: 85%;
	}
	
	.context-box .right{
		width:39%;
	}
	
	.context-box .right .right-bottom,.right-top{
		width: 100%;
		height: 30%;
	}
	
	.right-top{
		height: 70%;
	}
	
	.right-images-title{
		width: 80%;
		padding-top:28upx;
		margin: auto;
		height:80upx;
		font-size:27upx;
		color:#141821;
		
	}
	.right-images-price{
		text-align: center;
		height:85upx;
		font-size:30upx;
		line-height: 85upx;
	}
	
	.right-bottom-images{
		width: 49%;
		float: left;
		padding-top:15upx;
	}
	
	.right-images{
		width: 100%;
		height: 100%;
	}
	.images-buy{
		width:50upx;
		margin: auto;
		height: 50upx;
	}
	
	.parameter-box{
		width: 100%;

	}
	
	.parameter-box-title{
		width: 100%;
		height: 85upx;
		line-height: 85upx;
		font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
		font-size: 30upx;
	}
	
	.parameter-box .parameter-box-data{
		width: 94%;
		margin: auto; 
	 }
	
	.data-alls{
		width: 100%;
		height: 65upx;
	}
	
	.data-alls .data-alls-left,.data-alls-right{
		width: 40%;
		margin: auto;
		float: left;
		/* text-align: center; */
		height: 65upx;
		line-height: 65upx;
		font-size: 28upx;
		color:#3F536E;
	}
	.data-alls-right{
		width:60%;
		text-lndent:percentage;
		color:#646566;
		 text-align: right; 
	}
	
	.referral-box{
		width: 94%;
		margin: auto;
		/* border:1rpx solid red; */
	}
	
	.referral-box .referral-box-title{
		font-size: 20upx;
		width: 100%;
		height:85rpx;
		line-height: 85upx;
		font-family:'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
		font-size: 30upx;
	}
	
	.referral-box .referral-box-images{
		width:90%;
		margin: auto;
		height: 350upx;
	}
	
	.referral-box .referral-box-images image{
		width: 100%;
		height: 100%;
	}
	
	.referral-box .referral-box-font{
		font-size:26upx;
		font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
		color:#3B4144;
		letter-spacing:3upx;
	}
	
	.bottom-buy{
		width: 95%;
		margin: auto;
		position: fixed; 
		bottom: 0;
		 /* border: 1rpx solid red; */
		height: 110upx;
	}
	
	.ade{
	    width: 100%;
		height: 120upx; 
	}
	
	.button{
		background-color:green;
	}
</style>
