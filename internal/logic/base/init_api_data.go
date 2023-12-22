package base

import (
	"github.com/suyuan32/simple-admin-common/enum/common"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-core/rpc/types/core"
)

func (l *InitDatabaseLogic) insertApiData() (err error) {
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Pay"),
		Path:        pointy.GetPointer("/demo-order/page"),
		Description: pointy.GetPointer("获取支付demo列表"),
		ApiGroup:    pointy.GetPointer("demo-order"),
		Method:      pointy.GetPointer("POST"),
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Pay"),
		Path:        pointy.GetPointer("/demo-order/get"),
		Description: pointy.GetPointer("获取支付demo信息"),
		ApiGroup:    pointy.GetPointer("demo-order"),
		Method:      pointy.GetPointer("POST"),
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Pay"),
		Path:        pointy.GetPointer("/demo-order/create"),
		Description: pointy.GetPointer("创建示例订单"),
		ApiGroup:    pointy.GetPointer("demo-order"),
		Method:      pointy.GetPointer("POST"),
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Pay"),
		Path:        pointy.GetPointer("/demo-order/refund"),
		Description: pointy.GetPointer("退款示例订单"),
		ApiGroup:    pointy.GetPointer("demo-order"),
		Method:      pointy.GetPointer("POST"),
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Pay"),
		Path:        pointy.GetPointer("/order/get"),
		Description: pointy.GetPointer("获取支付信息"),
		ApiGroup:    pointy.GetPointer("order"),
		Method:      pointy.GetPointer("POST"),
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.CoreRpc.CreateApi(l.ctx, &core.ApiInfo{
		ServiceName: pointy.GetPointer("Pay"),
		Path:        pointy.GetPointer("/order/submit"),
		Description: pointy.GetPointer("提交支付"),
		ApiGroup:    pointy.GetPointer("order"),
		Method:      pointy.GetPointer("POST"),
	})
	return err
}

func (l *InitDatabaseLogic) insertMenuData() error {
	menuData, err := l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(common.DefaultParentId),
		Name:      pointy.GetPointer("PayManagementDirectory"),
		Component: pointy.GetPointer("LAYOUT"),
		Path:      pointy.GetPointer("/pay_dir"),
		Sort:      pointy.GetPointer(uint32(1)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("支付管理"),
			Icon:               pointy.GetPointer("ant-design:account-book-outlined"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(0)),
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(menuData.Id),
		Path:      pointy.GetPointer("/demo"),
		Name:      pointy.GetPointer("PayDemo"),
		Component: pointy.GetPointer("/pay/demo/index"),
		Sort:      pointy.GetPointer(uint32(1)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("支付demo"),
			Icon:               pointy.GetPointer("ant-design:account-book-filled"),
			HideMenu:           pointy.GetPointer(false),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(1)),
	})
	_, err = l.svcCtx.CoreRpc.CreateMenu(l.ctx, &core.MenuInfo{
		Level:     pointy.GetPointer(uint32(2)),
		ParentId:  pointy.GetPointer(menuData.Id),
		Path:      pointy.GetPointer("/cashier"),
		Name:      pointy.GetPointer("Cashier"),
		Component: pointy.GetPointer("/pay/cashier/index"),
		Sort:      pointy.GetPointer(uint32(2)),
		Disabled:  pointy.GetPointer(false),
		Meta: &core.Meta{
			Title:              pointy.GetPointer("收银台"),
			Icon:               pointy.GetPointer("ant-design:account-book-twotone"),
			HideMenu:           pointy.GetPointer(true),
			HideBreadcrumb:     pointy.GetPointer(false),
			IgnoreKeepAlive:    pointy.GetPointer(false),
			HideTab:            pointy.GetPointer(false),
			CarryParam:         pointy.GetPointer(false),
			HideChildrenInMenu: pointy.GetPointer(false),
			Affix:              pointy.GetPointer(false),
		},
		MenuType: pointy.GetPointer(uint32(1)),
	})
	if err != nil {
		return err
	}
	return nil
}
