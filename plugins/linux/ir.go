package linux

import (
	"gitlab.mpi-sws.org/cld/blueprint/blueprint/pkg/blueprint"
	"gitlab.mpi-sws.org/cld/blueprint/blueprint/pkg/core"
)

/*
The base IRNode interface for linux processes
*/
type Process interface {
	core.ProcessNode

	ImplementsLinuxProcess()
}

/*
Code and artifact generation interfaces that IRNodes can implement to provide linux
processes.
*/
type (

	/*
		For process nodes that want to provide code or other artifacts for their process.
		Methods on the `builder` argument are used for collecting the artifacts
	*/
	ProvidesProcessArtifacts interface {
		AddProcessArtifacts(target ProcessWorkspace) error
	}

	/*
		For process nodes that can be instantiated.
		Methods on the `builder` argument are used for declaring commands to start processes
	*/
	InstantiableProcess interface {
		AddProcessInstance(target ProcessWorkspace) error
	}
)

/*
Builders used by the above code and artifact generation interfaces
*/
type (
	/*
		A process workspace has commands for adding artifacts to the workspace and
		instantiating processes in a run.sh method.

		Other plugins can extend this workspace with additional methods.  For example,
		the Docker plugin extends the workspace to allow custom Dockerfile build
		commands.
	*/
	ProcessWorkspace interface {
		blueprint.BuildContext

		Info() ProcessWorkspaceInfo

		/*
			Creates a subdirectory in the workspace dir for a process node to collect
			its artifacts.
			Returns a fully qualified path on the local filesystem where artifacts will be
			collected.
		*/
		CreateProcessDir(name string) (string, error)

		/*
			Provides a build script that may be invoked to further collect or build process
			dependencies.
			This will typically be invoked from e.g. within a Container (e.g a Dockerfile),
			rather than on the host machine environment.

			path must refer to a script that resides within a process dir in this workspace;
			if not an error will be returned.

			When it does get invoked, the script will be invoked from the process dir in
			which it resides.
		*/
		AddBuildScript(path string) error

		/*
			A plugin can provide the shell command(s) to run its process.

			Name is just the name of the IRNode representing the process.  Other IRNodes
			that want to instantiate the process will use this name to look it up.

			If the process has dependencies on other IRNodes, they can be provided with
			the deps argument.  The generated code will ensure that the dependencies
			get instantiated first before the runfunc is executed.

			runfunc is a bash function declaration for running the process.
			The runfunc should adhere to the following:
			 - should be defined with syntax like function my_func() { ... }
			 - for any dependencies (config values, addresses, pids, etc.) they can be
			   accessed from environment variable with the corresponding name.  e.g.
			   a.grpc.addr will be in A_GRPC_ADDR.  The mapping from node name to
			   env variable name is implemented by process.EnvVar(name)
			 - the function must set an environment variable for Name with the result
			   of the runfunc.  Typically, this means setting the PID of a started process
			   e.g. MY_GOLANG_PROC=$!
			 - the function must return a return code that will be checked
			 - when it is invoked, the runfunc will be invoked from the root of the
			   proc workspace
			 - the runfunc will be renamed to prevent name clashes between IRNodes
		*/
		DeclareRunCommand(name string, runfunc string, deps ...blueprint.IRNode) error

		/*
			Indicates that we have completed building the workspace, and any finalization tasks
			(e.g. generating build scripts) can run.

			Only the plugin that created the workspace builder should call this method.
		*/
		Finish() error

		ImplementsProcessWorkspace()
	}

	ProcessWorkspaceInfo struct {
		Path   string // fully-qualified path on the filesystem to the workspace
		Target string // the type of workspace being built
	}
)