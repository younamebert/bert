<template>
	<view class="uni-popup-share">
		<!-- <view class="uni-share-title"><text class="uni-share-title-text">{{title}}</text></view> -->
		<view class="uni-share-content">
			<view class="uni-share-content-box">
				
				<view class="size-buy">
					<view class="size-top">
						<view class="size-top-images">
							<image :src=" '/static/image/goodsimage/'+reciveUserInfo[0][0].productImg"></image>
						</view>
						<view class="size-top-details">
							<view class="right-price">
							  ¥{{price}}
							</view>
							
							<view class="right-num">
							  {{num}}件
							</view>
					    </view>
					</view>
					
					<view class="uni-padding-wrap uni-common-mt">
						<view>
							<scroll-view :scroll-top="scrollTop" scroll-y="true" class="scroll-Y" @scrolltoupper="upper" @scrolltolower="lower"
							@scroll="scroll">
								<view class="size-list" >
										<view class="size-list-alls" v-for="(item,index) in reciveUserInfo" :key="index">
										 <view class="size-single" @click="redio(index,indexs)" v-for="(items,indexs) in item" :key="indexs">
											 <view class="configuration">{{items.attr}}</view>
											 <view class="price">¥{{items.price}}</view>
										 </view>
										</view>
								</view>
							</scroll-view>
						</view>
					</view>					
				</view>
			</view>
		</view>
		<view class="uni-share-button-box">
			<button class="uni-share-button button-bert">确认购买</button>
		</view>
		<view class="uni-share-button-box">
			<button class="uni-share-button" @click="close">取消</button>
		</view>
	</view>
</template>

<script>
	// import details from './details.vue'
	export default {
		name: 'UniPopupShare',
		props: ["reciveUserInfo"],
		
		thisbert:0,
		inject: ['popup'],
		methons:{
			select(item,index){
				this.$emit('select',{
					item,
					index,
				},()=>{
					this.popup.close()
				})
			}
		},
		close(){
			this.popup.close()
		},
		data() {
			return {
				scrollTop:0,
				 num:'--',
				 price:'--',
				 version:'',
				 img:'',
				 versionData: [
					 {
						goods:[
							{
								'size':'36',
								'price':'399',
								'num':122
							},
							{
								'size':'37',
								'price':'499',
								'num':22
							},
							{
								'size':'38',
								'price':'599',
								'num':563
							},
							{
								'size':'41.5',
								'price':'1299',
								'num':12
							}
						]
					},
					{
						goods:[
							{
								'size':'42',
								'price':'929',
								"num":12
							},
							{
								'size':'43',
								'price':'799',
								"num":1
							},
							{
								'size':'44',
								'price':'699',
								"num":2
							},
							{
								'size':'44.5',
								'price':'1599',
								"num":2
							}
						]
					}
					
				]
			}
		},
		onLoad() {
			this.img = this.reciveUserInfo[0][0].productImg
		},
		created() {},
		methods: {
			/**
			 * 选择内容
			 */
			select(item, index) {
				this.$emit('select', {
					item,
					index
				}, () => {
					this.popup.close()
				})
			},
			/**
			 * 关闭窗口
			 */
			close() {
				this.popup.close()
			},
			redio(e,es){
			  this.price = this.reciveUserInfo[e][es].price;
			  this.num = this.reciveUserInfo[e][es].num;
			}
		}
	}
</script>
<style lang="scss" scoped>
	.uni-popup-share {
		background-color: #fff;
	}
	.uni-share-title {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		align-items: center;
		justify-content: center;
		height: 40px;
	}
	.uni-share-title-text {
		font-size: 14px;
		color: #666;
	}
	.uni-share-content {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		justify-content: center;
		padding-top: 10px;
	}
	
	.uni-share-content-box {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		flex-wrap: wrap;
		width: 94%;
		margin: auto;
	}
	
	.uni-share-content-item {
		width: 90px;
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: column;
		justify-content: center;
		padding: 10px 0;
		align-items: center;
	}
	
	.uni-share-content-item:active {
		background-color: #f5f5f5;
	}
	
	.uni-share-image {
		width: 30px;
		height: 30px;
	}
	
	.uni-share-text {
		margin-top: 10px;
		font-size: 14px;
		color: #3B4144;
	}
	
	.uni-share-button-box {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		padding: 10px 15px;
	}
	
	.uni-share-button {
		flex: 1;
		border-radius: 50px;
		color: #666;
		font-size: 16px;
	}
	
	.uni-share-button::after {
		border-radius: 50px;
	}
	
	.size-buy{
		width: 100%;
		height:710upx;
		// border: 1upx solid red;
	}
	.scroll-Y{
		height:495upx;
		// bro
		// border: 1rpx solid blue;
	}
	.size-buy .size-top{
		width: 100%;
		height:30%;
		border-bottom: 1upx solid #C0C0C0;
	}
	.size-top .size-top-images{
		width:30%;
		height: 100%;
		float: left;
	}
	
	.size-top .size-top-images image{
		width: 100%;
		height: 100%;
	}
	
	.size-top .size-top-details{
		float:left;
		width: 40%;
		height: 100%;
	}
	 .size-top-details .right-num,.right-price{
		width: 100%;
		height:20%;
		padding-top: 40upx;
		text-align: right;
		// line-height: 300%;
		// border: 1upx solid red;
	}
	.size-top-details .right-num{
		font-size: 30upx;
		color: #C0C0C0;
	}
	
	.size-list{
		padding-top:30upx ;
		width: 100%;
		height: 160upx;
		margin: auto;
	}
	
	.size-list-alls{
		width: 100%;
		height:165upx;
		justify-content:center;
		display: flex;
		float: left;
	}
	
	.size-list-alls .size-single{
		width:30%;
		height: 140upx;
		float: left;
		margin-right:30upx;
		// margin-left:30upx;
		// border: 5upx solid #141821;
	}
	
	.size-single-buy{
		// border: 1upx solid red;
		border: 5upx solid #141821;
	}
	.size-list-alls .size-single:first-child{
		margin-left:30upx;
	}
	
	.size-single .configuration,.price{
		width: 100%;
		height: 50%;
		text-align:center;
		line-height: 200%;
	}
	.size-single .price{
		color:#CCCCCC;
		font-size: 30upx;
	}
	.button-bert{
		background-color:#2C405A;
		color:#FFFFFF;
	}
	
</style>