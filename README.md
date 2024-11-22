# geohashbench
Benchmarks to compare golang geohash implementations.

## Results

### String Encoding

```
                            │ results/EncodeString.out │
                            │          sec/op          │
MmcloughlinEncodeString-10                25.88n ± ∞ ¹
GansiduiEncodeString-10                   400.1n ± ∞ ¹
TomihiltunenEncodeString-10               403.3n ± ∞ ¹
PierrreEncodeString-10                    415.9n ± ∞ ¹
CodeforEncodeString-10                    424.3n ± ∞ ¹
FanixkEncodeString-10                     432.9n ± ∞ ¹
BroadyEncodeString-10                     750.5n ± ∞ ¹
geomean                                   303.9n
¹ need >= 6 samples for confidence interval at level 0.95

```

### Integer Encoding

```
                        │ results/EncodeInt.out │
                        │        sec/op         │
BsmEncodeInt-10                    4.651n ± ∞ ¹
MmcloughlinEncodeInt-10            4.940n ± ∞ ¹
EzzkoramEncodeInt-10               5.208n ± ∞ ¹
CorscEncodeInt52-10                288.3n ± ∞ ¹
geomean                            13.63n
¹ need >= 6 samples for confidence interval at level 0.95

```

### String Decoding

```
                            │ results/DecodeString.out │
                            │          sec/op          │
MmcloughlinDecodeString-10                39.76n ± ∞ ¹
PierrreDecodeString-10                    231.5n ± ∞ ¹
CodeforDecodeString-10                    358.0n ± ∞ ¹
BroadyDecodeString-10                     375.8n ± ∞ ¹
TomihiltunenDecodeString-10               437.2n ± ∞ ¹
FanixkDecodeString-10                     483.4n ± ∞ ¹
geomean                                   252.9n
¹ need >= 6 samples for confidence interval at level 0.95

```

### Meta

```
$ date
Fri Nov 22 17:57:04 EST 2024
$ go version
go version go1.23.3 darwin/arm64
$ sysctl -n machdep.cpu.brand_string
Apple M1 Max

```

## Packages Tested

* [mmcloughlin/geohash](https://github.com/mmcloughlin/geohash) ([godoc](https://godoc.org/github.com/mmcloughlin/geohash))
* [TomiHiltunen/geohash-golang](https://github.com/TomiHiltunen/geohash-golang) ([godoc](https://godoc.org/github.com/TomiHiltunen/geohash-golang))
* [gansidui/geohash](https://github.com/gansidui/geohash) ([godoc](https://godoc.org/github.com/gansidui/geohash))
* [pierrre/geohash](https://github.com/pierrre/geohash) ([godoc](https://godoc.org/github.com/pierrre/geohash))
* [broady/gogeohash](https://github.com/broady/gogeohash) ([godoc](https://godoc.org/github.com/broady/gogeohash))
* [bsm/geohashi](https://github.com/bsm/geohashi) ([godoc](https://godoc.org/github.com/bsm/geohashi))
* [corsc/go-geohash](https://github.com/corsc/go-geohash) ([godoc](https://godoc.org/github.com/corsc/go-geohash))
* [ezzkoram/geohash](https://github.com/ezzkoram/geohash) ([godoc](https://godoc.org/github.com/ezzkoram/geohash))
* [Codefor/geohash](https://github.com/Codefor/geohash) ([godoc](https://godoc.org/github.com/Codefor/geohash))
* [fanixk/geohash](https://github.com/fanixk/geohash) ([godoc](https://godoc.org/github.com/fanixk/geohash))
