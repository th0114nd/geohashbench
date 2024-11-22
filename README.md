# geohashbench
Benchmarks to compare golang geohash implementations.

## Results

### String Encoding

```
                           │ results/EncodeString.out │
                           │          sec/op          │
MmcloughlinEncodeString-10               25.95n ± ∞ ¹
PierrreEncodeString-10                   397.4n ± ∞ ¹
FanixkEncodeString-10                    405.1n ± ∞ ¹
GansiduiEncodeString-10                  417.9n ± ∞ ¹
BroadyEncodeString-10                    750.2n ± ∞ ¹
geomean                                  265.1n
¹ need >= 6 samples for confidence interval at level 0.95

```

### Integer Encoding

```
                        │ results/EncodeInt.out │
                        │        sec/op         │
BsmEncodeInt-10                    4.262n ± ∞ ¹
MmcloughlinEncodeInt-10            4.942n ± ∞ ¹
EzzkoramEncodeInt-10               5.250n ± ∞ ¹
CorscEncodeInt52-10                296.2n ± ∞ ¹
geomean                            13.45n
¹ need >= 6 samples for confidence interval at level 0.95

```

### String Decoding

```
                           │ results/DecodeString.out │
                           │          sec/op          │
MmcloughlinDecodeString-10               39.45n ± ∞ ¹
PierrreDecodeString-10                   231.6n ± ∞ ¹
BroadyDecodeString-10                    367.7n ± ∞ ¹
FanixkDecodeString-10                    475.7n ± ∞ ¹
geomean                                  199.9n
¹ need >= 6 samples for confidence interval at level 0.95

```

### Meta

```
$ date
Fri Nov 22 17:35:31 EST 2024
$ go version
go version go1.23.3 darwin/arm64
$ sysctl -n machdep.cpu.brand_string
Apple M1 Max

```

## Packages Tested

* [mmcloughlin/geohash](https://github.com/mmcloughlin/geohash) ([godoc](https://godoc.org/github.com/mmcloughlin/geohash))
* [gansidui/geohash](https://github.com/gansidui/geohash) ([godoc](https://godoc.org/github.com/gansidui/geohash))
* [pierrre/geohash](https://github.com/pierrre/geohash) ([godoc](https://godoc.org/github.com/pierrre/geohash))
* [broady/gogeohash](https://github.com/broady/gogeohash) ([godoc](https://godoc.org/github.com/broady/gogeohash))
* [bsm/geohashi](https://github.com/bsm/geohashi) ([godoc](https://godoc.org/github.com/bsm/geohashi))
* [corsc/go-geohash](https://github.com/corsc/go-geohash) ([godoc](https://godoc.org/github.com/corsc/go-geohash))
* [ezzkoram/geohash](https://github.com/ezzkoram/geohash) ([godoc](https://godoc.org/github.com/ezzkoram/geohash))
* [fanixk/geohash](https://github.com/fanixk/geohash) ([godoc](https://godoc.org/github.com/fanixk/geohash))
