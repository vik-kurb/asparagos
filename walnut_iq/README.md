# Walnuts Take the IQ Test 🌰

Walnuts believe they are the smartest among all the plants (though not everyone agrees - especially spinach). Each walnut also believes that they are the smartest among all walnuts. Now, they want to convince everyone of their genius.

All walnuts have taken an IQ test. We have the IQ results for each of them. Now, they want to know if they are smarter than the average IQ of all the other walnuts - excluding themselves.

But that's not all. After [gaining power](../fruit_sabotage/README.md), the fruits passed a new law: **No negativity!** Everyone should be happy and support the new King! So, subtraction is now banned in the kingdom. We must solve the problem without using the `-` operator.

---

### Input 🥦
A slice of integers — each integer represents the result of one walnut's IQ test. There are at least two elements in the slice.

### Output 🥕
A slice of integers — each integer represents the average IQ score of all the walnuts except the one at that index. The average should be calculated as the sum of the other results divided by their count, rounded down.

---

### Examples 🥒:

- **Example 1**

    Input: `[14, 75, 24, 86]`

    Output: `[61, 41, 58, 37]`

    For index 0: average of [75, 24, 86] = (75 + 24 + 86) / 3 = 61.66... → 61
    For index 1: average of [14, 24, 86] = 41.33... → 41
    For index 2: average of [14, 75, 86] = 58.33... → 58
    For index 3: average of [14, 75, 24] = 37.66... → 37

- **Example 2**

    Input: `[15, 23, 48]`

    Output: `[35, 31, 19]`

    For index 0: average of [23, 48] = (23 + 48) / 2 = 35.5 → 35
    For index 1: average of [15, 48] = 31.5 → 31
    For index 2: average of [15, 23] = 19

- **Example 3**

    Input: `[15, 23]`

    Output: `[23, 15]`

    Each index simply returns the other value.

---

### Solution 💡

🧠 Solution code: [walnut_iq.go](./walnut_iq.go)

🧪 Tests: [walnut_iq_test.go](./walnut_iq_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-walnuts-take-the-iq-test-8dl)
