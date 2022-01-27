泛型
=====

### 场景

假设需要写一个列表总数计算的函数，根据不同数据类型，我们可能分别要写 `SumInts(data []int)`，`SumFloats(data []float64)` 等函数.

若使用泛型的话，我们只需要写一个函数 `SumIntsOrFloats[T int | float64](data []T)` 即可.

### 参考
[Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
