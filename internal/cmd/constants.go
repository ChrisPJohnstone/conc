package cmd

const (
	cmdName           string = "conc"
	cmdShortHelp      string = "Conventional Commit Helper"
	verboseFlag       string = "verbose"
	commitTypeFlag    string = "type"
	scopeFlag         string = "scope"
	scopeFlagDefault  string = ""
	bodyFlag          string = "body"
	bodyFlagDefault   string = ""
	footerFlag        string = "footer"
	footerFlagDefault string = ""

	scopeFallback      string = ""
	scopeBranchPattern string = `^[^/]*/([A-Z]*-[0-9]*).*`
)
