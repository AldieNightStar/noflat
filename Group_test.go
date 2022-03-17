package noflat

import "testing"

// ====================

type TestGS struct {
	mp map[string]string
}

func (t *TestGS) Get(k string) string {
	return t.mp[k]
}

func (t *TestGS) Set(k string, v string) {
	t.mp[k] = v
}

// ====================

func TestGroup(t *testing.T) {
	getterSetter := &TestGS{make(map[string]string, 8)}
	g := Init(getterSetter).Group("super")
	g.Set("ddd", "aaa")
	g.Set("sss", "___")

	if g.Get("ddd") != "aaa" || g.Get("sss") != "___" {
		t.Fatal("Keys should be as it assigned")
	}
	if g.Get("#keys") != "ddd;sss" {
		t.Fatal("#keys are wrong!")
	}
	if getterSetter.Get("super_ddd") != "aaa" || getterSetter.Get("super_sss") != "___" {
		t.Fatal("Prefixes works wrong!")
	}
}

func TestSubGroup(t *testing.T) {
	getterSetter := &TestGS{make(map[string]string, 8)}
	g := Init(getterSetter).Group("super").SubGroup("duper")

	g.Set("a", "b")

	if g.Get("a") != "b" {
		t.Fatal("Setters/Getters wrong")
	}
	if getterSetter.Get("super_duper_a") != "b" {
		t.Fatal("Prefixes is wrong!")
	}
}

func TestKeys(t *testing.T) {
	getterSetter := &TestGS{make(map[string]string, 8)}
	superG := Init(getterSetter).Group("super")
	duperG := superG.SubGroup("duper")

	superG.Set("a", "c")
	duperG.Set("x", "y")

	if len(duperG.Keys()) != 1 || duperG.Keys()[0] != "x" {
		t.Fatal("Keys in super group is bad")
	}
}
