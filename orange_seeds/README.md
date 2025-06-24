# Oranges seeds 🍊

After the successful election victory, the fruits decided to expand their power and capture as much territory as possible (you can find the backstory of how fruits gained power in the Veggie Kingdom [here](../fruit_sabotage/README.md)).

The oranges were given an important mission: they should plant orange trees along the entire alley. The alley is a narrow strip of land divided into sectors. Only one tree can be planted per sector (orange trees need their privacy, they are definitely introverts).

Each tree can throw seeds to nearby sectors, but only within its maximum throwing distance. For example, if a tree stands at sector N and its maximum throwing distance is K, it can throw seeds to every sector in the interval [N +1, N + K]. A new tree grows in each sector that receives a seed.

The goal of the orange trees is to plant a tree at the last sector of the alley. They start at the first sector. It's guaranteed that they'll succeed, because oranges never fail their missions.

**The question is: what's the minimum number of throws they need to complete the mission?**

---

### Input 🍎
A slice of integers — each integer represents the maximum distance a seed can be thrown from that sector.  

---

### Examples 🍉:

- **Example 1**

    Input: `[3, 1, 5, 1, 1]`

    Output: `2`: the tree at sector `0` throws to sector `2`, and the tree at sector `2` throws to the last sector.

- **Example 2**

    Input: `[1, 3, 16, 1, 1, 5, 1]`

    Output: `3`: the tree at sector `0` throws to sector `1`, the tree at sector `1` throws to sector `2`, and the tree ar sector `2` throws to the last sector.

- **Example 3**

    Input: `[1]`

    Output: `0`: we're already at the last sector.

---

### Solution 💡

🧠 Solution code: [orange_seeds.go](./orange_seeds.go)

🧪 Tests: [orange_seeds_test.go](./orange_seeds_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-can-oranges-sow-the-seeds-of-discord-in-linear-time-1643)
