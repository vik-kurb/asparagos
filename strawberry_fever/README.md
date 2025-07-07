# Strawberry Fever 🍓

There’s one boy who is crazy about strawberries (and honestly, who isn’t?). One day, he decided to eat all the strawberries in the city.

There is a bunch of stores arranged in a circle where strawberries can be bought. The boy wants to visit all of them in a clockwise direction and buy every strawberry he can. For each store, we know:
- how many strawberries can be bought there,
- how many strawberries he needs to eat to have enough energy to move from that store to the next one.

We don't know where the boy will start his route, but he must finish it at the same store where he started.

Can the boy visit all the stores if he has unlimited money and can eat all the strawberries he collects along the way?

It's guaranteed that if completing the route is possible, there is exactly one valid starting point. In that case, we should find this starting point — and stop his evil strawberry plan!

---

### Input 🍒
`storeStock []int` — each `storeStock[i]` represents the number of strawberries available in the `i`th store.

`wayCost []int` — each `wayCost[i]` represents the number of strawberries the boy needs to eat to travel from store `i` to store `i+1`.

### Output 🍇
An integer — the index of the store where the boy should start his route to complete his plan. Return `-1` if no such starting point exists.

---

### Examples 🫐:

- **Example 1**

    Input: `storeStock = [1, 2, 3, 4, 5], wayCost = [3, 4, 5, 1, 2]`
    Output: `3`

    He starts at store `3` and eats 4 strawberries. His energy reserve is 4 strawberries.
    He travels to store `4`, spends 1 strawberry of energy, and eats 5 strawberries. His energy reserve is 4 - 1 + 5 = 8 strawberries.
    He travels to store `0`, spends 2 strawberries, and eats 1 strawberry. His energy reserve is 8 - 2 + 1 = 7.
    He travels to store `1`, spends 3 strawberries, eats 2 strawberries. His energy reserve is 7 - 3 + 2 = 6.
    He travels to store `2`, spends 4 strawberries, eats 3 strawberries. His energy reserve is 6 - 4 + 3 = 5.
    He travels to store `3`, spends 5 strawberries. Mission completed.


- **Example 2**

    Input: `storeStock = [2, 3, 4], wayCost = [3, 4, 3]`

    Output: `-1`

    If he starts at store `0` he eats 2 strawberries, but he needs 3 strawberries to get to the next store.

    If he starts at store `1` he eats 3 strawberries, but he needs 4 strawberries to get to the next store.

    If he starts at store `2` he eats 4 strawberries. Then, he travels to store `0`, spends 3 strawberries and eats 2 strawberries. His energy reserve is 4 - 3 + 2 = 3 strawberries.
    He travels to store `1`, spends 3 strawberries and eats 3 strawberries. His energy reserve is 3 - 3 + 3 = 3 strawberries. But he needs 4 strawberries to finish the circle and travel to the store `2`. So it's impossible.

---

### Solution 💡

🧠 Solution code: [strawberry_fever.go](./strawberry_fever.go)

🧪 Tests: [strawberry_fever_test.go](./strawberry_fever_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-are-strawberries-in-danger-nab)