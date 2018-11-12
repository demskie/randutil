package randutil

import "math"

// https://gist.github.com/tommyettinger/e6d3e8816da79b45bfe582384c2fe14a

func UniformUint32(z uint32) uint32 {
	z = z + 0x6d2b79f5
	z = (z ^ (z >> 15)) * (z | 1)
	z ^= z + (z^(z>>7))*(z|61)
	z ^= z >> 14
	return z
}

func UniformInt32Range(seed int64, min, max int32) int32 {
	val := UniformUint32(uint32(seed))
	diff := math.MaxUint32 / uint32(max-min)
	return int32(val/diff) + min
}

func UniformFloat64(seed int64) float64 {
	return float64(UniformUint32(uint32(seed))) / math.MaxUint32
}

func NormalUint32(z uint32) uint32 {
	in := z + 0x6d2b79f5

	z = (in ^ (in >> 15)) * (in | 1)
	z ^= z + (z^(z>>7))*(z|61)
	z ^= z >> 14
	out := uint64(z)

	in++

	z = (in ^ (in >> 15)) * (in | 1)
	z ^= z + (z^(z>>7))*(z|61)
	z ^= z >> 14
	out += uint64(z)

	in -= 2

	z = (in ^ (in >> 15)) * (in | 1)
	z ^= z + (z^(z>>7))*(z|61)
	z ^= z >> 14
	out += uint64(z)

	return uint32(out / 3)
}

func NormalInt32Range(seed int64, min, max int32) int32 {
	val := NormalUint32(uint32(seed))
	diff := math.MaxUint32 / uint32(max-min)
	return int32(val/diff) + min
}

func NormalFloat64(seed int64) float64 {
	return float64(NormalUint32(uint32(seed))) / math.MaxUint32
}
