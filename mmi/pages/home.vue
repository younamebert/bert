<template>
	<view>
	
		 <view class="index-top"><i>bert小店</i></view>
		 
		
		 <view class="context">   
		  
		    <view class="index-image">
           <view>
                   <view class="uni-padding-wrap">
                       <view class="page-section swiper">
                           <view class="page-section-spacing">
                               <swiper class="swiper" :indicator-dots="indicatorDots" :autoplay="autoplay" :interval="interval" :duration="duration">
                                   <swiper-item v-for="(item,index) in image" :key="index">
                                       <image style="width: 100%;height:100%;" :src="item.imgurl" mode="aspectFill"></image>
                                   </swiper-item>    
                               </swiper>
                           </view>
                       </view>
                   </view>
               </view>
		   </view>
		 </view>
		 
		 <view class="context-top">
			 <view class="context-icon">
				  <view class="top-float" v-for="(item,index) in topicon" :key='index'>
					  <view class="top-float-temp">
						<image class="icon" :src="item.icon" mode="">{{item.text}}</image>
					  </view>
				  </view>
			 </view>
		 </view>
		 
		
		 
		 
		 <view class="context-class">
			<view class="context-class-data" v-for="(item,index) in topclass" :key="index" v-if="index<4">
				<view class="images"><image :src="item.image"></image></view>
			    <view class="images-test">{{item.text}}</view>
			</view>
		 </view>
		 
		 
		 <view class="context-class" style="border-bottom:5upx solid #C0C0C0;">
		 			<view class="context-class-data" v-for="(item,index) in topclass.slice(3)" :key="index">
		 				<view class="images"><image :src="item.image"></image></view>
		 			    <view class="images-test">{{item.text}}</view>
		 			</view>
		 </view>
		 
		 
		 
		 
		 
		 <view class="context-goods">
			
		 		<view class="goods-box">
					<image class="goods-images" src="/static/image/3s.jpeg"></image>
				</view>
				
				
				<view class="goods-box-bottom">
					<view class="goods-title">DG 超级潮牌AJ COKRIS bert</view>
					<view class="goods-message">
						<view class="goods-price">¥399</view>
						<view class="goods-buysum">200人付款</view>
					</view>
				</view>
		 </view>
		 
		 
		 <view class="context-goods">
		
		 		<view class="goods-box">
		 					<image class="goods-images" src="/static/image/2s.jpeg"></image>
		 				</view>
		 				
		 				
		 				<view class="goods-box-bottom">
		 					<view class="goods-title">DG 超级潮牌衣服单品bert</view>
		 					<view class="goods-message">
		 						<view class="goods-price">¥399</view>
		 						<view class="goods-buysum">200人付款</view>
		 					</view>
		 				</view>
		 </view>
       <view class="ades"></view>
		 <view class="context-bottom">
			
			<view class="example-body">
				<view class="word-btn draw-cotrol-btn" hover-class="word-btn--hover" :hover-start-time="20" :hover-stay-time="70" @click="showDrawer('showLeft')"><text class="word-btn-white">菜单栏</text></view>
				<uni-drawer ref="showLeft" mode="left" :width="320" @change="change($event,'showLeft')">
					<!-- #ifndef MP-BAIDU || MP-ALIPAY || MP-TOUTIAO -->
					<uni-list>
						<uni-list-item title="搜索" @tap="jumpToDetails"/>
						<uni-list-item title="分类"/>
						<uni-list-item title="每日限购" :show-badge="true" badge-text="12" />
						<uni-list-item title="购物车"/>
						<uni-list-item title="个人"/>
					</uni-list>
					<!-- #endif -->
					<!-- #ifdef MP-BAIDU || MP-ALIPAY || MP-TOUTIAO -->
					<view class="uni-list">
						<uni-list-item title="搜索" @tap="jumpToDetails"/>
						<uni-list-item title="分类" />
						<uni-list-item title="每日限购" :show-badge="true" badge-text="12"/>
						<uni-list-item title="购物车"/>
						<uni-list-item title="个人" />
					</view>
					<!-- #endif -->
					<view class="close">
						<view class="word-btn" hover-class="word-btn--hover" :hover-start-time="20" :hover-stay-time="70" @click="closeDrawer('showLeft')"><text class="word-btn-white">关闭菜单</text></view>
					</view>
				</uni-drawer>
			</view>
			
		</view>
	
	
	</view>
