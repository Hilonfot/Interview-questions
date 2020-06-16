import Vue from 'vue'
import App from './App'

import basics from './pages/basics/home.vue'
Vue.component('basics',basics)


import about from './pages/about/home.vue'
Vue.component('about',about)

import tongji from './pages/tongji/home.vue'
Vue.component('tongji',tongji)

import cuCustom from './colorui/components/cu-custom.vue'
Vue.component('cu-custom',cuCustom)

// 全局挂载后使用
import api from '@/common/vmeitime-http/'
Vue.prototype.$api = api

Vue.config.productionTip = false

App.mpType = 'app'


Vue.prototype.$isLogin = function(){
	let token = ''
	try{
		token = uni.getStorageSync("STOREKEY_TOKEN")
		}catch(e){
		//TODO handle the exception
	}
	
	// console.log(token)
	
	if(token == ''){
			return false;
		}else{
			return true;
	}
	
};


const app = new Vue({
    ...App
})
app.$mount()

 



