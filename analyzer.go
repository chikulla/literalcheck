package literalcheck

import (
	"go/ast"
	"go/token"
	"go/types"

	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("literalcheck", New)
}

type LiteralcheckPlugin struct{}

func New(conf any) (register.LinterPlugin, error) {
	return &LiteralcheckPlugin{}, nil
}

func (l *LiteralcheckPlugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		Analyzer,
	}, nil
}

func (l *LiteralcheckPlugin) GetLoadMode() string {
	return register.LoadModeSyntax
}

// Analyzer is the definition of the custom linter
var Analyzer = &analysis.Analyzer{
	Name: "literalcheck",
	Doc:  "Check for passing raw string literals to specific type aliases",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		InspectFile(file, pass)
	}
	return nil, nil
}

func InspectFile(file *ast.File, pass *analysis.Pass) {
	ins := nodeInspector{pass}
	ast.Inspect(file, ins.inspect)
}

type nodeInspector struct {
	pass *analysis.Pass
}

func (i *nodeInspector) inspect(n ast.Node) bool {
	callExpr, ok := n.(*ast.CallExpr)
	if !ok {
		return true
	}
	for _, arg := range callExpr.Args {
		i.inspectExpr(arg)
	}

	return true
}

func (i *nodeInspector) inspectExpr(arg ast.Expr) {
	basicLit, ok := arg.(*ast.BasicLit)
	if !ok {
		return
	}

	// Check the kind of literal
	switch basicLit.Kind {
	case token.STRING, token.INT, token.FLOAT:
		i.checkLiteral(arg, basicLit)
	}
}

func (i *nodeInspector) checkLiteral(arg ast.Expr, lit *ast.BasicLit) {
	// Get the type of the argument
	argType := i.pass.TypesInfo.Types[arg].Type
	if namedType, ok := argType.(*types.Named); ok {
		i.pass.Reportf(lit.Pos(), "raw literal (%s) passed to type alias (%s), use a constant instead", lit.Kind, namedType.String())
	}
}
