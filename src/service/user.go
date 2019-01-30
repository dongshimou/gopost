package service

import (
	"model"
	"orm"
	"utility"
)

func SignIn(req *model.REQSignin) (*model.RESSignIn, error) {
	var err error
	user := model.User{}
	//验证 req.Username
	if err = utility.VerifyUsername(req.Username); err != nil {
		return nil, err
	}
	//加密 req.Password
	pass := utility.EncryptPassword(req.Password)

	user.Name = req.Username
	db := orm.Get()

	if err = db.Model(&user).Where(&user).First(&user).Error; err != nil {
		return nil, err
	}
	if pass != user.Password {
		return nil, utility.NewError(utility.ERROR_AUTH_CODE, utility.ERROR_MSG_ERROR_PASSWORD)
	}
	user.SignInIP = req.IP
	if err = db.Model(&user).Where(model.User{Name: user.Name}).Update(user).Error; err != nil {
		return nil, err
	}
	res := model.RESSignIn{}
	res.Username = user.Name
	res.Permission = user.Permission
	res.UpdateDatetime = formatDatetime(user.UpdatedAt)
	res.CreateDatetime = formatDatetime(user.CreatedAt)
	if res.Token, err = utility.GenerateToken(user.Name); err != nil {
		return nil, err
	}
	return &res, nil
}
func SignUp(req *model.REQSignUp) (*model.RESSignUp, error) {
	var err error
	user := model.User{}
	//验证 req.Username
	if err = utility.VerifyUsername(req.Username); err != nil {
		return nil, err
	}
	//加密 req.Password
	pass := utility.EncryptPassword(req.Password)

	user.Name = req.Username
	user.Password = pass
	user.Email = req.Email
	user.SignUpIP = req.IP

	user.Permission = utility.CreatePermission(
		model.Article_Read,
		model.Replay_Read,
		model.Replay_Create,
		model.Replay_Update,
	)
	db := orm.Get()

	count := 0
	if err = db.Model(&user).Where(&user).Count(&count).Error; err != nil {
		return nil, err
	}
	if count != 0 {
		return nil, utility.NewError(utility.ERROR_REQUEST_CODE, utility.ERROR_MSG_EXIST_USERNAME)
	}

	tx := db.Begin()
	if err = tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return nil, nil
}
