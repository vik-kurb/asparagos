# Watermelon Football 🍉

Welcome to The Great Watermelon Football Final! The watermelons are ready, and the stakes are high. There’s just one task left: we must split the players into two teams.

The watermelons are standing in a row, each with a rating. We must remove exactly one watermelon (don’t worry, this one becomes the referee). The remaining watermelons are split into two teams by index in the remaining row:

Team Even: players at even indices

Team Odd: players at odd indices

The total rating of both teams must be equal.

How many choices of the referee (how many indices to remove) make the teams' total ratings equal?

---

### Input 🍈

`rating []int` - the ratings of the watermelons, `len(rating) > 2`.

### Output 🥭

An integer representing the number of ways to choose one watermelon to remove (the referee) so that the remaining players can be split by index into two teams with equal total rating.

---

### Examples 🥝:

- **Example 1**

    Input: `rating = [1, 3, 2, 1, 2, 2]`

    Output: `2`

    We can remove the 2nd watermelon (rating `3`), remaining players are `[1, 2, 1, 2, 2]`.
    Team Even: `1 + 1 + 2 = 4`, Team Odd: `2 + 2 = 4`.

    We can remove the 4th watermelon (rating `1`), remaining players are `[1, 3, 2, 2, 2]`.
    Team Even: `1 + 2 + 2 = 5`, Team Odd: `3 + 2 = 5`.


- **Example 2**

    Input: `rating = [5, 7, 3, 8, 1]`

    Output: `1`

    We can remove the 4th watermelon (rating `8`), remaining players are `[5, 7, 3, 1]`.
    Team Even: `5 + 3 = 8`, Team Odd: `7 + 1 = 8`.


- **Example 3**

    Input: `rating = [1, 2, 3, 4]`

    Output: `0`

    There are no ways to remove one watermelon and obtain equal team totals.

    
---

### Solution 💡

🧠 Solution code: [watermelon_football.go](./watermelon_football.go)

🧪 Tests: [watermelon_football_test.go](./watermelon_football_test.go)

📖 Full explanation: Blog post on [Dev.to](https://dev.to/asparagos/go-coding-with-asparagos-fair-play-in-watermelon-football-55io)
