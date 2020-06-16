<template>
	<view>
		<basics v-if="PageCur=='basics'"></basics>
		
		<about v-if="PageCur=='about'"></about>
		
		<tongji v-if="PageCur=='tongji'"></tongji>
		
		<view class="cu-bar tabbar bg-white shadow foot">
			
			
			<view class="action" @click="NavChange" data-cur="tongji">
				<view class='cuIcon-cu-image'>
					<image :src="'/static/tabbar/tongji' + [PageCur == 'tongji'?'_cur':''] + '.png'"></image>
				</view>
				<view :class="PageCur=='tongji'?'text-blue':'text-gray'">统计</view>
			</view>
			
			<view class="action" @click="NavChange" data-cur="basics">
				<view class='cuIcon-cu-image'>
					<image :src="'/static/tabbar/basics' + [PageCur == 'basics'?'_cur':''] + '.png'"></image>
				</view>
				<view :class="PageCur=='basics'?'text-blue':'text-gray'">功能中心</view>
			</view>
			
			<view class="action" @click="NavChange" data-cur="about">
				<view class='cuIcon-cu-image'>
					<image :src="'/static/tabbar/about' + [PageCur == 'about'?'_cur':''] + '.png'"></image>
				</view>
				<view :class="PageCur=='about'?'text-blue':'text-gray'">我的</view>
			</view>
			
			
		</view>
	</view>
</template>

<script>
	export default {
		data() {
		return {
				PageCur: 'tongji'
			}
		},
		methods: {
			NavChange: function(e) {
				this.PageCur = e.currentTarget.dataset.cur
			}
		}
	}
</script>

<style>

</style>
