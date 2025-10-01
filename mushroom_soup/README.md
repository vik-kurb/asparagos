# Mushroom Soup 🍄

Autumn is a great time for soup, so the Veggie Kingdom is gathering mushrooms for a creamy mushroom soup. But there’s one problem: lactose intolerance. What's the solution? Cook two pots: one with dairy cream and one with coconut cream.

The amount of soup should be the same in both pots, so we need to split the mushrooms evenly by weight. Each mushroom has its own weight, so we can’t just divide them by count.

How can we quickly check whether it’s possible to partition the mushrooms into two groups with equal total weight?

---

### Input 🥔

`mushrooms []int` - weights of the mushrooms (each element is the weight of one mushroom).

### Output 🧅

Return `true` if it’s possible to partition the mushrooms evenly by weight (into two groups with equal total weight), otherwise return `false`.

---

### Examples 🧄:

- **Example 1**

    Input: `mushrooms = [1,5,11,5]`

    Output: `true`

    The first pot contains mushrooms with weights `1, 5, 5`, and the second pot contains a single mushroom with weight `11`.

- **Example 2**

    Input: `mushrooms = [1, 3, 10]`

    Output: `false`

    It’s impossible to split the mushrooms into two groups with equal total weight.

---

### Solution 💡

🧠 Solution code: [mushroom_soup.go](./mushroom_soup.go)

🧪 Tests: [mushroom_soup_test.go](./mushroom_soup_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-mushroom-soup-for-everyone-55a3)
