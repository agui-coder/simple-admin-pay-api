package model

import (
	"context"
	"encoding/json"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

// 支付方式
const (
	Url       = "url"
	Iframe    = "iframe"
	From      = "from"
	QrCode    = "qr_code"
	QrCodeUrl = "qr_code_url"
	BarCode   = "bar_code"
	App       = "app"
)

// 支付状态
const (
	WAITING uint8 = (iota + 1) * 10
	SUCCESS
	REFUND
	CLOSED
	ERROR
)

// 支付渠道
const (
	Wx       = "Wx"
	WxPub    = "WxPub"
	WxLite   = "WxLite"
	WxApp    = "WxApp"
	WxNative = "WxNative"

	Ali       = "Ali"
	AlipayPc  = "AlipayPc"
	AlipayWap = "AlipayWap"
	AlipayApp = "AlipayApp"
	AlipayQr  = "AlipayQr"
	AlipayBar = "AlipayBar"
	Mock      = "mock"
	Wallet    = "wallet"
)

type (
	// ClientConfig 支付客户端配置
	ClientConfig interface {
		Validate() error
	}

	// Properties 支付配置
	Properties struct {
		OrderNotifyUrl  string
		RefundNotifyUrl string
		OrderNoPrefix   string
		RefundNoPrefix  string
	}

	// Client 支付客户端
	Client interface {
		// Init 初始化
		Init() error
		// UnifiedOrder 统一下单
		UnifiedOrder(context.Context, OrderUnifiedReq) (*OrderResp, error)
		// GetId 获取渠道id
		GetId() uint64
		// GetOrder 获得支付订单
		GetOrder(context.Context, string) (*OrderResp, error)
		// Refresh 刷新配置
		Refresh(config ClientConfig) error
		// UnifiedRefund 退款
		UnifiedRefund(context.Context, RefundUnifiedReq) (*RefundResp, error)
		// ParseOrderNotify 解析支付回调
		ParseOrderNotify(map[string][]string, []byte) (*OrderResp, error)
	}

	OrderUnifiedReq struct {
		UserIp        string
		OutTradeNo    string
		Subject       string
		Body          string
		NotifyUrl     string
		ReturnUrl     string
		Price         int32
		ExpireTime    time.Time
		ChannelExtras map[string]string
		DisplayMode   string
	}

	OrderResp struct {
		Status           uint8
		OutTradeNo       string
		ChannelOrderNo   string
		ChannelUserId    *string
		SuccessTime      time.Time
		RawData          any
		DisplayMode      *string
		DisplayContent   *string
		ChannelErrorCode *string
		ChannelErrorMsg  *string
	}
	RefundUnifiedReq struct {
		OutTradeNo  string
		OutRefundNo string
		Reason      string
		PayPrice    int32
		RefundPrice int32
		NotifyUrl   string
	}

	RefundResp struct {
		Status           uint8
		OutRefundNo      string
		ChannelRefundNo  string
		SuccessTime      time.Time
		RawData          any
		ChannelErrorCode string
		ChannelErrorMsg  string
	}
)

// WaitingOf 创建等待支付订单
func WaitingOf(displayMode *string, outTradeNo string, rawData any) (*OrderResp, error) {
	logx.Debugf("wxRsp: %#v", rawData)
	jsonData, err := json.Marshal(rawData)
	if err != nil {
		return nil, err
	}
	return &OrderResp{
		Status:         WAITING,
		DisplayMode:    displayMode,
		DisplayContent: pointy.GetPointer(string(jsonData)),
		OutTradeNo:     outTradeNo,
		RawData:        rawData,
	}, nil
}

// SuccessOf 创建支付成功订单
func SuccessOf(channelOrderNo string, channelUserId string, successTime time.Time, outTradeNo string, rawData any) *OrderResp {
	return &OrderResp{
		Status:         SUCCESS,
		ChannelOrderNo: channelOrderNo,
		ChannelUserId:  &channelUserId,
		SuccessTime:    successTime,
		OutTradeNo:     outTradeNo,
		RawData:        rawData,
	}
}

// Of 创建支付订单
func Of(status uint8, channelOrderNo string, channelUserId *string, successTime time.Time, outTradeNo string, rawData any) *OrderResp {
	return &OrderResp{
		Status:         status,
		ChannelOrderNo: channelOrderNo,
		ChannelUserId:  channelUserId,
		SuccessTime:    successTime,
		OutTradeNo:     outTradeNo,
		RawData:        rawData,
	}
}
