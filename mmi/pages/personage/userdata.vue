<template>
	<view>
		<view class="context">
		            <form @submit="formSubmit" @reset="formReset">
		                   
							<!-- <view class="input-buy">
							<label>昵称</label>
		                    <input class="uni-input" :value="userOne.name" name="username"/>
							</view> -->
							
							
							
		
                            <uni-list>
                            	<uni-list-item :title="'名称: '+userOne.name" @click="confirmDialog('name')" message="成功消息" :duration="2000" :before-close="true" @close="close" @confirm="confirm"></uni-list-item>
                            	<uni-list-item :title="'年龄: '+userOne.age"></uni-list-item>
								<uni-list-item title="性别: 男" v-if="userOne.gender==1" @click="gender(1)"></uni-list-item>
								<uni-list-item title="性别: 女" v-if="userOne.gender==0" @click="gender(0)"></uni-list-item>
								<uni-list-item title="换绑电话" type="input" @click="confirmDialog('tel')" message="成功消息" :duration="2000" :before-close="true" @close="close" @confirm="confirm"></uni-list-item>
                            </uni-list>
							<uni-popup ref="dialogInput" type="dialog">
								<uni-popup-dialog mode="input" :title="msg" :value="tel" placeholder="新电话" @confirm="dialogInputConfirm"></uni-popup-dialog>
							</uni-popup>
							
                           <!-- <uni-list>
                           	<uni-list-item v-for="(item,index) in userOne" :key="index" :title="item[]"></uni-list-item>
                           </uni-list> -->
		                    
                      
		               <!-- <view class="uni-btn-v">
		                    <button form-type="submit">保存</button>
		                </view> -->
		            </form>
		        </view>
	</view>
</template>

<script>
	import tools from '@/tools/ade.js'
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
				userOne:"",
				tel:"",
				msg:"",
			}
		},
		onLoad() {
			// uni.clearStorage()
			let _this = this
			uni.getStorage({
				key:"personage",
				success:(res)=>{
					_this.userOne = JSON.parse(res.data);
					
					
				 let age = tools.timetoday("y") - tools.formatDates(_this.userOne.age,"y")
				 if(tools.timetoday("m") < tools.formatDates(_this.userOne.age,"m")){
					 age = age -1;
				 }
                 console.log(_this.userOne)
				 _this.userOne.age = age;
					 },fail() {
					uni.navigateTo({
					    url: '/pages/login/login'
					});
				}
			})
		},
		methods: {
		 	handleGetRegion(region){
		 	    this.region = region
		 	},
		 	confirmDialog(e) {
				
				if(e=="name"){
					this.msg = "name";
					this.tel = this.userOne.name;
				    this.$refs.dialogInput.open()
				}else if(e=="tel"){
					this.msg = "tel";
					this.tel = "";
					this.$refs.dialogInput.open()
				}
		 	    
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
				let urls = tools.api + "/login";
				
				
				let tokens = "";
				if(this.msg=="name"){
					tokens = JSON.stringify({"query":"userdata","Nickname":val})
				}else if(this.msg=="tel"){
					tokens = JSON.stringify({"query":"userdata","Tel":val})
				}
				
				console.log(done)
				uni.getStorage({
					key:"userdata",
					success:(res)=>{
						uni.request({
							url:urls,
							data:{
							"userdata":JSON.parse(res.data).data,
							"token":tools.de_bert(tokens),
							"pages":"userdata",
							}
	                    })
					}
			    })
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
		 	
		}
	}
</script>

<style>
.context{
	width: 94%;
	height: 900upx;
	margin: auto;
	padding-top: 40upx;
	/* border: 1rpx solid red; */
}
.context label{
	font-size: 33upx;
}
.context input{
	color:#2C405A;
	font-size: 28upx;
}
.uni-input{
 border-bottom:1upx solid #CCCCCC;
}
.input-buy{
	width:97%;
	margin: auto;
	height:130upx;
	/* border: 1upx solid red; */
}
.uni-btn-v{
	padding-top: 30upx;
	width: 98%;
	margin: auto;
}
</style>

