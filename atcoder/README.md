# 知見まとめ

<!-- TODO: 目次を挿入する-->

## modは計算途中で何度挟んでもよい

答えが大きくなる場合に「その量を 1000000007 で割ったあまりを求めよ」という問題が出されることがある。  
割る数は`1000000007`以外でもよく、ただその数が手頃な素数であるよう。  

### 解説

#### 足し算・引き算・掛け算の場合

シンプルに解けるのは、この3つの演算のみ  
**計算の途中過程で積極的にあまりをとってよい**

[playground](https://play.golang.org/p/TMqY4ZR9yMe)

```go
const mod = 1000000007

func main() {

    a := 11111
    b := 12345
    c := 98765

    fmt.Println("まとめてmod: ", a*b*c%mod)    // 130265846
    fmt.Println("都度mod　 : ", a*b%mod*c%mod) // 130265846
}
```

### 該当する問題

- [abc177c](https://atcoder.jp/contests/abc177/tasks/abc177_c)

### 参考文献

- [「1000000007 で割ったあまり」の求め方を総特集！ 〜 逆元から離散対数まで 〜](https://qiita.com/drken/items/3b4fdf0a78e7a138cd9a)

## 累積和

総和を求めるときに泥臭い方法だと制限時間に達してしまうときなどに必要になる方法  

### 解説

> 適切な前処理をしておくことで、配列上の区間の総和を求めるクエリを爆速で処理できるようになる手法

`a[0]...a[n-1]` に対して `s[0]...s[n]` を以下のように定めたもの

```shell
s[0]=0
s[1]=s[0]+a[0]
s[2]=s[1]+a[1]
...
s[n]=s[n-1]+a[n-1]
```

#### 例

`a={3,6,8,2}` のときに、`s={0,3,9,17,19}` となる  

これを用いると、例えば配列aの区間`[1,3)` (2は含まない)の総和は `s[3]-s[1]` で計算できる  
実際にやってみると、`17 - 3 = 14` で手で泥臭く数えたときの `6 + 8 = 14`と同じになることがわかる

[playground](https://play.golang.org/p/H7SyNkvRbae)

```go
func main() {
    a := []int{3,6,8,2}
    s := make([]int, len(a)+1)
    for i, v := range a {
        s[i+1] = s[i] + v
    }
    fmt.Println(s[3]-s[1])  // 14
}
```

#### 特徴

- 前処理に O(N)O(N) だけの時間をかける
- 記憶容量は O(N)O(N) でよい
- 前処理をしておけば、毎回のクエリに O(1)O(1) で爆速で答えることができる


### 該当する問題

- [abc177c](https://atcoder.jp/contests/abc177/tasks/abc177_c)

### 参考文献

- [累積和を何も考えずに書けるようにする！](https://qiita.com/drken/items/56a6b68edef8fc605821)

## 優先度付きキュー

要素１つ１つに優先度が割り当てられており、その優先度に従った順に取り出される。  
整数を格納する優先度付きキューであれば、その要素の値自体を優先度と考え取り出す。  
ヒープでの実装を用いた場合は、追加が`O(logN)`で、参照は`O(1)`となる。  

### 解説

[playground](https://play.golang.org/p/B_FF1GCtaCS)

```go
func main() {
    h := &IntHeap{2, 5, 1}
    heap.Init(h)
    heap.Push(h, 7)
    for h.Len() > 0 {
        fmt.Printf("%d ", heap.Pop(h))
    }
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}
```

### 該当する問題

- [abc173d](https://atcoder.jp/contests/abc173/tasks/abc173_d)

### 参考文献

- [優先度付きキュー](https://programming-place.net/ppp/contents/algorithm/data_struct/010.html)
- [Package heap](https://golang.org/pkg/container/heap/)

## 二重ループだけど`O(N^2)`ではなく`O(M+N)`

例えば、`a[0..n]`と`b[0..m]`の配列があり、それぞれを2重ループで舐める必要があるとする。ただ、`n`と`m`の値が大きい場合に、実行時間が長くなってしまいます。  
もし、問題の特性上、ループのある時点の`b`の値を保持/更新して実際には一度舐めるだけで十分の場合は、`O(M+N)`で処理することが可能

[playground](https://play.golang.org/p/5AXBz4F9i7k)

### 該当する問題

- [abc172c](https://atcoder.jp/contests/abc172/tasks/abc172_c)