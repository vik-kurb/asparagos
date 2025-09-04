# Coconut Race 🥥

Coconuts are having a race! Just for fun. After all, there’s no real need to hurry when you’re a coconut.

Each coconut starts from a certain position and has its own speed. They’re all rolling in the same direction toward a common finish line.

Whenever one coconut catches up with another, they merge into a coconut bouquet and continue rolling together at the speed of the slowest coconut in the group.

How many bouquets will reach the finish line?

---

### Input 🥭

`position []int` - the starting positions of the coconuts.

`speed []int` - the speeds of the coconuts.

`target int` - the finish line; all coconuts are rolling toward this point.

### Output 🍍

An integer representing the number of coconut bouquets that reach the target.

---

### Examples 🍌:

- **Example 1**

    Input: `position = [1, 3, 18], speed = [5, 1, 3], target = 20`

    Output: `2`

    The first coconut (from position `1`, speed `5`) catches up to the second one (position `3`, speed `1`) before reaching the target.

    They form a bouquet and move together at speed `1`.

    The third coconut (position `18`, speed `3`) reaches the target before the bouquet catches up.

    So, there are `2` bouquets at the finish.

- **Example 2**

    Input: `position = [2, 0, 4], speed = [3, 7, 1], target = 10`

    Output: `1`

    The coconut starting from `0` (speed `7`) quickly catches up to the one at `2` (speed `3`).

    Together they move at speed `3` and catch up to the coconut at `4` (speed `1`) before reaching the target.

    All three coconuts finish as a single bouquet.

---

### Solution 💡

🧠 Solution code: [coconut_race.go](./coconut_race.go)

🧪 Tests: [coconut_race_test.go](./coconut_race_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-coconuts-never-roll-alone-8f8)
