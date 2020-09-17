// Copyright 2020, Honeycomb, Hound Technology, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package honeycomb contains samplers for use with Honeycomb
package honeycomb

import (
	"crypto/sha1"
	"errors"
	"math"

	"go.opentelemetry.io/otel/sdk/trace"
)

var (
	ErrInvalidSampleRate = errors.New("sample rate must be >= 1")
)

type DeterministicSampler struct {
	SampleRate int
	upperBound uint32
}

func NewDeterministicSampler(sampleRate uint) (*DeterministicSampler, error) {
	if sampleRate < 1 {
		return nil, ErrInvalidSampleRate
	}

	// Get the actual upper bound - the largest possible value divided by
	// the sample rate. In the case where the sample rate is 1, this should
	// sample every value.
	upperBound := math.MaxUint32 / uint32(sampleRate)
	return &DeterministicSampler{
		SampleRate: int(sampleRate),
		upperBound: upperBound,
	}, nil
}

// bytesToUint32 takes a slice of 4 bytes representing a big endian 32 bit
// unsigned value and returns the equivalent uint32.
func bytesToUint32be(b []byte) uint32 {
	return uint32(b[3]) | (uint32(b[2]) << 8) | (uint32(b[1]) << 16) | (uint32(b[0]) << 24)
}

func (ds *DeterministicSampler) ShouldSample(p trace.SamplingParameters) trace.SamplingResult {
	if ds.SampleRate == 1 {
		return trace.SamplingResult{
			Decision: trace.RecordAndSampled,
		}
	}
	determinant := []byte(p.TraceID[:])
	sum := sha1.Sum([]byte(determinant))
	v := bytesToUint32be(sum[:4])

	var decision trace.SamplingDecision
	if v <= ds.upperBound {
		decision = trace.RecordAndSampled
	} else {
		decision = trace.NotRecord
	}

	return trace.SamplingResult{
		Decision: decision,
	}
}

func (ds *DeterministicSampler) Description() string {
	return "A determistic head sampler for use with Honeycomb"
}
