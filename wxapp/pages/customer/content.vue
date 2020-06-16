<template>
	
	<view>
		
		<cu-custom bgColor="bg-gradual-blue" :isBack="true">
			<block slot="backText">返回</block>
			<block slot="content" class="con-title">{{title}}</block>
		</cu-custom>
		
		<view class="cu-bar bg-white search fixed" :style="[{top:CustomBar + 'px'}]">
			<view class="search-form round">
				<text class="cuIcon-search"></text>
				<input type="text" placeholder="姓名/电话" confirm-type="search"  v-model="searchKey"></input>
			</view>
			<view class="action" @tap="searchCustomer">
				<button class="cu-btn bg-gradual-blue shadow-blur round">搜索</button>
			</view>
		</view>
		
		<scroll-view scroll-y  :style="[{height:'calc(100vh - '+ CustomBar + 'px - 50px)'}]" :scroll-with-animation="true" :enable-back-to-top="true"  @scrolltolower="scrolltolower">
		
		<view class="cu-list menu-avatar">
			
			
			<view class="cu-item padding-xl radius shadow bg-white" :class="modalName=='move-box-'+ index?'move-cur':''"  v-for="(item,index) in list" :key="index" @touchstart="ListTouchStart" @touchmove="ListTouchMove" @touchend="ListTouchEnd" :data-target="'move-box-' + index">
					<!-- <view class="cu-avatar round lg" style="background-image:url(https://ossweb-img.qq.com/images/lol/web201310/skin/big10001.jpg);"></view> -->
					<view class="cu-avatar round lg" style="background-color: #ADADAD;" :data-phone="item.phone" @tap="call" :data-detail="item.id"  :data-name="item.name" :data-ctype="item.link_type">{{item.name.substr(0,1)}}
					
					
					<view class="cu-tag badge" :class="item.sex==1?'cuIcon-female bg-pink':'cuIcon-male bg-blue'"></view>
					</view>
					<view class="content">
						
						<view class="text-grey">
							<view class="text-cut" @tap="showModal" :data-detail="item.id"   data-target="userDetail">{{item.name}} </view>
							
							<view class="cu-tag round line-blue sm">{{source_array[item.source]}}</view>
							<view class="cu-tag round line-blue sm">{{model_array[item.intention_house_model]}}</view>
							<view class="cu-tag round line-blue sm">{{degree_array[item.intention_house_degree]}}</view>
						</view>
						
						
						<view class="text-gray text-sm flex">
							
							<view class="text-cut">
								<text class="cuIcon-form text-grey margin-right-xs"></text>
								{{item.other_info}} 
							</view> 
							
						</view>
						
						<!-- <view class="text-gray text-sm flex">
							
							<view class="text-cut">
								<text class="cuIcon-mobilefill text-grey margin-right-xs"></text>
								{{item.phone}} 
								
								<text class="cuIcon-locationfill text-grey margin-right-xs margin-left"></text>
								{{item.addr}} 
							</view> 
							
						</view> -->
						
						<view class="text-gray text-sm flex">
							
							<view class="text-cut">
								
								<text class="cuIcon-time text-grey margin-right-xs"></text>
								{{item.created_at}}
								
								
							</view> 
							
						</view>
						
					</view>
					<view class="action">
						<view class="text-grey text-xs"><text class="text-grey margin-right-xs"></text></view>
						<!-- <view class="cu-tag round bg-grey sm">5</view> cuIcon-right-->
					</view>
					
					
					<view class="move radius shadow ">
						<view class="bg-grey"  @tap="follow"  :data-detail="item.id" :data-name="item.name" :data-ctype="item.link_type">跟进</view>
						<view class="bg-red"   @tap="deal"    :data-detail="item.id" :data-name="item.name">成交</view>
					</view>
					
					
			</view>
			
			<view class="noMoreData bg-white text-center padding" v-if="noMoreData">没有更多记录了</view>
			
		</view>
		
		</scroll-view>
		
	</view>
</template>

