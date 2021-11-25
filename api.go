package main

type Volume struct {
	Capacity int    `json:"capacity"`// 容量大小
	Region   string `json:"region"`// 所在局点
	VolumeID string `json:"volumeId"`// volume id
}
