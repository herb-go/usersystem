package usersystem

import (
	"testing"

	"github.com/herb-go/herbsecurity/authority"
)

var testType = DataType("test")
var testType2 = DataType("test2")

func testSession(id string) *Session {
	p := authority.NewPayloads()
	p.Set(PayloadUID, []byte(id))
	return NewSession().WithType("test").WithPayloads(p)
}

func TestUserSystem(t *testing.T) {
	s := New()
	if GetUsersystem(s.context) != s {
		t.Fatal(s)
	}
	if GetUID(UIDContext(s.context, "test")) != "test" {
		t.Fatal(s)
	}
	session := testSession("test")
	if GetSession(SessionContext(s.context, session)) != session {
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
