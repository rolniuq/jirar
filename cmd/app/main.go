package main

import (
	"encoding/json"
	"fmt"
	"jirar/configs"
	"net/http"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

// --- STRUCTS (Để map JSON từ Jira về Go) ---
type JiraResponse struct {
	Issues []Issue `json:"issues"`
}

type Issue struct {
	Key    string      `json:"key"`
	Fields IssueFields `json:"fields"`
}

type IssueFields struct {
	Summary string `json:"summary"`
	Status  Status `json:"status"`
	Updated string `json:"updated"` // Dạng string ISO8601
}

type Status struct {
	Name string `json:"name"`
}

func main() {
	appConfigs := configs.NewAppConfig()
	if err := appConfigs.Load(); err != nil {
		panic(err)
	}

	jiraConfigs := appConfigs.GetJiraConfigs()

	// 1. Kiểm tra cấu hình
	if jiraConfigs.Domain == "" || jiraConfigs.Email == "" || jiraConfigs.Token == "" {
		fmt.Println("❌ Lỗi: Thiếu biến môi trường.")
		fmt.Println("Vui lòng set: JIRA_DOMAIN, JIRA_EMAIL, JIRA_TOKEN")
		os.Exit(1)
	}

	// 2. Chuẩn bị Request
	// JQL: Lấy ticket liên quan đến tôi, update trong 24h qua
	jql := "updated >= -24h AND (assignee = currentUser() OR watcher = currentUser() OR text ~ currentUser()) ORDER BY updated DESC"

	apiURL := fmt.Sprintf("%s/rest/api/3/search", jiraConfigs.Domain)
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		panic(err)
	}

	// Thêm Query Params
	q := req.URL.Query()
	q.Add("jql", jql)
	q.Add("fields", "summary,status,updated") // Chỉ lấy field cần thiết cho nhẹ
	q.Add("maxResults", "10")
	req.URL.RawQuery = q.Encode()

	// Thêm Auth & Headers
	req.SetBasicAuth(jiraConfigs.Email, jiraConfigs.Token)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	// 3. Gọi API
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❌ Lỗi kết nối: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("❌ Lỗi API: %s (Check lại Token/Email)\n", resp.Status)
		os.Exit(1)
	}

	// 4. Parse JSON
	var data JiraResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("❌ Lỗi đọc dữ liệu: %v\n", err)
		os.Exit(1)
	}

	if len(data.Issues) == 0 {
		fmt.Println("✅ Không có thông báo mới nào trong 24h qua!")
		return
	}

	renderTable(jiraConfigs, data.Issues)
}

func renderTable(jiraConfigs *configs.JiraConfigs, issues []Issue) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Key", "Status", "Updated", "Summary", "Link"})

	// Style cho bảng đẹp hơn
	// table.SetBorder(false)
	// table.SetHeaderLine(false)
	// table.SetRowLine(false)
	// table.SetColumnSeparator("  ")
	// table.SetHeaderAlignment(tablewriter.ALIGN_LEFT)
	// table.SetAlignment(tablewriter.ALIGN_LEFT)

	for _, issue := range issues {
		// Format thời gian cho dễ nhìn (HH:MM)
		parsedTime, _ := time.Parse("2006-01-02T15:04:05.000-0700", issue.Fields.Updated)
		timeStr := parsedTime.Format("15:04 02/01")

		// Tạo link để click (trên terminal hỗ trợ)
		link := fmt.Sprintf("%s/browse/%s", jiraConfigs.Domain, issue.Key)

		// Thêm màu sắc (Optional)
		statusColor := ""
		if issue.Fields.Status.Name == "Done" || issue.Fields.Status.Name == "Resolved" {
			statusColor = "✅ " // Icon cho đẹp
		} else if issue.Fields.Status.Name == "In Progress" {
			statusColor = "🔥 "
		} else {
			statusColor = "todo "
		}

		row := []string{
			issue.Key,
			statusColor + issue.Fields.Status.Name,
			timeStr,
			issue.Fields.Summary,
			link,
		}
		table.Append(row)
	}
	table.Render()
}
