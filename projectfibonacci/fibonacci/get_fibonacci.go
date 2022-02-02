package fibonacci

func FibWrap(caheArg ...bool) (func(uint) uint64, error) {

	var mapFibNumCache = make(map[uint]uint64)
	var FibNum func(uint) uint64
	withCacheArg := len(caheArg) > 0 && caheArg[0]

	FibNum = func(n uint) (result uint64) {

		if withCacheArg {

			if result, ok := mapFibNumCache[n]; ok {
				return result
			}

			defer func() {
				mapFibNumCache[n] = result
			}()
		}

		if n <= 1 {
			return uint64(n)
		} else {
			return FibNum(n-1) + FibNum(n-2)
		}
	}
	return FibNum, nil
}
