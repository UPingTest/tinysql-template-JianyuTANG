// Copyright 2019 PingCAP, Inc.
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

// Code generated by go generate in expression/generator; DO NOT EDIT.

package expression

import (
	"github.com/pingcap/tidb/types"
	"github.com/pingcap/tidb/util/chunk"
)

func (b *builtinIfNullIntSig) vecEvalInt(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	if err := b.args[0].VecEvalInt(b.ctx, input, result); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalInt(b.ctx, input, buf1); err != nil {
		return err
	}

	arg0 := result.Int64s()
	arg1 := buf1.Int64s()
	for i := 0; i < n; i++ {
		if result.IsNull(i) && !buf1.IsNull(i) {
			result.SetNull(i, false)
			arg0[i] = arg1[i]
		}
	}
	return nil
}

func (b *builtinIfNullIntSig) vectorized() bool {
	return true
}

func (b *builtinIfNullRealSig) vecEvalReal(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	if err := b.args[0].VecEvalReal(b.ctx, input, result); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETReal, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalReal(b.ctx, input, buf1); err != nil {
		return err
	}

	arg0 := result.Float64s()
	arg1 := buf1.Float64s()
	for i := 0; i < n; i++ {
		if result.IsNull(i) && !buf1.IsNull(i) {
			result.SetNull(i, false)
			arg0[i] = arg1[i]
		}
	}
	return nil
}

func (b *builtinIfNullRealSig) vectorized() bool {
	return true
}

func (b *builtinIfNullStringSig) vecEvalString(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETString, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalString(b.ctx, input, buf0); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETString, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(b.ctx, input, buf1); err != nil {
		return err
	}

	result.ReserveString(n)
	for i := 0; i < n; i++ {
		if !buf0.IsNull(i) {
			result.AppendString(buf0.GetString(i))
		} else if !buf1.IsNull(i) {
			result.AppendString(buf1.GetString(i))
		} else {
			result.AppendNull()
		}
	}
	return nil
}

func (b *builtinIfNullStringSig) vectorized() bool {
	return true
}

func (b *builtinIfIntSig) vecEvalInt(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}
	if err := b.args[1].VecEvalInt(b.ctx, input, result); err != nil {
		return err
	}
	buf2, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalInt(b.ctx, input, buf2); err != nil {
		return err
	}

	arg0 := buf0.Int64s()
	arg2 := buf2.Int64s()
	rs := result.Int64s()
	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:
			if buf2.IsNull(i) {
				result.SetNull(i, true)
			} else {
				result.SetNull(i, false)
				rs[i] = arg2[i]
			}
		case arg != 0:
		}
	}
	return nil
}

func (b *builtinIfIntSig) vectorized() bool {
	return true
}

func (b *builtinIfRealSig) vecEvalReal(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}
	if err := b.args[1].VecEvalReal(b.ctx, input, result); err != nil {
		return err
	}
	buf2, err := b.bufAllocator.get(types.ETReal, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalReal(b.ctx, input, buf2); err != nil {
		return err
	}

	arg0 := buf0.Int64s()
	arg2 := buf2.Float64s()
	rs := result.Float64s()
	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:
			if buf2.IsNull(i) {
				result.SetNull(i, true)
			} else {
				result.SetNull(i, false)
				rs[i] = arg2[i]
			}
		case arg != 0:
		}
	}
	return nil
}

func (b *builtinIfRealSig) vectorized() bool {
	return true
}

func (b *builtinIfStringSig) vecEvalString(input *chunk.Chunk, result *chunk.Column) error {
	n := input.NumRows()
	buf0, err := b.bufAllocator.get(types.ETInt, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf0)
	if err := b.args[0].VecEvalInt(b.ctx, input, buf0); err != nil {
		return err
	}
	buf1, err := b.bufAllocator.get(types.ETString, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf1)
	if err := b.args[1].VecEvalString(b.ctx, input, buf1); err != nil {
		return err
	}
	buf2, err := b.bufAllocator.get(types.ETString, n)
	if err != nil {
		return err
	}
	defer b.bufAllocator.put(buf2)
	if err := b.args[2].VecEvalString(b.ctx, input, buf2); err != nil {
		return err
	}

	result.ReserveString(n)
	arg0 := buf0.Int64s()
	for i := 0; i < n; i++ {
		arg := arg0[i]
		isNull0 := buf0.IsNull(i)
		switch {
		case isNull0 || arg == 0:
			if buf2.IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendString(buf2.GetString(i))
			}
		case arg != 0:
			if buf1.IsNull(i) {
				result.AppendNull()
			} else {
				result.AppendString(buf1.GetString(i))
			}
		}
	}
	return nil
}

func (b *builtinIfStringSig) vectorized() bool {
	return true
}
