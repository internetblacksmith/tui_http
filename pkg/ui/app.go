package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"restless/pkg/http"
	"restless/pkg/models"
)

type tabType int

const (
	requestTab tabType = iota
	responseTab
	historyTab
)

type Model struct {
	activeTab tabType
	width     int
	height    int

	methodInput      textinput.Model
	urlInput         textinput.Model
	bodyInput        textinput.Model
	headerKeyInput   textinput.Model
	headerValueInput textinput.Model

	request  *models.Request
	response *models.Response
	client   *http.Client
	history  []*models.Request

	isLoading bool
	error     string

	currentMethod  int
	methods        []models.HTTPMethod
	headers        []models.Header
	selectedHeader int

	focusedInput int
}

type requestCompleteMsg struct {
	response *models.Response
	error    error
}

func NewModel() *Model {
	methods := []models.HTTPMethod{
		models.GET, models.POST, models.PUT, models.DELETE,
		models.PATCH, models.HEAD,
	}

	methodInput := textinput.New()
	methodInput.Placeholder = "GET"
	methodInput.Focus()
	methodInput.CharLimit = 156
	methodInput.Width = 20

	urlInput := textinput.New()
	urlInput.Placeholder = "https://jsonplaceholder.typicode.com/posts/1"
	urlInput.CharLimit = 200
	urlInput.Width = 60

	bodyInput := textinput.New()
	bodyInput.Placeholder = "Request body (JSON, XML, etc.)"
	bodyInput.CharLimit = 1000
	bodyInput.Width = 60

	headerKeyInput := textinput.New()
	headerKeyInput.Placeholder = "Header Key"
	headerKeyInput.Width = 25

	headerValueInput := textinput.New()
	headerValueInput.Placeholder = "Header Value"
	headerValueInput.Width = 35

	return &Model{
		activeTab:        requestTab,
		methodInput:      methodInput,
		urlInput:         urlInput,
		bodyInput:        bodyInput,
		headerKeyInput:   headerKeyInput,
		headerValueInput: headerValueInput,
		request:          &models.Request{},
		client:           http.NewClient(),
		methods:          methods,
		headers:          []models.Header{},
		currentMethod:    0,
		focusedInput:     0,
	}
}

func (m *Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "tab":
			switch m.activeTab {
			case requestTab:
				m.activeTab = responseTab
			case responseTab:
				m.activeTab = historyTab
			case historyTab:
				m.activeTab = requestTab
			}
			return m, nil

		case "enter":
			if m.activeTab == requestTab && !m.isLoading {
				if m.focusedInput == 3 || m.focusedInput == 4 {
					m.addHeader()
				} else {
					return m, m.executeRequest()
				}
			}

		case "up":
			if m.activeTab == requestTab {
				if m.focusedInput > 0 {
					m.focusedInput--
					m.updateInputFocus()
				}
			}

		case "down":
			if m.activeTab == requestTab {
				if m.focusedInput < 4 {
					m.focusedInput++
					m.updateInputFocus()
				}
			}

		case "left":
			if m.activeTab == requestTab && m.focusedInput == 0 {
				if m.currentMethod > 0 {
					m.currentMethod--
					m.methodInput.SetValue(string(m.methods[m.currentMethod]))
				}
			}

		case "right":
			if m.activeTab == requestTab && m.focusedInput == 0 {
				if m.currentMethod < len(m.methods)-1 {
					m.currentMethod++
					m.methodInput.SetValue(string(m.methods[m.currentMethod]))
				}
			}
		}

	case requestCompleteMsg:
		m.isLoading = false
		if msg.error != nil {
			m.error = msg.error.Error()
		} else {
			m.response = msg.response
			m.error = ""
			m.activeTab = responseTab
		}
		return m, nil
	}

	switch m.focusedInput {
	case 0:
		m.methodInput, cmd = m.methodInput.Update(msg)
	case 1:
		m.urlInput, cmd = m.urlInput.Update(msg)
	case 2:
		m.bodyInput, cmd = m.bodyInput.Update(msg)
	case 3:
		m.headerKeyInput, cmd = m.headerKeyInput.Update(msg)
	case 4:
		m.headerValueInput, cmd = m.headerValueInput.Update(msg)
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m *Model) updateInputFocus() {
	m.methodInput.Blur()
	m.urlInput.Blur()
	m.bodyInput.Blur()
	m.headerKeyInput.Blur()
	m.headerValueInput.Blur()

	switch m.focusedInput {
	case 0:
		m.methodInput.Focus()
	case 1:
		m.urlInput.Focus()
	case 2:
		m.bodyInput.Focus()
	case 3:
		m.headerKeyInput.Focus()
	case 4:
		m.headerValueInput.Focus()
	}
}

func (m *Model) addHeader() {
	if m.headerKeyInput.Value() != "" && m.headerValueInput.Value() != "" {
		m.headers = append(m.headers, models.Header{
			Key:   m.headerKeyInput.Value(),
			Value: m.headerValueInput.Value(),
		})
		m.headerKeyInput.SetValue("")
		m.headerValueInput.SetValue("")
	}
}

