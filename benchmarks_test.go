package geohashbench

import (
	"testing"

	codefor "github.com/Codefor/geohash"
	tomihiltunen "github.com/TomiHiltunen/geohash-golang"
	broady "github.com/broady/gogeohash"
	bsm "github.com/bsm/geohashi"
	corsc "github.com/corsc/go-geohash"
	ezzkoram "github.com/ezzkoram/geohash"
	fanixk "github.com/fanixk/geohash"
	gansidui "github.com/gansidui/geohash"
	mmcloughlin "github.com/mmcloughlin/geohash"
	pierrre "github.com/pierrre/geohash"
	tidwall "github.com/tidwall/tile38/geojson/geohash"
)

func BenchmarkMmcloughlinEncodeInt(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mmcloughlin.EncodeInt(points[i][0], points[i][1])
	}
}

func BenchmarkMmcloughlinEncodeString(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mmcloughlin.Encode(points[i][0], points[i][1])
	}
}

func BenchmarkMmcloughlinDecodeString(b *testing.B) {
	geohashes := RandomStringGeohashes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mmcloughlin.Decode(geohashes[i])
	}
}

func BenchmarkTomihiltunenEncodeString(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tomihiltunen.Encode(points[i][0], points[i][1])
	}
}

func BenchmarkTomihiltunenDecodeString(b *testing.B) {
	geohashes := RandomStringGeohashes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tomihiltunen.Decode(geohashes[i])
	}
}

func BenchmarkGansiduiEncodeString(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gansidui.Encode(points[i][0], points[i][1], 12)
	}
}

func BenchmarkPierrreEncodeString(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pierrre.Encode(points[i][0], points[i][1], 12)
	}
}

func BenchmarkPierrreDecodeString(b *testing.B) {
	geohashes := RandomStringGeohashes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pierrre.Decode(geohashes[i])
	}
}

func BenchmarkBroadyEncodeString(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		broady.Encode(points[i][0], points[i][1])
	}
}

func BenchmarkBroadyDecodeString(b *testing.B) {
	geohashes := RandomStringGeohashes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		broady.Decode(geohashes[i])
	}
}

func BenchmarkBsmEncodeInt(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bsm.Encode(points[i][0], points[i][1])
	}
}

func BenchmarkCorscEncodeInt52(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		corsc.EncodeInt(points[i][0], points[i][1], 52)
	}
}

func BenchmarkEzzkoramEncodeInt(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ezzkoram.FromCoordinates(points[i][0], points[i][1]).Hash()
	}
}

func BenchmarkCodeforEncodeString(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codefor.Encode(points[i][0], points[i][1])
	}
}

func BenchmarkCodeforDecodeString(b *testing.B) {
	geohashes := RandomStringGeohashes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codefor.Decode(geohashes[i])
	}
}

func BenchmarkFanixkEncodeString(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fanixk.Encode(points[i][0], points[i][1])
	}
}

func BenchmarkFanixkDecodeString(b *testing.B) {
	geohashes := RandomStringGeohashes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fanixk.Decode(geohashes[i])
	}
}

func BenchmarkTidwallEncodeString(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tidwall.Encode(points[i][0], points[i][1], 12)
	}
}

func BenchmarkTidwallDecodeString(b *testing.B) {
	geohashes := RandomStringGeohashes(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tidwall.Decode(geohashes[i])
	}
}
