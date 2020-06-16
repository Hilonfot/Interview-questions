<template>
	<view class="register">
	
		<view class="content">
			<!-- 头部logo -->
			<view class="header">
				<image :src="logoImage"></image>
			</view>
			<!-- 主体 -->
			<view class="main">
				<wInput
					v-model="phoneData"
					type="text"
					maxlength="11"
					placeholder="手机号"
				></wInput>
				<wInput
					v-model="passData"
					type="password"
					maxlength="11"
					placeholder="登录密码"
					isShowPass
				></wInput>
				<wInput
					v-model="passDataRe"
					type="password"
					maxlength="11"
					placeholder="新密码"
					isShowPass
				></wInput>
				<wInput
					v-model="passDataRe"
					type="password"
					maxlength="11"
					placeholder="重复新密码"
					isShowPass
				></wInput>
				<!-- <wInput
					v-model="verCode"
					type="number"
					maxlength="4"
					placeholder="验证码"
					
					isShowCode
					ref="runCode"
					@setCode="getVerCode()"
				></wInput> -->
					
			</view>
				
			<wButton 
				text="提 交"
				:rotate="isRotate" 
				@click.native="startReg()"
			></wButton>
			
			<!-- 底部信息 -->
			<view class="footer">
				<text 
					@tap="isShowAgree" 
					class="cuIcon"
					:class="showAgree?'cuIcon-radiobox':'cuIcon-round'"
				>同意</text>
				<!-- 协议地址 -->
				<navigator url="" open-type="navigate">《协议》</navigator>
			</view>
		</view>
	</view>
</template>

<script>
	var _this;
	import wInput from '../../components/watch-login/watch-input.vue' //input
	import wButton from '../../components/watch-login/watch-button.vue' //button
	
	export default {
		data() {
			return {
				//logo图片 base64
				logoImage:'data:image/png;base64,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',
				phoneData:'', // 用户/电话
				passData:'', //密码
				verCode:"", //验证码
				showAgree:true, //协议是否选择
				isRotate: false, //是否加载旋转
			}
		},
		components:{
			wInput,
			wButton,
		},
		mounted() {
			_this= this;
		},
		methods: {
			isShowAgree(){
				//是否选择协议
				_this.showAgree = !_this.showAgree;
			},
			getVerCode(){
				//获取验证码
				if (_this.phoneData.length != 11) {
				     uni.showToast({
				        icon: 'none',
						position: 'bottom',
				        title: '手机号不正确'
				    });
				    return false;
				}
				console.log("获取验证码")
				this.$refs.runCode.$emit('runCode'); //触发倒计时（一般用于请求成功验证码后调用）
				uni.showToast({
				    icon: 'none',
					position: 'bottom',
				    title: '模拟倒计时触发'
				});
				
				setTimeout(function(){
					_this.$refs.runCode.$emit('runCode',0); //假装模拟下需要 终止倒计时
					uni.showToast({
					    icon: 'none',
						position: 'bottom',
					    title: '模拟倒计时终止'
					});
				},3000)
			},
		    startReg() {
				//注册
				if(this.isRotate){
					//判断是否加载中，避免重复点击请求
					return false;
				}
				if (this.showAgree == false) {
				    uni.showToast({
				        icon: 'none',
						position: 'bottom',
				        title: '请先同意《协议》'
				    });
				    return false;
				}
				if (this.phoneData.length !=11) {
				    uni.showToast({
				        icon: 'none',
						position: 'bottom',
				        title: '手机号不正确'
				    });
				    return false;
				}
		        if (this.passData.length < 6) {
		            uni.showToast({
		                icon: 'none',
						position: 'bottom',
		                title: '密码不正确'
		            });
		            return false;
		        }
				if (this.verCode.length != 4) {
				    uni.showToast({
				        icon: 'none',
						position: 'bottom',
				        title: '验证码不正确'
				    });
				    return false;
				}
				console.log("注册成功")
				_this.isRotate=true
				setTimeout(function(){
					_this.isRotate=false
				},3000)
		    }
		}
	}
</script>

<style>
	@import url("../../components/watch-login/css/icon.css");
	@import url("./css/main.css");
</style>