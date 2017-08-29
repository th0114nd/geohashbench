# geohashbench
Benchmarks to compare golang geohash implementations.

## Results

### String Encoding

```
name                        time/op
MmcloughlinEncodeString-4   76.2ns ± 0%
PierrreEncodeString-4        649ns ± 0%
CodeforEncodeString-4        714ns ± 0%
FanixkEncodeString-4         775ns ± 0%
TomihiltunenEncodeString-4   811ns ± 0%
GansiduiEncodeString-4       832ns ± 0%
TidwallEncodeString-4        905ns ± 0%
BroadyEncodeString-4        1.40µs ± 0%
```

### Integer Encoding

```
name                    time/op
MmcloughlinEncodeInt-4  12.9ns ± 0%
BsmEncodeInt-4          16.3ns ± 0%
EzzkoramEncodeInt-4     43.4ns ± 0%
CorscEncodeInt52-4       511ns ± 0%
```

### String Decoding

```
name                        time/op
MmcloughlinDecodeString-4   190ns ± 0%
PierrreDecodeString-4       362ns ± 0%
CodeforDecodeString-4       453ns ± 0%
BroadyDecodeString-4        534ns ± 0%
TidwallDecodeString-4       553ns ± 0%
TomihiltunenDecodeString-4  610ns ± 0%
FanixkDecodeString-4        742ns ± 0%
```

### Meta

```
$ date
Mon Aug 28 21:53:30 PDT 2017
$ go version
go version go1.8.1 darwin/amd64
$ sysctl -n machdep.cpu.brand_string
Intel(R) Core(TM) i7-6660U CPU @ 2.40GHz
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
* [tidwall/tile38/geojson/geohash](https://github.com/tidwall/tile38/geojson/geohash) ([godoc](https://godoc.org/github.com/tidwall/tile38/geojson/geohash))