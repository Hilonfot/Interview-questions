<template>
	<view>
		<cu-custom bgColor="bg-gradual-blue" :isBack="true">
			<block slot="backText">返回</block>
			<block slot="content" class="con-title">成交信息-{{dealData.title}}</block>
		</cu-custom>
		
				<!-- 成交信息填写开始 -->
			
				<form class="padding">
					<view class="cu-form-group">
						<view class="title">日期：</view>
						<picker mode="date" :value="dealData.date" @change="DateChangeDeal">
							<view class="picker">
								{{dealData.date}}
							</view>
						</picker>
						<picker mode="time" :value="dealData.time"  @change="TimeChangeDeal">
							<view class="picker">
								{{dealData.time}}
							</view>
						</picker>
					</view>
					
					<view class="cu-form-group">
						<view class="title">房号：</view>
						<input name="input" v-model="dealData.number"></input>
					</view>
					<view class="cu-form-group">
						<view class="title">面积：</view>
						<input name="input" v-model="dealData.size"></input>
						<view class="cu-tag bg-blue">
							平方
						</view>
					</view>
					<view class="cu-form-group">
						<view class="title">单价：</view>
						<input name="input" v-model="dealData.price"></input>
					</view>
					
					<view class="cu-form-group">
						<view class="title">总金额：</view>
						<input name="input" v-model="dealData.total"></input>
					</view>
					
					<view class="cu-form-group">
						<textarea maxlength="-1"  placeholder="其他信息备注" v-model="dealData.other_info"></textarea>
					</view>
					
					<view  class="padding-xl" @tap="submitDeal">
						<button class="cu-btn block bg-gradual-blue margin-tb-sm lg"><text>提&nbsp;&nbsp;交</text></button>
					</view>
					
				</form>
			
				<!-- 成交信息填写结束 -->
	</view>
</template>

<script>
	export default {
		data() {
			return {
				dealData:{
					title: "",
					c_id: 0,
					date: "",
					time: "",
					number: "",
					size: "",
					price: "",
					total: "",
					other_info: "",
				},
			}
		},
		
		onLoad(options) {
			// 重置数据
			this.dealData = {
				title: options.name,
				c_id: options.detail,
				date: this.getNowData(),
				time: this.getNowTime(),
				
				number: "",
				size: "",
				price: "",
				total: "",
				other_info: "",
			}	
		},
		methods: {
			submitDeal(){
				var _this = this
				var add_time = _this.dealData.date+ " "+ _this.dealData.time
				var data = {
					c_id: parseInt(_this.dealData.c_id),
					add_time:add_time,
					number: _this.dealData.number,
					size: _this.dealData.size,
					price: _this.dealData.price,
					total: _this.dealData.total,
					other_info: _this.dealData.other_info,
				}
				_this.$api.DealStore(data).then((res)=>{
					if (res.code == 200){
						uni.navigateBack()
					}
					uni.showToast({
					    title: res.msg,
					    icon: 'none',
						duration: 2000
					});
				})
			},
			TimeChangeDeal(e) {
				this.dealData.time = e.detail.value
			},
			DateChangeDeal(e) {
				this.dealData.date = e.detail.value
			},
			getNowData(){
				var date1 = new Date()
				return date1.getFullYear()+"-"+this.twoNum(date1.getMonth()+1)+"-"+this.twoNum(date1.getDate()) //time1表示当前时间
			},
			getNowTime(){
				var date1 = new Date()
				return 	this.twoNum(date1.getHours())+":"+this.twoNum(date1.getMinutes())
			},
			twoNum(num){
				if (num <= 9) {
				    num = "0" + num;
				 }
				 return num
			},
		}
	}
</script>

<style>

</style>