</template>
<script>
	export default {
		data() {
			return {
				showRight: false,
				showLeft: false,
				imagekey:0,
				image:[
					{
					  colorClass:"uni-bg-blue",
					  imgurl:"/static/image/1.jpg"
					},
					{
					  colorClass:"uni-bg-blue",
					  imgurl:'/static/image/2.jpg'
					}
				],
				background: ['color1', 'color2', 'color3'],
				indicatorDots: true,
				autoplay: true,
				interval: 2000,
				duration: 500,
				topicon:[
						{
						 icon:"/static/icon/by.png",
						 text:"全场包邮"
					    },
						{
						 icon:"/static/icon/zp.png",
						 text:"正品保障"
						},
						{
						 icon:"/static/icon/yz.png",
						 text:"随时可证"
						}
				],
				topclass:[
					   {
						  image:"/static/image/2s.jpeg",
						  text:"品牌衣服",
					   },
					   {
						  image:"/static/image/dp.jpg",
						  text:"潮流单品",
					   },
					   {
						  image:"/static/image/1s.jpg",
						  text:"电子数码",
					   },
					   {
						  image:"/static/image/3s.jpeg",
						  text:"aj",
					   },
					   {
						  image:"/static/image/3s.jpeg",
						  text:"aj",
					   },
					   {
						  image:"/static/image/3s.jpeg",
						  text:"aj",
					   },
					   {
						  image:"/static/image/alls.png",
						  text:"全部",
					   }
					   
				],
				
				
			}
		},
		onLoad() {
		     
		},
		methods: {
			confirm() {},
			// 打开窗口
			showDrawer(e) {
				this.$refs[e].open()
			},
			// 关闭窗口
			closeDrawer(e) {
				this.$refs[e].close()
			},
			// 抽屉状态发生变化触发
			change(e, type) {
				console.log((type === 'showLeft' ? '左窗口' : '右窗口') + (e ? '打开' : '关闭'));
				this[type] = e
			}
		},
		onNavigationBarButtonTap(e) {
			if (this.showLeft) {
				this.$refs.showLeft.close()
			} else {
				this.$refs.showLeft.open()
			}
		},
		// app端拦截返回事件 ，仅app端生效
		onBackPress() {
			if (this.showRight || this.showLeft) {
				this.$refs.showLeft.close()
				this.$refs.showRight.close()
				return true
			}      
		},
		 jumpToDetails(e){
		                console.log(e)
		                // var newsid = e.currentTarget.dataset.newsid;
		               
		            }
		
	}
</script>

