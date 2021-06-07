// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

package generator

import "math/rand"

// Generator generates a sequence of values, following some distribution (Uniform, Zipfian, etc.).
type Generator interface {
	// Next generates the next value in the distribution.
	Next(r *rand.Rand) int64

	// Last returns the previous value generated by the distribution, e.g. the
	// last Next call.
	// Calling Last should not advance the distribution or have any side effect.
	// If Next has not been called, Last should return something reasonable.
	Last() int64
}