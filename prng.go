package randutil

import "math"

// https://gist.github.com/tommyettinger/e6d3e8816da79b45bfe582384c2fe14a

func UniformUint32(seed uint32) uint32 {
	z := seed + 0x6d2b79f5
	z = (z ^ (z >> 15)) * (z | 1)
	z ^= z + (z^(z>>7))*(z|61)
	z ^= z >> 14
	return z
}

func UniformInt32Range(seed int64, min, max int32) int32 {
	val := UniformUint32(uint32(seed))
	diff := int64(math.MaxUint32) / int64(max-min)
	return int32(int64(val)/diff) + int32(min)
}

func NormalInt32Range(seed int64, min, max int32) int32 {
	val := UniformUint32(uint32(seed))
	val += UniformUint32(uint32(seed + 1))
	val += UniformUint32(uint32(seed - 1))
	diff := int64(math.MaxUint32) / int64(max-min)
	return int32(int64(val/3)/diff) + int32(min)
}
