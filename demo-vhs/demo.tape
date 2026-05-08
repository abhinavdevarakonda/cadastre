Output cadr-demo.gif

Set Shell "zsh"
Set FontSize 16
Set FontFamily "JetBrainsMono Nerd Font"
Set Width 1400
Set Height 850
Set Padding 30
Set Theme "Jellybeans"
Set TypingSpeed 60ms

# Setup clean tmux environment (Hiding setup)
Hide
  Type "tmux -f /dev/null -L test new-session -- bash" Enter
  Type "tmux split-window -d -h -- bash && \" Enter
  Type "tmux set status && \" Enter
  Type 'tmux setw pane-border-style "fg=0" && \' Enter
  Type 'tmux setw pane-active-border-style "fg=0"' Enter
  Sleep 0.5
  
  # Prepare directories in both panes
  Type "cd /home/abee/code/projects/demo/cafe-test" Enter
  Ctrl+B
  Type "o"
  Type "cd /home/abee/code/projects/demo/cafe-test" Enter
  Type "clear" Enter
  
  # Back to Left Pane
  Ctrl+B
  Type "o"
  
  Ctrl+L
  Sleep 1
Show

# --- PANE 1: THE BRAIN (Start TUI first) ---
Type "cadr" 
Sleep 1
Enter
Sleep 2
Type "t"
Sleep 0.5
Type "t"

# --- PANE 2: THE RUNTIME ---
Ctrl+B
Type "o"
Sleep 0.5
Type "cadr run 'python main.py'" 
Sleep 1
Enter
Sleep 4

# --- SWITCH BACK TO TUI ---
Ctrl+B
Type "o"
Sleep 2

# --- INTERACTION ---
# Watch heatmap
Sleep 4

# Navigate
Type "i"
Space
Sleep 1

Type "k"
Sleep 0.3
Type "k"
Sleep 0.3
Type "k"
Sleep 3

# Open Source
Enter
Sleep 0.8
Type "V"
Sleep 3
Escape
Sleep 1

# Exit Editor (nvim)
Type ":q" Enter
Sleep 1

# Resume
Type "j"
Sleep 0.3
Type "j"
Sleep 0.3
Type "j"
Sleep 0.3
Type "j"

Sleep 2
Type "t"
Sleep 2
Enter

# Next file (to show impact)
Sleep 0.8
Type "V"
Sleep 3
Escape
Sleep 0.5
Type ":q" Enter


# Cleanup
Hide
  Ctrl+B
  Type "o"
  Ctrl+C
  Type "exit" Enter
  Sleep 1
  Type "exit" Enter
Show

Sleep 2
