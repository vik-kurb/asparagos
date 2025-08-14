# Salsa Festival 🥗

There’s a salsa festival in the Veggie Kingdom this weekend. All vegetables are invited, free of charge!

But veggies are picky: each one will only come if certain friends arrive first.
For example, bell pepper will only show up after his brother, chili pepper. A potato will only come if both pumpkin and corn are already there (this trio became best friends after the lecture “Introduction to the Carbohydrate Diet”).

Nobody can untangle these complicated relationships. (Did you know that apple is trying to sue pineapple for using its name?)
So, can we figure out whether all the veggies will come to the salsa festival?

---

### Input 🥒
`vegNum int` - the total number of vegetables in the kingdom.

`requirements [][]int` - each element `requirements[i] = [a, b]` means that vegetable `a` will come to the festival only if vegetable `b` comes first.

### Output 🍅
`true` if it’s possible for all vegetables to attend the festival,

`false` if some requirements create a deadlock (i.e. a cycle).

---

### Examples 🌶️:

- **Example 1**

    Input: `vegNum = 2, requirements = [[1, 0]]`

    Output: `true`

    Vegetable `0` can come without any conditions. Vegetable `1` depends on `0`, and since `0` can come, `1` can come too.

- **Example 2**

    Input: `vegNum = 2, requirements = [[1, 0], [0, 1]]`

    Output: `false`

    Vegetable `1` depends on `0`, and `0` depends on `1`. This creates a cycle, neither can come first, so no one comes.

---

### Solution 💡

🧠 Solution code: [salsa_festival.go](./salsa_festival.go)

🧪 Tests: [salsa_festival_test.go](./salsa_festival_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-will-graph-cycles-spoil-the-salsa-festival-2eod)
