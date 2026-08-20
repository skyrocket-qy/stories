20240820
Software

# Google Wire vs. Uber Fx: What Time Do You Want to Panic?

**TL;DR:** Choose **Google Wire** unless you truly need Uber Fx's application lifecycle hooks. Your future self will thank you.

---

## ☕ The Morning Coffee Fable

Every morning in our office starts the same way: making a cup of pour-over coffee.

To make coffee, you need a few things working together:
1. Coffee beans
2. A grinder (which needs the beans)
3. Hot water
4. A dripper (which takes the ground coffee and the hot water)

In software, passing the beans to the grinder and the hot water to the dripper is all that **"Dependency Injection" (DI)** really is:

```go
beans := NewBeans()
grinder := NewGrinder(beans)
water := NewHotWater()
coffee := NewCoffee(grinder, water)
```

It sounds simple enough. But as our office grew, we added milk frothers, syrup pumps, timers, and water filters. Manually wiring all these parts together by hand became tedious.

So, the team decided to automate the coffee-making process. That’s when the argument began.

---

## 🦉 The Owl’s Solution: The Blindfolded Robotic Chef (Uber Fx)

The Owl—our Chief Architect—perched on top of a stack of software architecture books, adjusted its gold wire-rimmed spectacles, and announced:

> *"According to clean inversion-of-control paradigms, we should not manually connect the water to the dripper. We need a dynamic runtime container that discovers dependencies on the fly."*

The Owl installed a high-tech robotic chef (**Uber Fx**) in the kitchen. 

