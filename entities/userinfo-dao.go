package entities


// Save .
func Save(u *UserInfo) error {
    d ,err:= GetDB().Last(u)
    if (err != nil) {
        return err
    }
    u.UID = d.UID+1;
    GetDB().Create(u)
}

// FindAll .
func FindAll() []UserInfo {
    ulist := make([]UserInfo, 0, 0)
    GetDB().Find(&ulist)
    return ulist
}

// FindByID .
func FindByID(id int) *UserInfo {
    u = UserInfo{}
    GetDB().Where(&UserInfo{UID: id}).First(&u)
    return &u
}