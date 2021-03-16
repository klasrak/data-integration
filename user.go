package data_integration

// User represents User domain type
type User struct {
	ID       string `json:"id,omitempty" bson:"_id,omitempty"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
