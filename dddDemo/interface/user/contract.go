package user

import (
    "context"
    "encoding/json"
    "time"
)

const UserKey = "user"


type Service interface {

    // Login 登录相关，使用用户名密码登录，获取完成User信息
    Login(ctx context.Context, user *User) (*User, error)
}


// User 代表一个用户，注意这里的用户信息字段在不同接口和参数可能为空
type User struct {
    ID        int64     `gorm:"column:id;primaryKey"` // 代表用户id, 只有注册成功之后才有这个id，唯一表示一个用户
    UserName  string    `gorm:"column:username"`
    Password  string    `gorm:"column:password"`
    Email     string    `gorm:"column:email"`
    CreatedAt time.Time `gorm:"column:created_at"`

    Token string `gorm:"-"` // token 可以用作注册token或者登录token
}

// MarshalBinary 实现BinaryMarshaler 接口
func (b *User) MarshalBinary() ([]byte, error) {
    return json.Marshal(b)
}

// UnmarshalBinary 实现 BinaryUnMarshaler 接口
func (b *User) UnmarshalBinary(bt []byte) error {
    return json.Unmarshal(bt, b)
}

