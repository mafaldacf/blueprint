package main

import (
	"fmt"
	"go/token"
	"go/types"
	"log"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types/typeutil"
)

// -------------------------------------
// ------------- CONSTANTS -------------
// -------------------------------------
const (
	inputpackagepath = "../workflow/digota/"
	outfilename      = "./ssa-simple.out"
)

// -------------------------------------

// -------------------------------------
// --------------- GRAPH ---------------
// -------------------------------------

type Graph struct {
	nodes []*Node
	edges []*Edge
	pos   map[string]*Node
}

func NewAbstractGraph() *Graph {
	return &Graph{pos: make(map[string]*Node)}
}

func (graph *Graph) String() string {
	keys := make([]string, 0, len(graph.pos))
	for k := range graph.pos {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		// remove the "t" prefix and convert to int
		ni, _ := strconv.Atoi(strings.TrimPrefix(keys[i], "t"))
		nj, _ := strconv.Atoi(strings.TrimPrefix(keys[j], "t"))
		return ni < nj
	})

	str := ""
	for _, pos := range keys {
		node := graph.pos[pos]
		if node.tainted {
			str += "*"
		}
		str += fmt.Sprintf("%s (%s): %s\n", node.name, strings.Join(node.pointedBy, ","), node.str)
		for _, edge := range graph.GetEdges() {
			if edge.HasFromNode(node) {
				str += fmt.Sprintf("\t|--%s--> %s (%s): %s\n", edge.Type(), edge.to.name, strings.Join(edge.to.pointedBy, ","), edge.to.str)
			}
		}
		str += "\n"
	}

	return str
}

func (graph *Graph) GetNodeByPosIfExists(pos string) *Node {
	node, ok := graph.pos[pos]
	if ok {
		return node
	}
	return graph.pos["_"+pos]
}

func (graph *Graph) GetNodeByPos(pos string) *Node {
	node, ok := graph.pos[pos]
	if ok {
		return node
	}
	node, ok = graph.pos["_"+pos]
	if !ok {
		log.Fatalf("could not find node for pos (%s)", pos)
	}
	return node
}

func (graph *Graph) AddExistingNodeToNewPos(node *Node, pos string) {
	graph.pos[pos] = node
}

func (graph *Graph) AddNode(node *Node, pos string) {
	graph.nodes = append(graph.nodes, node)
	graph.pos[pos] = node
}

func (graph *Graph) AddEdge(edge *Edge) {
	graph.edges = append(graph.edges, edge)
}

func (graph *Graph) EdgeExists(from *Node, to *Node, edgeType EdgeType, fieldVal int, indexVal string) bool {
	for _, edge := range graph.edges {
		if edge.from.String() == from.String() && edge.to.String() == to.String() && (edge.edgeType == edgeType) && (edge.param == fieldVal || edge.index == indexVal) {
			return true
		}
	}
	return false
}

func (graph *Graph) GetNodes() []*Node {
	return graph.nodes
}

func (graph *Graph) GetEdges() []*Edge {
	return graph.edges
}

