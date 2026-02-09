package tui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sort"
	
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/abhinavdevarakonda/maplet/internal/graph"
)

var (
	selectedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("170")).Bold(true)
	normalStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	headerStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
)

type Model struct {
	graph       *graph.Graph 
	currentView string
	items       []TreeItem
	selected    int
	height      int
	itemToOpen  *TreeItem
}

type TreeItem struct {
	Name  string
	Path  string
	Line  int
	Depth int
	Type  string // file/function/method
}

func NewModel(g *graph.Graph) Model {
	items := buildStructureTree(g)
	return Model{
		graph:       g,
		currentView: "structure",
		items:       items,
		selected:    0,
		itemToOpen:  nil,
	}
}

func buildStructureTree(g *graph.Graph) []TreeItem {
	var items []TreeItem
	fileMap := make(map[string][]*graph.Node)
	
	for _, edge := range g.Edges {
		if edge.Type == graph.ContainsEdge {
			fromNode := g.Nodes[edge.From]
			toNode := g.Nodes[edge.To]
			
			if fromNode != nil && toNode != nil && 
			   fromNode.Type == graph.FileNode && 
			   toNode.Type == graph.FunctionNode {
				fileMap[edge.From] = append(fileMap[edge.From], toNode)
			}
		}
	}
	
	var fileNodes []*graph.Node
	for fileID := range fileMap {
		if node, exists := g.Nodes[fileID]; exists {
			fileNodes = append(fileNodes, node)
		}
	}
	
	sort.Slice(fileNodes, func(i, j int) bool {
		return fileNodes[i].Path < fileNodes[j].Path
	})
	
	for _, fileNode := range fileNodes {
		items = append(items, TreeItem{
			Name:  fileNode.Path,
			Path:  fileNode.Path,
			Line:  1,
			Depth: 0,
			Type:  "file",
		})
		
		functions := fileMap[fileNode.ID]
		
		sort.Slice(functions, func(i, j int) bool {
			return functions[i].Line < functions[j].Line
		})
		
		for _, fn := range functions {
			items = append(items, TreeItem{
				Name:  fn.Name,
				Path:  fileNode.Path,
				Line:  fn.Line, 
				Depth: 1,
				Type:  "function",
			})
		}
	}
	
	return items
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "j", "down":
			if m.selected < len(m.items)-1 {
				m.selected++
			}
		case "k", "up":
			if m.selected > 0 {
				m.selected--
			}
		case "h":
			m.currentView = "structure"
			m.items = buildStructureTree(m.graph)
			m.selected = 0
		case "l":
			m.currentView = "calls"
			// TODO: Build call tree here
			m.selected = 0
		case "enter":
			// Store the item to open and quit
			m.itemToOpen = &m.items[m.selected]
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
	}
	return m, nil
}

func (m Model) View() string {
	var s strings.Builder

	header := fmt.Sprintf("Maplet Navigator - %s view (h/l to switch, j/k to navigate, enter to open, q to quit)", m.currentView)
	s.WriteString(headerStyle.Render(header))
	s.WriteString("\n\n")

	start := 0
	end := len(m.items)
	if m.height > 0 && end > m.height-5 {
		if m.selected > m.height-7 {
			start = m.selected - m.height + 7
		}
		end = start + m.height - 5
		if end > len(m.items) {
			end = len(m.items)
		}
	}

	for i := start; i < end; i++ {
		item := m.items[i]
		indent := strings.Repeat("  ", item.Depth)
		
		var icon string
		switch item.Type {
		case "file":
			icon = "📄"
		case "function":
			icon = "ƒ"
		case "method":
			icon = "→"
		default:
			icon = "•"
		}

		var line string
		if item.Type == "function" && item.Line > 0 {
			line = fmt.Sprintf("%s%s %s (line %d)", indent, icon, item.Name, item.Line)
		} else {
			line = fmt.Sprintf("%s%s %s", indent, icon, item.Name)
		}
		
		if i == m.selected {
			s.WriteString(selectedStyle.Render("> " + line))
		} else {
			s.WriteString(normalStyle.Render("  " + line))
		}
		s.WriteString("\n")
	}

	return s.String()
}

func Start(g *graph.Graph) error {
	for {
		m := NewModel(g)
		p := tea.NewProgram(m, tea.WithAltScreen())
		
		finalModel, err := p.Run()
		if err != nil {
			return err
		}
		
		// Check if we need to open an editor
		if finalModel, ok := finalModel.(Model); ok {
			if finalModel.itemToOpen != nil {
				// Save terminal state
				sttyOutput, _ := exec.Command("stty", "-g").Output()
				
				// Open the editor
				editor := os.Getenv("EDITOR")
				if editor == "" {
					editor = "nvim"
				}
				
				var cmd *exec.Cmd
				if strings.Contains(editor, "vim") || strings.Contains(editor, "nvim") {
					cmd = exec.Command(editor, fmt.Sprintf("+%d", finalModel.itemToOpen.Line), finalModel.itemToOpen.Path)
				} else if strings.Contains(editor, "code") {
					cmd = exec.Command(editor, "-g", fmt.Sprintf("%s:%d", finalModel.itemToOpen.Path, finalModel.itemToOpen.Line))
				} else {
					cmd = exec.Command(editor, finalModel.itemToOpen.Path)
				}
				
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				
				if err := cmd.Run(); err != nil {
					fmt.Fprintf(os.Stderr, "Error running editor: %v\n", err)
				}
				
				// Restore terminal state
				if len(sttyOutput) > 0 {
					exec.Command("stty", string(sttyOutput)).Run()
				}
				
				// After editor closes, loop back to TUI
				continue
			}
			
			// If no item to open, user pressed 'q', so exit
			return nil
		}
		
		return nil
	}
}
