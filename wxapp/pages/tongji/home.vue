<template>
	<view class="bg-white padding-top">
		<!-- <cu-custom bgColor="bg-gradual-blue" :isBack="false">
			<block slot="backText">返回</block>
			<block slot="content" class="con-title">统计</block>
		</cu-custom> -->
			
		<view  class="padding">
			<view>
				
				<picker @change="PickerChange" :value="user_index" :range="user_info" range-key="nick_name">
					<view class="picker">
						<view class="cu-avatar round lg" style="background-color: #ADADAD;" v-if="user_info[user_index].nick_name">{{user_info[user_index].nick_name.substr(0,1)}}</view>
						
						<text class="select-name padding-left">{{user_info[user_index].nick_name}}<text class="lg text-gray cuIcon-right select-name-icon"></text></text>
					</view>
				</picker>
				
			</view>
			
			
			<scroll-view scroll-x class="bg-white nav margin-top" scroll-with-animation :scroll-left="scrollLeft">
				<view class="cu-item" :class="1==tab_cur?'text-blue cur':''"  @tap="tabSelect" :data-id="1">
					<text class="nav-title">数据概览</text>
				</view>
				
				<view class="cu-item" :class="2==tab_cur?'text-blue cur':''"  @tap="tabSelect" :data-id="2">
					<text class="nav-title">数据统计</text>
				</view>
				
			</scroll-view>
			
			<view class="margin-top" v-if="tab_cur == 1">
				<view class="tap1-title">
					<text class="t1">今日概况</text>
					<text class="t2">数据更新至 {{data.today}}</text>
				</view>
				<view class="flex margin-top">
					<view class="flex-sub  margin-xs flex-box">
						<view>到访人数</view>
						<view class="margin-top-xs box-num">{{data.today_count}}</view>
					</view>
					
					<view class="flex-sub margin-xs">
						<view>成交数量</view>
						
						<view class="margin-top-xs box-num">{{data.today_deal}}</view>
					</view>
				</view>
				
				<view class="progress-show-title margin-top cuIcon-rankfill">客户来源分布</view>
				
					<view>
					<view class="bg-white progress-box">
						
						
					<view class="flex margin-top-xs" v-for="(item,index) in source_array" :key="index">
						<view class="progress-title basis-sm">{{item}}: {{data.today_source_count[index].num?data.today_source_count[index].num:"" }}</view>
						<view class="cu-progress radius">
							<view class="bg-gradual-blue" :style="[{ width:loading?data.today_source_count[index].percent+'%':''}]">{{data.today_source_count[index].percent}}%</view>
						</view>
					</view>
						
						
						
						
					</view>
				</view>
				
				<view class="progress-show-title margin-top cuIcon-sponsorfill">意向程度分布</view>
				
				<view>
					<view class="bg-white progress-box">
						
						<view class="flex margin-top-xs" v-for="(item,index) in degree_array" :key="index">
							<view class="progress-title basis-sm">{{item}}: {{data.today_degree_count[index].num?data.today_degree_count[index].num:"" }}</view>
							<view class="cu-progress radius">
								<view class="bg-gradual-blue" :style="[{ width:loading?data.today_degree_count[index].percent+'%':''}]">{{data.today_degree_count[index].percent}}%</view>
							</view>
						</view>
						
					</view>
				</view>
				
			</view>
			
			
			<view class="margin-top" v-if="tab_cur == 2">
				
				<text class="data-title" data-param="0">日期范围</text> 
				<view class="cu-tag round margin-right" :class="time_index==1?'line-blue':''"  data-param="1"  @tap="selectTime">当月</view>
				<view class="cu-tag round margin-right" :class="time_index==-7?'line-blue':''" data-param="-7" @tap="selectTime">近7天</view>
				<view class="cu-tag round margin-right" :class="time_index==-3?'line-blue':''" data-param="-3" @tap="selectTime">近3天</view>
				<view class="cu-tag round margin-right" :class="time_index==-4?'line-blue':''" data-param="-4" @tap="selectTime">当天</view>
				
				
				<view class="flex flex-wrap margin-top">
					
					<view class="basis-sm">
						<picker mode="date" :value="start_time"  @change="StartTimeChange">
							<view class="picker">
								<text class="text-time">{{start_time}}</text><text class="cuIcon-triangledownfill"></text>
								<text class="text-dot">到</text>
							</view>
						</picker>
					</view>
					
					<view class="basis-sm">
						<picker mode="date" :value="end_time" @change="EndTimeChange">
							<view class="picker">
								<text class="text-time">{{end_time}}</text>
								<text class="cuIcon-triangledownfill"></text>
							</view>
						</picker>
					</view>
				</view>
				
				
				
				<view class="progress-show-title margin-top cuIcon-selectionfill">销售数据</view>
				<view>
					<view class="bg-white progress-box">
						
						<view class="flex margin-top-xs">
							<view class="progress-title basis-sm">总共:{{data.deal_total?data.deal_total:""}}</view>
							<view class="cu-progress radius">
								<view class="bg-gradual-blue" :style="[{ width:loading?'100%':''}]">100%</view>
							</view>
						</view>
						
						<view class="flex margin-top-xs" v-for="(item,index) in data.deal_count" :key="index">
							<view class="progress-title basis-sm">{{item.name}}: {{item.num==0?"":item.num}}</view>
							<view class="cu-progress radius">
								<view class="bg-gradual-blue" :style="[{ width:loading?item.percent+'%':''}]">{{item.percent}}%</view>
							</view>
						</view>
						
					</view>
				</view>
				
				
				
				<view class="progress-show-title margin-top cuIcon-rankfill">客户来源分布</view>
					<view>
					<view class="bg-white progress-box">
						
						<view class="flex margin-top-xs" v-for="(item,index) in source_array" :key="index">
							<view class="progress-title basis-sm">{{item}}: {{data.source_count[index].num?data.source_count[index].num:"" }}</view>
							<view class="cu-progress radius">
								<view class="bg-gradual-blue" :style="[{ width:loading?data.source_count[index].percent+'%':''}]">{{data.source_count[index].percent}}%</view>
							</view>
						</view>
						
						
						
					</view>
				</view>
				
				<view class="progress-show-title margin-top cuIcon-sponsorfill">意向程度分布</view>
				<view>
					<view class="bg-white progress-box">
						
						<view class="flex margin-top-xs" v-for="(item,index) in degree_array" :key="index">
							<view class="progress-title basis-sm">{{item}}: {{data.degree_count[index].num?data.degree_count[index].num:"" }}</view>
							<view class="cu-progress radius">
								<view class="bg-gradual-blue" :style="[{ width:loading?data.degree_count[index].percent+'%':''}]">{{data.degree_count[index].percent}}%</view>
							</view>
						</view>
						
						
					</view>
				</view>
				
				
				<view class="progress-show-title margin-top cuIcon-friendaddfill">客户详细数据({{data.count}}人)</view>
				<view class="box margin-top-xs">
					
						<t-table @change="change">
							<t-tr>
								<t-th>日期</t-th>
								<t-th>客户</t-th>
								<t-th>来源</t-th>
								<t-th>操作</t-th>
							</t-tr>
							
							<scroll-view scroll-y  :style="[{height:'750rpx'}]" :scroll-with-animation="true" :enable-back-to-top="true"  @scrolltolower="scrolltolower">
							
								<t-tr v-for="item in data.customer_data" :key="value">
									<t-td>{{ item.created_at}}</t-td>
									<t-td>{{ item.name}} {{ degree_array[item.intention_house_degree]}}</t-td>
									<t-td>{{ source_array[item.source] }}</t-td>
									<t-td><text class="show" :data-detail="item.id" @tap="showModal">查看</text></t-td>
								</t-tr>
								
								<t-tr  v-if="noMoreData">
									<t-td>没有更多记录了</t-td>
								</t-tr>
							
							</scroll-view>
								
						</t-table>
						
				</view>
				
				
				
			</view>
			
			<view class="footer"></view>
			
			
		</view>
	</view>
