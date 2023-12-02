import React, { useEffect, useState } from 'react'
import {View ,Text,Picker, ScrollView,Input} from '@tarojs/components'
import { AtButton, AtList, AtListItem} from 'taro-ui'
import {useDidShow, useDidHide } from '@tarojs/taro'
import './index.scss'
import Taro from '@tarojs/taro'; 
import {createDemoOrder, DemoOrderInfo, DemoOrderPageResponse, DemoOrderVO,getDemoOrderPage,PageParam, submitOrder, SubmitOrderVO,SubmitOrderResponse, DisplayContent} from '@/api/demo/pay'

export default function () {

  const [demoOrderList, setDemoOrderList] = useState<DemoOrderPageResponse | null>(null);
  const [openid, setOpenid] = useState('');
  const [isEditing, setIsEditing] = useState(false);
  interface Product {
    id: number;
    name: string;
  }
  
  const productList: Product[] = [
    {
      id: 1,
      name: '华为手机,价格：0.01元'
    },
    {
      id: 2,
      name: '小米电视,价格：0.1元'
    },
    {
      id: 3,
      name: '苹果手表,价格：1元'
    },
    {
      id: 4,
      name: '华硕笔记本,价格：10元'
    },
    {
      id: 5,
      name: '蔚来汽车,价格：2000元'
    }
  ];
  
  const [selectedProduct, setSelectedProduct] = useState<Product | undefined>(undefined);

  const handlePickerChange = (e) => {
    const selectedValue = e.detail.value;
    setSelectedProduct(productList[selectedValue]);
  };

  const orderConfirm=async()=>{
    if (!selectedProduct) {
      Taro.showToast({
        title: '请先选择商品',
        icon: 'none',
        duration: 2000,
      });
      return
    }
    const param: DemoOrderVO = {spuId: selectedProduct.id}; // 合法，空对象
    createDemoOrder(param).then(response => {
        console.info(response)
        fetchData();
    })
      
    
  }

  // 可以使用所有的 React Hooks
  useEffect(() => {})
  useDidShow(() => {
    fetchData();
  });
  const fetchData = () => {
    const query: PageParam = {
      pageSize: 10,
      page: 1,
    };

    getDemoOrderPage(query)
      .then((response: DemoOrderPageResponse) => {
        // 成功获取数据时更新状态
        setDemoOrderList(response);
      })
      .catch((error) => {
        // 处理请求失败的情况
        console.error('请求失败:', error);
      });
  };
  // 对应 onHide
  useDidHide(() => {})

  const handleGoToPayment = (demoOrder: DemoOrderInfo) => {
    Taro.showModal({
      title : '通知',
      content:'确认支付?',
      complete:async (res)=>{
        if (res.cancel) {
          return
        }
        if (res.confirm) {
          if (!openid) {
            Taro.showToast({
              title: '请填入openid' 
            })
            return
          }
          const data:SubmitOrderVO = {
            id:demoOrder.payOrderId,
            channelCode:'WxPub',
            channelExtras:{
                openid : openid
            }
          }
          const res: SubmitOrderResponse = await submitOrder(data)
          const payContent:DisplayContent = JSON.parse(res.displayContent);
          console.info(payContent)
          Taro.requestPayment({
            timeStamp:payContent.TimeStamp,
            nonceStr:payContent.NonceStr,
            package:payContent.PackageValue,
            signType:payContent.SignType,
            paySign:payContent.PaySign,
            success:function(response){
                if(response.errMsg=="requestPayment:ok"){
                    Taro.navigateTo({
                        url:'../paymentFinish/paymentFinish'
                    })
                }
            },
            fail:function(response){
                Taro.showToast({
                    icon:'error',
                    title:'您取消了支付'
                })
            }
        })
        }
      }
    })
  }

  const handleInputChange = (e) => {
    setOpenid(e.detail.value)
  }

  const handleEditClick = () => {
    setIsEditing(true);
  };

  const handleConfirmClick = () => {
    setIsEditing(false);
  };

  return (
    <View className='container'>
      <View className='page-section'>
        <View>
          <Text>请填入openid:</Text>
          <Input
            type='text'
            placeholder='openid'
            disabled={!isEditing} // 根据编辑状态禁用/启用输入框
            onInput={handleInputChange}
          />
          {isEditing ? (
            <AtButton onClick={handleConfirmClick}>确认</AtButton>
          ) : (
            <AtButton onClick={handleEditClick}>编辑</AtButton>
          )}
        </View>
        <Picker mode='selector' range={productList} range-key='name' onChange={(e) => handlePickerChange(e)}>
          <View className='picker'>
            商品选择:
            {selectedProduct ? (
              <View className='selected-info'>
                <Text>{selectedProduct.name}</Text>
                <Text className='price'>价格: {selectedProduct.price}</Text>
              </View>
            ) : '未选择'}
          </View>
        </Picker>
      </View>
      <AtButton onClick={orderConfirm}>
        创建新订单
      </AtButton>
      <View className='order-list'>
        {/* 在这里使用 demoOrderList 渲染页面 */}
        {demoOrderList && (
          <AtList>
            <Text className='order-count'>总订单数: {demoOrderList.total}</Text>
            <ScrollView className='order-scroll-view'>
              {demoOrderList.data.map((demoOrder) => (
                <View>
                  <AtListItem onClick={() => handleGoToPayment(demoOrder)} arrow='right' title={demoOrder.spuName} note={
                    `价格: ${demoOrder.price}分, 
                    支付渠道: ${demoOrder.payChannelCode}, 
                    支付状态: ${demoOrder.payStatus ? '已支付' : '未支付'}`
                  }>
                  </AtListItem>
                </View>
              ))}
            </ScrollView>
          </AtList>
        )}
      </View>
    </View>
  )
    
}