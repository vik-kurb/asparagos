# Fruit Sabotage 🍍

The election day has come to the Veggie Kingdom. All vegetables and fruits are casting their votes for the next King — each one voting for a single candidate.

Just before the election, a secret plan was revealed:
all fruits secretly agreed to vote for the same fruit!

Honestly, we get it — fruits are clearly tired of the long-running vegetable monarchy. More and more vegetables are being added to smoothies, carrot cakes are gaining popularity... that would drive anyone bananas.

We don’t know which fruit they’ve chosen as their future King, but we do know this: they hold the majority. The total number of votes is N, and the number of fruit votes is greater than ⌊N / 2⌋.

Can we figure out who the next fruit King will be in the Veggie Kingdom?

A small note: there's a bit of a crisis in the Kingdom — all the budget was spent on saving potatoes from bugs, so we’re only allowed to use constant memory.

---

### Input 🥦
A slice of integers — each integer represents a vote for some candidate.

Although we are vegetables and fruits, this is a serious election: all votes are anonymous — and there is at least one vote (mine).

### Output 🥕
An integer — the value that occurs more than ⌊N / 2⌋ times, where N is the length of the input slice.
It’s guaranteed that such a value exists.

---

### Examples 🥒:

- **Example 1**

    Input: `[1, 2, 2]`

    Output: `2`

- **Example 2**

    Input: `[3, 2, 3, 4, 3, 4, 3, 3]`

    Output: `3`

- **Example 3**

    Input: `[10]`

    Output: `10`

---

### Solution 💡

🧠 Solution code: [fruit_sabotage.go](./fruit_sabotage.go)

🧪 Tests: [fruit_sabotage_test.go](./fruit_sabotage_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-can-we-find-the-king-in-o1-space-4fd4)
