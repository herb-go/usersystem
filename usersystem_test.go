package usersystem

import "testing"

var testType = Datatype("test")
var testType2 = Datatype("test2")

func TestUserSystem(t *testing.T) {
	s := New()
	if GetUsersystem(s.Context) != s {
		t.Fatal(s)
	}
	if GetUID(UIDContext(s.Context, "test")) != "test" {
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
