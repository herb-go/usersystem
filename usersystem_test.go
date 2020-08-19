package usersystem

import (
	"testing"

	"github.com/herb-go/herbsecurity/authority"
)

var testType = DataType("test")
var testType2 = DataType("test2")

type testSession string

func (s testSession) ID() string {
	return ""
}
func (s testSession) Type() SessionType {
	return ""
}
func (s testSession) UID() (string, error) {
	return string(s), nil
}
func (s testSession) SaveUID(string) error {
	return nil
}
func (s testSession) Payloads() (*authority.Payloads, error) {
	return nil, nil
}
func (s testSession) SavePayloads(*authority.Payloads) error {
	return nil
}
func (s testSession) Destory() (bool, error) {
	return false, nil
}
func (s testSession) Save(key string, v interface{}) error {
	return nil
}
func (s testSession) Load(key string, v interface{}) error {
	return nil
}
func (s testSession) Remove(key string) error {
	return nil
}
func (s testSession) IsNotFoundError(err error) bool {
	return false
}
func TestUserSystem(t *testing.T) {
	s := New()
	if GetUsersystem(s.Context) != s {
		t.Fatal(s)
	}
	if GetUID(UIDContext(s.Context, "test")) != "test" {
		t.Fatal(s)
	}
	session := testSession("test")
	if GetSession(SessionContext(s.Context, session)) != session {
		t.Fatal(s)
	}
	ds := NewPlainDataset()
	ds.InitType(testType)
	ds.InitType(testType2)
	v, ok := ds.Get(testType, "test")
	if v != nil || ok != false {
		t.Fatal(ds)
	}
	ds.Set(testType, "test", "testvalue")
	v, ok = ds.Get(testType, "test")
	if v == nil || v.(string) != "testvalue" || ok != true {
		t.Fatal(ds)
	}
	ds.Delete(testType, "test")
	v, ok = ds.Get(testType, "test")
	if v != nil || ok != false {
		t.Fatal(ds)
	}
	ds.Set(testType, "test", "testvalue")
	ds.Set(testType2, "test", "testvalue2")
	v, ok = ds.Get(testType, "test")
	if v == nil || v.(string) != "testvalue" || ok != true {
		t.Fatal(ds)
	}
	v, ok = ds.Get(testType2, "test")
	if v == nil || v.(string) != "testvalue2" || ok != true {
		t.Fatal(ds)
	}
	ds.Flush("test")
	v, ok = ds.Get(testType, "test")
	if v != nil || ok != false {
		t.Fatal(ds)
	}
	v, ok = ds.Get(testType2, "test")
	if v != nil || ok != false {
		t.Fatal(ds)
	}
}