</template>

<script>
	
	import tTable from '@/components/t-table/t-table.vue';
	import tTh from '@/components/t-table/t-th.vue';
	import tTr from '@/components/t-table/t-tr.vue';
	import tTd from '@/components/t-table/t-td.vue';

	export default {
			components: {
		            tTable,
		            tTh,
		            tTr,
		            tTd
		        },
			data() {   
				return {
					source_array: ['渠道','老带新','全民经纪人','客户自来','其他'],
					degree_array: ['A','B','C','D'],
					
					time_index: 1,
					start_time: '',
					end_time: '',
					
					user_id: -1,
					user_index: 0,
					user_info: [
						{nick_name: "-"}
					],
					
					tab_cur: 1,
					scrollLeft: 0,
					loading: false,
					data: [],
					
					page: 0,
					pageSize: 30,
					noMoreData: false,
				}
			},
			created(){
				
				let that = this;
			
				setTimeout(function() {
					that.loading = true
				}, 500)
				
				this.start_time = this.getMonthFirstDay()
				this.end_time =  this.getNowtime()
				
				this.getData()
				
			},
			methods: {
				showModal(e) {
					var detail = e.currentTarget.dataset.detail
					
					uni.navigateTo({
						url:  "../customer/info/info?detail="+detail
					});
					
				},
				scrolltolower(){
					
					if(!this.noMoreData){
						this.page = this.page + 1
						this.getData()
					}
					
				},
				tabSelect(e) {
					this.tab_cur = e.currentTarget.dataset.id;
					this.page = 0;
					this.getData()
					this.scrollLeft = (e.currentTarget.dataset.id - 1) * 60
				},
				PickerChange(e) {
					this.user_index = e.detail.value
					this.user_id = this.user_info[this.user_index].id
					this.page = 0;
					this.getData()
				},
				StartTimeChange(e){
					this.start_time = e.detail.value
					this.page = 0;
					this.getData()
				},
				EndTimeChange(e){
					this.end_time = e.detail.value
					this.page = 0;
					this.getData()
				},
				getNowtime(){
					var date1 = new Date()
					return date1.getFullYear()+"-"+this.twoNum(date1.getMonth()+1)+"-"+this.twoNum(date1.getDate()) //time1表示当前时间
				},
				getTime(m){
					var date1 = new Date()
					var date2 = new Date(date1)
					date2.setDate(date1.getDate()+m)
					return  date2.getFullYear()+"-"+this.twoNum(date2.getMonth()+1)+"-"+this.twoNum(date2.getDate())
				},
				getMonthFirstDay(){
					var date1 = new Date()
					return date1.getFullYear()+"-"+this.twoNum(date1.getMonth()+1)+"-01"
				},
				twoNum(num){
					if (num <= 9) {
					    num = "0" + num;
					 }
					 return num
				},
				selectTime(e){
					var _this = this
					var param = e.currentTarget.dataset.param
					_this.time_index =  param
					// 更高时间区间
					switch(param){
						case  "1": _this.start_time = _this.getMonthFirstDay();break;
						case "-7": _this.start_time = _this.getTime(-7);break;	
						case "-3": _this.start_time = _this.getTime(-3);break;	
						case "-4": _this.start_time = _this.getNowtime();break;	
					}
					
					_this.end_time =  _this.getNowtime();
					_this.page = 0;
					_this.getData();
					
				},
				getData(){
					let _this = this
					uni.showLoading({
						title: '请求数据中'
					});
					
					_this.$api.GetStatistics({page:_this.page,page_size:_this.pageSize,user_id: _this.user_id,start_time:_this.start_time,end_time:_this.end_time,tab_cur:parseInt(_this.tab_cur)}).then((res)=>{
						
						if(res.code == 200){
							
							if (_this.user_id == -1){
								// 将服务器user_info单独
								_this.user_info = res.data.user_info
							}
							
							if (_this.page == 0){
								_this.data = res.data
							}else{
								// 触底加载数据
								if (res.data.customer_data == null){
									_this.noMoreData = true
								}else{
									_this.data.customer_data = _this.data.customer_data.concat(res.data.customer_data)
								}
							}
						}else{
							uni.showToast({
							    title: res.msg,
							    icon: 'none',
								duration: 2000
							});
						}
						
					});
					
					uni.hideLoading();
				},
				
			}
	}
