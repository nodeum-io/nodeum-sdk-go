package legacy

import (
	"context"
	"fmt"
	"net/http"
)

type MountPointType int

const (
	SmbV1 MountPointType = iota
	NfsV3
	StorNextV5
	NfsV4
	SmbV21
	SmbV3
)

type MountPoint struct {
	Id           int32          `json:"id"`
	Name         string         `json:"name"`
	Type         MountPointType `json:"type"`
	Source       string         `json:"target"`
	Options      string         `json:"options"`
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	ScanInterval string         `json:"scan_interval"`
}

func (mp1 *MountPoint) Equal(mp2 *MountPoint) bool {
	return mp1.Id == mp2.Id &&
		mp1.Name == mp2.Name &&
		mp1.Type == mp2.Type &&
		mp1.Source == mp2.Source &&
		mp1.Options == mp2.Options &&
		mp1.Username == mp2.Username &&
		mp1.Password == mp2.Password &&
		mp1.ScanInterval == mp2.ScanInterval
}

type GetMountPointsResponse struct {
	StandardResponse
	MountPoints []MountPoint `json:"mount_points"`
}

func (api *ClientV1) GetMountPoints(ctx context.Context) (data *GetMountPointsResponse, err error) {
	url := "frontend/mount_points"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	err = api.Req(req.WithContext(ctx), &data)

	return
}

func (api *ClientV1) SyncMountPoint(ctx context.Context, name string, tx int) (data *StandardResponse, err error) {
	url := fmt.Sprintf(
		"frontend/mount_points/%s/sync?tx=%d",
		name,
		tx,
	)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return
	}

	err = api.Req(req.WithContext(ctx), &data)
	return
}
