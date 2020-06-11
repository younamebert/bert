import Vue from 'vue'
import App from './App'

Vue.config.productionTip = false

function api(){
	return "http://127.0.0.1/"
}
module.exports = api

App.mpType = 'app'

const app = new Vue({
	...App
})
app.$mount()
