# Messing around with heaps in Go

Some basic operations in heap structure viewed in a interactive manner, and also his performance. The `pkg` includes methods to `PushBack`, `PushFront`, `GoUp` and `GoDown` in a slice with a heap property. Also you can visualize it in a pretty manner with `prettyprint`, also inside `pkg`.

<center>
<img src="assets/BELGIUM.png" style ="width: 30%; height: auto;">
</center>

### PrettyPrint

The heap will be exhibited like this:

```
               50
               / \
       48               45
       / \               / \
   29       15       35       40
   / \       / \       / \       / \
 27   26   4   12   33   30   37   20
 / \   /
21 19 25
```

## Usage

- **Compile**:

  For building the interactive version.

  ```terminal
  make i
  ```

  > It will generate `interactive-heap*`

  For building the version with the performance tests.

  ```
  make p
  ```

  > It will generate `performance-heap*`

- **Running**:

  Simply execute the generated binary.

## Contributing

- Implement a `PopFront` and `PopBack` for the heap.
- Implement the `HeapSort`.

---

&copy; Lima
