package handler

const (
	// Common messages
	ErrFailedToEncodeResponse = "レスポンスのエンコードに失敗しました"
	ErrInternalServer         = "サーバー内部でエラーが発生しました"
	ErrUserNotFound           = "ユーザーが見つかりません"

	// Score related messages
	SuccessMsgScoreRegistered    = "スコアが正常に登録されました"
	ErrMsgInvalidRequestBody     = "リクエストボディが不正です"
	ErrMsgInvalidUserIDParameter = "ユーザーIDが不正です"
	ErrFailedToRegisterScore     = "スコアの登録に失敗しました"
	ErrMsgInvalidSortByParameter = "不正なソート対象のカラムです"
	ErrMsgInvalidStartParameter  = "不正なランキングの開始位置です"
	ErrMsgInvalidLimitParameter  = "不正なランキングの取得件数です"

	// User related messages
	ErrMsgStudentNumberRequired = "student_numberが指定されていません"
)
