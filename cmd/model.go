package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFF"))
	cursorStyle  = focusedStyle
	noStyle      = lipgloss.NewStyle()
)

type CmdModel struct {
	focusIndex int
	inputs     []cmdModelInput
	cursorMode cursor.Mode
}

type cmdModelInput struct {
	label string
	input textinput.Model
}

func generateModel(flags *PersistentFlags) CmdModel {
	m := CmdModel{
		inputs: []cmdModelInput{},
	}

	t := textinput.New()
	t.Placeholder = "Please insert your Github Personal Access Token"
	t.CharLimit = 0
	t.Cursor.Style = cursorStyle
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle
	t.Width = 100
	t.SetValue(flags.Token)
	t.Focus()
	m.inputs = append(m.inputs, cmdModelInput{
		label: "GitHub Personal Access Token",
		input: t,
	})

	t = textinput.New()
	t.Placeholder = "Please insert the output file path"
	t.CharLimit = 0
	t.Cursor.Style = cursorStyle
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle
	t.Width = 100
	t.SetValue(flags.OutputFile)
	m.inputs = append(m.inputs, cmdModelInput{
		label: "Output File Path",
		input: t,
	})

	t = textinput.New()
	t.Placeholder = "Please insert the users to backup (comma separated, empty if not needed)"
	t.CharLimit = 0
	t.Cursor.Style = cursorStyle
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle
	t.Width = 100
	t.SetValue(strings.Join(flags.Users, ","))
	m.inputs = append(m.inputs, cmdModelInput{
		label: "Users to Backup",
		input: t,
	})

	t = textinput.New()
	t.Placeholder = "Please insert the orgs to backup (comma separated, empty if not neeeded)"
	t.CharLimit = 0
	t.Cursor.Style = cursorStyle
	t.PromptStyle = focusedStyle
	t.TextStyle = focusedStyle
	t.Width = 100
	t.SetValue(strings.Join(flags.Orgs, ","))
	m.inputs = append(m.inputs, cmdModelInput{
		label: "Orgs to Backup",
		input: t,
	})

	return m
}

func (m CmdModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CmdModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > cursor.CursorHide {
				m.cursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].input.Cursor.SetMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs)-1 {
				return m, tea.Quit
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].input.Focus()
					m.inputs[i].input.PromptStyle = focusedStyle
					m.inputs[i].input.TextStyle = focusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].input.Blur()
				m.inputs[i].input.PromptStyle = noStyle
				m.inputs[i].input.TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m *CmdModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i].input, cmds[i] = m.inputs[i].input.Update(msg)
	}

	return tea.Batch(cmds...)
}

func (m CmdModel) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(fmt.Sprintf("%s\n", m.inputs[i].label))
		b.WriteString(m.inputs[i].input.View())
		b.WriteRune('\n')
		b.WriteRune('\n')
	}

	b.WriteString("(esc to quit)")
	b.WriteRune('\n')
	b.WriteRune('\n')

	return b.String()
}
