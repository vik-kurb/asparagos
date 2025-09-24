# Apple Escape 🍎

The season of Apple Pies has begun. All the apples are leaving their homes in search of a safer place. Nobody wants to end up in a pie anymore — it’s just not trendy. A smoothie is better, or at least a strudel.

There is a train going in one direction with a limited number of seats. Apples travel in groups: each group wants to board the train at point X and leave at point Y.

Can the train take every apple without exceeding its seats?

---

### Input 🍐

`trips []Trip` - information about group trips. Each `Trip` has a start point `from`, an end point `to`, and the group size `num`.

`capacity int` - the number of seats on the train.

### Output 🍊

`true` if the train can carry all apples without exceeding capacity at any point; `false` otherwise.

---

### Examples 🍋:

- **Example 1**

    Input: `trips = {{from: 2, to: 7, num: 1}, {from: 0, to: 2, num: 3}, {from: 1, to: 3, num: 2}}, capacity = 5`

    Output: `true`

    At point `0`, the train takes 3 apples.

    Then, at point `1`, it takes 2 more apples, so now all 5 seats are occupied.

    At point `2`, 3 apples leave the train. At the same time, 1 apple boards the train. Now there are 3 apples on board.

    At point `3`, 2 apples leave the train. Only 1 apple remains, which will leave at point `7`.

    The apples on board never exceeded the train’s capacity.

- **Example 2**

    Input: `trips = {{from: 0, to: 2, num: 3}, {from: 1, to: 3, num: 2}}, capacity = 4`

    Output: `false`

    At point `0`, the train takes 3 apples.

    At point `1`, it should take 2 more apples. That would require 5 seats, but the capacity is 4, so the train can't take them.

---

### Solution 💡

🧠 Solution code: [apple_escape.go](./apple_escape.go)

🧪 Tests: [apple_escape_test.go](./apple_escape_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-saving-apples-from-pies-58p8)
