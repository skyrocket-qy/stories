package main

import (
	"context"
	"testing"

	"go.uber.org/fx"
)

func BenchmarkWire(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := InitializeService()
		s.Do()
	}
}

func BenchmarkDirect(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s := NewService()
		s.Do()
	}
}

func BenchmarkFx(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		app := fx.New(
			fx.Provide(NewService),
			fx.Invoke(func(s Service) {
				s.Do()
			}),
			fx.NopLogger,
		)
		_ = app.Start(context.Background())
		_ = app.Stop(context.Background())
	}
}

func BenchmarkFxNewOnly(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fx.New(
			fx.Provide(NewService),
			fx.Invoke(func(s Service) {
				s.Do()
			}),
			fx.NopLogger,
		)
	}
}
