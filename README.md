# what is
行列をいい感じに処理する

# how to use
`go run main.go size matrix.file`

### ファイル形式
拡大行列形式`(A | b)`で書く
```
1 2 3 a
4 5 6 b
7 8 9 c
```

この例であれば`go run main.go 3 file`にて実行可

# to do
- 計算の高速化
 - もっとgoroutineがうまく使える？
 - 諸々最適化
- 各種処理の実装
 - 正規化とか

