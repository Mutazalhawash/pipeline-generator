package pipeline

type Plugins struct {
	AndroidLint   AndroidLint
	Findbugs      Findbugs
	Pmd           Pmd
	TaskPublisher TaskPublisher
	Violations    Violations
	HtmlReports   []Report
	ExtEmail      ExtEmail
	Testdroid     Testdroid
}

type AndroidLint struct {
	Pattern           string
	FailedTotalNormal int
	FailedTotalHigh   int
}

type Findbugs struct {
	Pattern string
}

type Pmd struct {
	Pattern string
}

type TaskPublisher struct {
	Pattern        string
	ExcludePattern string
}

type Violations struct {
	Checkstyle Checkstyle
}

type Checkstyle struct {
	Pattern string
}

type Report struct {
	ReportName  string
	ReportDir   string
	ReportFiles string
}

type ExtEmail struct {
	Content string
}

type Testdroid struct {
	AppPath       string
	DeviceGroupId string
	ProjectId     string
}
