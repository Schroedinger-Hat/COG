package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	utils "cog/utils"
)

type model struct {
	issues   []utils.Issue
	index    int
	width    int
	height   int
	spinner  spinner.Model
	progress progress.Model
	done     bool
}

var csvFilePath string

var (
	currentIssueNameStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("211"))
	doneStyle             = lipgloss.NewStyle().Margin(1, 2)
	checkMark             = lipgloss.NewStyle().Foreground(lipgloss.Color("42")).SetString("✓")
)

func newModel() model {
	p := progress.New(
		progress.WithDefaultGradient(),
		progress.WithWidth(40),
		progress.WithoutPercentage(),
	)
	s := spinner.New()
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
	return model{
		issues:   utils.GetIssue(csvFilePath),
		spinner:  s,
		progress: p,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Batch(parseAndCreate(m.issues[m.index]), m.spinner.Tick)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q":
			return m, tea.Quit
		}
	case createdIssueMsg:
		if m.index >= len(m.issues)-1 {
			m.done = true
			return m, tea.Quit
		}

		progressCmd := m.progress.SetPercent(float64(m.index) / float64(len(m.issues)-1))

		m.index++
		return m, tea.Batch(
			progressCmd,
			tea.Printf("%s %s", checkMark, m.issues[m.index]), // print success message above our program
			parseAndCreate(m.issues[m.index]),                 // create the next issue
		)
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	case progress.FrameMsg:
		newModel, cmd := m.progress.Update(msg)
		if newModel, ok := newModel.(progress.Model); ok {
			m.progress = newModel
		}
		return m, cmd
	}
	return m, nil
}

func (m model) View() string {
	n := len(m.issues) - 1
	w := lipgloss.Width(fmt.Sprintf("%d", n))

	if m.done {
		return doneStyle.Render(fmt.Sprintf("Done! Created %d Issue.\n", n))
	}

	issuesCount := fmt.Sprintf(" %*d/%*d", w, m.index, w, n-1)

	spin := m.spinner.View() + " "
	prog := m.progress.View()
	cellsAvail := max(0, m.width-lipgloss.Width(spin+prog+issuesCount))

	issueName := currentIssueNameStyle.Render(m.issues[m.index].Name)
	info := lipgloss.NewStyle().MaxWidth(cellsAvail).Render("Creating " + issueName)

	cellsRemaining := max(0, m.width-lipgloss.Width(spin+info+prog+issuesCount))
	gap := strings.Repeat(" ", cellsRemaining)

	return spin + info + gap + prog + issuesCount
}

type createdIssueMsg string

func parseAndCreate(issue utils.Issue) tea.Cmd {
	// change with the creating of the issue
	d := time.Millisecond * time.Duration(rand.Intn(500)) //nolint:gosec
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return createdIssueMsg(issue.Name + ":" + issue.Description)
	})
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	flag.StringVar(&csvFilePath, "csv", "", "The path of the csv with all the infos about issues to create")
	flag.Parse()
	if csvFilePath == "" {
		fmt.Println("Error: you must pass a csv file!")
		os.Exit(0)
	}

	if _, err := tea.NewProgram(newModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
