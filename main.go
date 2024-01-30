package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

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
var githubUser string
var githubRepository string
var GHToken string

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
	url := "https://api.github.com/repos/" + githubUser + "/" + githubRepository + "/issues"

	title := "\"title\":\"" + issue.Name + "\""
	description := "\"body\":\"" + issue.Description + "\""
	labels := "\"labels\":["
	for index, label := range issue.Labels {
		labels += "\"" + label + "\""
		if index < len(issue.Labels)-1 {
			labels += ","
		}
	}
	labels += "]"
	payload := []byte("{" + title + "," + description + "," + labels + "}")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("Error creating request:", err)
		os.Exit(0)
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("Authorization", "Bearer "+GHToken)
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(0)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		return func() tea.Msg {
			return createdIssueMsg(issue.Name + ":" + issue.Description)
		}
	} else {
		return nil
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	envGHTokenName := "GHTOKEN"
	envGHTokenValue, exists := os.LookupEnv(envGHTokenName)
	if !exists {
		fmt.Println("Environment variable " + envGHTokenName + " not set!")
	}
	GHToken = envGHTokenValue

	flag.StringVar(&csvFilePath, "csv", "", "The path of the csv with all the infos about issues to create")
	flag.StringVar(&githubUser, "gh-user", "", "The user of the repository where we want to create the issues")
	flag.StringVar(&githubRepository, "gh-repository", "", "The repository where we want to create the issues")
	flag.Parse()
	if csvFilePath == "" || githubRepository == "" || githubUser == "" {
		fmt.Println("Error: you must pass all the arguments!")
		os.Exit(0)
	}

	if _, err := tea.NewProgram(newModel()).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
