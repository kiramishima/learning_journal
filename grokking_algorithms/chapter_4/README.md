# EXERCISES
## 4.1
Write out the code for the earlier sum function.

My answer:
```python
# the book uses python, this my python version
def sum(arr):
    if len(arr) == 0:
        return 0
    else:
        return arr[0] + sum(arr[1:])

print(sum([1,2,3,4,5])) # prints 15

```

Book answer:
```python
def sum(list):
    if list == []:
        return 0
    return list[0] + sum(list[1:])
```

## 4.2
Write a recursive function to count the number of items in a list.

My answer:
```python
def count(arr):
    if len(arr) == 0:
        return 0
    else:
        return 1 + count(arr[1:])
    
print(count([1,2]))
```

Book answer:
```python
def count(list):
    if list == []:
        return 0
    return 1 + count(list[1:])
```
## 4.3
Find the maximum number in a list.


## 4.4
Remember binary search from chapter 1? It’s a divide-and-conquer algorithm, too. Can you come up with the base case and recursive case for binary search?

My answer:

# How long would each of these operations take in Big O notation?

## 4.5
Printing the value of each element in an array.

My answer: $O(n)$

Book answer: $O(n)$

## 4.6
Doubling the value of each element in an array.

My answer: $O(n)$

Book answer: $O(n)$

## 4.7
Doubling the value of just the first element in an array.

My answer: $O(1)$

Book answer: $O(1)$

## 4.8
Creating a multiplication table with all the elements in the array. So if your array is [2, 3, 7, 8, 10], you first multiply every element by 2, then multiply every element by 3, then by 7, and so on.

My answer: $O(n^2)$

Book answer: $O(n^2)$