<template>
	<view class="bw-cd">
		{{cdStr}}
	</view>
</template>

<script>
	let setStartTime;
	export default {
		data() {
			return {
				cdTime:0,
				cdStr:'',
			};
		},
		methods:{
			/*清空倒计时*/
			clearStartTime(){
				return new Promise( (res,rel) => {
					clearInterval(setStartTime);
					setStartTime = null; 
					this.cdStr = '';
					res();
				})
				
			},
			 formatSeconds(t) {
				let mi = 60,hh = mi*60,dd = hh*24;
				let d = this.formatBit( Math.floor(t/dd)),
						h = this.formatBit( Math.floor((t - d*dd)/hh)),
						m = this.formatBit( Math.floor((t - d*dd - h*hh)/mi)),
						s = this.formatBit( Math.floor((t - d*dd - h*hh - m*mi)));
				let tstr = d+'天'+h+'小时'+m+"分"+s+'秒';
				return tstr;
			},
			initTime(time){
				this.clearStartTime().then( ()=> {
					this.cdTime = Math.floor(time/1000);
					console.log(this.cdTime);
					setStartTime = setInterval(() => {
						this.cdTime--;
						this.cdStr = this.formatSeconds(this.cdTime);
						// console.log(this.cdStr);
						if (this.cdTime <= 0) {
							clearInterval(setStartTime)
							this.$emit('bwCountdown');
						}
					}, 1000)
				})
				
			},
			formatBit(v){
				v = +v
				return v > 9 ? v : '0' + v
			},
		}
	}
</script>

<style lang="scss">
.bw-cd{
	flex: 1;
	height: 100%;
	display: flex;
	align-items: center;
	justify-content: center;
}
</style>
