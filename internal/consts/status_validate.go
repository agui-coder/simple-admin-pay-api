package consts

import (
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/go-playground/validator/v10"
)

var EnumValidate = make(map[string]map[string]int32)

func init() {
	EnumValidate["pay_status"] = pay.PayStatus_value
	EnumValidate["pay_type"] = pay.PayType_value
}

// InEnumValidate 自定义验证器,校验状态值是否在EnumValidate 中
func InEnumValidate(fl validator.FieldLevel) bool {
	enumMap, ok := EnumValidate[fl.Param()]
	if !ok {
		// 如果找不到对应的枚举类型，验证失败
		return false
	}
	// 检查字段值是否在枚举数组中
	for _, val := range enumMap {
		if fl.Field().Uint() == uint64(val) {
			return true
		}
	}
	return false
}
