# Potato Bugs 🥔

Yet another challenging day in the Veggie Kingdom. To save the harvest, all potatoes must be treated for bugs. That’s why they’re lining up at the Fight for Insect-Free Organics (FIFO) office.

But potatoes aren’t the sharpest veggies — they might mess things up and accidentally form a cycle in the queue. We need to check whether their queue contains a cycle.

The Kingdom is going through hard times, so we must solve this using only constant memory.

---

### Input 🥦
A head of a linked list of potatoes. Each potato has a Name and a pointer to the next one:

```go
type PotatoNode struct {
    Name string
    Next *PotatoNode
}
 ```

### Output 🥕
Return `true` if the list contains a cycle, or `false` otherwise.

---

### Examples 🥒:

- **Example 1**

    Input:
    ```
    Potato Jack -> Potato George -> Potato Arthur -> nil
    ```
    Output: `false`

- **Example 2**

    Input:
    ```
    Potato Jack -> Potato George -> Potato Arthur
                        ↑                ↓
                        ← ← ← ← ← ← ← ← ←
    ```
    Output: `true`

---

### Solution 💡

🧠 Solution code: [potato_bugs.go](./potato_bugs.go)

🧪 Tests: [potato_bugs_test.go](./potato_bugs_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/asparagos-vs-potato-bugs-can-he-detect-the-cycle-in-o1-space-1dll)
