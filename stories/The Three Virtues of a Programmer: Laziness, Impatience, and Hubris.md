# The Three Virtues of a Programmer: Laziness, Impatience, and Hubris

![The Three Virtues: Impatience Chihuahua, Hubris Owl, and Lazy Panda](../../assets/three_virtues_trio.jpg)

## 🎯 The Task: Processing a Batch of User Events

### 0. The Naïve Starting Code

```go
func ProcessEvents(events []string) []string {
    var out []string
    for _, e := range events {
        out = append(out, "processed: "+e)
    }
    return out
}
```

---

### 1. 🐕 The Chihuahua’s Impatience: *"Make it fast! I hate waiting!"*

The Chihuahua sees dynamic slice resizing.

* **Virtue (Impatience):** Refuses to tolerate latency. Pre-allocates slice capacity to make the hot path run instantly.

```go
// ⚡ Impatience: Pre-allocate capacity -> Zero reallocation overhead
func ProcessEvents(events []string) []string {
    out := make([]string, 0, len(events))
    for _, e := range events {
        out = append(out, "processed: "+e)
    }
    return out
}
```

---

### 2. 🦉 The Owl’s Hubris: *"Make it bulletproof! I take pride in my code!"*

The Owl sees missing validation.

* **Virtue (Hubris):** Demands flawless craftsmanship.

```go
// 🛡️ Hubris: Strict correctness & zero hidden panics
func ProcessEvents(events []string) ([]string, error) {
    out := make([]string, 0, len(events))
    for _, e := range events {
        if e == "" {
            return nil, errors.New("event payload cannot be empty")
        }
        out = append(out, "processed: "+e)
    }
    return out, nil
}
```

---

### 3. 🐼 The Panda’s Laziness: *"Make it run itself! I want to sleep."*

The Panda realizes: *"Wait... why are we manually calling this function every time? If an error happens, we'll get paged at 3 AM."*

* **Virtue (Laziness):** Refuses repetitive manual toil. Wraps the logic into an **automated, self-healing background worker** that consumes events from a channel and handles errors on its own.

```go
// 😴 Laziness: Self-driving background worker -> Zero manual toil, zero 3 AM pages
func StartEventWorker(ctx context.Context, in <-chan []string) {
    go func() {
        for {
            select {
            case <-ctx.Done():
                return
            case batch := <-in:
                if _, err := ProcessEvents(ctx, batch); err != nil {
                    log.Printf("[Auto-Heal] Dropped bad batch: %v", err)
                }
            }
        }
    }()
}
```

---

## 🧩 How the 3 Virtues Combine

```go
func main() {
    eventQueue := make(chan []string, 100)

    // 🐼 Panda (Laziness): Runs by itself in the background
    StartEventWorker(eventQueue)

    // 🐕 Chihuahua (Impatience) & 🦉 Owl (Hubris):
    // Lightning fast, pre-allocated, zero panics!
    eventQueue <- []string{"login", "purchase", "logout"}
}
```

---

## 📊 Summary: The Three Virtues in Practice

> 1. **Be Impatient:** Always benchmark and optimize the hot path.  
> 2. **Have Hubris:** Write code so robust, typed, and error-safe that you are proud to put your name on it.  
> 3. **Be Lazy:** Don't repeat yourself.
