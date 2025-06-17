# Twin salads 🥗

There are two chefs who have to prepare twin salads. In front of them lies a straight line of vegetables, arranged in a row. The chefs are standing on opposite ends of this line.

Each time, both chefs pick one vegetable — the one closest to them — and add it to their salad. Obviously, the order in which the vegetables are added is crucial (you know, when a cucumber joins a close-knit team of tomato and pepper, it's not the same as when the pepper crashes the cucumber-tomato party).

If, in the end, there’s only one vegetable left in the middle, the chefs can split it evenly between themselves. If there are no vegetables at all, the salads can be considered as twins.

**The question is: are the chefs able to make identical salads?**

---

### Input 🥦
A slice of strings — each string represents a vegetable.  

**Case doesn't matter**: it doesn't matter whether you call yourself `Broccoli` or `BrOcCoLi` — you're just a broccoli.

---

### Examples 🥒:

- **Example 1**

    Input: `["cucumber", "tomato", "pepper", "pepper", "tomato", "cucumber"]`

    Output: `true`

- **Example 2**

    Input: `["cucumber", "tomato", "pepper", "lettuce"]`

    Output: `false`

- **Example 3**

    Input: `["cucumber", "tomato", "pepper", "tomato", "cucumber"]`

    Output: `true`

- **Example 4**

    Input: `["cucumber", "tomato", "cucumber", "tomato"]`

    Output: `false`

- **Example 5**

    Input: `[]`

    Output: `true`

- **Example 6**

    Input: `["cucumber"]`

    Output: `true`

- **Example 7**

    Input: `["cucumber", "tomato", "ToMato", "CUCUMBER"]`

    Output: `true`

---

### Solution 💡

🧠 Solution code: [twin_salads.go](./twin_salads.go)

🧪 Tests: [twin_salads_test.go](./twin_salads_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-can-two-chefs-make-twin-salads-4de8)
