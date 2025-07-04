package main

import (
	"math/rand"
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// PercentBalancer implements a percent load balancing technique.
type percentBalancer struct {
	commonBalancer
	random           *rand.Rand
	migrationPercent int
}

// NewRandomBalancer returns a random proxy balancer.
func NewMigrationBalancer(targets []*middleware.ProxyTarget, migrationPercent int) middleware.ProxyBalancer {
	b := percentBalancer{}
	b.targets = targets
	b.migrationPercent = migrationPercent
	b.random = rand.New(rand.NewSource(int64(time.Now().Nanosecond())))
	return &b
}

// Next randomly returns an upstream target.
//
// Note: `nil` is returned in case upstream target list is empty.
func (b *percentBalancer) Next(c echo.Context) *middleware.ProxyTarget {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if len(b.targets) == 0 {
		return nil
	} else if len(b.targets) == 1 {
		return b.targets[0]
	}

	if b.random.Intn(100) < b.migrationPercent {
		return b.targets[1]
	}

	return b.targets[0]
}

type commonBalancer struct {
	targets []*middleware.ProxyTarget
	mutex   sync.Mutex
}

// AddTarget adds an upstream target to the list and returns `true`.
//
// However, if a target with the same name already exists then the operation is aborted returning `false`.
func (b *commonBalancer) AddTarget(target *middleware.ProxyTarget) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for _, t := range b.targets {
		if t.Name == target.Name {
			return false
		}
	}
	b.targets = append(b.targets, target)
	return true
}

// RemoveTarget removes an upstream target from the list by name.
//
// Returns `true` on success, `false` if no target with the name is found.
func (b *commonBalancer) RemoveTarget(name string) bool {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	for i, t := range b.targets {
		if t.Name == name {
			b.targets = append(b.targets[:i], b.targets[i+1:]...)
			return true
		}
	}
	return false
}
