# geohashbench
Benchmarks to compare golang geohash implementations.

## Results

### String Encoding

```
                            │ results/EncodeString.out │
                            │          sec/op          │
MmcloughlinEncodeString-10                26.60n ± ∞ ¹
CodeforEncodeString-10                    402.7n ± ∞ ¹
PierrreEncodeString-10                    410.1n ± ∞ ¹
FanixkEncodeString-10                     412.7n ± ∞ ¹
TomihiltunenEncodeString-10               413.6n ± ∞ ¹
GansiduiEncodeString-10                   417.9n ± ∞ ¹
BroadyEncodeString-10                     770.5n ± ∞ ¹
geomean                                   304.3n
¹ need >= 6 samples for confidence interval at level 0.95

```

### Integer Encoding

```
                        │ results/EncodeInt.out │
                        │        sec/op         │
BsmEncodeInt-10                    4.247n ± ∞ ¹
MmcloughlinEncodeInt-10            4.936n ± ∞ ¹
EzzkoramEncodeInt-10               5.452n ± ∞ ¹
CorscEncodeInt52-10                303.1n ± ∞ ¹
geomean                            13.64n
¹ need >= 6 samples for confidence interval at level 0.95

```

### String Decoding

```
                            │ results/DecodeString.out │
                            │          sec/op          │
MmcloughlinDecodeString-10                40.25n ± ∞ ¹
PierrreDecodeString-10                    229.1n ± ∞ ¹
CodeforDecodeString-10                    358.5n ± ∞ ¹
BroadyDecodeString-10                     376.3n ± ∞ ¹
TomihiltunenDecodeString-10               441.6n ± ∞ ¹
FanixkDecodeString-10                     480.4n ± ∞ ¹
geomean                                   253.3n
¹ need >= 6 samples for confidence interval at level 0.95

```

### Meta

```
$ date
Fri Nov 22 17:52:54 EST 2024
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
