<template>
    <view>
        <view class="top">
            <view class="top-tab">
              <view :class="{'top-tab-buy':true,'checkeboxsStyle':index==checkbox}" @click="checkboxs(index)" v-for="(items,index) in toptab" :key="index">
				{{items.topdata}}
			  </view>
            </view>
        </view>
    </view>
</template>

<script>
	export default {
		data() {
			return {
			 checkbox:0,
			 toptab:[
				{
				 topdata:'待付款',
				 indexNumder:1,
			    },
				{
				 topdata:'待发货',
				 indexNumder:2,
				},
				{
				 topdata:'待收货',
				 indexNumder:3,	
				},
				{
				 topdata:'已完成',
				 indexNumder:4,	
				}
			   ]
			}
		},
		computed:{
		            regionName(){
		                // 转为字符串
		              if(this.region.map(item=>item.name).join(' ')){
						    this.address =  this.region.map(item=>item.name).join(' ')
					  }
		            }
		 },
		methods: {
		radioChange: function(evt) {
		            for (let i = 0; i < this.items.length; i++) {
		                if (this.items[i].value === evt.target.value) {
		                    this.current = i;
		                    break;
		                }
		            }
		        },
	    checkboxs:function(e){
			  this.checkbox = e;
		}
		}
	}
</script>

<style>
.top{
	width: 98%;
	height: 150upx;
	margin: auto;
	font-size: 30upx;
	font-family: 'Trebuchet MS', 'Lucida Sans Unicode', 'Lucida Grande', 'Lucida Sans', Arial, sans-serif;
	  border: 1rpx solid blue; 
}
.top-tab{
	width: 90%;
	padding-top: 15upx;
	height: 130upx;
	margin: auto;
	font-family: 'Gill Sans', 'Gill Sans MT', Calibri, 'Trebuchet MS', sans-serif;
    justify-content:center;
    display: flex;
    float: left;
}
.top-tab .top-tab-buy{
	width: 140upx;
	margin-left:10%;
	text-align: center;
	font-size: 28upx;
	line-height: 95upx;
	/* border-radius: 10upx; */
	height: 95upx;
	border-radius:50%;
	 /* border: 1upx solid #2C405A; */
	/* background-color: #141821; */
}

.checkeboxsStyle{
	border-bottom: 4upx solid #3B4246;
	border-top: 4upx solid #3B4246; 
	/* background-color: #2C405A; */
    /* color:#fff; */
	/* color:#0000FF; */
}

.top-tab .top-tab-buy image{
	width: 100%;
	height:100%;
}
</style>
