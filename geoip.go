package als

import (
	"github.com/funkygao/geoip"
)

type GeoPoint struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

func LoadGeoDb(geodbfile string) (err error) {
	geo, err = geoip.Open(geodbfile)
	return
}

func GeoEnabled() bool {
	return geo != nil
}

func IpToGeoPoint(ip string) (gp GeoPoint) {
	if !GeoEnabled() {
		panic("must LoadGeoDb before IpToGeoPoint")
	}

	if rec := geo.GetRecord(ip); rec != nil {
		gp = GeoPoint{Lat: rec.Latitude, Lon: rec.Longitude}
	}

	return
}

// Return 2 letter country name, e,g. US
func IpToCountry(ip string) string {
	if !GeoEnabled() {
		panic("must LoadGeoDb before IpToCountry")
	}

	country, _ := geo.GetCountry(ip)
	return country
}
