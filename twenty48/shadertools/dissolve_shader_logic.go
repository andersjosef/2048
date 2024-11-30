// Copyright 2020 The Ebiten Authors
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

// Modifications made by [Your Name or Organization]:
// - Simplified the limit calculation by removing the triangle wave function.
// - Changed the source image for `level` from `imageSrc3` to `imageSrc1`.
// - Removed the unused `Cursor` variable to reduce global scope complexity.

//go:build ignore

//kage:unit pixels

package shadertools

var Time float

func Fragment(dstPos vec4, srcPos vec2, color vec4) vec4 {
	limit := Time
	level := imageSrc1UnsafeAt(srcPos).x

	// Add a white border
	if limit-0.1 < level && level < limit {
		alpha := imageSrc0UnsafeAt(srcPos).w
		return vec4(alpha)
	}

	return step(limit, level) * imageSrc0UnsafeAt(srcPos)
}
