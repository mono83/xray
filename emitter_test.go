package xray

import "testing"

func BenchmarkSyncEmitter_NoListeners1000(b *testing.B) {
	ee := NewSyncEmitter()
	evt := &logEvent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			ee.Emit(evt)
		}
	}
}

func BenchmarkSyncEmitter_OneListeners1000(b *testing.B) {
	count := 0
	listener := func(events ...Event) {
		count += len(events)
	}

	ee := NewSyncEmitter()
	ee.On(listener)
	evt := &logEvent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			ee.Emit(evt)
		}
	}

	if b.N*1000 != count {
		b.Errorf("Expected %d but counted %d", b.N, count)
	}
}

func BenchmarkSyncEmitter_TenListeners1000(b *testing.B) {
	count := 0
	listener := func(events ...Event) {
		count += len(events)
	}

	ee := NewSyncEmitter()
	ee.On(listener)
	ee.On(listener)
	ee.On(listener)
	ee.On(listener)
	ee.On(listener)
	ee.On(listener)
	ee.On(listener)
	ee.On(listener)
	ee.On(listener)
	ee.On(listener)
	evt := &logEvent{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for j := 0; j < 1000; j++ {
			ee.Emit(evt)
		}
	}

	if b.N*1000*10 != count {
		b.Errorf("Expected %d but counted %d", b.N, count)
	}
}
