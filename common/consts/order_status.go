package consts

const (
	WAITING uint8 = iota * 10
	SUCCESS
	REFUND
	CLOSED
	FAILURE
	RequestSuccess
	RequestFailure
)

const (
	OrderType = iota + 1
	RefundType
)
