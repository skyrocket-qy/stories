# Because composition is superior to inheritance

We have a duck and a chicken, and they work well. We noticed that ducks and chickens share many similarities—they both fly, they each have two legs and two wings, and they make sounds like "quack" or "cluck." So, we abstracted a parent class called Bird, from which both Duck and Chicken inherited. Then, when we wanted to add extra behavior for flying, we only needed to change the Bird class; the code was reduced and seemed easier to maintain. The Bird class worked great.

But as our business expanded, penguins appeared. Penguins can’t fly, though they can swim. Suddenly, the design of Bird ran into problems, so we “pushed down” the flying functionality into a subclass for birds that do fly, with penguins inheriting from a non-flying bird class. Next came the rubber duck—the debate sparked: Is it a bird or not? People wasted a lot of time arguing about what exactly a bird is and what a bird should do.

However, there are no birds in our everyday life (please note this important point); Bird is just an abstraction. In real life, we have chickens and ducks. We thought they shared some common characteristics, so we labeled those commonalities as “bird,” but we can never be sure if the next animal we encounter should be considered a bird or whether the definition of bird should be changed. This is why inheritance becomes problematic. Let’s consider what composition would look like:

We identified the common factors of chickens and ducks—flying, two legs, two wings, and the ability to make noise. These properties, plus other characteristics, can be composed to form a chicken or a duck. The flying capability can be extracted and used everywhere flying is required. When I encounter a penguin, I don’t compose it with flying. The ability to fly shouldn’t be something inherited by chickens or ducks from a parent class; instead, it should be one of the composable parts that make up a chicken or a duck.

This issue arises from poor abstraction skills and falling into one’s own logical trap. Bird is, by nature, a symbol representing a combination of specific functions. So why force a penguin, which doesn’t fly, to inherit from Bird? That isn’t the “original sin” of inheritance—in fact, the JDK code doesn’t have that same flaw. It’s just that as projects evolve, inheritance needs timely refactoring, which is very hard to achieve, leading to its eventual abuse and a result that is increasingly unrecognizable.

In a sense, inheritance is the original sin—it reverses the relationship between things and the characteristics they share. Instead of one thing deriving from the form of another, it should be that the common characteristics of things are summarized into an abstract form. Inheritance then simply makes things inherit from a base class (an abstract form), which is a reversal of the natural order.

The principle of inheritance is: only when something truly is a particular thing should it inherit from that thing (the is-a relationship). Is a penguin a bird? Is a square a rectangle? These are classic debates, yet in actual development, you waste a lot of time trying to decide whether something really is an is-a relationship, often making wrong assumptions.

Where do all these reasons come from? Even if composition is indeed superior to inheritance, why not provide both composition and inheritance, letting programmers choose for themselves instead of having the language designer decide based on personal taste?

In reality, language designers are often lazy—they don’t want to provide too many features or let programmers decide. For example, Java’s implicit message is: “You lot won’t master pointers, so I won’t provide them (even though Go does).” Similarly, operator overloading is left out because, in their view, you won’t learn to use it properly; and if you can’t handle multiple inheritance, then stick with single inheritance, and so on.

Rust takes a different approach: since you’re likely to make memory leaks, it enforces strict lifetimes to pin you down; if you can’t manage inheritance properly (something Java still allows), then Rust simply forbids it; and when it comes to multithreading, since programmers are prone to mistakes, it enforces strict lifetimes to prevent common errors.

C/C++, on the other hand, goes to the opposite extreme: “You’re all experienced programmers, so I’ll give you every feature you ask for. The more functionality, the better. I’ll provide every concept imaginable—if you run into problems composing those concepts, that’s on you.” Even syntactic sugar has increased over time (for example, C++17’s structured bindings are a clear example).

In real-world projects, I’ve seen many developers implement a “dummy” flying function for penguins just to get the code running, and then they tell testers, “Well, penguins can’t fly; even if we let them fly, nothing will happen.”

Complexity never disappears—it merely shifts. With inheritance, you might end up with a base class containing millions of interfaces. With composition, you only implement the minimal interfaces you need; and if reusability is possible, you don’t even need to implement them, nor do you have to worry about issues like multiple inheritance or the diamond problem.

The problem with object-oriented programming languages is that they always come with an implicit context. You ask for a banana, but what you get is a gorilla holding the banana—and the entire jungle along with it.
— Joe Armstrong (creator of Erlang)


Here is an honest review:

🎯 The Verdict: Great Core Idea, but Too Long & Distracted
1. What Works Great (The First 4 Paragraphs) 🌟
The Duck & Penguin Dilemma: The realization that "there are no 'birds' in everyday life—'Bird' is just an abstraction" is a brilliant, intuitive insight.
The Rubber Duck argument is instantly recognizable to anyone who has argued over OOP hierarchies.
2. Why It Feels Too Long and Distracted ⚠️
It Digresses into Language Rants (Paragraphs 6–10):
Talking about Java pointers, C++ operator overloading, and Rust lifetimes distracts from the main topic (Inheritance vs. Composition). It reads like an essay on programming languages rather than a focused story.
No Character Roles (Missing the Owl, Chihuahua, and Panda):
It lacks the fun dynamic you have in The Owl's Factory.
No Code Snippets:
It explains the idea purely in prose, without showing the 5-line Go code comparison that makes readers say: "Aha! That's how it looks in real life."
💡 How to Make It Punchy & Fun (Like the Table Story)
Keep the Duck / Penguin / Rubber Duck idea, but structure it like this:

The Owl’s Grand Hierarchy:
The Owl builds class Bird { Fly() } because ducks and chickens fly.
The Chihuahua's Panic:
The Chihuahua needs to ship a Penguin and a Rubber Duck.
The Penguin inherits .Fly(), crashes into the ocean, and the Chihuahua adds a dummy Fly() { panic("can't fly") }.
The Panda's Clean Composition:
Break it down: Flying is just a capability, swimming is a capability. Compose what you need.
The Punchline:
End with Joe Armstrong's classic quote:
"You wanted a banana, but what you got was a gorilla holding the banana—and the entire jungle."

Would you like to trim down the tangents and rewrite this into a short, character-driven story with code?