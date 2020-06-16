<template>
	<view>
		<cu-custom bgColor="bg-gradual-blue" :isBack="true">
			<block slot="backText">返回</block>
			<block slot="content" class="con-title">客户信息-{{customerInfo.customer.name}}</block>
		</cu-custom>
		
		
		<view class="cu-timeline">
			<!-- 备注 -->
			<view class="cu-item">
				<view class="content">
					备注：{{customerInfo.customer.other_info}}
				</view>
			</view>
			
			<!-- 成交记录 -->
			
			<view class="cu-item"   v-for="(item,index) in customerInfo.deal_info" :key="index">
				<view class="content">
					<view>
						{{item.add_time}} <text class="deal">已购房</text>
					</view>
					<view>
						<view>房号： {{item.number}}</view>
						<view>面积： {{item.size}}</view>
						<view>单价： {{item.price}}</view>
						<view>总金额： {{item.total}}</view>
						<view>其他信息： {{item.other_info}}</view>
					</view>
				</view>
			</view>
			
			
			<!-- 跟进记录 -->
			
			<view class="cu-item"   v-for="(item,index) in customerInfo.follow_info" :key="index">
				<view class="content">
					<view>
						{{item.add_time}}  {{followType[item.type]}}
					</view>
					<view>
						{{item.info}}
					</view>
				</view>
			</view>
		</view>
		<view style="height: 200rpx;" class="bg-white"></view>
	</view>
</template>

<script>
	export default {
		data() {
			return {
				customerInfo:[],
				
				followType: ["电话呼出","电话接入","拜访","来访","其他"],
			}
		},
		onLoad(options) {
			// 弹窗查询
			this.$api.CustomerShow(options.detail).then((res)=>{
				
				this.customerInfo = res.data
				
			})
			
		},
		
		methods: {
			
		}
	}
</script>

<style>
	.deal{
		color: red;
		padding-left: 10rpx;
	}
</style>
