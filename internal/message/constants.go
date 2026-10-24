package message

const (
	CommitTypeFallback string = "feat"
	ScopeFlagSentinel  string = "-"
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
	scopeBranchPattern string = `^[^/]*/([A-Z]*-[0-9]*).*`
	scopeFallback      string = ""
)
