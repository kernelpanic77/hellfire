package internal

import "testing"

// manages the entire lifecycle of the test

type Manager struct {
	// 
	testState   *TestState   // test state
}

func (m *Manager) Init(main *testing.M) {
	// initialize test state
}

func (m *Manager) Start() {
	// start goroutines to manager summary
	// start goroutines for managing thresholds
	// start goroutine for EndPoints
	// 
}

// func (m *)