# The Smartest Peanut in the Neighborhood 🥜

Remember how the walnuts once tried to find out who was the [smartest](../walnut_iq/README.md)? Now it's the peanuts' turn, though their approach is a bit different.

Peanuts don't care about being the smartest among all nuts. They just want to feel smarter than their nearby neighbors.

You’re given the cleverness levels of peanuts planted in a row. For each group of `k` consecutive peanuts, determine which peanut is the smartest.

Is it a hard nut to crack?

---

### Input 🌰
- A slice of integers — each integer represents the cleverness level of a peanut in the row.
- `k` — an integer that represents the size of a group (a sliding window).

### Output 🌱
- A slice of integers — each representing the maximum cleverness level within a sliding window of size `k`.

---

### Examples 🥒:

- **Example 1**

    Input: `[1, 6, 3, 4, 1, 2, 6], k = 3`

    Output: `[6, 6, 4, 4, 6]`

    The first group is `[1, 6, 3]`, its maximum is `6`.
    The second group is `[6, 3, 4]`, its maximum is `6`.
    The third group is `[3, 4, 1]`, its maximum is `4`.
    The fourth group is `[4, 1, 2]`, its maximum is `4`.
    The last group is `[1, 2, 6]`, its maximum is `6`.

- **Example 2**

    Input: `[1, 2, 3, 2, 1], k = 2`

    Output: `[2, 3, 3, 2]`

    The first group is `[1, 2]`, its maximum is `2`.
    The second group is `[2, 3]`, its maximum is `3`.
    The third group is `[3, 2]`, its maximum is `3`.
    The last group is `[2, 1]`, its maximum is `2`.

---

### Solution 💡

🧠 Solution code: [smartest_peanut.go](./smartest_peanut.go)

🧪 Tests: [smartest_peanut_test.go](./smartest_peanut_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-the-smartest-peanut-in-the-neighborhood-3o92)
