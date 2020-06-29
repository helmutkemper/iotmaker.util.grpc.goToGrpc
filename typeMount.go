package iotmaker_util_grpc_goToGrpc

type Mount struct {
	Type          Type
	Source        string
	Target        string
	ReadOnly      bool
	Consistency   Consistency
	BindOptions   BindOptions
	VolumeOptions VolumeOptions
	TmpfsOptions  TmpfsOptions
}
