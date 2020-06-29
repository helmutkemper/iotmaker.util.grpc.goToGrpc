package iotmaker_util_grpc_goToGrpc

type Location struct {
	Name       string
	Zone       []Zone
	Tx         []ZoneTrans
	CacheStart int64
	CacheEnd   int64
	CacheZone  Zone
}
