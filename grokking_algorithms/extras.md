This appendix discusses the performance of AVL trees, which were introduced in chapter 8. You will need to read that chapter before reading this.

Remember that AVL trees offer $O(log n)$ search performance. But there is something misleading going on. Here are two trees. Both offer $O(log n)$ search performance, but their heights are different!

<table>
    <tr>
        <td>
            <img src="fig_1.png">
        </td>
        <td>
            <img src="fig_2.png">
        </td>
    </tr>
</table>

(Dashed nodes show the holes in the tree.)

AVL trees allow a difference of one in heights. That’s why, even though both these trees have 15 nodes, the perfectly balanced tree is height 3, but the AVL tree is height 4. The perfectly balanced tree is what we might picture a balanced tree to look like, where each level is completely filled with nodes before a new level is added. But the AVL tree is also considered “balanced,” even though it has holes—gaps where a node could be.

Remember that in a tree, performance is closely related to height. How can these trees offer the same performance if their heights are different? Well, we never discussed what the base in log n is!

The perfectly balanced tree has performance O(log n), where the “log” is log base 2, just like binary search. We can see that in the picture. Each new level doubles the nodes plus 1. So a perfectly balanced tree of height 1 has 3 nodes, of height 2 has 7 nodes (32 + 1), of height 3 has 15 nodes (72 + 1), etc. You could also think of it as each layer adds a number of nodes equal to a power of 2.

![](fig_3.png)

So the perfectly balanced tree has performance O(log n), where the “log” is log base 2.

The AVL tree has some gaps. In an AVL tree, each new layer adds less than double the nodes. It turns out that an AVL offers performance O(log n), but the “log” is log base phi (aka the golden ratio, aka ~1.618).

This is a small but interesting difference—AVL trees offer performance that is not quite as good as perfectly balanced trees since the base is different. But the performance is still very close, since both are O(log n) after all. Just know it’s not exactly the same.

# NP-hard problems

Both the set-covering and traveling salesperson problems have something in common: they are hard to solve. You have to check every possible iteration to find the smallest set cover or the shortest route.

Both of these problems are NP-hard. The terms NP, NP-hard, and NP-complete can cause a lot of confusion. They certainly confused me. I’ll try to explain what all these terms mean, but I need to explain some other concepts first. Here is a roadmap of the things we will learn and how they depend on each other:

![](fig_4.png)

But, first, I need to explain what a decision problem is because all the problems we will look at in the rest of this appendix are decision problems.

## Decision problems

NP-complete problems are always decision problems. A decision problem has a yes-or-no answer. The traveling salesperson problem is not a decision problem. It’s asking you to find the shortest path, which is an optimization problem.

**Note**
> I know I was talking about NP-hard problems in the introduction, and now I’m talking about NP-complete problems. I’ll explain what the difference is soon.

![](fig_5.png)

Here’s a decision version of the traveling salesperson problem.

![](fig_6.png)

Notice how this question has a yes-or-no answer: is there a path of length 5? I wanted to talk about decision problems up front because all NP-complete problems are decision problems. So all *the problems I discuss in the rest of this will be decision problems*. So when you see “traveling salesperson” mentioned in the rest of this appendix, I mean the *decision* version of the traveling salesperson problem.

Now let’s start learning what NP-complete actually means! The first step is to learn about the satisfiability (SAT) problem.

## The satisfiability problem

![](fig_7.png)

Jerry, George, Elaine, and Kramer are all ordering pizza.

“Ooh, let’s get pepperoni!” says Elaine.

“Pepperoni is good. Sausage is good. We could get pepperoni or sausage,” says Jerry.

“Get me an olive pizza to maintain my complexion,” says Kramer. “Lots of olives. Or, onions.”

“I can do any pizza but no onions,” says George. “I can’t take any more onions, Jerry!”

“Oh boy. OK, let me figure this out. So what toppings do I need again?” says Jerry.

Can you help him out? Here are everyone’s requirements:

- Pepperoni (Elaine)
- Pepperoni or sausage (Jerry)
- Olives or onions (Kramer)
- No onions (George)

See if you can figure out what toppings the pizza should have before moving on.

Did you get it? A pepperoni and olive pizza satisfies all the requirements. This is an example of a SAT problem. In pseudocode, I could write it like this. First, I have four boolean variables:

```
pepperoni = ?
sausage = ?
olives = ?
onions = ?
```

Then I write out a boolean formula: 
```
(pepperoni) and (pepperoni or sausage) and (olives or onions) and (not onions)
```

This formula contains the requirements for each person in the form of boolean logic. The SAT problem asks the question: Can you set these variables to some values so that the statement evaluates to `true`?

The SAT problem is famous because it is the first NP-complete problem, described in 1971 (although I don’t think the authors used Seinfeld as an example). Before this, the concept of an NP-complete problem did not exist. Here is how the SAT problem works. You start with a boolean
formula:

```
if (pepperoni) and (olives or onions):
    print("pizza")
```

Then you ask, is there some way we can assign our variables so that code prints `pizza`?

This example is pretty easy, so we can solve it ourselves. If `pepperoni` and `onions` are `true`, this code will print `pizza`. So the answer would be *yes*.

Here’s one where the answer would be *no*:

```
if (olives or onions) and (not olives) and (not onions):
    print("pizza")
```

There is nothing you can set the variables to so this code will print `pizza`!

The SAT problem always looks for a yes-or-no answer, so it is a decision problem.

SAT is actually a pretty hard problem. Here is a tougher example just to give you an idea:

```
if (pepperoni or not olives) and (onions or not pepperoni) and (not olives or not pepperoni):
    print("pizza")
```

![](fig_8.png)

You don’t need to solve this one. I’m just showing it as an example so you can appreciate how hard this problem can get. You can have any number of variables and any number of clauses, and the problems get pretty hard pretty quickly.

With $n$ toppings, there are $2^n$ possible pizzas. If you list them all out and check each one, you get something called a truth table. Here’s the truth table for `pepperoni and (olives or onions)`.

Sometimes you need to list every option, just like the set-covering problem and the traveling salesperson problem! In fact, the SAT problem is just as hard as these two problems. It has a big O run time of $O(2^n)$.

## Hard to solve, quick to verify

