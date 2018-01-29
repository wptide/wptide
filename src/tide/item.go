package tide

import "reflect"

type ResultSet struct {
	Results []Item
}

type Item struct {
	Title       string                 `json:"title"`
	Version     string                 `json:"version"`
	Checksum    string                 `json:"checksum"`
	Visibility  string                 `json:"visibility"`
	ProjectType string                 `json:"project_type"`
	SourceUrl   string                 `json:"source_url"`
	SourceType  string                 `json:"source_type"`
	CodeInfo    CodeInfo               `json:"code_info,omitempty"`
	Results     map[string]AuditResult `json:"results,omitempty"`
}

type CodeInfo struct {
	Type    string                `json:"type"`
	Details []InfoDetails         `json:"details"`
	Cloc    map[string]ClocResult `json:"cloc"`
}

type InfoDetails struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type InfoDetailsSimple struct {
	Name        string
	PluginURI   string
	ThemeURI    string
	Version     string
	Description string
	Author      string
	AuthorURI   string
	TextDomain  string
}

type ClocResult struct {
	Blank   int `json:"blank"`
	Comment int `json:"comment"`
	Code    int `json:"code"`
	NFiles  int `json:"n_files"`
}

type AuditResult struct {
	Full struct {
		Type       string `json:"type,omitempty"`
		Key        string `json:"key,omitempty"`
		BucketName string `json:"bucket_name,omitempty"`
	} `json:"full,omitempty"`
	Details struct {
		Type       string `json:"type,omitempty"`
		Key        string `json:"key,omitempty"`
		BucketName string `json:"bucket_name,omitempty"`
		*PhpcsResults
		*LighthouseResults
	} `json:"details,omitempty"`
	Summary            *AuditSummary `json:"summary,omitempty"`
	CompatibleVersions []string      `json:"compatible_versions,omitempty"`
	Error              string        `json:"error,omitempty"`
}

type PhpcsResults struct {
	Totals struct {
		Errors   int `json:"errors,omitempty"`
		Warnings int `json:"warnings,omitempty"`
	} `json:"totals,omitempty"`
	Files map[string]struct {
		Errors   int `json:"errors, omitempty"`
		Warnings int `json:"warnings,omitempty"`
		Messages []struct {
			Message  string `json:"message"`
			Source   string `json:"source"`
			Type     string `json:"type"`
			Line     int    `json:"line"`
			Column   int    `json:"column"`
			Severity int    `json:"severity,omitempty"`
		} `json:"messages,omitempty"`
	} `json:"files,omitempty"`
}

type AuditSummary struct {
	*PhpcsSummary
	*LighthouseSummary
}

type PhpcsSummary struct {
	Files map[string]struct {
		Errors   int `json:"errors"`
		Warnings int `json:"warnings"`
	} `json:"files,omitempty"`
	FilesCount    int `json:"files_count,omitempty"`
	ErrorsCount   int `json:"errors_count,omitempty"`
	WarningsCount int `json:"warnings_count,omitempty"`
}

// @todo Define this later
type LighthouseResults struct{}

type LighthouseSummary struct {
	ReportCategories []LighthouseCategory `json:"reportCategories,omitempty"`
}

type LighthouseCategory struct {
	Name        string  `json:"name"`
	Weight      float32 `json:"weight"`
	Description string  `json:"description"`
	Id          string  `json:"id"`
	Score       float32 `json:"score"`
}

func SimplifyCodeDetails(details []InfoDetails) *InfoDetailsSimple {
	simple := &InfoDetailsSimple{}

	sV := reflect.ValueOf(simple).Elem()

	for _, item := range details {
		sV.FieldByName(item.Key).Set(reflect.ValueOf(item.Value))
	}
	return simple
}

func ComplexifyCodeDetails(simple *InfoDetailsSimple) []InfoDetails {

	sV := reflect.ValueOf(*simple)

	details := []InfoDetails{}

	for i := 0; i < sV.NumField(); i++ {
		if sV.Field(i).String() != "" {
			details = append(details, InfoDetails{
				sV.Type().Field(i).Name,
				sV.Field(i).String(),
			})
		}
	}

	return details
}