</script>

<style>
	.con-title{
		font-weight: 500;
	}
	.select-name,.select-name-icon{
		color: rgb(0, 0, 0);
		font-size: 40rpx;
	}
	.nav-title{
		font-weight: 400;
		color: rgb(0,0,0);
	}
	.nav .cu-item {
		height: 70rpx;
		display: inline-block;
		line-height: 70rpx;
		margin: 0 40rpx 0 0;
		padding: 0rpx;
}
.flex-box{
	border-right: 1rpx solid rgb(216, 216, 216);
}
.tap1-title .t1{
	font-size: 35rpx;
	font-weight: 400;
	color: rgb(0,0,0);
	
}

.tap1-title .t2{
	float: right;
}

.nav .cu-item.cur .nav-title{
	color: #0081FF;
}
.box-num{
	color: rgb(0,0,0);
	font-size: 40rpx;
}
.progress-title{
	border-right: 1rpx solid rgb(216, 216, 216);
}
.cu-progress{
	position: relative;
	top:6rpx;
}
.progress-show-title{
	font-size: 30rpx;
	font-weight: 400;
	color: rgb(0,0,0);
}
.data-title{
	font-size: 30rpx;
	padding-right: 26rpx;
	position: relative;
	top: 3rpx;
}
.picker{
	font-size: 30rpx;
}
.footer{height: 100rpx;}

.text-time{
	font-size: 32rpx;
	color: rgb(0,0,0);
}
.text-dot{
	padding-left: 25rpx;
}
.show{
	color: rgb(40, 176, 195);
}
</style>
