# The World’s Biggest Gazpacho 🍅

Everyone is suffering from the summer heat. One chef decides to surprise everyone by making the biggest gazpacho ever.

He needs as many tomatoes as possible. There are several tomato beds in the garden, but he can't gather them all. He wants to keep it a secret -- nobody should notice anything suspicious for as long as possible. So he comes up with a strategy: he’ll collect all the tomatoes from some beds and leave the others untouched — but he won’t leave two adjacent beds empty. Otherwise, someone might suspect something (especially the lemons -- they’re always sticking their nose where it doesn’t belong).

Time is running out, the sun is burning, and we all want that delicious cold soup as soon as possible. How can the chef gather the maximum number of tomatoes? And how many tomatoes can he collect in total?

---

### Input 🥒
A slice of integers representing tomato beds —- each integer is the number of tomatoes in that bed.

### Output 🫑
An integer —- the maximum number of tomatoes the chef can collect without leaving two adjacent beds empty.

---

### Examples 🧅:

- **Example 1**

    Input:
    ```
    [6, 3, 5, 10]
    ```
    Output: `16`

    1. If we take tomatoes from the first bed, we have two options:

    a. Take tomatoes from the third bed: 6 + 5 = 11
    b. Take tomatoes from the last bed: 6 + 10 = 16

    2. If we take tomatoes from the second bed, we can only take tomatoes from the last bed: 3 + 10 = 13

    The best option gives us 16 tomatoes.

- **Example 2**

     Input:
    ```
    [3, 7, 2]
    ```
    Output: `7`

    1. If we take tomatoes from the first bed, we can also take tomatoes from the last bed: 3 + 2 = 5
    2. If we take tomatoes from the second bed, we can’t take any others, so we get 7 tomatoes.
    The best option gives us 7 tomatoes.

---

### Solution 💡

🧠 Solution code: [worlds_biggest_gazpacho.go](./worlds_biggest_gazpacho.go)

🧪 Tests: [worlds_biggest_gazpacho_test.go](./worlds_biggest_gazpacho_test.go)

📖 Full explanation: Blog post on [Dev.to]()
