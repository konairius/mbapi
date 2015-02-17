package mbapi

import "testing"

var b *Backend

const hostname = "mediaserver-nerdcave/mediabrowser"

func init() {
	b = NewBackend(hostname)
	// b = NewBackend("mediaserver-sulzbach:8096/mediabrowser")
}

func Test_GetUsers(t *testing.T) {
	users, err := b.GetUsers()
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Logf("Users: %#v", len(users))
	for _, u := range users {
		t.Logf("User %v: Id=%v, LastLoginDate=%v", u.Name, u.Id, u.LastLoginDate)
	}
}

func Test_Login(t *testing.T) {
	if err := b.Login("konsti", "kodi"); err != nil {
		t.Errorf("%v\n", err)
	}
}

func Test_Movies(t *testing.T) {
	movies, err := b.Movies()
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Logf("Movies: %v\n", len(movies))
	// for _, m := range movies {
	// 	t.Logf("Movie: %v \n", m.Name)
	// }
}

func Test_Shows(t *testing.T) {
	shows, err := b.TVShows()
	if err != nil {
		t.Errorf("%v\n", err)
	}
	t.Logf("Shows: %v\n", len(shows))
	// for _, m := range shows {
	// 	t.Logf("Show: %v \n", m.Id)
	// }
}

func Test_ShowById(t *testing.T) {
	s, err := b.ShowById("063c13bdf76c49c5d5db750d31a6d8a9")
	if err != nil {
		t.Errorf("%v\n", err)
	}

	t.Logf("Got TV-Show: %s\n", s.Name)
}

// func Test_Stream(t *testing.T) {
// 	movies, err := b.Movies()
// 	if err != nil {
// 		t.Errorf("%v\n", err)
// 	}
// 	s, err := b.Stream(movies[0].Id)
// 	if err != nil {
// 		t.Errorf("%v\n", err)
// 	}
// 	p := make([]byte, 128)
// 	// buf := bytes.NewBuffer(p)
// 	n, err := s.ReadAt(0, p)
// 	n, err = s.ReadAt(128, p)
// 	if err != nil {
// 		t.Errorf("%v\n", err)
// 	}
// 	t.Logf("Read %v byte", n)
// }
