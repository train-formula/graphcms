package models

import "math"

func PtrIntToInt32(p *int) *int32 {

	t := *p

	var r int32

	if t > math.MaxInt32 {
		r = int32(math.MaxInt32)
	} else if t < math.MinInt32 {
		r = int32(math.MinInt32)
	} else {
		r = int32(t)
	}

	return &r
}
