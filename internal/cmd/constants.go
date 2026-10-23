package cmd

const (
	cmdName            string = "conc"
	cmdShortHelp       string = "Conventional Commit Helper"
	verboseFlag        string = "verbose"
	commitTypeFlag     string = "type"
	commitTypeFallback string = "feat"
	commitTypeBuild    string = "build"
	commitTypeChore    string = "chore"
	commitTypeCi       string = "ci"
	commitTypeDocs     string = "docs"
	commitTypeFeat     string = "feat"
	commitTypeFix      string = "fix"
	commitTypePerf     string = "perf"
	commitTypeRefactor string = "refactor"
	commitTypeRevert   string = "revert"
	commitTypeStyle    string = "style"
	commitTypeTest     string = "test"
	scopeFlag          string = "scope"
	scopeFlagDefault   string = ""
	scopeFlagSentinel  string = "-"
)

var commitTypes []string = []string{
	commitTypeBuild,
	commitTypeChore,
	commitTypeCi,
	commitTypeDocs,
	commitTypeFeat,
	commitTypeFix,
	commitTypePerf,
	commitTypeRefactor,
	commitTypeRevert,
	commitTypeStyle,
	commitTypeTest,
}
