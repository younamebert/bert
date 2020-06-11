<template>
	<view>	
	
		<view class="context">
			<pick-regions :defaultRegion="defaultRegionCode" @getRegion="handleGetRegion">
			    <uni-list>
			    	<uni-list-item :title="address" note=""></uni-list-item>
					{{regionName}}
			    </uni-list>
			</pick-regions>
		
			<uni-list>
				<uni-list-item :title="tel" note="电话" type="input" @click="confirmDialog" message="成功消息" :duration="2000" :before-close="true" @close="close" @confirm="confirm" :showArrow="false"></uni-list-item>
				
			</uni-list>
			<uni-popup ref="dialogInput" type="dialog">
				<uni-popup-dialog mode="input" title="输入电话" :value="tel" placeholder="请输入内容" @confirm="dialogInputConfirm"></uni-popup-dialog>
			</uni-popup>
			
            <view class="uni-textarea">
                <textarea placeholder="详细地址"/>
            </view>
			
			<view class="bottom">
			    <button type="primary" style="background-color: #2C405A;">保存</button>
			</view>
        </view>
		</view>
	</view>
</template>
<script>
import pickRegions from '@/components/pick-regions/pick-regions.vue'
import uniPopupDialog from '@/components/uni-popup/uni-popup-dialog.vue'
import uniPopupShare from '@/components/uni-popup/uni-popup-share.vue'
	export default {
		components:{ pickRegions,
		uniPopupDialog,
		uniPopupShare
		},
		data() {
			return {
			 region:[],
			 defaultRegion:['广东省','广州市','番禺区'],
			 defaultRegionCode:'440113',
			 address:'选择地址',
			 tel:'13975515540',
			 type: 'top',
			 msgType: 'success',
			 message: '这是一条成功消息提示',
			 value: '默认输入的内容'
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
		 onReady() {
		 },
		onLoad(e) {
			console.log(e);
		},
		methods: {
			handleGetRegion(region){
			    this.region = region
			},
			confirmDialog() {
			    this.$refs.dialogInput.open()
			},
			dialogConfirm(done) {
			 	this.$refs.popupMessage.open()
			 	console.log('点击确认');
			 	done()
			 	},
	
			dialogInputConfirm(done, val) {
			 	uni.showLoading({
			 		title: '保存成功'
			 	})
				this.tel = val
				setTimeout(() => {
					uni.hideLoading()
					// 关闭窗口后，恢复默认内容
					done()
				}, 1000)
			 	},
			 	dialogClose(done) {
			 		this.msgType = 'info'
			 		this.message = '点击了对话框的取消按钮'
			 		this.$refs.popupMessage.open()
			 		// 需要执行 done 才能关闭对话框
			 		done()
			 	}
			 },
			 /**
			  * 拦截应用返回事件，仅仅 app 端生效
			  */
			 onBackPress() {
			 
			 }
			}
	
	
</script>

<style lang="scss" scoped>
.context{
	width: 94%;
	height: 600upx;
	margin: auto;
	// border: 1rpx solid red;
}
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

	.example-info {
		padding: 15px;
		color: #3b4144;
		background: #ffffff;
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


	.example-body {
		/* #ifndef APP-NVUE */
		display: block;
		/* #endif */
		padding: 5px 15px;
	}


	.popup-content {
		background-color: #fff;
		padding: 15px;
	}

	/* #ifndef APP-NVUE */
	.button {
		margin: 0;
		padding: 0;
		line-height: 1;
	}

	/* #endif */

	.button {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		align-items: center;
		justify-content: center;
		flex: 1;
		height: 40px;
		margin: 0 15px;
	}

	.button-text {
		color: #fff;
		font-size: 14px;
	}

	.message {
		/* #ifndef APP-NVUE */
		display: flex;
		/* #endif */
		flex-direction: row;
		justify-content: center;
		padding: 15px 0;
	}

	.dialog {
		padding: 15px 0;
	}

	.dialog-box {
		margin-left: 15px;
		margin-bottom: 15px;
	}

	.dialog-text {
		font-size: 14px;
		color: #666;
	}

	.share {
		padding: 15px 0;
	}

	.popup__success {
		color: #fff;
		background-color: #4cd964;
	}

	.popup__warn {
		color: #fff;
		background-color: #f0ad4e;
	}

	.popup__error {
		color: #fff;
		background-color: #dd524d;
	}

	.popup__info {
		color: #fff;
		background-color: #909399;
	}
	.uni-textarea{
		width: 98%;
		margin: auto;
		height: 200upx;
		// border-top:1rpx solid #C0C0C0 ;
		border: 1rpx solid #C0C0C0;
		border-radius:5upx ;
	}
	.bottom{
		width: 100%;
		height: 80upx;
		margin:auto;
		
		padding-top: 30upx;
		// border: 1rpx solid red;
	}
</style>
