package service

import (
    "AF-Excel/interface/user"
    "af-go-frame/core/container"
    "context"
)

type UserService struct {
    container container.Container
}



func (u *UserService) Login(ctx context.Context, user1 *user.User) (*user.User, error) {
    //ormService := u.container.MustMake(contract.ORMKey).(contract.ORMService)
    //db, err := ormService.GetDB()
    //if err != nil {
    //    return nil, err
    //}

    userDB := &user.User{}
    //if err := db.Where("username=?", user.UserName).First(userDB).Error; err != nil {
    //    if err == gorm.ErrRecordNotFound {
    //        return nil, nil
    //    }
    //    return nil, err
    //}
    //
    //if err := bcrypt.CompareHashAndPassword([]byte(userDB.Password), []byte(user.Password)); err != nil {
    //    return nil, err
    //}

    userDB.Password = ""
    // 缓存session
    //cacheService := u.container.MustMake(contract.CacheKey).(contract.CacheService)
    token := "kingnet"
    //key := fmt.Sprintf("session:%v", token)
    //if err := cacheService.SetObj(ctx, key, userDB, 24*time.Hour); err != nil {
    //    return nil, err
    //}

    userDB.Token = token
    return userDB, nil
}


func NewUserService(params ...interface{}) (interface{}, error) {
    container := params[0].(container.Container)
    //logger := container.MustMake(contract.LogKey).(contract.Log)
    //configer := container.MustMake(contract.ConfigKey).(contract.Config)
    return &UserService{container: container}, nil
}
