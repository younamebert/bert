<template>
	<view>
		<view class="context">
			<view class="shopcart-buy">
				<view class="top">
					<checkbox-group name="checkbox" @change="checkboxChange">
					                        <label>
					                            <checkbox value="checkbox1" color="blue" style="transform:scale(0.8)"/>
					                        </label>
					                    </checkbox-group>
				</view>
				<view class="left">
					<view class="left-images">
						<image src="/static/image/2s.jpeg"></image>
					</view>
				</view>
				<view class="right">
					<uni-list>
						<uni-list-item class="list-rights" title="潮流单品" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="XL码" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="¥200" :show-arrow="false"></uni-list-item>
					</uni-list>
				</view>
			</view>
		</view>
		
		<view class="context">
			<view class="shopcart-buy">
				<view class="top">
					<checkbox-group name="checkbox" @change="checkboxChange">
					                        <label>
					                            <checkbox value="checkbox1" color="blue" style="transform:scale(0.8)"/>
					                        </label>
					                    </checkbox-group>
				</view>
				<view class="left">
					<view class="left-images">
						<image src="/static/image/2s.jpeg"></image>
					</view>
				</view>
				<view class="right">
					<uni-list>
						<uni-list-item class="list-rights" title="潮流单品" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="XL码" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="¥200" :show-arrow="false"></uni-list-item>
					</uni-list>
				</view>
			</view>
		</view>
		
		<view class="context">
			<view class="shopcart-buy">
				<view class="top">
					<checkbox-group name="checkbox" @change="checkboxChange">
					                        <label>
					                            <checkbox value="checkbox1" color="blue" style="transform:scale(0.8)"/>
					                        </label>
					                    </checkbox-group>
				</view>
				<view class="left">
					<view class="left-images">
						<image src="/static/image/2s.jpeg"></image>
					</view>
				</view>
				<view class="right">
					<uni-list>
						<uni-list-item class="list-rights" title="潮流单品" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="XL码" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="¥200" :show-arrow="false"></uni-list-item>
					</uni-list>
				</view>
			</view>
		</view>
		
		<view class="context">
			<view class="shopcart-buy">
				<view class="top">
					<checkbox-group name="checkbox" @change="checkboxChange">
					                        <label>
					                            <checkbox value="checkbox1" color="blue" style="transform:scale(0.8)"/>
					                        </label>
					                    </checkbox-group>
				</view>
				<view class="left">
					<view class="left-images">
						<image src="/static/image/2s.jpeg"></image>
					</view>
				</view>
				<view class="right">
					<uni-list>
						<uni-list-item class="list-rights" title="潮流单品" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="XL码" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="¥200" :show-arrow="false"></uni-list-item>
					</uni-list>
				</view>
			</view>
		</view>
		
		<view class="context">
			<view class="shopcart-buy">
				<view class="top">
					<checkbox-group name="checkbox" @change="checkboxChange">
					                        <label>
					                            <checkbox value="checkbox1" color="blue" style="transform:scale(0.8)"/>
					                        </label>
					                    </checkbox-group>
				</view>
				<view class="left">
					<view class="left-images">
						<image src="/static/image/2s.jpeg"></image>
					</view>
				</view>
				<view class="right">
					<uni-list>
						<uni-list-item class="list-rights" title="潮流单品" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="XL码" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="¥200" :show-arrow="false"></uni-list-item>
					</uni-list>
				</view>
			</view>
		</view>
		
		<view class="context">
			<view class="shopcart-buy">
				<view class="top">
					<checkbox-group name="checkbox" @change="checkboxChange">
					                        <label>
					                            <checkbox value="checkbox1" color="blue" style="transform:scale(0.8)"/>
					                        </label>
					                    </checkbox-group>
				</view>
				<view class="left">
					<view class="left-images">
						<image src="/static/image/2s.jpeg"></image>
					</view>
				</view>
				<view class="right">
					<uni-list>
						<uni-list-item class="list-rights" title="潮流单品" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="XL码" :show-arrow="false"></uni-list-item>
						<uni-list-item class="list-rights" title="¥200" :show-arrow="false"></uni-list-item>
					</uni-list>
				</view>
			</view>
		</view>
		<view class="ade"></view>
		<view class="shopcart-bottom">
			<view class="shopcart-bottom-top">
				    <checkbox value="checkbox1" color="red" style="transform:scale(0.9)"/>
				总计:¥{{sumprice}}
				</view>
			<view class="shopcart-bottom-submit">
				<button type="primary">确认支付</button>
			</view>
		</view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				sumprice:100
			}
		},
		methods: {
		checkboxChange(){
			
		}	
		}
	}
</script>

<style>
.context{
	width:94%;
	margin: auto;
	padding-top: 30upx;
	height: 360upx;
}

.context .shopcart-buy{
	padding-top: 20upx;
	width: 96%;
	margin:auto;
	border-radius: 6upx;
	height: 310upx;
	border: 1upx solid #CCCCCC;
}

.shopcart-buy .left{
	width: 45%;
	height: 210upx;
	float: left;
}
.shopcart-buy .left .left-images{
	width: 90%;
	margin: auto;
	height: 210upx;
}
.shopcart-buy .left .left-images image{
	width: 100%;
	height: 100%;
}

.shopcart-buy .right{
	width:54%;
	float: left;
	height: 210upx;
}

.shopcart-buy .right .list-rights{
	height: 70upx;
}

.top{
	width: 90%;
	height: 70upx;
	margin: auto;
	/* border-bottom: 1upx solid #CCCCCC; */
}
.shopcart-bottom{
	width: 100%;
	/* display: none; */
	margin: auto;
	position: fixed;
	background-color: #FFF; 
	border-top:1upx solid #F1F1F1;
	bottom: 0;
	border-radius: 6upx;
	height: 220upx;
}

.shopcart-bottom .shopcart-bottom-top{
	width: 90%;
	margin: auto;
	font-size:35upx;
	height:100upx;
	vertical-align: middle;
	color:#141821;
	line-height: 100upx;
}
.shopcart-bottom .shopcart-bottom-submit{
	width: 95%;
	height: 100upx;
	margin: auto;
}
.ade{
	width: 100%;
	height: 220upx;
}
</style>