func (graph *Graph) GetEdgesFromNode(node *Node) []*Edge {
	var edges []*Edge
	for _, edge := range graph.edges {
		if edge.from == node {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (graph *Graph) GetValueEdgesFromNode(node *Node) []*Edge {
	var edges []*Edge
	for _, edge := range graph.edges {
		if edge.from == node && edge.edgeType == EDGE_VALUE {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (graph *Graph) GetFieldEdgesFromNode(node *Node) []*Edge {
	var edges []*Edge
	for _, edge := range graph.edges {
		if edge.from == node && edge.edgeType == EDGE_FIELD {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (graph *Graph) GetVersionEdgesFromNode(node *Node) []*Edge {
	var edges []*Edge
	for _, edge := range graph.edges {
		if edge.from == node && edge.edgeType == EDGE_VERSION {
			edges = append(edges, edge)
		}
	}
	return edges
}

func (graph *Graph) NodeHasValueEdge(node *Node) bool {
	for _, edge := range graph.GetEdgesFromNode(node) {
		if edge.edgeType == EDGE_VALUE {
			return true
		}
	}
	return false
}

func (graph *Graph) GetNodeForSSAValue(v ssa.Value) *Node {
	for _, node := range graph.nodes {
		if (node.name == v.Name() || slices.Contains(node.pointedBy, v.Name())) && node.str == v.String() {
			return node
		}
	}
	return nil
}

func (graph *Graph) GetNodeForSSAValue2(v ssa.Value) *Node {
	for _, node := range graph.nodes {
		if node.str == v.String() {
			return node
		}
	}
	return nil
}

// -------------------------------------
// ---------------- NODE ---------------
// -------------------------------------

type NodeType int

const (
	NODE_DEFAULT NodeType = iota
	NODE_FUNCTION
	NODE_PLACEHOLDER
	NODE_PARAMETER
	NODE_PHI
)

type Node struct { // objects
	name      string
	str       string
	fn        string
	pointedBy []string //variables
	nodeType  NodeType
	tainted   bool
}

func (node *Node) String() string {
	return fmt.Sprintf("(%s) %s: %s", node.fn, node.name, node.str)
}

func (node *Node) NewNodeVersion(name string) *Node {
	return &Node{
		name:      name,
		pointedBy: []string{name},
		str:       node.str,
		fn:        node.fn,
		nodeType:  NODE_PARAMETER, // is this possible?
	}
}

func NewNode(val ssa.Value, nodeType NodeType) *Node {
	var fn string
	if val.Parent() != nil { // may be nil for parameter node
		fn = val.Parent().Name()
	}

	return &Node{
		name:      val.Name(),
		str:       val.String(),
		fn:        fn,
		pointedBy: []string{val.Name()},
		nodeType:  nodeType,
	}
}

func (node *Node) AddToPointedBy(pointedBy string) {
	node.pointedBy = append(node.pointedBy, pointedBy)
}

func (node *Node) HasName(name string) bool {
	return node.name == name
}

func (node *Node) HasFn(fn string) bool {
	return node.fn == fn
}

// -------------------------------------
// ---------------- EDGE ---------------
// -------------------------------------

type EdgeType int

const (
	EDGE_FIELD EdgeType = iota
	EDGE_INDEX
	EDGE_PARAMETER
	EDGE_VALUE
	EDGE_VERSION
	EDGE_INTERFACE
	EDGE_CONVERTED
	EDGE_PHI
	EDGE_RETURN
	EDGE_COPY
)

type Edge struct {
	from *Node
	to   *Node

	edgeType EdgeType

	index string
	param int
}

func (edge *Edge) Type() string {
	switch edge.edgeType {
	case EDGE_FIELD:
		return fmt.Sprintf("field(%d)", edge.param)
	case EDGE_INDEX:
		return fmt.Sprintf("idx(%s)", edge.index)
	case EDGE_PARAMETER:
		return fmt.Sprintf("param(%d)", edge.param)
	case EDGE_VALUE:
		return "value"
	case EDGE_VERSION:
		return "version"
	case EDGE_INTERFACE:
		return "interface of"
	case EDGE_CONVERTED:
		return "converted to"
	case EDGE_PHI:
		return "phi"
	case EDGE_RETURN:
		return fmt.Sprintf("ret(%d)", edge.param)
	case EDGE_COPY:
		return "copy"
	default:
		log.Fatalf("unknown edge type (%v) for edge: %v", edge.edgeType, edge)
		return ""
	}
}

func NewEdge(from *Node, to *Node, edgeType EdgeType, index string, param int) *Edge {
	return &Edge{
		from:     from,
		to:       to,
		edgeType: edgeType,
		index:    index,
		param:    param,
	}
}

func (edge *Edge) HasFromNode(node *Node) bool {
	return edge.from == node
}

// -------------------------------------
// -------------- PARSER ---------------
// -------------------------------------

func ParseInstr(graph *Graph, instr ssa.Instruction, idx int) {
	if val, ok := instr.(ssa.Value); ok {
		parseValue(graph, instr, idx, val)
		return
	}

	switch t := instr.(type) {
	case *ssa.Store:
		// 04 [store] *t1 = currency
		fmt.Printf("%02d [store] %v\n", idx, instr.String())
		addrNode := parseValue(graph, instr, idx, t.Addr)
		valNode := parseValue(graph, instr, idx, t.Val)

		if graph.NodeHasValueEdge(addrNode) {
			newAddrNode := addrNode.NewNodeVersion(t.Addr.Name())
			graph.AddNode(newAddrNode, t.Addr.Name())
			edge := NewEdge(addrNode, newAddrNode, EDGE_VERSION, "", 0)
			graph.AddEdge(edge)
			addrNode = newAddrNode
		}

		edge := NewEdge(addrNode, valNode, EDGE_VALUE, "", 0)
		graph.AddEdge(edge)
	case *ssa.Return:
		fmt.Printf("[A] skipping... %02d [%T] %v\n", idx, instr, instr.String())
	case *ssa.Jump:
		fmt.Printf("[A] skipping... %02d [%T] %v\n", idx, instr, instr.String())
	default:
		fmt.Printf("[1] ignoring... %02d [%T] %v\n", idx, instr, instr.String())
	}
}

func parseValue(graph *Graph, instr ssa.Instruction, idx int, val ssa.Value) *Node {
	if node := graph.GetNodeByPosIfExists(val.Name()); node != nil {
		return node
	}

	switch t := val.(type) {
	case *ssa.Call:
		fmt.Printf("%02d [call] %s = %v\n", idx, val.Name(), val.String())
		node := NewNode(val, NODE_FUNCTION)

		graph.AddNode(node, val.Name())
		var argNodes []*Node
		for i, arg := range t.Call.Args {
			argNode := parseValue(graph, instr, idx, arg)
			argNodes = append(argNodes, argNode)
			edge := NewEdge(node, argNode, EDGE_PARAMETER, "", i)
			graph.AddEdge(edge)
		}

		if fn, ok := t.Call.Value.(*ssa.Builtin); ok {
			fmt.Printf("1. CALLING: %v\n", fn.Name())
			for _, argNode := range argNodes {
				fmt.Printf("\t ARGNODE = %v\n", argNode)
			}
		} else if t.Call.Method != nil &&
			t.Call.Method.Signature().Recv().Type().String() == "github.com/blueprint-uservices/blueprint/runtime/core/backend.NoSQLCollection" &&
			t.Call.Method.Name() == "InsertOne" {
			fmt.Printf("2. CALLING: [%T] %v //[%T] %v\n", t.Call.Method, t.Call.Method, t.Call.Value, t.Call.Value)
			for _, argNode := range argNodes {
				recurseTaint(graph, argNode)
			}
		}

		return node
	case *ssa.Alloc:
		fmt.Printf("%02d [alloc] %s = %v\n", idx, val.Name(), val.String())
		node := NewNode(t, NODE_DEFAULT)
		graph.AddNode(node, t.Name())
		return node
	case *ssa.Slice:
		fmt.Printf("%02d [slice] %s = %v\n", idx, val.Name(), val.String())
		node := parseValue(graph, instr, idx, t.X)
		node.AddToPointedBy(val.String())
		return node
	case *ssa.FieldAddr:
		// 00 [field] t27 = &t0.Items [#3]
		fmt.Printf("%02d [field] %s = %v\n", idx, val.Name(), val.String())
		node := graph.GetNodeForSSAValue2(val)
		if node == nil {
			node = NewNode(t, NODE_DEFAULT)
			graph.AddNode(node, val.Name())
		} else {
			node.AddToPointedBy(val.Name())
			graph.AddExistingNodeToNewPos(node, val.Name())
		}
		topNode := graph.GetNodeByPosIfExists(t.X.Name())
		if topNode == nil {
			// e.g. 00 [field] t36 = &s.skuService [#0]
			// ignore for now
			return nil
		}
		if !graph.EdgeExists(topNode, node, EDGE_FIELD, t.Field, "") {
			edge := NewEdge(topNode, node, EDGE_FIELD, "", t.Field)
			graph.AddEdge(edge)
		}
		return node
	case *ssa.IndexAddr:
		fmt.Printf("%02d [index] %s = %v\n", idx, val.Name(), val.String())
		node := graph.GetNodeForSSAValue2(val)
		if node == nil {
			node = NewNode(t, NODE_DEFAULT)
			graph.AddNode(node, t.Name())
		} else {
			node.AddToPointedBy(t.Name())
			graph.AddExistingNodeToNewPos(node, t.Name())
		}
		topNode := graph.GetNodeByPosIfExists(t.X.Name())
		if topNode == nil {
			// ignore for now
			return nil
		}
		if !graph.EdgeExists(topNode, node, EDGE_INDEX, 0, t.Index.String()) {
			edge := NewEdge(topNode, node, EDGE_INDEX, t.Index.String(), 0)
			graph.AddEdge(edge)
		}
		return node
	case *ssa.UnOp:
		// e.g.,
		// 01 [unary] t14 = *t13
		// 05 [unary] t31 = *t30
		fmt.Printf("%02d [unary] %s = %v\n", idx, val.Name(), val.String())
		targetNode := parseValue(graph, instr, idx, t.X)

		node := NewNode(t, NODE_DEFAULT)
		graph.AddNode(node, t.Name())

		if graph.GetEdgesFromNode(targetNode) != nil {
			/* versionEdges := graph.GetVersionEdgesFromNode(targetNode)
			if versionEdges != nil {
				lastVersionEdge := versionEdges[len(versionEdges)-1]
				targetNode = lastVersionEdge.to
			} */
			valueEdges := graph.GetValueEdgesFromNode(targetNode)
			if valueEdges != nil {
				// 1. get the current value of the current address
				// 2. create new address and assign that value
				targetNodeValue := valueEdges[0].to
				targetNodeValue.AddToPointedBy(val.Name())
				graph.AddExistingNodeToNewPos(targetNodeValue, val.Name())
				//edge := NewEdge(node, targetNodeValue, EDGE_VALUE, "", 0)
				//graph.AddEdge(edge)
				return node
			}
			// create copy edge since target already has some value or fields
			/*edge := NewEdge(node, targetNode, EDGE_COPY, "", 0) */
			edge := NewEdge(targetNode, node, EDGE_COPY, "", 0)
			graph.AddEdge(edge)
		} else {
			// assign value for the first time if it does not exist yet
			edge := NewEdge(targetNode, node, EDGE_VALUE, "", 0)
			graph.AddEdge(edge)
		}
		return node
	case *ssa.MakeInterface: // same as *ssa.UnOp
		fmt.Printf("%02d [interface] %s = %v\n", idx, val.Name(), val.String())
		targetNode := parseValue(graph, instr, idx, t.X)

		node := NewNode(t, NODE_DEFAULT)
		graph.AddNode(node, t.Name())

		edge := NewEdge(node, targetNode, EDGE_INTERFACE, "", 0)
		graph.AddEdge(edge)
		return node
	case *ssa.Convert: // same as *ssa.UnOp and *ssa.MakeInterface
		fmt.Printf("%02d [convert] %s = %v\n", idx, val.Name(), val.String())
		targetNode := parseValue(graph, instr, idx, t.X)

		node := NewNode(t, NODE_DEFAULT)
		graph.AddNode(node, t.Name())

		edge := NewEdge(node, targetNode, EDGE_CONVERTED, "", 0)
		graph.AddEdge(edge)
		return node
	case *ssa.Parameter: // dynamic
		fmt.Printf("%02d [parameter] %s = %v\n", idx, val.Name(), val.String())
		node := NewNode(val, NODE_PLACEHOLDER)
		node.name = "_" + node.name
		graph.AddNode(node, "_"+val.Name())
		return node
	case *ssa.Const:
		fmt.Printf("%02d [const] %s = %v\n", idx, val.Name(), val.String())
		return NewNode(val, NODE_PARAMETER)
	case *ssa.Phi:
		fmt.Printf("%02d [phi] %s = %v\n", idx, val.Name(), val.String())
		node := NewNode(val, NODE_PHI)
		graph.AddNode(node, val.Name())
		for _, phiEdge := range t.Edges {
			fmt.Printf("HERE FOR PHI EDGE: %v\n", phiEdge)
			otherNode := parseValue(graph, instr, idx, phiEdge)
			edge := NewEdge(node, otherNode, EDGE_PHI, "", 0)
			graph.AddEdge(edge)
		}
		return node
	case *ssa.Extract:
		fmt.Printf("%02d [extract] %s = %v\n", idx, val.Name(), val.String())
		extractFromNode := parseValue(graph, instr, idx, t.Tuple)
		node := NewNode(t, NODE_DEFAULT)
		graph.AddNode(node, val.Name())
		edge := NewEdge(extractFromNode, node, EDGE_RETURN, "", t.Index)
		graph.AddEdge(edge)
		return node

	case *ssa.BinOp, *ssa.Global: //FIXME
		fmt.Printf("[B] skipping... %02d [%T] %s = %v\n", idx, val, val.Name(), val.String())
		node := NewNode(val, NODE_DEFAULT)
		graph.AddNode(node, val.Name())
		return node
	default:
		fmt.Printf("[2] ignoring... %02d [%T] %s = %v\n", idx, val, val.Name(), val.String())
	}
	log.Fatal("returning nil node")
	return nil
}

func recurseTaint(graph *Graph, node *Node) {
	fmt.Printf("visiting node: %v\n", node.String())
	if node.tainted == false {
		node.tainted = true
		fmt.Printf("\ttainting node: %v\n", node.String())
		for _, edge := range graph.GetEdgesFromNode(node) {
			recurseTaint(graph, edge.to)
		}
	}
}

// -------------------------------------
// ---------------- MAIN ---------------
// -------------------------------------

var ssaPkgs map[*packages.Package]bool

func recurse(prog *ssa.Program, pkg *packages.Package) {
	if _, ok := ssaPkgs[pkg]; ok {
		return
	}
	prog.CreatePackage(pkg.Types, pkg.Syntax, pkg.TypesInfo, false)
	ssaPkgs[pkg] = true
	for _, impt := range pkg.Imports {
		recurse(prog, impt)
	}
}

/* var driver neo4j.DriverWithContext
var ctx context.Context */

func main() {
	cfg := &packages.Config{Mode: packages.LoadAllSyntax}
	pkgs, err := packages.Load(cfg, inputpackagepath)
	if err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	//prog := ssa.NewProgram(fset, ssa.PrintFunctions)
	prog := ssa.NewProgram(fset, 0)

	ssaPkgs = make(map[*packages.Package]bool)
	ssaPkgsFiltered := make([]*ssa.Package, len(pkgs))
	for i, pkg := range pkgs {
		if _, ok := ssaPkgs[pkg]; ok {
			continue
		}
		ssaPkgsFiltered[i] = prog.CreatePackage(pkg.Types, pkg.Syntax, pkg.TypesInfo, false)
		ssaPkgs[pkg] = true
		for _, impt := range pkg.Imports {
			recurse(prog, impt)
		}
	}

	prog.Build()

	var appPkgs []*ssa.Package
	for _, ssaPkg := range ssaPkgsFiltered {
		if ssaPkg == nil || ssaPkg.Pkg == nil {
			continue
		}
		if ssaPkg.Pkg.Name() != "digota" {
			continue
		}
		/* if ssaPkg.Func("main") == nil && ssaPkg.Func("init") == nil {
			continue
		} */
		appPkgs = append(appPkgs, ssaPkg)
	}

	/* uri := "bolt://localhost:7687"
	username := "neo4j"
	password := "password"

	driver, err = createNeo4jDriver(uri, username, password)
	if err != nil {
		log.Fatalf("failed to connect to Neo4j: %v", err)
	}
	ctx = context.Background()
	defer driver.Close(ctx) */

	ssaAnalysis(prog, appPkgs)
}

/* func createNeo4jDriver(uri, username, password string) (neo4j.DriverWithContext, error) {
	return neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
}

func saveGraphToNeo4j(graph *Graph) error {
	session := driver.NewSession(ctx, neo4j.SessionConfig{AccessMode: neo4j.AccessModeWrite})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
		// create nodes
		for _, node := range graph.nodes {
			_, err := tx.Run(ctx,
				`MERGE (n:Node {name: $name, fn: $fn, str: $str, type: $type, tainted: $tainted})`,
				map[string]interface{}{
					"name":    node.name,
					"fn":      node.fn,
					"str":     node.str,
					"type":    int(node.nodeType),
					"tainted": node.tainted,
				},
			)
			if err != nil {
				return nil, err
			}
		}

		// create edges
		for _, edge := range graph.edges {
			_, err := tx.Run(ctx,
				`MATCH (from:Node {name: $from}), (to:Node {name: $to})
				 MERGE (from)-[:`+edge.Type()+` {index: $index, param: $param}]->(to)`,
				map[string]interface{}{
					"from":  edge.from.name,
					"to":    edge.to.name,
					"index": edge.index,
					"param": edge.param,
				},
			)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
	})
	return err
} */

func iterateFunc(outFile *os.File, fn *ssa.Function, memberType types.Type) {
	var graph = NewAbstractGraph()

	var filename string
	namedMemberType, ok := memberType.(*types.Named)

	if ok && namedMemberType.Obj().Name() != "SkuServiceImpl" && namedMemberType.Obj().Name() != "OrderServiceImpl" {
		return
	}

	fmt.Printf("=============================\n")
	if ok {
		filename = fmt.Sprintf("%s/%s.graph", namedMemberType.Obj().Name(), fn.Name())
		fmt.Printf("%s.%s()\n", namedMemberType.Obj().Name(), fn.Name())
	} else {
		filename = fmt.Sprintf("%s.graph", fn.Name())
		fmt.Printf("%s()\n", fn.Name())
	}
	fmt.Printf("=============================\n")

	fmt.Fprintf(outFile, "Function: %s\n", fn.Name())
	fmt.Printf("\n--------------- Function: %s\n", fn.Name())
	for i, block := range fn.Blocks {
		fmt.Fprintf(outFile, "Block #%d: %s.%s\n", i, fn.Name(), block.Comment)
		fmt.Printf("----- Block #%d: %s.%s\n", i, fn.Name(), block.Comment)

		for j, instr := range block.Instrs {
			if val, ok := instr.(ssa.Value); ok {
				fmt.Fprintf(outFile, "\t\t\t%02d: %s = %s\n", j, val.Name(), instr.String())
			} else {
				fmt.Fprintf(outFile, "\t\t\t%02d: %s\n", j, instr.String())
			}
			if filename == "OrderServiceImpl/New.graph" || filename == "OrderServiceImpl/New2.graph" {
				ParseInstr(graph, instr, j)
			}
		}
	}

	if ok {
		outfile, err := os.Create(fmt.Sprintf("out/%s", filename))
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}
		outfile.WriteString(graph.String())
		defer outfile.Close()

	}
	fmt.Println()
	fmt.Println()
	/* saveGraphToNeo4j(graph) */
}

func ssaAnalysis(prog *ssa.Program, pkgs []*ssa.Package) {
	outFile, err := os.Create(outfilename)
	if err != nil {
		log.Fatal(err)
	}
	defer outFile.Close()

	for _, ssaPkg := range pkgs {
		outfile, err := os.Create(fmt.Sprintf("%s.ssa", ssaPkg.Pkg.Name()))
		if err != nil {
			log.Fatalf("failed to create output file: %v", err)
		}
		defer outfile.Close()
		ssaPkg.WriteTo(outfile)

		for _, member := range ssaPkg.Members {
			switch m := member.(type) {
			case *ssa.Function:
				iterateFunc(outFile, m, nil)

			case *ssa.Global:
				fmt.Fprintf(outFile, "\tGlobal: %s, Type: %s\n", m.Name(), m.Type().String())

			case *ssa.Type:
				fmt.Fprintf(outFile, "\tType: %s\n", m.Type())

				// this logic was copied from
				// package: golang.org/x/tools/go/ssa
				// file: print.go
				// function: func (p *Package) WriteTo(w io.Writer) (int64, error)
				for _, sel := range typeutil.IntuitiveMethodSet(m.Type(), &prog.MethodSets) {
					method := prog.MethodValue(sel)
					fmt.Fprintf(outFile, "\tMethod: %v\n", sel.Obj().Type())
					if method != nil {
						iterateFunc(outFile, method, m.Type())
					}
				}

				methods := prog.MethodSets.MethodSet(m.Type().Underlying())
				for i := 0; i < methods.Len(); i++ {
					sel := methods.At(i)
					fmt.Fprintf(outFile, "\tMethod: %v\n", sel.Obj().Type())
					method := prog.MethodValue(sel)
					if method != nil {
						iterateFunc(outFile, method, m.Type())
					}
				}

			default:
				fmt.Fprintf(outFile, "\tUnknown member type: %T\n", m)
			}
		}
	}
}
