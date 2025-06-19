# Vitamin deficiency 🫑

We’re facing a severe vitamin C deficiency.
Luckily, vegetables with the highest vitamin C content are here to help.

We have a set of vegetables. Each of them contains a certain amount of vitamin C.
We also have a target amount of vitamin C we want to consume as a snack.
We are allowed to choose exactly two different vegetables.

Can we get the exact target amount by combining two of them?

---

### Input 🥦
A slice of integers — each integer represents the amount of vitamin C a vegetable contains.
A target sum — an integer we need to reach by summing two different elements of the slice.

### Output 🥕
A slice of two integers — indices of the elements whose values sum up to the target.
If no such pair exists, return nil.
If there are multiple valid pairs, return any of them.

---

### Examples 🥒:

- **Example 1**

    Input: `[26, 73, 14, 61, 100], 134`

    Output: `[1, 3]` (73 + 61 == 134)

- **Example 2**

    Input: `[26, 73, 14, 61, 85], 99`

    Output: `[0, 1]` or `[2, 4]` (26 + 73 == 14 + 85 == 99)

- **Example 3**

    Input: `[26, 73, 14, 61, 100], 30`

    Output: `nil`

- **Example 4**

    Input: `[26, 73, 14, 61, 100], 52`

    Output: `nil` (26 + 26 == 52, but we can't take one element twice)

- **Example 5**

    Input: `[26, 73, 14, 26, 100], 52`

    Output: `[0, 3]` (26 + 26 == 52, these are different elements, even though their values are the same)

- **Example 6**

    Input: `[], 15`

    Output: `nil`

---

### Solution 💡

🧠 Solution code: [vitamin_deficiency.go](./vitamin_deficiency.go)

🧪 Tests: [vitamin_deficiency_test.go](./vitamin_deficiency_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-can-two-veggies-cure-a-vitamin-crisis-56jl)