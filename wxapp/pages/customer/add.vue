<template>
	<view>
		<cu-custom bgColor="bg-gradual-blue" :isBack="true">
			<block slot="backText">返回</block>
			<block slot="content" class="con-title">新增客户</block>
		</cu-custom>
		
		<form>
			
			<view class="cu-form-group margin-top">
				<view class="title">姓名:</view>
				<input  placeholder="客户姓名" name="input" v-model="name"></input>
			</view>
			
			<radio-group class="block" @change=RadioChange>
				<view class="cu-form-group">
					<view class="title">性别:</view>
					<view class="padding-right">
						
						<label><radio class="margin-left blue" :class="sex==0?'checked':''" :checked="sex==0?true:false" value=0></radio>先生</label>
						<label><radio class="margin-left blue" :class="sex==1?'checked':''" :checked="sex==1?true:false" value=1></radio>女士</label>
						
					</view>
				</view>
			</radio-group>
				
			
			<view class="cu-form-group">
				<view class="title">客户来源:</view>
				<picker @change="sourceChange" v-model="source" :range="source_array">
					<view class="picker">
						{{source_array[source]}}
					</view>
				</picker>
			</view>
			
			<view class="cu-form-group">
				<view class="title">客户分组:</view>
				<picker @change="typeChange" v-model=link_type :range=type_array>
					<view class="picker">
						{{type_array[link_type]}}
					</view>
				</picker>
			</view>
			
			<!-- <view class="cu-form-group">
				<view class="title">身份证:</view>
				<input  placeholder="身份证登记" name="input" v-model="id_card"></input>
			</view> -->
			
			<view class="cu-form-group">
				<view class="title">手机号码:</view>
				<input placeholder="客户手机号码" name="input" v-model="phone"></input>
				<view class="cu-capsule radius">
					<view class='cu-tag bg-blue '>
						+86
					</view>
					<view class="cu-tag bg-blue">
						中国大陆
					</view>
				</view>
			</view>
			
			<view class="cu-form-group">
				<view class="title">地址:</view>
				<input placeholder="客户地址" name="input" v-model="addr"></input>
			</view>
			
				
				
			<view class="cu-form-group margin-top">
				<view class="title">购房预算:</view>
				<picker @change="modelChange" v-model="intention_house_model" :range="model_array">
					<view class="picker">
						{{model_array[intention_house_model]}}
					</view>
				</picker>
			</view>
			
			<view class="cu-form-group">
				<view class="title">意向程度:</view>
				<picker @change="degreeChange" v-model="intention_house_degree" :range="degree_array">
					<view class="picker">
						{{degree_array[intention_house_degree]}}
					</view>
				</picker>
			</view>
			
			
			<!-- !!!!! placeholder 在ios表现有偏移 建议使用 第一种样式 -->
			<view class="cu-form-group margin-top">
				<textarea maxlength="-1"  placeholder="其他信息备注" v-model="other_info"></textarea>
			</view>
			
			
			<view  class="padding-xl" @tap="submit">
				<button class="cu-btn block bg-gradual-blue margin-tb-sm lg"><text>提&nbsp;&nbsp;交</text></button>
					 <!-- <text class="cuIcon-loading2 cuIconfont-spin"></text> -->
			</view>
			
			
		</form>
		
		
	</view>	
	
	
	
</template>

<script>
	export default {
		data(){
			return {
				
				source_array: ['渠道','老带新','全民经纪人','客户自来','其他'],
				model_array: ['20万以下', '20~40万', '40~60万','60~100万','100以上'],
				type_array: ['待联系','已联系','复核','已成交','公共池','无效客户'],
				degree_array: ['意向程度:A级','意向程度:B级','意向程度:C级','意向程度:D级'],
				
				name:'',
				sex: 0,
				source: 0,
				link_type: 0,
				id_card:'',
				phone:'',
				addr:'',
			
				intention_house_model: 0,
				intention_house_degree: 0,
			
				other_info:''
				
				}
		},
	    onLoad(option) {
			//打印出上个页面传递的参数。
	        // console.log(option.param); 
	    },
		methods: {
			modelChange(e) {
				this.intention_house_model = e.detail.value
			},
			sourceChange(e){
				this.source = e.detail.value
			},
			degreeChange(e){
				this.intention_house_degree = e.detail.value
			},
			typeChange(e){
				this.link_type = e.detail.value
			},
			RadioChange(e) {
				this.sex = e.detail.value
			},
			submit(){
				let _this = this
				uni.showLoading({
					title: '提交中'
				});
				_this.$api.CustomerAdd({name: _this.name,sex:parseInt(_this.sex),source:parseInt(_this.source),id_card:_this.id_card,phone:_this.phone,addr:_this.addr,link_type:parseInt(_this.link_type),intention_house_model:parseInt(_this.intention_house_model),intention_house_degree:parseInt(_this.intention_house_degree),other_info:_this.other_info}).then((res)=>{
					uni.hideLoading();
					
					uni.showToast({
					    title: res.msg,
						icon: 'none',
					    mask: true
					});
					if(res.code == 200){
						
						uni.reLaunch({
									url: '../index/index',
								});
					
					}
				})
			}
		}
	}
</script>

<style>
	.bg-gradual-cyan {
		background-image: linear-gradient(45deg, #76EEC6, #48D1CC);
		color: #ffffff;
	}
	.con-title{
		font-weight: 500;
	}
</style>
