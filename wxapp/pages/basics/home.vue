<template name="basics">
	
	<view class="bg-white">
		
		<image src="/static/101.png"
		 mode="widthFix" class="response"></image>
		 
		
		<view class="bg-title">
			<view class="action act-title">
				<!-- <text class="cuIcon-title text-orange"></text> -->
				<text>客户管理</text>
			</view>
		</view>
		
		<view class="cu-list grid" :class="['col-' + gridCol,gridBorder?'':'no-border', 'card-menu']">
			
			<view class="cu-item" v-for="(item,index) in cuIconList" :key="index">
			<!-- v-if="index<gridCol*2" -->
				<!-- <view :class="['cuIcon-' + item.cuIcon,'text-' + item.color]"> -->
					<!-- <view class="cu-tag badge" v-if="item.badge!=0">
						<block v-if="item.badge!=1">{{item.badge>99?'99+':item.badge}}</block>
					</view> -->
				<!-- </view> -->
				<view  :data-nav="item.nav"  :data-param="item.param" :data-title="item.name"     @tap="nav">
					<view>
						<image class="image-grid" :src="item.cuIcon" ></image>
						<view class="cu-tag badge" v-if="item.badge!=0">
							<block v-if="item.badge!=1">{{item.badge>99?'99+':item.badge}}</block>
						</view> 
					</view>
					<text>{{item.name}}</text>
				</view>
				
				
			</view>
		</view>
		
		
		
			
	</view>
</template>

<script>
	export default {
		name: "basics",
		data() {
			return {
				cuIconList: [{
					cuIcon: '/static/renwu/renwu_0.png',
					color: 'red',
					badge: 120,
					name: '待联系',
					nav: "../customer/content",
					param: 0
				}, {
					cuIcon: '/static/renwu/renwu_1.png',
					color: 'orange',
					badge: 1,
					name: '已联系',
					nav: "../customer/content",
					param: 1
				}, {
					cuIcon: '/static/renwu/renwu_2.png',
					color: 'yellow',
					badge: 0,
					name: '复核',
					nav: "../customer/content",
					param: 2
				}, {
					cuIcon: '/static/renwu/renwu_3.png',
					color: 'olive',
					badge: 22,
					name: '已成交',
					nav: "../customer/content",
					param: 3
				}, {
					cuIcon: '/static/renwu/renwu_4.png',
					color: 'cyan',
					badge: 0,
					name: '公共池',
					nav: "../customer/content",
					param: 4
				},
				{
					cuIcon: '/static/renwu/renwu_5.png',
					color: 'cyan',
					badge: 0,
					name: '无效客户',
					nav: "../customer/content",
					param: 5
				},
				{
					cuIcon: '/static/renwu/renwu_6.png',
					color: 'cyan',
					badge: 0,
					name: '新增客户',
					nav: "../customer/add",
					param: ""
				}],
				gridCol: 4,
				gridBorder: false,
				
				
				
			};
		},
		onLoad() {
			let _this = this
			this.TowerSwiper('swiperList');
		},
		methods: {
			nav(e){
				if (this.$isLogin()){
					let param = e.currentTarget.dataset.param
					let nav = e.currentTarget.dataset.nav
					let title = e.currentTarget.dataset.title
					//在起始页面跳转到test.vue页面并传递参数
					uni.navigateTo({
					    url:  nav+"?param="+param+ "&title="+title
					});
					
				}else{
					
					uni.showToast({
					    title: "未登录或非法访问",
					    icon: 'none',
						duration: 2000
					});
					
				}
			},
			
		}
	}
</script>

<style>
	.act-title{
		font-size: 30rpx;
		color: white;
		font-weight: 600;
		padding-bottom: 10rpx;
		text-align: center;}
		.image-grid{width: 70rpx;height: 70rpx;
	}
	.act-title text{
		position: relative;
		top: 10rpx;
	
	}
	.bg-title{
		background-image:url(../../static/bg-title.png);
		background-size: cover;
		width: 750rpx;
		height: 60rpx;
	}
	.tower-swiper .tower-item {
		transform: scale(calc(0.5 + var(--index) / 10));
		margin-left: calc(var(--left) * 100upx - 150upx);
		z-index: var(--index);
	}
</style>