<script>
	export default {
			data() {
				return {
					source_array: ['渠道','老带新','全民经纪人','客户自来','其他'],
					model_array: ['20万以下', '20~40万', '40~60万','60~100万','100以上'],
					type_array: ['待联系','已联系','复核','已成交','公共池','无效客户'],
					degree_array: ['A级','B级','C级','D级'],
					title: '',
					type: 0,
					page: 0,
					pageSize: 30,
					list: [],
					searchKey: "",
					flag: false,
					CustomBar: this.CustomBar,
						
					noMoreData: false,
					
					modalName: null,
					listTouchStart: 0,
					listTouchDirection: null,
					
					modalName: null,
				};
			},
			onLoad(options) {
				
				this.title = options.title
				this.type = options.param
				
				this.list = []
				this.page = 0
				this.noMoreData = false
				this.getCustList()
			},
			methods:{
			
				hideModal(e) {
					this.modalName = null
				},
				
				// ListTouch触摸开始
				ListTouchStart(e) {
					this.listTouchStart = e.touches[0].pageX
				},
				
				// ListTouch计算方向
				ListTouchMove(e) {
					// this.listTouchDirection = e.touches[0].pageX - this.listTouchStart > 50 ? 'right' : 'left'
				
				var offset =  e.touches[0].pageX - this.listTouchStart
				
					if (offset > 80) {
					 
						this.listTouchDirection = 'right'
					  
					} else if (offset < -80) {
					  
					   this.listTouchDirection = 'left'
					 
					}
				
				},
				
				// ListTouch计算滚动
				ListTouchEnd(e) {
					if (this.listTouchDirection == 'left') {
						this.modalName = e.currentTarget.dataset.target
					} else {
						this.modalName = null
					}
					this.listTouchDirection = null
				},
				
				searchCustomer(){
					this.list = []
					this.page = 0
					this.noMoreData = false
					this.getCustList()
				},
				getCustList(){
					let _this = this
					uni.showLoading({
						title: '获取客户信息中'
					});
					_this.$api.GetCustomerInfo({page:_this.page,pageSize:_this.pageSize,search:_this.searchKey,type:_this.type}).then((res)=>{
						uni.hideLoading();
							
							
						if(res.code == 200){
							
							if(res.data == null){
								_this.noMoreData = true
							}else{
								_this.list = _this.list.concat(res.data)
							
								console.log(res.data)
							}  
							
						}else{
							
							uni.showToast({
							    title: res.msg,
							    icon: 'none',
								duration: 2000
							});
							
						}
					})
				},
				
				call(e){
					var _this = this
					let phone = e.currentTarget.dataset.phone
					uni.makePhoneCall({
					    phoneNumber: phone,
						success:function(){
							//  弹框填写跟进信息
							var name = e.currentTarget.dataset.name
							var detail = e.currentTarget.dataset.detail
							var ctype = e.currentTarget.dataset.ctype
							uni.navigateTo({
								url:  "./follow/follow?name="+name+ "&detail="+detail+"&ctype="+ctype
							});
							
						},
						fail:function(){
								uni.showToast({
								    title: "取消了拨打电话",
								    icon: 'none',
									duration: 2000
								});
						},
						complete:function(){
						}
					});
					
				},
				
				// 当每次滑动到底部再次请求接口，实现上啦加载更多刷新
				scrolltolower(){
					
					if(!this.noMoreData){
						this.page = this.page + 1
						this.getCustList()
					}
						
				},
				
				deal(e){
					var name = e.currentTarget.dataset.name
					var detail = e.currentTarget.dataset.detail
					
					uni.navigateTo({
					    url:  "./deal/deal?name="+name+ "&detail="+detail
					});
				},
					
				follow(e){
					var name = e.currentTarget.dataset.name
					var detail = e.currentTarget.dataset.detail
					var ctype = e.currentTarget.dataset.ctype
					uni.navigateTo({
						url:  "./follow/follow?name="+name+ "&detail="+detail+"&ctype="+ctype
					});
				},
				showModal(e) {
					var detail = e.currentTarget.dataset.detail
					
					uni.navigateTo({
						url:  "./info/info?detail="+detail
					});
					
				},
			}
	}
</script>

<style>
	page {
			padding-top: 100upx;
		}
	
	.info-scroll{max-height: 950rpx;}
	.box-modal{
		padding: 0rpx 40rpx 0rpx 40rpx;
		text-align: left;
	}
	.box-modal-btn{
		margin-top: 40rpx;
		margin-bottom: 20rpx;
	}
	.move{font-size: 35rpx;}	
	.box view .cu-bar {
		margin-top: 20upx;
	}
	
	.cu-list.menu-avatar>.cu-item .content .cu-tag.sm {
		font-size: 25rpx;
		height: 33rpx;
		line-height: 33rpx;
	}
	
	.cu-list.menu-avatar>.cu-item {
		margin-top: 15rpx;
		height: 160rpx;
	}
	
	/* 回到顶部 */
	.top {
		position: relative;
		display: none; /* 先将元素隐藏 */
	}
	
	.topc {
		position: fixed;
		right: 0;
		background: #F0F0F0;
		top: 70%;
		height: 50px;
		line-height: 50px;
	}
	.noMoreData{
		text-align: center;
	}
	.con-title{
		font-weight: 500;
	}
</style>