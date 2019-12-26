// Copyright 2018 The UniChain Team Authors
// This file is part of the unichain project.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package prque

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSstack(t *testing.T) {
	// Create some initial data
	size := 16 * blockSize
	data := make([]*item, size)
	for i := 0; i < size; i++ {
		data[i] = &item{rand.Int(), rand.Int63()}
	}
	stack := newSstack(nil)
	for rep := 0; rep < 2; rep++ {
		// Push all the data into the stack, pop out every second
		secs := []*item{}
		for i := 0; i < size; i++ {
			stack.Push(data[i])
			if i%2 == 0 {
				secs = append(secs, stack.Pop().(*item))
			}
		}
		rest := []*item{}
		for stack.Len() > 0 {
			rest = append(rest, stack.Pop().(*item))
		}
		// Make sure the contents of the resulting slices are ok
		for i := 0; i < size; i++ {
			if i%2 == 0 && data[i] != secs[i/2] {
				t.Errorf("push/pop mismatch: have %v, want %v.", secs[i/2], data[i])
			}
			if i%2 == 1 && data[i] != rest[len(rest)-i/2-1] {
				t.Errorf("push/pop mismatch: have %v, want %v.", rest[len(rest)-i/2-1], data[i])
			}
		}
	}
}

func TestSstackSort(t *testing.T) {
	// Create some initial data
	size := 16 * blockSize
	data := make([]*item, size)
	for i := 0; i < size; i++ {
		data[i] = &item{rand.Int(), int64(i)}
	}
	// Push all the data into the stack
	stack := newSstack(nil)
	for _, val := range data {
		stack.Push(val)
	}
	// Sort and pop the stack contents (should reverse the order)
	sort.Sort(stack)
	for _, val := range data {
		out := stack.Pop()
		if out != val {
			t.Errorf("push/pop mismatch after sort: have %v, want %v.", out, val)
		}
	}
}

func TestSstackReset(t *testing.T) {
	// Push some stuff onto the stack
	size := 16 * blockSize
	stack := newSstack(nil)
	for i := 0; i < size; i++ {
		stack.Push(&item{i, int64(i)})
	}
	// Clear and verify
	stack.Reset()
	if stack.Len() != 0 {
		t.Errorf("stack not empty after reset: %v", stack)
	}
}
