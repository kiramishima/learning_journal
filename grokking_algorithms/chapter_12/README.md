# EXERCISES
## 12.1
In the Netflix example, you calculated the distance between two different users using the distance formula. But not all users rate movies the same way. Suppose you have two users, Yogi and Pinky, who have the same taste in movies. But Yogi rates any movie he likes as a 5, whereas Pinky is choosier and reserves the 5s for only the best. They’re well matched, but according to the distance algorithm, they aren’t neighbors. How would you take their
different rating strategies into account?

![](ex_1.png)

My answer: Only consider the doesn't like and likes it or get the average by user.


Book answer:

```
Answer: You could use something called normalization. You
look at the average rating for each person and use it to scale their
ratings. For example, you might notice that Pinky’s average rating
is 3, whereas Yogi’s average rating is 3.5. So you bump up Pinky’s
ratings a little until her average rating is 3.5 as well. Then you can
compare their ratings on the same scale.
```

## 12.2
Suppose Netflix nominates a group of “influencers.” For example, Quentin Tarantino and Wes Anderson are influencers on Netflix, so their ratings count for more than a normal user’s. How would you change the recommendations system so it’s biased toward the ratings of influencers?

My answer:

```
Don´t consider the outliers in this case the influencers or normalize the value of influencers.
```

Book answer:

```
Answer: You could give more weight to the ratings of the
influencers when using KNN. Suppose you have three neighbors:
Joe, Dave, and Wes Anderson (an influencer). They rated
Caddyshack 3, 4, and 5, respectively. Instead of just taking the
average of their ratings (3 + 4 + 5 / 3 = 4 stars), you could give Wes
Anderson’s rating more weight: 3 + 4 + 5 + 5 + 5 / 5 = 4.4 stars.
```

## 12.3
Netflix has millions of users. The earlier example looked at the five closest neighbors for building the recommendations system. Is this too low? Too high?

My answer: Low

Book Answer: It’s too low. If you look at fewer neighbors, there’s a bigger
chance that the results will be skewed. A good rule of thumb is
that if you have N users, you should look at sqrt(N) neighbors.