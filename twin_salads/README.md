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

You can find the solution code [here](./twin_salads.go).

1. First of all, if the input has 0 or 1 elements, the salads are automatically identical, and there's nothing else to check.

2. We can make twin salads if the first vegetable is the same as the last, the second is the same as the second to last, and so on.

    We'll use two indices `leftInd` and `rightInd`.
    `leftInd` points to the vegetable that the first chef should pick up next.
    `rightInd` points to the vegetable that the second chef should pick up next.

    The first chef goes from left to right, so `leftInd` will be increased at each iteration. The second chef goes from right to left, so `rightInd` will be decreased.

3. At each iteration we check if the vegetables are equal. We use `strings.ToLower()`, because case doesn't matter.

    If the left and right vegetables are not equal, it means we can't make identical salads.

    And if we've checked all the pairs without finding any differences, we can be sure that salads are twins.

[Here](./twin_salads_test.go) you can find tests.
