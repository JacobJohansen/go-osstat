// +build darwin

package cpu

import (
	"testing"
)

func TestGetCpu(t *testing.T) {
	cpu, err := Get()
	if err != nil {
		t.Fatalf("error should be nil but got: %v", err)
	}
	if cpu.User <= 0 || cpu.System <= 0 || cpu.Idle <= 0 || cpu.Total <= 0 {
		t.Errorf("invalid cpu value: %+v", cpu)
	}
}