Here is how the robotic chef works:
1. You press the **"Brew Coffee"** button.
2. The robot wakes up, puts on a blindfold, and begins scanning every cabinet in the kitchen using an infrared camera (Go's `reflect` package).
3. It reads the labels on every jar, builds an internal mental map of all ingredients, and figures out how to connect the grinder to the dripper.
4. If everything is in place, it brews the coffee.

It looked magnificent on the Owl's architecture diagram. 

Until Monday morning.

The Chihuahua ran into the kitchen at 8:55 AM, vibrating from pure adrenaline. *"I added cinnamon to the coffee!"* it barked, slamming a new ingredient onto the counter, bypassing the recipe review, and hitting the big green **"Brew"** button.

The code compiled fine. The binary started running. But halfway through the morning rush, the robotic chef reached into the cabinet, found that someone forgot to connect the cinnamon shaker, panicked, dropped the boiling kettle onto the floor, and shouted:

```text
panic: could not build arguments for function "main".NewCoffee: 
  missing dependency: *kitchen.Cinnamon is not provided in container
```

The kitchen was covered in boiling water. The morning was ruined.

---

## 🐼 The Panda’s Solution: The Printed Recipe Card (Google Wire)

I slapped a cold fever gel patch across my forehead, poured myself a cup of plain water from the tap, and looked at the mess. 

*"Stop touching things,"* I sighed. *"Let me show you a better way."*

I introduced **Google Wire**.

Instead of hiring a blindfolded robot to search cabinets while the stove is already on fire, Google Wire acts like a **pre-flight recipe checker**:

1. Before you ever turn on the stove, Wire looks at your ingredients list while you are sitting at your desk (`go generate` / `wire`).
2. It traces the connections ahead of time: *Beans go to Grinder, Grinder goes to Dripper, Water goes to Dripper.*
3. It writes down a plain, simple recipe card in clean Go code (`wire_gen.go`).
4. **Crucially:** If the Chihuahua forgets to supply the cinnamon, Wire stops you right there on your laptop:
   ```text
   wire: no provider found for *kitchen.Cinnamon
   ```
   **The program refuses to even compile.**

You catch the mistake immediately while writing code in broad daylight, instead of having your application explode in front of real users.

---

## ⚡ The Core Question: *When* Does the Computation Happen?

At their core, both tools do the same job—they wire your structs together. But they do the math at completely different times:

* **Uber Fx (Runtime):** Waits until the program starts up, then uses runtime reflection to figure out what fits where. If something is missing, it crashes at runtime.
* **Google Wire (Compile-time):** Figures out what fits where ahead of time, generates plain standard Go code, and catches missing parts at compile time.

---

## 📊 The Hard Proof: Benchmarking Wire vs. Fx

Is there a real performance cost to having a robot scan your kitchen with reflection every time?

Let's test it with a simple Go benchmark (available in [`benchmarks/fx-vs-wire`](../benchmarks/fx-vs-wire)):

```go
type Service interface {
    Do()
}

type MyService struct{}

func (s *MyService) Do() {}

func NewService() Service {
    return &MyService{}
}
```

We compared:
1. **Direct:** Standard manual Go initialization (`NewService()`).
2. **Google Wire:** Static generated code (`InitializeService()`).
3. **Uber Fx:** Dynamic reflection container (`fx.New(...)`).

Here are the reproducible benchmark results:

```text
BenchmarkWire-10         	1000000000	      0.2271 ns/op	       0 B/op	       0 allocs/op
BenchmarkDirect-10       	1000000000	      0.2254 ns/op	       0 B/op	       0 allocs/op
BenchmarkFx-10           	     39981	   30595.0000 ns/op	   35621 B/op	     511 allocs/op
BenchmarkFxNewOnly-10    	     43504	   27659.0000 ns/op	   35355 B/op	     499 allocs/op
```

### What does this data prove?

1. **Wire is 135,000× faster with Zero Memory Allocations:**  
   Wire runs in **`0.22 nanoseconds`** with **`0 heap allocations`**. Because Wire generated plain Go code ahead of time, the Go compiler can inline the entire constructor into a single CPU instruction—identical to writing manual code.

2. **The Reflection Tax:**  
   Uber Fx takes **`~30 microseconds`** and consumes **`511 heap allocations`** (~35.6 KB) just to create a single trivial object. Over 90% of that time is spent purely inside `fx.New` doing reflection lookups.

While 30 microseconds might not sound huge for a single server startup, it becomes massive in:
* **Serverless / Lambda Functions:** Where cold-start latency directly slows down user requests.
* **CLI Tools:** Where users expect sub-second responsiveness.
* **Unit Tests:** When running hundreds of test suites that each spin up application contexts.

---

## 🥊 Feature-by-Feature Breakdown

### 1. How It Works
* **🦉 Uber Fx:** Uses reflection at runtime to inspect constructors and build an in-memory graph.
* **🐼 Google Wire:** Analyzes dependencies ahead of time and generates plain Go code (`wire_gen.go`).

### 2. Missing Dependency Error
* **🦉 Uber Fx:** 💥 Application panics and crashes on startup in production.
* **🐼 Google Wire:** 🛡️ `go build` fails immediately on your laptop before you can even run it.

### 3. Speed & Memory
* **🦉 Uber Fx:** ~30 µs startup latency and 500+ heap allocations per container.
* **🐼 Google Wire:** **~0.22 ns and 0 heap allocations** (fully inlined by the compiler).

### 4. Debugging & Visibility
* **🦉 Uber Fx:** Buried inside dynamic reflection stack traces.
* **🐼 Google Wire:** Plain, readable Go code you can step through line-by-line with standard debuggers.

### 5. Lifecycle Management
* **🦉 Uber Fx:** Built-in application lifecycle hooks (`OnStart`, `OnStop`, graceful shutdown).
* **🐼 Google Wire:** Manual (relies on standard Go cleanup functions and closures).

---

## When Does Uber Fx Make Sense?

To be fair to the Owl: Uber Fx is not just a DI tool; it is an **application lifecycle framework**.

If you are building a huge application with dozens of modules that need coordinated startup and shutdown sequences (e.g. *"Stop HTTP server -> flush message queues -> close database pool"*), Fx’s built-in `OnStart`/`OnStop` hooks are very convenient.

However, if your primary goal is simply clean Dependency Injection without magic, Fx adds runtime risks and reflection overhead that you don't need.

---

## 🐼 The Panda's Verdict

As the maintainer who has to keep the lights on:

1. **Explicit is better than magical.** Go's philosophy has always favored clear, readable, compile-time safety over runtime reflection magic.
2. **A compile error is a blessing.** Every mistake caught by `go build` while writing code is an outage that never happened.
3. **Keep it simple.** Use **Google Wire**. Generate your wiring, review the generated Go code, and enjoy a quiet, uninterrupted cup of coffee.

---

*Reproducible benchmark code is available in this repository under [`benchmarks/fx-vs-wire/`](../benchmarks/fx-vs-wire).*
