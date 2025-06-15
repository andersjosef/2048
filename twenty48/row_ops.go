package twenty48

import "github.com/andersjosef/2048/twenty48/animations"

func processRow(rowIndex int, row [4]int) (outRow [4]int, deltas []animations.MoveDelta, scoreGain int) {
	compacted, deltas := compactRow(rowIndex, row, true)
	merged, scoreGain := mergeRow(compacted)
	outRow, _ = compactRow(rowIndex, merged, false)
	return
}