func (m *Model) executeRequest() tea.Cmd {
	method := models.HTTPMethod(m.methodInput.Value())
	if method == "" {
		method = models.GET
	}

	req := &models.Request{
		ID:        fmt.Sprintf("req_%d", time.Now().Unix()),
		Method:    method,
		URL:       m.urlInput.Value(),
		Body:      m.bodyInput.Value(),
		Headers:   m.headers,
		CreatedAt: time.Now(),
	}

	m.isLoading = true
	m.history = append(m.history, req)

	return func() tea.Msg {
		resp, err := m.client.ExecuteRequest(req)
		return requestCompleteMsg{response: resp, error: err}
	}
}

func (m *Model) View() string {
	var b strings.Builder

	title := titleStyle.Render("🌊 Restless - The Cheeky HTTP Client")
	b.WriteString(title)
	b.WriteString("\n\n")

	tabs := []string{
		m.renderTab("Request", requestTab),
		m.renderTab("Response", responseTab),
		m.renderTab("History", historyTab),
	}
	b.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, tabs...))
	b.WriteString("\n\n")

	switch m.activeTab {
	case requestTab:
		b.WriteString(m.renderRequestTab())
	case responseTab:
		b.WriteString(m.renderResponseTab())
	case historyTab:
		b.WriteString(m.renderHistoryTab())
	}

	b.WriteString("\n\n")
	b.WriteString("Tab: Switch tabs • Enter: Send request • ↑/↓: Navigate • ←/→: Change method • q: Quit")

	return b.String()
}

func (m *Model) renderTab(name string, tab tabType) string {
	if m.activeTab == tab {
		return activeTabStyle.Render(name)
	}
	return tabStyle.Render(name)
}

func (m *Model) renderRequestTab() string {
	var b strings.Builder

	methodDisplay := methodStyle[string(m.methods[m.currentMethod])].Render(string(m.methods[m.currentMethod]))
	b.WriteString(fmt.Sprintf("Method: %s (use ←/→ to change)\n", methodDisplay))
	b.WriteString(fmt.Sprintf("URL: %s\n", inputStyle.Render(m.urlInput.View())))
	b.WriteString(fmt.Sprintf("Body: %s\n", inputStyle.Render(m.bodyInput.View())))

	b.WriteString("\nHeaders:\n")
	for i, header := range m.headers {
		style := lipgloss.NewStyle().Foreground(lipgloss.Color("#874BFD"))
		if i == m.selectedHeader {
			style = style.Background(lipgloss.Color("#F25D94"))
		}
		b.WriteString(style.Render(fmt.Sprintf("  %s: %s\n", header.Key, header.Value)))
	}

	b.WriteString(fmt.Sprintf("Add Header: %s : %s\n",
		m.headerKeyInput.View(), m.headerValueInput.View()))

	if m.isLoading {
		b.WriteString("\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("#FFBD2E")).Render("🚀 Sending request..."))
	}

	if m.error != "" {
		b.WriteString("\n" + errorStyle.Render("❌ Error: "+m.error))
	}

	return panelStyle.Render(b.String())
}

func (m *Model) renderResponseTab() string {
	if m.response == nil {
		return panelStyle.Render("No response yet. Send a request first!")
	}

	var b strings.Builder

	statusColor := successStyle
	if m.response.StatusCode >= 400 {
		statusColor = errorStyle
	}

	b.WriteString(fmt.Sprintf("Status: %s\n", statusColor.Render(m.response.Status)))
	b.WriteString(fmt.Sprintf("Duration: %v\n", m.response.Duration))
	b.WriteString(fmt.Sprintf("Size: %d bytes\n", m.response.Size))
	b.WriteString(fmt.Sprintf("Timestamp: %s\n", m.response.Timestamp.Format("15:04:05")))

	b.WriteString("\nHeaders:\n")
	for key, value := range m.response.Headers {
		b.WriteString(fmt.Sprintf("  %s: %s\n",
			lipgloss.NewStyle().Foreground(lipgloss.Color("#874BFD")).Render(key), value))
	}

	b.WriteString("\nBody:\n")
	maxBodyLength := 500
	body := m.response.Body
	if len(body) > maxBodyLength {
		body = body[:maxBodyLength] + "..."
	}
	b.WriteString(body)

	return responseStyle.Render(b.String())
}

func (m *Model) renderHistoryTab() string {
	if len(m.history) == 0 {
		return panelStyle.Render("No requests in history yet.")
	}

	var b strings.Builder
	b.WriteString("Request History:\n\n")

	for i, req := range m.history {
		methodColor := methodStyle[string(req.Method)]
		b.WriteString(fmt.Sprintf("%d. %s %s\n",
			i+1, methodColor.Render(string(req.Method)), req.URL))
		b.WriteString(fmt.Sprintf("   %s\n\n", req.CreatedAt.Format("15:04:05")))
	}

	return panelStyle.Render(b.String())
}
