# EXERCISES

Give the run time for each of these scenarios in terms of Big O.

1.3 You have a name, and you want to find the person’s phone number in the phone book.

My answer:  **$O(log_n)$**

Book answer: **$O(log_n)$**

1.4 You have a phone number, and you want to find the person’s name in the phone book. (Hint: You’ll have to search through the whole book!) 

My answer: **$O(n)$**

Book answer: **O(n)**

1.5 You want to read the numbers of every person in the phone book. 

My answer: **$O(n)$**

Book answer: **O(n)**

1.6 You want to read the numbers of just the As. (This is a tricky one! It involves concepts that are covered more in chapter 4. Read the answer—you may be surprised!)

My answer: **$O(n)$**

Book answer: **O(n)**
> You may think, “I’m only doing this for 1 out
of 26 characters, so the run time should be O(n/26).” A simple rule to remember is, ignore numbers that are added, subtracted, multiplied, or divided. None of these are correct Big O run times: O(n + 26), O(n - 26), O(n * 26), O(n / 26). They’re all the same as O(n)! Why? If you’re curious, flip to “Big O notation revisited,” in chapter 4, and read up on constants in Big O notation (a constant
is just a number; 26 was the constant in this question).