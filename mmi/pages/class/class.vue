<template>
    <view>
		
		<view class="search"></view>
		<view>
			<view class="class-left">
				<view class="uni-padding-wrap uni-common-mt">
				   <view>
				        <scroll-view :scroll-top="scrollTop" scroll-y="true" class="scroll-Y" :style="{height:allsHeight + 'px'}" @scrolltoupper="upper" @scrolltolower="lower" @scroll="scroll">
									   
									   <view v-for="(item,index) in leftData" :key="index">
										      <view :class="{'left-buy':true,'white':index==white}" @click="unlink(index,item.class_id)" :style="{height:buyHeight + 'px','line-height':buyHeight + 'px'}">{{item.class_name}}</view>
				                        </view> 
				                       </scroll-view>
				    </view>
				</view>
			</view>
			<view class="class-right">
				<view class="right-buy">
					<scroll-view :scroll-top="scrollTop" show-scrollbar="false" scroll-y="true" class="scroll-Y" :style="{height:allsHeight + 'px'}" @scrolltoupper="upper" @scrolltolower="lower"
					               @scroll="scroll">
				<view v-for="(item,index) in rightData" :key="index">				  
				<view class="buy-title">{{item.brand_value}}</view>

					<view class="buy" :style="{'height':r(item.obj) + 'rpx'}">		
							<view class="buy-text-oneline" v-for="(items,indexs) in item.obj" :key="indexs">
								<view class="buy-text-image"><image class="images" :src="items.series_img"></image></view>
								<view class="buy-text-data">{{items.series_name}}</view>
							</view>
					</view>

				</view>
					
	
					</scroll-view>
				</view>
			</view>
		</view>
    </view>
</template>

<script>
	import tools from "@/tools/ade.js"
export default {
    data() {
        return {
            scrollTop: 0,
            old: {
                scrollTop: 0
            },
			allsHeight:0,
			buyHeight:0,
			white:0,
			leftData:"",
			brandpath:"/static/image/goodsimage",
			rightData:"",
			brand:'',
			series:'',
			data:'',
        
		}
    },
    methods: {
         upper: function(e) {
                    console.log(e)
                },
                lower: function(e) {
                    console.log(e)
                },
                scroll: function(e) {
                    console.log(e)
                    this.old.scrollTop = e.detail.scrollTop
                },
                goTop: function(e) {
                    this.scrollTop = this.old.scrollTop
                    this.$nextTick(function() {
                        this.scrollTop = 0
                    });
                    uni.showToast({
                        icon:"none",
                        title:"纵向滚动 scrollTop 值已被修改为 0"
                    })
                },
				unlink:function(e,classid){
					this.white = e;
					this.rightData = this.leftData[e].obj;
					// console.log(e)
					// this.rightData = this.rightData.obj
					console.log(this.rightData)
					// console.log(this.leftData[e])
				},
				r:(e)=>{
					var len = e.length

					var start = 3;
					var startHeight = 250;
	
						for(var i=1;i<10;i++){
						start = start*i;
						 if(len<start || len==start){
							 console.log(startHeight);
							 return  startHeight * i;
						 }
						}
		
				}
    },
	onLoad() {
		
		
		let _this = this;
		let urls = tools.api + "/class";
		let tokens = tools.en_bert(JSON.stringify({"query":"class"}));
	 
	
	uni.getStorage({
		key:"userdata",
		success: (res) => {
			let data = JSON.parse(res.data);
		  uni.request({
		  	url:urls,
		  		data:{"token":tokens,"usertoken":data.data,"pages":"class"},
		  		success:(res)=>{
		  		 if(res.data.msg != 'cache'){
		  			 console.log("bert")
		  			 
		  			 _this.brand = res.data.brand;
		  			 _this.series = res.data.series;
		  			 _this.leftData = res.data.data;
		  			 
		  			 
		  			 for (var index in _this.brand){
		  			 	 _this.brand[index].obj = [];
		  			 	for (var indexs in _this.series){
		  			 		if(_this.brand[index].brand_id == _this.series[indexs].brand_id){
		  			 			_this.series[indexs].series_img = _this.brandpath + _this.series[indexs].series_img;
		  			 			 _this.brand[index].obj.push(_this.series[indexs])
		  			 		}
		  			 	}  
		  			 }
		  			 
		  			 for(var i in _this.leftData){
		  			 	_this.leftData[i].obj = [];
		  			 	for (var is in _this.brand){
		  			 		if(_this.leftData[i].class_id == _this.brand[is].class_id){
		  			 			_this.leftData[i].obj.push(_this.brand[is])
		  			 		}
		  			 	}
		  			 }
		  			
		  		 uni.setStorage({
		  		 	key:"class",
		  			data:JSON.stringify(res.data),
		  		 })
		  		 
		  		  _this.rightData = _this.leftData[0].obj;
		  		  
		  		  
		  		   uni.getSystemInfo({
		  		        success: function (res) {
		  		          _this.allsHeight = res.windowHeight;
		  		        }
		  		      })
		  		  	
		  		  _this.buyHeight = _this.allsHeight / _this.leftData.length;
		  		  _this.buyHeight = _this.buyHeight.toFixed();
		  		 }
		  		}
		  })
	    },
		})



	 
		  
	
		 


	},
	mounted() {
		var _this = this;
		 uni.getSystemInfo({
		      success: function (res) {
		        _this.allsHeight = res.windowHeight;
		      }
		    })
			
		this.buyHeight = this.allsHeight / this.leftData.length;
		// console.log("bertlenth"+this.leftData.length);
		this.buyHeight = this.buyHeight.toFixed();
	}
}
</script>

<style>
/* scroll-Y */
.scroll-Y{
	width: 100%;
	overflow:hidden;
	height: 800rpx;
	/* border: 1px solid red; */
}

.class-left{
	width:21%;
	float: left;
	
	/* border: 1px solid red; */
}

scroll-view ::-webkit-scrollbar {

    width: 0;

    height: 0;

    background-color: transparent;

}


.left-buy{
	width: 100%;
	text-align: center;
	font-size:29upx;
	/* border:; */
	/* color: */
    background-color: #F7F7F7;
	/* border-: 1upx solid #333333; */
	font-family: 'Courier New', Courier, monospace;
	/* border: 1px solid red; */
}


.class-right{
	width:79%;
	float: left;
     /* border: 1px solid red; */
}
.class-right .right-buy{
	width: 100%;
	/* border: 1rpx solid red; */
	/* height: 900rpx; */
}


.buy-title{
	width: 90%;
	height: 80upx;
	line-height: 80upx;
	text-align: center;
	margin: auto;
	float:none;
	/* border: 1rpx solid red; */
	border-bottom: 2upx solid #C0C0C0;
}

.buy-text-oneline{
	width: 33%;
	float:left;
	padding-top: 50upx;
	height:200upx;
	/* border: 1px solid red; */
}

.buy-text-oneline .buy-text-image{
	width: 90%;
	margin: auto;
	height:120upx;
}

.buy{

	width: 90%;
	/* height:510rpx; */
	 /* border: 1rpx solid blue; */
	margin: auto;
}

	
.buy-text-oneline .buy-text-image .images{
	width: 100%;
	height:100upx;
}
.buy-text-oneline .buy-text-data{
	width: 100%;
	padding-top:10upx;
	margin: auto;
	font-size: 20upx;
	font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
	text-align: center;
}

.white{
	background-color: #FFFFFF;
}
</style>
