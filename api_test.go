package mbapi

import "testing"

func Test_RenderQuery(t *testing.T) {

	const url = "/Users/12/Items?PersonTypes=BadPeople,GoodPeople&IsMissing=true&SortOrder=Ascending"

	q := &ItemQuery{}
	q.IsMissing = true
	q.UserId = "12"
	q.PersonTypes = []string{"BadPeople", "GoodPeople"}
	q.SortOrder = "Ascending"
	// q.Person = 89
	u, err := RenderQuery(q)
	if err != nil {
		t.Error(err)
	}
	if u != url {
		t.Errorf("RenderQuery did not return the Expected Result\n")
	}

	t.Log(u)
}
