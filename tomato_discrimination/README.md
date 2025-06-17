# Tomato discrimination 🍅

Vegetables are waiting in a queue to be added to the soup.
Suddenly, the Botanical Government passes a new law: tomatoes are now officially fruits!

Thankfully, everybody still loves tomatoes in soups, so they won’t be removed.
But there’s a catch: all tomatoes must now be moved to the end of the queue, no matter where they were before the law. The order of all other vegetables must remain unchanged.

Be careful: some of them will try to mislead you by calling themselves pomodoro.
Don’t fall for it — anyone who calls themselves "tomato" or "pomodoro" must go to the end.
Also, to avoid further confusion, rename all "pomodoro" to "tomato" before adding them to the end.

---

### Input 🥦
A slice of strings — each string represents a vegetable.
Only lowercase letters are allowed.
Only "tomato" and "pomodoro" should be discriminated.

### Output 🥕
A slice of strings — each string represents a vegetable.
The order of all non-tomatoes stays the same.
All "tomato" and "pomodoro" (renamed to "tomato") appear at the end of the queue.

---

### Examples 🥒:

- **Example 1**

    Input: `["cucumber", "tomato", "pepper", "tomato", "lettuce"]`

    Output: `["cucumber", "pepper", "lettuce", "tomato", "tomato"]`

- **Example 2**

    Input: `["tomato", "pomodoro", "lettuce"]`

    Output: `["lettuce", "tomato", "tomato"]`

- **Example 3**

    Input: `["cucumber", "pepper", "lettuce"]`

    Output: `["cucumber", "pepper", "lettuce"]`

- **Example 4**

    Input: `[]`

    Output: `[]`

---

### Solution 💡

🧠 Solution code: [tomato_discrimination.go](./tomato_discrimination.go)

🧪 Tests: [tomato_discrimination_test.go](./tomato_discrimination_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-tomato-discrimination-in-the-soup-queue-4i7)