package cache

type errors struct {
	AnErrorOnExist string
	AnErrorOnGet   string
	AnErrorOnSet   string
	NotRunnable    string
}

var errorMessages = errors{
	AnErrorOnExist: "cache_an_error_on_exist",
	AnErrorOnGet:   "cache_an_error_on_get",
	AnErrorOnSet:   "cache_an_error_on_set",
	NotRunnable:    "cache_not_runnable",
}
