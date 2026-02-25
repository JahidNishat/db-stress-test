package ui

import (
	"db-stress/internal/runner"
	"fmt"
	"sort"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	statsCh   <-chan runner.Stats
	startTime time.Time

	successCount int64
	failCount    int64
	totalLatency time.Duration

	done      bool
	endTime   time.Time
	latencies []float64
}

func InitialModel(ch <-chan runner.Stats) Model {
	return Model{
		statsCh:   ch,
		startTime: time.Now(),
		latencies: make([]float64, 0, 10000),
	}
}

type statMsg runner.Stats
type testFinishedMsg struct{}

func waitForStats(ch <-chan runner.Stats) tea.Cmd {
	return func() tea.Msg {
		s, ok := <-ch
		if !ok {
			return testFinishedMsg{}
		}
		return statMsg(s)
	}
}

func (m Model) Init() tea.Cmd {
	return waitForStats(m.statsCh)
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case statMsg:
		if msg.IsError {
			m.failCount++
		} else {
			m.successCount++
			m.totalLatency += msg.Duration
			m.latencies = append(m.latencies, msg.Duration.Seconds())
		}
		return m, waitForStats(m.statsCh)
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	case testFinishedMsg:
		m.done = true
		m.endTime = time.Now()
		return m, tea.Quit
	}
	return m, nil
}

func (m Model) View() string {
	var duration time.Duration
	if m.done {
		duration = m.endTime.Sub(m.startTime)
	} else {
		duration = time.Since(m.startTime)
	}

	elapsed := duration.Seconds()
	if elapsed == 0 {
		elapsed = 1
	}
	rps := float64(m.successCount) / elapsed

	sort.Float64s(m.latencies)
	p50 := calculatePercentile(m.latencies, 0.50)
	p99 := calculatePercentile(m.latencies, 0.99)

	return fmt.Sprintf(`
⚡ DB Stress Test Running...

Duration:	%.2fs
Success:	%d
Failed:		%d
RPS:		%.2f req/s
P50:		%.2f s
P99:		%.2f s

Press 'q' to quit.`, elapsed, m.successCount, m.failCount, rps, p50, p99)
}

func calculatePercentile(sorted []float64, p float64) float64 {
	if len(sorted) == 0 {
		return 0
	}
	idx := int(float64(len(sorted)-1) * p)
	return sorted[idx]
}
