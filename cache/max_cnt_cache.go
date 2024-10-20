package cache

type MaxCntCache struct {
	*BuildInMapCache

	cnt    int32
	maxCnt int32
}

func NewMaxCntCache(maxCnt int32) *MaxCntCache {
	return &MaxCntCache{
		BuildInMapCache: NewBuildInMapCache(),
		cnt:             0,
		maxCnt:          maxCnt,
	}
}
