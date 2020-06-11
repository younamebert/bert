<template>
	<view>
		<view class="header">
			<view class="input-view">
				<uni-icons type="search" size="22" color="#666666" />
				<input confirm-type="search" class="input" type="text" placeholder="输入搜索关键词" @confirm="confirm">
			</view>
		</view>

	   <view class="ranking">
		   <view>
			  
				 
			   <uni-list v-for="(item,index) in kanking" :key="index">
<uni-list-item :title="	item.list" :show-arrow="false" class="list-kanking"></uni-list-item>

			   </uni-list>
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
				kanking:[
					{
					  list:'aj'  
					},
					{
					  list:'bert衣服'
					},
					{
					  list:'DG菊'
					}
				]
			}
		},
		methods: {
		
		}
	}
</script>

<style>


	.header {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		padding: 10px 15px;
		align-items: center;
		border-top-width: 1px;
		border-top-color: #f5f5f5;
		border-top-style: solid;
		background-color: #ffffff;
	}

	.input-view {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		align-items: center;
		flex-direction: row;
		background-color: #e7e7e7;
		height: 30px;
		border-radius: 15px;
		padding: 0 10px;
		flex: 1;
		background-color: #f5f5f5;
	}

	
	.input {
		flex: 1;
		padding: 0 5px;
		height: 24px;
		line-height: 24px;
		font-size: 14px;
		background-color: transparent;
	}
    .ranking{
		width: 100%;
		/* border: 1upx solid red; */
		padding-top: 160upx;
	}
	
	.list-kanking{
		width: 80%;
		margin: auto;
		border-bottom: 1upx solid #C0C0C0;
	}
</style>