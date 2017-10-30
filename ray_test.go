package xray

import (
	"github.com/mono83/xray/args"
	"testing"
)

// BenchmarkRay_Common_Pattern creates 5 level-depth rays and emits 10 events
func BenchmarkRay_Common_Pattern(b *testing.B) {
	root := ROOT.Fork()

	// Forking
	level1 := root.Fork().With(args.AppName("foo"))
	level2 := level1.Fork().With(args.Type("bar"))
	level3 := level2.Fork().With(args.String{N: "zzz", V: "xxx"})
	level4 := level3.Fork().With(args.String{N: "zzz2", V: "xxx2"})
	level5 := level4.Fork().With(args.Type("baz"))

	// Attaching loggers
	cntTrace := 0
	cntDebug := 0
	cntInfo := 0
	cntError := 0
	root.On(
		func(events ...Event) {
			for _, event := range events {
				l, ok := event.(LogEvent)
				if ok && l.GetLevel() == TRACE {
					cntTrace++
				}
			}
		},
	)
	root.On(
		func(events ...Event) {
			for _, event := range events {
				l, ok := event.(LogEvent)
				if ok && l.GetLevel() == DEBUG {
					cntDebug++
				}
			}
		},
	)
	root.On(
		func(events ...Event) {
			for _, event := range events {
				l, ok := event.(LogEvent)
				if ok && l.GetLevel() == INFO {
					cntInfo++
				}
			}
		},
	)
	root.On(
		func(events ...Event) {
			for _, event := range events {
				l, ok := event.(LogEvent)
				if ok && l.GetLevel() == ERROR {
					cntError++
				}
			}
		},
	)

	// Sending
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		level5.Trace("This is trace :rayId")
		level5.Trace("This is trace :rayId")
		level5.Debug("This is debug :rayId")
		level5.Debug("This is debug :rayId")
		level5.Debug("This is debug :rayId")
		level5.Debug("This is debug :rayId")
		level5.Info("This is info :rayId")
		level5.Info("This is info :rayId")
		level5.Info("This is info :rayId")
		level5.Error("This is error :rayId")
	}

	if cntTrace != b.N*2 {
		b.Fail()
	}
	if cntDebug != b.N*4 {
		b.Fail()
	}
	if cntInfo != b.N*3 {
		b.Fail()
	}
	if cntError != b.N {
		b.Fail()
	}
}
