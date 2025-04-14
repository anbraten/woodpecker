package wasm_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.woodpecker-ci.org/woodpecker/v3/server/wasm"
)

func TestWasm(t *testing.T) {
	ctx := context.Background()

	reg, err := wasm.NewRegistryService(wasm.RegistryConfig{
		RegistryURLs: []string{
			"https://gist.githubusercontent.com/anbraten/7eadb77c1119bf02d08215132891df07/raw/c43f514f2366cb03161c4d8f212e1773c8a49f7d/registry.json",
		},
	})
	require.NoError(t, err)

	err = reg.UpdateIndex(context.Background(), true)
	require.NoError(t, err)

	module, err := reg.GetModuleInfo("compiler.lua")
	require.NoError(t, err)

	require.Equal(t, "compiler.lua", module.ID)
	require.NotEmpty(t, module.Versions)

	bin, err := reg.GetBinary(ctx, module, module.Tags["latest"])
	require.NoError(t, err)
	require.NotEmpty(t, bin)

	runtime := wasm.NewRuntime(reg)
	require.NotNil(t, runtime)

	mod, err := runtime.LoadModule(ctx, module.ID, module.Tags["latest"], nil)
	require.NoError(t, err)
	require.NotNil(t, mod)

	// Test the module
	arg1 := "123"
	arg2 := "456"

	_, err = mod.ExportedFunction("compile").Call(ctx, arg1, arg2)
	require.NoError(t, err)

}
