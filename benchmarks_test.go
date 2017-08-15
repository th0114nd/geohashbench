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
)

func BenchmarkMmcloughlinEncode(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mmcloughlin.Encode(points[i][0], points[i][1])
	}
}

func BenchmarkMmcloughlinEncodeInt(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mmcloughlin.EncodeInt(points[i][0], points[i][1])
	}
}

func BenchmarkTomihiltunenEncode(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tomihiltunen.Encode(points[i][0], points[i][1])
	}
}

func BenchmarkGansiduiEncode(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gansidui.Encode(points[i][0], points[i][1], 12)
	}
}

func BenchmarkPierrreEncode(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pierrre.Encode(points[i][0], points[i][1], 12)
	}
}

func BenchmarkBroadyEncode(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		broady.Encode(points[i][0], points[i][1])
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

func BenchmarkCodeforEncode(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		codefor.Encode(points[i][0], points[i][1])
	}
}

func BenchmarkFanixkEncode(b *testing.B) {
	points := RandomPoints(b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		fanixk.Encode(points[i][0], points[i][1])
	}
}
