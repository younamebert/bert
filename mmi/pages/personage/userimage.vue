<template>
	<view class="context">
	<easy-upload
	types="image"
	:dataList="imageList"
	uploadUrl="http://localhost:3000/upload"
	:types="category"
	deleteUrl='http://localhost:3000/upload'
	:uploadCount="1"
	uploadIcon="/static/icon/up.png"
	@successImage="successImage"
	@successVideo="successvideo"
	></easy-upload>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				imageList: [],
				category: 'image'
			}
		},
		onLoad() {
			uni.request({
				url: 'http://localhost:3000/upload',
				method: 'GET',
				data: {
					category: this.category
				},
				success: res => {
					this.imageList = res.data
				},
				fail(err) {
					console.log(err)
				}
			});
		},
		methods: {
			successImage(e){
				uni.showModal({
					content : '上传成功,详细信息请看控制台'
				})
				console.log(e)
			},
			successvideo(e){
				console.log(e)
			}
		}
	}
</script>

<style>
	.context{
		width: 94%;
		margin:auto;
		/* height: ; */
	}
</style>
