# Week 11: Bubbletea (TUI) Architecture

## 1. Declarative vs Imperative UI
- **Imperative (Old):** "Print line 1", "Print line 2". This flickers and is hard to update.
- **Declarative (Bubbletea):** "Here is the Model (State). Re-draw everything when it changes."

## 2. The MVU (Model-View-Update) Pattern
1.  **Model:** The Data (`Count`, `Duration`, `Stats`).
2.  **View:** The HTML/String (`fmt.Sprintf("Count: %d", m.Count)`).
3.  **Update:** The Logic (`case msg: m.Count++`).

## 3. Channels and The "Pipe"
We connect our Worker Pool (Imperative Logic) to the UI (Declarative) using a `chan Stats`.
- **Runner:** Slides stats down the pipe (`chan<-`).
- **UI:** Picks them up (`<-chan`) using `waitForStats` command.
- **Safety:** Using `chan<-` ensures the Runner can't steal messages from the UI.

## 4. The "Ghost" Bug (Double Rendering)
If you print to `stdout` (`fmt.Println`) while Bubbletea is running, the text will glitch.
**Solution:** Remove all `fmt.Println` from the runner logic. Let the UI handle 100% of the screen.
