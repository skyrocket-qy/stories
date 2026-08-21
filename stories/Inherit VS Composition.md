# Inherit vs. Composition

## 🦆 The Grand Zoo Taxonomy

The Panda is put in charge of the city zoo's animal registry.

🦆 Duck & 🐔 Chicken arrive. They both have wings and can fly.

🐼: Create a Bird class (wings, fly). Perfect for both.

🐧 Penguin arrives, it has wings, but can't fly.

🐼: Create a NoFlyBird class to classify it.

🥝 Kiwi arrives, it has legs, but no wings.

🐼: Is it not a bird? I am confused.

🐤 Rubber Duck arrives...

## What is the problem?

we try to find common characteristics and put them in a base class(Bird), then put chicken and Duck into it.

However, we never know what animal will appear next, and break the existing hierarchy.

If such thing occurs, we have to refactor the Bird class, add more properties or methods, or break the class into pieces, and regroup them, and regroup them, again and again.

And how to name the class? we will spend a lot of time thinking about this.


## How Composition solves this?

Inheritance forces an 'is-a' relationship, while composition is a 'has-a/can-do' relationship.

For duck and chicken, we just give two wings and fly method.
For penguins, we only give two wings.
For Kiwi, we may give two legs.

> *"We don't care what they are, just what they can do or have."*
