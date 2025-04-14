package wasm

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

type Runtime struct {
	registry *registryService
}

func NewRuntime(registry *registryService) *Runtime {
	return &Runtime{
		registry: registry,
	}
}

func (r *Runtime) LoadModule(ctx context.Context, moduleID, version string, config wazero.ModuleConfig) (api.Module, error) {
	// TODO: add compilation cache

	// limit memory to 64 pages (4MB)
	runtime := wazero.NewRuntimeWithConfig(ctx, wazero.NewRuntimeConfigInterpreter().WithMemoryLimitPages(64))
	defer runtime.Close(ctx)

	// Get the module info
	module, err := r.registry.GetModuleInfo(moduleID)
	if err != nil {
		return nil, err
	}

	bin, err := r.registry.GetBinary(ctx, module, version)
	if err != nil {
		return nil, err
	}

	if config == nil {
		config = wazero.NewModuleConfig()
	}

	// Set the module name and disable the start function
	config = config.WithName(module.ID).WithStartFunctions()

	return runtime.InstantiateWithConfig(ctx, bin, config)
}
