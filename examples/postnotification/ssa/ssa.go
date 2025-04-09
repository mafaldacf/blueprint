package main

import (
	"fmt"
	"go/token"
	"log"
	"os"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
)

var createdPkgs map[*packages.Package]bool

func recurse(prog *ssa.Program, pkg *packages.Package) {
	if _, ok := createdPkgs[pkg]; ok {
		return
	}
	prog.CreatePackage(pkg.Types, pkg.Syntax, pkg.TypesInfo, false)
	createdPkgs[pkg] = true
	for _, impt := range pkg.Imports {
		recurse(prog, impt)
	}
}

func main() {
	createdPkgs = make(map[*packages.Package]bool)
	cfg := &packages.Config{Mode: packages.LoadAllSyntax}
	pkgs, err := packages.Load(cfg, "../workflow/postnotification/...")
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	prog := ssa.NewProgram(fset, ssa.PrintFunctions)

	ssaPkgs := make([]*ssa.Package, len(pkgs))
	for i, pkg := range pkgs {
		if _, ok := createdPkgs[pkg]; ok {
			return
		}
		ssaPkgs[i] = prog.CreatePackage(pkg.Types, pkg.Syntax, pkg.TypesInfo, false)
		createdPkgs[pkg] = true
		for _, impt := range pkg.Imports {
			recurse(prog, impt)
		}
	}
	
	prog.Build()

	outFile, err := os.Create("ssa.out")
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	for _, ssaPkg := range ssaPkgs {
		fmt.Fprintf(outFile, "Package: %s\n", ssaPkg.Pkg.Path())

		for name, member := range ssaPkg.Members {
			fmt.Fprintf(outFile, "  Member: %s (%T)\n", name, member)

			switch m := member.(type) {
			case *ssa.Function:
				fmt.Fprintf(outFile, "    Function: %s: %v\n", m.Name(), m.Blocks)

				for i, b := range m.Blocks {
					fmt.Fprintf(outFile, "block [%d]: %v\n", i, b)
					for j, inst := range b.Instrs {
						fmt.Fprintf(outFile, "\t inst #%d: %v\n", j, inst)
					}
				}
		
			case *ssa.Global:
				fmt.Fprintf(outFile, "    Global: %s, Type: %s\n", m.Name(), m.Type().String())
		
			case *ssa.Type:
				fmt.Fprintf(outFile, "    Type: %s\n", m.Type().String())
		
			default:
				fmt.Fprintf(outFile, "    Unknown member type: %T\n", m)
			}
			fmt.Fprintf(outFile, "\n")
		}
		fmt.Fprintf(outFile, "\n----------------\n\n")
	}

	/* for _, ssaPkg := range ssaPkgs {
		fmt.Printf("Package: %s\n", ssaPkg.Pkg.Path())

		for name, member := range ssaPkg.Members {
			fmt.Printf("  Member: %s (%T)\n", name, member)
		
			switch m := member.(type) {
			case *ssa.Function:
				fmt.Printf("    Function: %s\n", m.Name())
				printFunction(m)
		
			case *ssa.Global:
				fmt.Printf("    Global: %s, Type: %s\n", m.Name(), m.Type().String())
		
			case *ssa.Type:
				fmt.Printf("    Type: %s\n", m.Type().String())
		
			default:
				fmt.Printf("    Unknown member type: %T\n", m)
			}
			fmt.Println()
		}

		fmt.Println()
		fmt.Println("----------------")
	} */
}

func printFunction(fn *ssa.Function) {
	fmt.Printf("  Function: %s\n", fn.Name())

	for _, block := range fn.Blocks {
		fmt.Printf("    Block: %s\n", block.Comment)

		for _, instr := range block.Instrs {
			fmt.Printf("      Instr: %s\n", instr.String())
		}
	}
}
