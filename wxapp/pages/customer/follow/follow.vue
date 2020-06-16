<template>
	<view>
		<cu-custom bgColor="bg-gradual-blue" :isBack="true">
			<block slot="backText">返回</block>
			<block slot="content" class="con-title">跟进客户-{{followData.title}}</block>
		</cu-custom>
		
		

								
				<form class="padding">
					
					<view class="cu-form-group">
						<view class="title">日期：</view>
						<picker mode="date" :value="followData.date" @change="DateChangeFollow">
							<view class="picker">
								{{followData.date}}
							</view>
						</picker>
						<picker mode="time" :value="followData.time"  @change="TimeChangeFollow">
							<view class="picker">
								{{followData.time}}
							</view>
						</picker>
					</view>
					
					<view class="cu-form-group ">
						<view class="title">方式：</view>
						<picker :range="followType"  @change="followTypeChange">
							<view class="picker">
								{{followType[followData.type]}}
							</view>
						</picker>
					</view>
					
					
					<view class="cu-form-group ">
						<view class="title">客户分组：</view>
						<picker :range="type_array"  @change="followCTypeChange">
							<view class="picker">
								{{type_array[followData.c_type]}}
							</view>
						</picker>
					</view>
					
					
					<view class="cu-form-group">
						<textarea maxlength="-1"  placeholder="跟进内容" v-model="followData.info"></textarea>
					</view>
					
					
					<view  class="padding-xl" @tap="submitFollow">
						<button class="cu-btn block bg-gradual-blue margin-tb-sm lg"><text>提&nbsp;&nbsp;交</text></button>
					</view>
					
				</form>
								
				<!-- 跟进信息填写结束 -->
	</view>
</template>

<script>
	export default {
		data() {
			return {
				type_array: ['待联系','已联系','复核','已成交','公共池','无效客户'],
				followType: ["电话呼出","电话接入","拜访","来访","其他"],
				followData: {
					title: "",
					c_id: 0,
					date: "",
					time: "",
					c_type: 0,
					type: 0,
					info: "",
				},
			}
		},
		onLoad(options) {
			this.followData = {
				title: options.name,
				c_id: options.detail,
				date: this.getNowData(),
				time: this.getNowTime(),
				
				c_type: options.ctype,
				type: 0,
				info: "",
			}
		},
		methods: {
			submitFollow(){
				var _this = this
				var add_time = _this.followData.date+ " "+ _this.followData.time
				var data = {
					c_id: parseInt(_this.followData.c_id),
					add_time:add_time,
					type: parseInt(_this.followData.type),
					info: _this.followData.info,
					c_type: parseInt(_this.followData.c_type),
				}
				_this.$api.FollowStore(data).then((res)=>{
					
					if (res.code == 200){
						uni.navigateBack()
					}
					
					uni.showToast({
					    title: res.msg,
					    icon: 'none',
					    mask: true,
						duration: 2000
					});
				})
			},
			followTypeChange(e){
				this.followData.type = e.detail.value
			},
			followCTypeChange(e){
				this.followData.c_type = e.detail.value
			},
			TimeChangeFollow(e) {
				this.followData.time = e.detail.value
			},
			DateChangeFollow(e) {
				this.followData.date = e.detail.value
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
