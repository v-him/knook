package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/v-him/knook/internal/api"
	"github.com/v-him/knook/internal/cmd/config"
)

type StatusOptions struct{
	Quiet bool
}

type errMsg error

type userInfo struct {
	username string
	email string
	token string
}

type statusModel struct {
	info userInfo
	spinner spinner.Model
	apiDone bool
	err error
}

func initialStatusModel() statusModel {
	s := spinner.New()
	s.Spinner = spinner.Ellipsis
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("69"))
	return statusModel{ spinner: s }
}

func (m statusModel) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, fetchUserInfo)
}

func (m statusModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		default:
			return m, nil
		}

	case errMsg:
		m.err = msg
		return m, nil

	case *userInfo:
		m.info = *msg
		m.apiDone = true
		return m, tea.Quit

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}
}

func (m statusModel) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	checkmark := lipgloss.NewStyle().SetString("✓").Foreground(lipgloss.Color("142")).Bold(true)
	username := lipgloss.NewStyle().SetString(m.info.username).Foreground(lipgloss.Color("108")).Bold(true)
	email := lipgloss.NewStyle().SetString(m.info.email).Foreground(lipgloss.Color("175")).Bold(true)

	var hiddenToken string

	if len(m.info.token) >= 4 && m.info.token[:4] == "lip_" {
		hiddenToken = m.info.token[:4] + strings.Repeat("*", len(m.info.token) - 4 )
	} else {
		hiddenToken = strings.Repeat("*", len(m.info.token))
	}

	token := lipgloss.NewStyle().SetString(hiddenToken).Foreground(lipgloss.Color("243")).Bold(true)

	var str string

	if m.apiDone {
		str = fmt.Sprintf("   %s Logged in to lichess.org account %s\n", checkmark, username)
		str += fmt.Sprintf("   - Email: %s\n", email)
		str += fmt.Sprintf("   - Token: %s\n", token)
	} else {
		str = fmt.Sprintf("\n   %s Loading\n\n", m.spinner.View())
	}

	return str
}


func fetchUserInfo() tea.Msg {
	client := &http.Client{}
	cfg, err := config.Read()
	if err != nil {
		return errMsg(err)
	}

	profile, err := api.GetProfile(cfg.Token, client)
	if err != nil {
		return errMsg(err)
	}

	if profile.Username == "" {
		return errMsg(errors.New("Token is valid but could not determine your username"))
	}

	email, err := api.GetEmail(cfg.Token, client)
	if err != nil {
		return errMsg(err)
	}

	if email == "" {
		return errMsg(errors.New("Token is valid but could not determine your email"))
	}
	return &userInfo{
		username: profile.Username,
		email: email,
		token: cfg.Token,
	}
}

func Status(opts StatusOptions) {
	p := tea.NewProgram(initialStatusModel())
	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
