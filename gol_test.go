package gol

import (
	"github.com/andresperezl/gol/region"
	"testing"
)

func TestNewClient(t *testing.T) {
	key := "testkey"
	reg := region.NA1
	c := NewClient("testkey", region.NA1)
	if got := c.CurrentAPIKey; got != key {
		t.Errorf("Expected client to have key %s but got %s", key, got)
	}
	if got := c.CurrentRegion; got != reg {
		t.Errorf("Expected client to have region %s but got %s", reg, got)
	}
	newkey := "newkey"
	newreg := region.BR1
	c.SetAPIKey(newkey)
	c.SetRegion(newreg)
	if got := c.CurrentAPIKey; got != newkey {
		t.Errorf("Expected client to have key %s but got %s", newkey, got)
	}
	if got := c.CurrentRegion; got != newreg {
		t.Errorf("Expected client to have region %s but got %s", newreg, got)
	}
}
