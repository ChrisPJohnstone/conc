package cmd

const (
	cmdName               string = "conc"
	cmdShortHelp          string = "Conventional Commit Helper"
	commitTypeFlag        string = "type"
	commitTypeFlagDefault string = "feat"
	commitTypeBuild       string = "build"
	commitTypeChore       string = "chore"
	commitTypeCi          string = "ci"
	commitTypeDocs        string = "docs"
	commitTypeFeat        string = "feat"
	commitTypeFix         string = "fix"
	commitTypePerf        string = "perf"
	commitTypeRefactor    string = "refactor"
	commitTypeRevert      string = "revert"
	commitTypeStyle       string = "style"
	commitTypeTest        string = "test"
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
