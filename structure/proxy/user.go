package proxy

import (
	"errors"
)

type User struct {
	ID int32
}

type UserFinder interface {
	FindUser(id int32) (User, error)
}

type UserList []User

func (l UserList) Find(id int32) *User {
	for _, v := range l {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

type UserListProxy struct {
	MockedDatabase      *UserList
	StackCache          UserList
	StackSize           int
	LastSearchUsedCache bool
}

func (u *UserListProxy) FindUser(id int32) (User, error) {

	foundCache := u.StackCache.Find(id)
	if foundCache != nil {
		u.LastSearchUsedCache = true
		return *foundCache, nil
	}

	u.LastSearchUsedCache = false
	found := u.MockedDatabase.Find(id)
	if found != nil {
		u.StackCache = append(u.StackCache, *found)
		if len(u.StackCache) > u.StackSize {
			u.StackCache = u.StackCache[1:]
		}
		return *found, nil
	}

	return User{}, errors.New("User Not found")
}
