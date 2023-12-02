import defHttp from "../request"


export interface DemoOrderVO {
  spuId: number
  createTime?: number
}

export interface DemoOrderInfo {
  createAt: number;
  id: number;
  payChannelCode: string;
  payOrderId: number;
  payRefundId: number;
  payStatus: boolean;
  payTime: number;
  price: number;
  refundPrice: number;
  refundTime: number;
  spuId: number;
  spuName: string;
  updateAt: number;
}

export interface DemoOrderPageResponse {
  total: number;
  data: DemoOrderInfo[];
}

export interface PageParam {
  pageSize?: number
  page?: number
}

export interface SubmitOrderVO {
  id:number
  channelCode: string
  channelExtras: any
}

export interface DisplayContent {
    Appid: string;
    TimeStamp: string;
    NonceStr: string;
    PackageValue: string;
    SignType: any;
    PaySign: string;
  }

export interface SubmitOrderResponse {
    status: number;
    displayMode: string;
    displayContent: string; 
    parsedDisplayContent?: DisplayContent;
  }

// 创建示例订单
export function createDemoOrder(data: DemoOrderVO) {
  return defHttp.post({
    url: '/pay/demo-order/create',
    data: data
  })
}

// 获得示例订单
export function getDemoOrder(id: number) {
  return defHttp.get({
    url: '/pay/demo-order/get/' + id
  })
}

// 获得示例订单分页
export function getDemoOrderPage(query: PageParam): Promise<DemoOrderPageResponse> {
  return defHttp.get({
    url: '/pay/demo-order/page',
    params: query
  })
}
// 提交支付订单
export async function submitOrder(data:SubmitOrderVO): Promise<SubmitOrderResponse>{
  return await defHttp.post({ url: '/pay/order/submit', data })
}
