package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"sort"

	"github.com/decomp/exp/bin"
	_ "github.com/decomp/exp/bin/elf" // register ELF decoder
	_ "github.com/decomp/exp/bin/pe"  // register PE decoder
	"github.com/google/subcommands"
	"github.com/lapsang-boys/pippi/pkg/pi"
	"github.com/lapsang-boys/pippi/pkg/services"
	binpb "github.com/lapsang-boys/pippi/proto/bin"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

var (
	// Default gRPC address to listen on.
	defaultBinAddr = fmt.Sprintf("localhost:%d", services.BinPort)
)

// serverCmd is the command to launch a gRPC server processing parse binary
// requests.
type serverCmd struct {
	// gRPC address to listen on.
	binAddr string
}

func (*serverCmd) Name() string {
	return "server"
}

func (*serverCmd) Synopsis() string {
	return "launch gRPC server"
}

func (*serverCmd) Usage() string {
	const use = `
Reply to binary file parse request.

Usage:
	server [OPTION]...

Flags:
`
	return use[1:]
}

func (cmd *serverCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.binAddr, "addr", defaultBinAddr, "gRPC address to listen on")
}

func (cmd *serverCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	if err := listen(cmd.binAddr); err != nil {
		warn.Printf("listen failed; %+v", err)
		return subcommands.ExitFailure
	}
	return subcommands.ExitSuccess
}

// listen listens on the given gRPC address for incoming requests to parse binary
// files.
func listen(binAddr string) error {
	dbg.Printf("listening on %q", binAddr)
	// Launch gRPC server.
	l, err := net.Listen("tcp", binAddr)
	if err != nil {
		return errors.WithStack(err)
	}
	server := grpc.NewServer()
	// Register binary parser service.
	binpb.RegisterBinaryParserServer(server, &binParserServer{})
	if err := server.Serve(l); err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// binParserServer implements binpb.BinaryParserServer.
type binParserServer struct{}

// ParseBinary parses the given binary file.
func (s *binParserServer) ParseBinary(ctx context.Context, req *binpb.ParseBinaryRequest) (*binpb.File, error) {
	if err := pi.CheckBinID(req.BinId); err != nil {
		return nil, errors.WithStack(err)
	}
	dbg.Printf("parsing ID %q", req.BinId)
	// Parse binary file.
	binPath, err := pi.BinPath(req.BinId)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	file, err := bin.ParseFile(binPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	// Send reply.
	reply := &binpb.File{
		Arch:  archpb(file.Arch),
		Entry: uint64(file.Entry),
	}
	// Sections.
	for _, sect := range file.Sections {
		section := &binpb.Section{
			Name:     sect.Name,
			Addr:     uint64(sect.Addr),
			Offset:   sect.Offset,
			Length:   uint64(len(sect.Data)),
			FileSize: uint64(sect.FileSize),
			MemSize:  uint64(sect.MemSize),
		}
		if sect.Perm&bin.PermR != 0 {
			section.Perms = append(section.Perms, binpb.Perm_R)
		}
		if sect.Perm&bin.PermW != 0 {
			section.Perms = append(section.Perms, binpb.Perm_W)
		}
		if sect.Perm&bin.PermX != 0 {
			section.Perms = append(section.Perms, binpb.Perm_X)
		}
		reply.Sections = append(reply.Sections, section)
	}
	// Imports.
	for funcAddr, funcName := range file.Imports {
		fn := &binpb.Func{
			Addr: uint64(funcAddr),
			Name: funcName,
		}
		reply.Imports = append(reply.Imports, fn)
	}
	sort.Slice(reply.Imports, func(i, j int) bool {
		return reply.Imports[i].Addr < reply.Imports[j].Addr
	})
	// Exports.
	for funcAddr, funcName := range file.Exports {
		fn := &binpb.Func{
			Addr: uint64(funcAddr),
			Name: funcName,
		}
		reply.Exports = append(reply.Exports, fn)
	}
	sort.Slice(reply.Exports, func(i, j int) bool {
		return reply.Exports[i].Addr < reply.Exports[j].Addr
	})
	return reply, nil
}

// archpb converts the given Go machine architecture to protobuf format.
func archpb(arch bin.Arch) binpb.Arch {
	switch arch {
	case bin.ArchX86_32:
		return binpb.Arch_X86_32
	case bin.ArchX86_64:
		return binpb.Arch_X86_64
	case bin.ArchMIPS_32:
		return binpb.Arch_MIPS_32
	case bin.ArchPowerPC_32:
		return binpb.Arch_PowerPC_32
	default:
		panic(fmt.Errorf("support for arch %v not yet implemented", arch))
	}
}
