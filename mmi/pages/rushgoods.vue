<template>
	<view>
		<view class="text">
			<view class="context" v-for="(item,index) in rushdata" :key="index" @click="isStart(index)">
				<view class="buy">
					<view class="buy-top">
						<view class="buy-top-left">
							<view class="buy-top-left-images">
								<image :src="'/static/bute/'+item.goodsimg"></image>
							</view>
						</view>
						
						<view class="buy-top-right">
							<view class="right-list-data">
									<view class="list-rights" style="font-size: 28upx;color:#2C405A;">{{item.goodsname}}</view>
									<view class="list-rights" style="font-size: 28upx;color:#2C405A;">{{item.goodsattrbute}}</view>
									<view class="list-rights" style="font-size:30upx;color:#007AFF;">现价: ¥{{item.ruling_price}}</view>
									<view class="list-rights" style="font-size:23upx;text-decoration: line-through;font-style:italic;color:#C0C0C0;">原价: ¥{{item.original_price}}</view>
							</view>
						</view>
					</view>
					
					<view class="buy-bottom">
						
							<!-- <uni-countdown :show-day="false"  background-color='#007AFF' color="#fff" backgroundColor="blue" class="example-body-time" :second="10" @timeup="timeup" /> -->
						<uni-bertdown :second="ms(item.starttime)" @timeup="timeup"></uni-bertdown>
							  
					</view>
				</view>
			</view>

		</view>
	</view>
</template>

<script>
	import uniBertdown from '@/components/uni-bertdown/uni-bertdown.vue'
	import tools from '@/tools/ade.js'
	export default {
		components: {
			uniBertdown,
		},
		data() {
			return {
				testHour: 0,
				testMinute: 0,
				testSecond: 0,
				rushdata:"",
			}
		},
		created() {
			setTimeout(() => {
				this.testHour = 1
				this.testMinute = 1
				this.testSecond = 0
			}, 2000)	
			
		},
		onLoad() {
			
		let _this = this
			uni.getStorage({
				key:"userdata",
				success:(res)=>{
					let data = JSON.parse(res.data)
					let urls = tools.api + "/rushgoods";
					console.log(data)
					
					uni.request({
						url:urls,
						data:{
							usertoken:data.data,
							pages:"rushgoods",
						},
						success:(res)=>{
							if(res.data.status ==200){
								let data = JSON.parse(tools.de_bert(res.data.data))
								// console.log(data)
								_this.rushdata = data;
								console.log(data)
							}
						}
					})
				},
				fail() {
				
				uni.navigateTo({
				    url: '/pages/login/login'
				});	
				}
			})
		},
		methods: {
			
		   timeup() {
				
				console.log(1234);
			},
		   ms(e){
           var timestamp = e - (Date.parse(new Date()) / 1000);
		   // console.log(Date.parse(new Date()) / 1000)
		   // console.log(e)
		      return timestamp
		  },
		  isStart(e){
		 	let data = this.rushdata[e];
			var timestamp = e - (Date.parse(new Date()) / 1000);
		
			if(String(timestamp) >= String(data.starttime)){
				console.log(1)
			}else{
				
			}
		  }
		}
	}
</script>

<style scoped>
 .context{
	 width: 95%;
	 height: 430rpx;
	 margin:auto;
	 padding-top:30upx;
 }
 
 .context .buy{
	 width:95%;
	 height:386upx;
	 margin: auto;
	 /* border:; */
	 border-radius: 10upx;
	 background-color: #fff;
	 border: 1rpx solid #F1F1F1; 
 }
 .context .buy .buy-top{
	 width: 100%;
	 height: 240upx;
	  /* border: 1upx solid blue; */
	 
 }
 
 .context .buy .buy-top .buy-top-left{
	 width: 45%;
	 float: left;
	 padding-top: 30upx;
	 height: 200upx;
 }
 
 .buy-top-left .buy-top-left-images{
	 width: 90%;
	 height: 80%;
	 margin: auto;
 }
 
  .buy-top-left .buy-top-left-images image{
	width: 100%;
	height: 100%;
	border-radius:10upx ;
  }
  
  .buy-top-right{
	width: 54%;
	float: left;
	height: 230upx;
  }
  
   .right-list-data{
	   width: 90%;
	   float: right;
	   height: 230upx;
   }
   
   
  .right-list-data .list-rights{
	  width:100%;
	  height:24%;
	  line-height: 57.5upx;
	  padding-left:10upx;
	  font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
  }
  
  .buy-bottom{
	  width: 100%;
	  padding-top:20upx;
	  margin: auto;
	  height: 128upx;
	  border-radius: 0upx 0upx 10upx 10upx;
	  background-color:rgba(237,237,245,1);
  }
  
</style>
