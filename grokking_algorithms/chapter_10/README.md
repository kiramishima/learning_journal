# EXERCISES

## 10.1
You work for a furniture company, and you have to ship furniture all over the country. You need to pack your truck with boxes. All the boxes are of different sizes, and you’re trying to maximize the space you use in each truck. How would you pick boxes to maximize space? Come up with a greedy strategy. Will that give you the optimal solution?

My answer:

```
1- Get the size of the space.
2- Divide the space in slices just found the best size for the boxes.
```

Book answer:

```
Answer: A greedy strategy would be to pick the largest box that
will fit in the remaining space and repeat until you can’t pack any
more boxes. No, this won’t give you the optimal solution
```
## 10.2
You’re going to Europe, and you have seven days to see everything you can. You assign a point value to each item  (how much you want to see it) and estimate how long it takes. How can you maximize the point total (seeing all the things you really want to see) during your stay? Come up with a greedy strategy. Will that give you the optimal solution?

My answer:

```
Made a schedule with the hours and places, and how many time it costs visiting the place.
```

Book answer: 

```
Answer: Keep picking the activity with the highest point value that
you can still do in the time you have left. Stop when you can’t do
anything else. No, this won’t give you the optimal solution.
```