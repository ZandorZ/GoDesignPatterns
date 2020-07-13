package proxy

import (
	"math/rand"
	"testing"
)

func Test_UserListProxy(t *testing.T) {
	mockedDatabase := UserList{}

	rand.Seed(2312311)
	for i := 0; i < 100000; i++ {
		n := rand.Int31()
		mockedDatabase = append(mockedDatabase, User{n})
	}

	proxy := UserListProxy{
		MockedDatabase: &mockedDatabase,
		StackSize:      2,
		StackCache:     UserList{},
	}

	knownIDs := [3]int32{mockedDatabase[4].ID, mockedDatabase[14].ID, mockedDatabase[8].ID}

	t.Run("FindUser - Empty cache", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one sucessful search in an empty cache, the size of it must be one")
		}

		if proxy.LastSearchUsedCache == true {
			t.Error("No user can be returned from an empty cache")
		}
	})

	t.Run("FindUser - Ask for same user", func(t *testing.T) {
		user, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		if user.ID != knownIDs[0] {
			t.Error("Returned name doesn't match with expected")
		}

		if len(proxy.StackCache) != 1 {
			t.Error("After one sucessful search in an empty cache, the size of it must be one")
		}

		if !proxy.LastSearchUsedCache {
			t.Error("The user should have been returned from the cache")
		}
	})

	t.Run("FindUser - overflowing the stack", func(t *testing.T) {
		user1, err := proxy.FindUser(knownIDs[0])
		if err != nil {
			t.Fatal(err)
		}

		proxy.FindUser(knownIDs[1])
		if proxy.LastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		proxy.FindUser(knownIDs[2])
		if proxy.LastSearchUsedCache {
			t.Error("The user wasn't stored on the proxy cache yet")
		}

		if found1 := proxy.StackCache.Find(user1.ID); found1 != nil {
			t.Error("User that should be gone was found")
		}

		if len(proxy.StackCache) != 2 {
			t.Error("After inserting 3 users cache should no grow")
		}

	})

}
