# Inherit VS Composition

## The inheritance delima

I am a zoo keeper, my goal is to classify animals.

A duck and a chicken comes to my zoo. They both have two wings and can fly.

Ok, we create a `Bird`(class) which have two wings and can fly, and it is perfect for both of them.

A penguins appeared, it has two wings but can't fly, so I create a `NoFlyBird`(class) and classify it.

A Kiwi appeared, it has two legs but doesn't have wings, so, is it not a bird? I confused.

A rubber duck appeared...

## What is the problem?

we try to find common characteristics and put them in a base class(Bird), then put chicken and Duck into it.

However, we never know what animal will appear next, and whether or not it will broken the previous class.

If such thing occurs, we have to refactor the Bird class, add more properties or methods, or break the class into pieces, and regroup them, and regroup them, and regroup them, again and again.

And how to name the class? we will spend a lot of time thinking about this.


## How Composition solves this?

Essentially, inheritance means 'is-a' relationship, while composition is a 'has-a/can-do' relationship.

For duck and chicken, we just give two wings and fly method.
For penguins, we only give two wings.
For Kiwi, we may give two legs.

> *"We don't care what they are, just what they can do or have."*
