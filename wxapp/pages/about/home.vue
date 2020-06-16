<template name="about">
	<scroll-view scroll-y class="scrollPage">
		
	  <view class="UCenter-bg">
		  
		  
	    <view src="" class="userAvatarUrl">
				<open-data  type="userAvatarUrl"></open-data>
		</view>
	
	    <view class="text-xl userNickName">
			<open-data type="userNickName" class="userNickName" lang="zh_CN">
				
			</open-data>
	      <!-- <text class="text-df">v2.0</text> -->
	    </view>
	   <!-- <view class="margin-top-sm">
	      <text>By:文晓港</text>
	    </view> -->
	    <image src="https://blog.hilonfot.com/wave.gif" mode="scaleToFill" class="gif-wave"></image>
	  </view>
	  <view class="padding flex text-center text-grey bg-white shadow-warp">
	    <view class="flex flex-sub flex-direction solid-right">
	      <view class="text-xxl text-orange">12</view>
	      <view class="margin-top-sm">
	        <text class="cuIcon-attentionfill"></text>成交/套</view>
	    </view>
	    <view class="flex flex-sub flex-direction solid-right">
	      <view class="text-xxl text-blue">980.5</view>
	      <view class="margin-top-sm">
	        <text class="cuIcon-favorfill"></text>金额/万</view>
	    </view>
	    <view class="flex flex-sub flex-direction">
	      <view class="text-xxl text-green">17</view>
	      <view class="margin-top-sm">
	        <text class="cuIcon-fork"></text>客户/位</view>
	    </view>
	  </view>
	  <view class="cu-list menu card-menu margin-top-xl margin-bottom-xl shadow-lg radius">
		  
	    <!-- <view class="cu-item arrow">
	      <view class="content" bindtap="CopyLink" data-link="https://github.com/weilanwl/ColorUI">
	        <text class="cuIcon-github text-grey"></text>
	        <text class="text-grey">关于我</text>
	      </view>
	    </view> -->
	   <!-- <view class="cu-item arrow">
	      <navigator class="content" url="/pages/about/about/about" hover-class="none">
	        <image src="/static/logo.png" class="png" mode="aspectFit"></image>
	        <text class="text-grey">关于ColorUI组件库</text>
	      </navigator>
	    </view> -->
	    <!-- <view class="cu-item arrow">
	      <navigator class="content" url="/pages/about/log/log" hover-class="none">
	        <text class="cuIcon-formfill text-green"></text>
	        <text class="text-grey">日志</text>
	      </navigator>
	    </view> -->
	   <!-- <view class="cu-item arrow">
	      <view class="content" bindtap="showQrcode">
	        <text class="cuIcon-appreciatefill text-red"></text>
	        <text class="text-grey">赞赏支持</text>
	      </view>
	    </view> -->
		
		<!-- <button open-type="getUserInfo" lang="zh_CN" @getuserinfo="onGotUserInfo">授权账号信息</button> -->
		
		<view class="cu-item arrow" >
		  <navigator class="content" url="/pages/login/login" hover-class="none">
		    <text class="cuIcon-appreciatefill text-red"></text>
		    <text class="text-grey">登陆</text>
		  </navigator>
		</view>
		
		<view class="cu-item arrow" data-target="bottomModal" >
		  <navigator class="content" url="/pages/about/info" hover-class="none">
		    <text class="cuIcon-creativefill text-orange"></text>
		    <text class="text-grey">个人信息</text>
		  </navigator>
		</view>
		
		<view class="cu-item arrow">
		  <navigator class="content" url="/pages/login/register" hover-class="none">
		    <text class="cuIcon-formfill text-green"></text>
		    <text class="text-grey">修改密码</text>
		  </navigator>
		</view>
			
	  
		
		<view class="cu-item arrow">
		  <button class="cu-btn content" open-type="feedback">
		    <text class="cuIcon-writefill text-cyan"></text>
		    <text class="text-grey">意见反馈</text>
		  </button>
		</view>
		
		
		
		
	  </view>
	  <view class="cu-tabbar-height"></view>
	</scroll-view>
</template>

<script>
	
	export default {
		name: "about",
		data() {
			 return {
				userinfo: [],
				modalName: null,
			 };
		},

		onLoad(){
					console.log("页面打开？")
						let _this = this;
						uni.login({
							provider: 'weixin',
							success: function(loginRes) {
								console.log("登陆微信？")
								// 获取用户信息				
								uni.getUserInfo({
									provider: 'weixin',
									success: function(infoRes) {	
										console.log("获取信息？")
										_this.userinfo = info.userInfo
										console.log(_this.userinfo)
									}
								});
							}
						});
					},


		methods: {
			/**
			* 用户同意授权个人微信信息
			* @param {Object} e 用户的信息
			*/
			onGotUserInfo(e) {
			    console.info(e.detail)
			    console.info(e.detail.userInfo);
			},
		}
	}
</script>

<style>
	.UCenter-bg {
	  background-image: url(https://blog.hilonfot.com/aboutme.jpeg);
	  background-size: cover;
	  height: 550rpx;
	  display: flex;
	  justify-content: center;
	  padding-top: 40rpx;
	  overflow: hidden;
	  position: relative;
	  flex-direction: column;
	  align-items: center;
	  color: #fff;
	  font-weight: 300;
	  text-shadow: 0 0 3px rgba(0, 0, 0, 0.3);
	}
	
	.UCenter-bg text {
	  opacity: 0.8;
	}
	
	.UCenter-bg image {
	  width: 200rpx;
	  height: 200rpx;
	}
	
	.UCenter-bg .gif-wave{
	  position: absolute;
	  width: 100%;
	  bottom: 0;
	  left: 0;
	  z-index: 99;
	  mix-blend-mode: screen;  
	  height: 100rpx;   
	}
	
	map,.mapBox{
	  left: 0;
	  z-index: 99;
	  mix-blend-mode: screen;  
	  height: 100rpx;   
	}
	
	map,.mapBox{
	  width: 750rpx;
	  height: 300rpx;
	}
	.userAvatarUrl{
		width: 200rpx;
		
		height: 200rpx;
		
		border-radius: 50%;
		
		overflow: hidden;
		
		/* margin-top: 100rpx; */
		
		/* justify-content: center; */
	}
	.userNickName{
		color: #080808;
		font-size: 45rpx;
		font-weight: 300;
		}
	.userNickName{
		height: 45rpx;
	}
</style>
