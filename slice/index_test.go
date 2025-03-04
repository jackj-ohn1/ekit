// Copyright 2021 gotomicro
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want int
		name string
	}{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 0,
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{1, 3, 4, 2, 0},
			dst:  0,
			want: 4,
			name: "the ele locate the end",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, Index[int](test.src, test.dst))
	}
}

func TestIndexFunc(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want int
		name string
	}{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 0,
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{1, 3, 4, 2, 0},
			dst:  0,
			want: 4,
			name: "the ele locate the end",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, IndexFunc[int](test.src, test.dst, func(src, dst int) bool {
			return src == dst
		}))
	}
}

func TestLastIndex(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want int
		name string
	}{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 1,
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: 5,
			name: "normal test",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, LastIndex[int](test.src, test.dst))
	}
}

func TestLastIndexFunc(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want int
		name string
	}{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: 1,
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  1,
			want: -1,
			name: "the length of src is 0",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: -1,
			name: "dst not exist",
		},
		{
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: 5,
			name: "normal test",
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, LastIndexFunc[int](test.src, test.dst, func(src, dst int) bool {
			return src == dst
		}))
	}
}

func TestIndexAll(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want []int
		name string
	}{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: []int{0, 1},
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  1,
			want: []int{},
			name: "the length of src is 0",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: []int{},
			name: "dst not exist",
		},
		{
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: []int{0, 5},
			name: "normal test",
		},
	}
	for _, test := range tests {
		res := IndexAll[int](test.src, test.dst)
		assert.Equal(t, true, equal[int](res, test.want))
	}
}

func TestIndexAllFunc(t *testing.T) {
	tests := []struct {
		src  []int
		dst  int
		want []int
		name string
	}{
		{
			src:  []int{1, 1, 3, 5},
			dst:  1,
			want: []int{0, 1},
			name: "normal test",
		},
		{
			src:  []int{},
			dst:  1,
			want: []int{},
			name: "the length of src is 0",
		},
		{
			src:  []int{1, 4, 6},
			dst:  7,
			want: []int{},
			name: "dst not exist",
		},
		{
			src:  []int{0, 1, 3, 4, 2, 0},
			dst:  0,
			want: []int{0, 5},
			name: "normal test",
		},
	}
	for _, test := range tests {
		res := IndexAllFunc[int](test.src, test.dst, func(src, dst int) bool {
			return src == dst
		})
		assert.Equal(t, true, equal[int](res, test.want))
	}
}
