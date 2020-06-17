### Install

```
$ go get github.com/ikasamt/goslicer
```


### Dependencies

```
$ go get github.com/cheekybits/genny
$ go get github.com/ikasamt/zapp-jam
```

### features

- intSlice
- stringSlice
- anySlice

### functions 

- Max
- Min
- SortBy
- DistinctBy



### Usage

```
$ cd example
$ zapp-jam .
```

```

// example/city_test.go
// 北海道
hokkaido := cities.Where(func(x City)bool{return x.PrefectureName== `北海道`})

log.Println("total count: ",  hokkaido.Count())
log.Println("total population: ",  hokkaido.Select(`Population`).AsInt().Sum())
log.Println("sort desc population: ",  hokkaido.SortBy(func(a City,b City)bool{return a.Population > b.Population})[0])
log.Println("sort asc population: ",  hokkaido.SortBy(func(a City,b City)bool{return a.Population < b.Population})[0])

```