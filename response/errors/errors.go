package errors

type SelfError struct {
	Err string
}

func NewSelfError(err string) SelfError {
	return SelfError{Err: err}
}

func (s SelfError) Error() string {
	return s.Err
}

var (
	LoginAccountOrPasswordError = NewSelfError("账号或密码错误")

	ServerError = NewSelfError("服务器错误")

	NotFoundError = NewSelfError("未找到该资源")

	ParameterNotValid = NewSelfError("参数错误")

	// image

	ErrImageType  = NewSelfError("图片类型错误")
	GetImageError = NewSelfError("获取图片失败")

	// activity-type

	ErrActivityType      = NewSelfError("活动类型错误")
	GetActivityTypeError = NewSelfError("获取活动类型失败")

	// activity

	GetActivityError        = NewSelfError("获取活动类型失败")
	GetPopularActivityError = NewSelfError("获取热门活动失败")
)

var (
	ErrStatus = map[SelfError]int{
		LoginAccountOrPasswordError: 400,

		ServerError: 500,
	}
)

var (
	SelfErrStatus = map[SelfError]int{
		LoginAccountOrPasswordError: 400100,

		ServerError: 500100,
	}
)
