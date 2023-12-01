package consts

import (
	"github.com/go-playground/validator/v10"
)

var PayStatusEnum = make(map[string][]uint8)

func init() {
	PayStatusEnum["pay_notify_status"] = []uint8{WAITING, SUCCESS, REFUND, CLOSED, FAILURE, RequestSuccess, RequestFailure}
	PayStatusEnum["pay_notify_type"] = []uint8{OrderType, RefundType}
	PayStatusEnum["common_status"] = []uint8{Enable, Disable}
}

func AddEnum(key string, enum []uint8) {
	PayStatusEnum[key] = enum
}

func DeleteEnum(key string) {
	delete(PayStatusEnum, key)
}

// InEnumValidate 自定义验证器,校验状态值是否在PayStatusEnum 中
func InEnumValidate(fl validator.FieldLevel) bool {
	enumArray, ok := PayStatusEnum[fl.Param()]
	if !ok {
		// 如果找不到对应的枚举类型，验证失败
		return false
	}
	// 检查字段值是否在枚举数组中
	for _, val := range enumArray {
		if fl.Field().Uint() == uint64(val) {
			return true
		}
	}
	return false
}
