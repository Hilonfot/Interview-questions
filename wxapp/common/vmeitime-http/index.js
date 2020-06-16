import http from './interface'

/**
 * 将业务所有接口统一起来便于维护
 * 如果项目很大可以将 url 独立成文件，接口分成不同的模块
 * 
 */

// 存储跟进信息
export const FollowStore = (data) => {
    return http.request({
        url: '/v1/follow',
        method: 'POST', 
		data,
		// handle:true
    })
}

// 存储成交信息
export const DealStore = (data) => {
    return http.request({
        url: '/v1/deal',
        method: 'POST', 
		data,
		// handle:true
    })
}


// 获取首页信息展示
export const CustomerShow = (data) => {
    return http.request({
        url: '/v1/customer/'+data,
        method: 'GET', 
		// handle:true
    })
}


// 获取首页信息展示
export const GetStatistics = (data) => {
    return http.request({
        url: '/v1/statistics',
        method: 'POST', 
        data,
		// handle:true
    })
}



// 获取客户信息
export const GetCustomerInfo = (data) => {
    return http.request({
        url: '/v1/customer',
        method: 'GET', 
        data,
		// handle:true
    })
}


// 添加客户新客户
export const CustomerAdd = (data) => {
    return http.request({
        url: '/v1/customer',
        method: 'POST', 
        data,
		// handle:true
    })
}

// 登陆
export const Login = (data) => {
    return http.request({
        url: '/login',
        method: 'POST', 
        data,
		// handle:true
    })
}

// 单独导出(测试接口) import {test} from '@/common/vmeitime-http/'
export const test = (data) => {
	/* http.config.baseUrl = "http://localhost:8080/api/"
	//设置请求前拦截器
	http.interceptor.request = (config) => {
		config.header = {
			"token": "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		}
	} */
	//设置请求结束后拦截器
	http.interceptor.response = (response) => {
		console.log('个性化response....')
		//判断返回状态 执行相应操作
		return response;
	}
    return http.request({
		baseUrl: 'https://unidemo.dcloud.net.cn/',
        url: 'ajax/echo/text?name=uni-app',
		dataType: 'text',
        data,
    })
}

// 默认全部导出  import api from '@/common/vmeitime-http/'
export default {
	test,
	Login,
	CustomerAdd,
	GetCustomerInfo,
	GetStatistics,
	CustomerShow,
	FollowStore,
	DealStore,
}