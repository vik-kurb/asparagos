# Lemon Squeezy 🍋

Summer is over, lemonade is out of fashion, but the lemons refuse to waste time. They've formed a team called "Lemon Squeezy" and are aiming to win the annual pub quiz.

Now they must decide who joins the team. Each lemon has two attributes: IQ and age. They want a team with the maximum total IQ. But lemons have fragile citrus egos: no lemon will play in a team with a younger lemon who has a higher IQ.

Can they quickly figure out who should be in the team? And what will the total IQ of "Lemon Squeezy" be?

---

### Input 🍊

`iqs []int` - IQ of each lemon.
`ages []int` - age of each lemon.

### Output 🍐

An integer — the maximum total IQ of a team with no conflicts. A conflict exists if a younger lemon has a higher IQ than the older lemon. Lemons of the same age never conflict, regardless of IQ.

---

### Examples 🍈:

- **Example 1**

    Input: `iqs = [1, 2, 3], ages = [10, 15, 20]`

    Output: `6`

    All lemons can be in the same team.

- **Example 2**

    Input: `iqs = [4, 9, 20], ages = [20, 10, 15]`

    Output: `29`

    The team consists of the second and third lemons. The first lemon can’t be included because its `IQ = 4` is lower than the IQ of the second lemon, who is younger.

- **Example 2**

    Input: `iqs = [17, 24, 50], ages = [46, 37, 5]`

    Output: `50`

    The team consists of a single lemon — the last one. Any other combination would include a younger lemon with a higher IQ than an older lemon, causing a conflict.

---

### Solution 💡

🧠 Solution code: [lemon_squeezy.go](./lemon_squeezy.go)

🧪 Tests: [lemon_squeezy_test.go](./lemon_squeezy_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-lemon-squeezy-dynamic-easy-12lc)
