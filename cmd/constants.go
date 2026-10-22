package cmd

const (
	commitTypeFlag            string = "type"
	commitTypeFlagDefault     string = "feat"
	commitTypeFlagDescription string = "type of commit"
	commitTypeBuild           string = "build"
	commitTypeChore           string = "chore"
	commitTypeCi              string = "ci"
	commitTypeDocs            string = "docs"
	commitTypeFeat            string = "feat"
	commitTypeFix             string = "fix"
	commitTypePerf            string = "perf"
	commitTypeRefactor        string = "refactor"
	commitTypeRevert          string = "revert"
	commitTypeStyle           string = "style"
	commitTypeTest            string = "test"
)

var commitTypes = []string{
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
