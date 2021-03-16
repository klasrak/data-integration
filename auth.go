package data_integration

// AccessDetails represents AccessDetails domain type
type AccessDetails struct {
	TokenUUID string
	UserID    string
	Email     string
}

// TokenDetails represents TokenDetails domain type
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	TokenUUID    string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}