<style>
	/* 头条小程序组件内不能引入字体 */
	/* #ifdef MP-TOUTIAO */
	@font-face {
		font-family: uniicons;
		font-weight: normal;
		font-style: normal;
		src: url('~@/static/uni.ttf') format('truetype');
	}

	/* #endif */

	/* #ifndef APP-NVUE */
	page {
		display: flex;
		flex-direction: column;
		box-sizing: border-box;
		background-color: #fff;
		min-height: 100%;
		height: auto;
	}

	view {
		font-size: 14px;
		line-height: inherit;
	}

	.example {
		padding: 0 15px 15px;
	}

	.example-body {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		flex-wrap: wrap;
		justify-content: center;
		padding: 0;
		font-size: 14px;
		background-color: #ffffff;
		overflow: hidden;
	}

	/* #endif */
	.example {
		padding: 0 15px;
	}

	.example-info {
		/* #ifndef APP-NVUE */
		display: block;
		/* #endif */
		padding: 15px;
		color: #3b4144;
		background-color: #ffffff;
		font-size: 14px;
		line-height: 20px;
	}

	.example-info-text {
		font-size: 14px;
		line-height: 20px;
		color: #3b4144;
	}


	.example-body {
		flex-direction: column;
		padding: 15px;
		background-color: #ffffff;
	}

	.word-btn-white {
		font-size: 18px;
		color: #FFFFFF;
	}

	.word-btn {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		align-items: center;
		justify-content: center;
		border-radius: 6px;
		height: 48px;
		margin: 15px;
		background-color: #007AFF;
	}

	.word-btn--hover {
		background-color: #4ca2ff;
	}



	.uni-drawer-info {
		background-color: #ffffff;
		padding: 15px;
		padding-top: 5px;
		color: #3b4144;
	}

	.uni-padding-wrap {
		padding: 0 15px;
		line-height: 1.8;
	}

	.input {
		flex: 1;
		padding: 0 5px;
		height: 24px;
		line-height: 24px;
		font-size: 14px;
		background-color: transparent;
	}

	.close {
		padding: 15px;
	}

	.example-body {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		padding: 0;
	}

	.draw-cotrol-btn {
		flex: 1;
	}

	.info {
		padding: 15px;
		color: #666;
	}

	.info-text {
		font-size: 14px;
		color: #666;
	}

	.scroll-view {
		flex: 1
	}

	/* 处理抽屉内容滚动 */
	.scroll-view-box {
		flex: 1;
		position: absolute;
		top: 0;
		right: 0;
		bottom: 0;
		left: 0;
	}

	.info-content {
		padding: 5px 15px;
	}
	
	/* .context{
		border: 1px solid red;
		
		height: 100%;
	} */
	
	.context-bottom{
		  
		　　　left: 0;
		        position: fixed; 
		        bottom: 0; 
		        width: 100%; 
	
	}
	
	.index-top{
		width: 100%;
		height: 90upx;
		/* border: 1px solid red; */
		text-align: center;
        line-height:90upx ;
		font-size: 40upx;
		 color: #2C405A;
		 
		/* font-size: ; */
		background-color: #efeff4;
		font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
	}
	
	.index-image{
		width: 100%;
		margin: auto;
	 	/* height:1px; */
	}
	
	.index-image .index-image-alls{
		width: 100%;
		/* height:100%; */
		/* border-radius: 6upx; */
	}

	
   .context{
	 width: 100%;
	 /* height:333upx; */
	 /* border: 1upx solid red; */
	 /* border: 1upx solid red; */
   }
   .context-top{
	   width: 100%;
	   height: 100upx;
	   border-bottom: 3upx solid #C0C0C0;
	   
   }
   .context-icon .icon{
	   width:50upx;
	   height: 50upx;
	   vertical-align: middle;
   }
   
   .top-float{
	   float: left;
	   /* border: 1upx solid blue; */
	   /* height: auto; */
	   width: 33%;
	   height: 100upx;
	   
	   
   }
   
   .top-float .top-float-temp{
	   width: 80%;
	   margin:auto;
	   text-align: center;
	   padding-top: 30upx;
	   /* border: 1upx solid red; */
	   height: 64upx;
	   
	   /* padding-left:50%; */
	   line-height: 55upx;
   }
   
   .top-float .top-float-temp .icon{
   	   float: left;
	   /* margin: auto; */
   }
   
   .context-class{
	   width: 100%;
	   height: 180upx;
	   /* border: 1upx solid red; */
	   }
   
   .context-class .context-class-data{
	   width: 25%;
	   height: 140upx;
	   /* border: 1px solid red; */
	   float: left;
   }
   .context-class-data .images{
	   width:80upx;
	   padding-top:20upx;
	   height: 80upx;
	   /* border: 1px solid red; */
	   margin: auto;
	   
	   
   }
   .context-class-data .images image{
	   width: 100%;
	   height: 100%;
   }
   
   
   .context-class-data .images-test{
	  width: 100%;
	  height:55upx;
	  line-height: 55upx;
	  /* border: 1upx solid red; */
	  
	  text-align: center;
  }
  
  /* goods show */
  .context-goods{
	  width:90%;
	  padding-top: 80upx;
	  height: 580upx;
	  margin: auto;
	  border-bottom: 10upx solid #C0C0C0;
  }
  
  .context-goods .goods-box .goods-images{
	  width: 100%;
	  /* height: 100%; */
	  height: 100%;
  }
  .context-goods .goods-box{
	  width:90%;
	  height:400upx;
	  /* border: 1upx solid red; */
	  margin: auto;
	  /* border-bottom: 2upx solid #2C405A; */
  }
  .context-goods .goods-box-bottom{
	  width:90%;
	  height:170upx;
	  margin: auto;
	  /* border: 1upx solid red; */
  }
  
  .goods-box-bottom .goods-title{
	  width: 100%;
	  height: 80upx;
	  text-align: center;
	  color:#3C3E49;
	  line-height: 80upx;
	  font-family: 'Courier New', Courier, monospace;
	  /* border-bottom: 1upx solid #C0C0C0; */
	  font-size: 34upx;
  }
  
 .goods-box-bottom .goods-message{
	 width: 100%;
	 height: 80upx;
	 /* border: 1upx solid green; */
 }
 
 .goods-box-bottom .goods-message .goods-price, .goods-buysum{
	 width: 50%;
	 height: 70upx;
	 line-height: 70upx;
	 /* border: 1px solid red; */
	 float: left;
 }
 .goods-box-bottom .goods-message .goods-price{
	 color:#000;
	 font-size: 43upx;
 }
 .goods-box-bottom .goods-message .goods-buysum{
 	color: #555555;
	text-align: right;
 }     
	
 .ades{
	 width:100%;
	 height:190upx;
 }
</style>