package drone

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_createDrone(t *testing.T) {
	tests := []struct {
		name string
		want func() (Drone, error)
	}{
		{
			name: "1",
			want: func() (Drone, error) {
				return Drone{}, nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := createDrone
			drone, err := got()
			require.NoError(t, err)
			wantedDrone, _ := tt.want()

			require.NotEqual(t, wantedDrone.ID, drone.ID)
			require.NotEqual(t, wantedDrone.Lat, drone.Lat)
			require.NotEqual(t, wantedDrone.Lon, drone.Lon)
		})
	}
}

func TestDrone_ScanEnemy(t *testing.T) {
	type fields struct {
		ID  uint64
		Lat float64
		Lon float64
	}
	tests := []struct {
		name   string
		fields fields
		want   func() *Squad
	}{
		{
			name: "smoke test",
			fields: fields{
				ID:  1,
				Lat: 1,
				Lon: 1,
			},
			want: func() *Squad {
				return &Squad{[]Enemy{
					{
						ID: "",
					},
				}}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Drone{
				ID:  tt.fields.ID,
				Lat: tt.fields.Lat,
				Lon: tt.fields.Lon,
			}
			d.ScanEnemy()
			for _, s := range d.Scans.Enemies {
				want := tt.want().Enemies[0]
				require.NotEqual(t, s.ID, want.ID)
				fmt.Println(s.GetInfo())
			}
		})
	}
}
