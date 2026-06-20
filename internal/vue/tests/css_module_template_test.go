package vue_tests

import (
	"testing"

	"github.com/microsoft/typescript-go/shim/fourslash"
	"github.com/microsoft/typescript-go/shim/lsp/lsproto"
)

// A named CSS module (<style module="css">) must expose its binding (`css`) to
// the template, just like Volar does. Before the fix the template reference
// `css.card` produced TS2339 "Property 'css' does not exist".
func TestCssModuleNamedBindingInTemplate(t *testing.T) {
	runFourslashTest(t, `// @filename: file.vue
<script setup lang="ts"></script>

<template>
	<div :class="css.card">x</div>
</template>

<style module="css">
.card { color: red; }
</style>`, func(t *testing.T, f *fourslash.FourslashTest, version vueVersion) {
		f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{})
	})
}

// A bare <style module> exposes the default `$style` binding to the template.
func TestCssModuleDefaultBindingInTemplate(t *testing.T) {
	runFourslashTest(t, `// @filename: file.vue
<script setup lang="ts"></script>

<template>
	<div :class="$style.card">x</div>
</template>

<style module>
.card { color: red; }
</style>`, func(t *testing.T, f *fourslash.FourslashTest, version vueVersion) {
		f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{})
	})
}
