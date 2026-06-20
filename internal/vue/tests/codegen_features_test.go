package vue_tests

import (
	"testing"

	"github.com/microsoft/typescript-go/shim/fourslash"
	"github.com/microsoft/typescript-go/shim/lsp/lsproto"
)

// A generic SFC's own body must type-check: the `generic="T extends ..."` type
// parameters have to be in scope for defineProps / script usage. Before the fix
// `T` was undeclared, producing TS2304 "Cannot find name 'T'".
func TestSetupGenericOwnFileTypeChecks(t *testing.T) {
	runFourslashTest(t, `// @filename: file.vue
// @strict: true
<script lang="ts" setup generic="T extends string">
const props = defineProps<{ items: T[]; current: T }>()
const first: T = props.items[0]
</script>

<template>
	<div>{{ current }} {{ first }}</div>
</template>`, func(t *testing.T, f *fourslash.FourslashTest, version vueVersion) {
		// Generic SFCs (`generic="..."`) are a Vue 3.3+ feature; the props type
		// shape on 3.2 differs and isn't meaningful to assert here.
		if version == vue_3_2 {
			return
		}
		f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{})
	})
}

// A variable referenced only by a CSS v-bind() must be treated as used (and
// type-checked). With noUnusedLocals, before the fix it was reported unused
// (TS6133) because v-bind expressions weren't emitted.
func TestCssVBindMarksVariableUsed(t *testing.T) {
	runFourslashTest(t, `// @filename: file.vue
// @strict: true
// @noUnusedLocals: true
<script lang="ts" setup>
const themeColor = 'red'
</script>

<template><div>hi</div></template>

<style scoped>
.box { color: v-bind(themeColor); }
</style>`, func(t *testing.T, f *fourslash.FourslashTest, version vueVersion) {
		f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{})
	})
}

// A component with multiple named slots must produce a valid __VLS_Slots
// intersection type. Before the fix a `;` was emitted after each member, which
// terminated the type after the first slot and made the rest a syntax error.
func TestMultipleSlotsNoSyntaxError(t *testing.T) {
	runFourslashTest(t, `// @filename: file.vue
// @strict: true
<script lang="ts" setup>
const title = 'x'
</script>

<template>
	<div>
		<slot name="header" :data="title" />
		<slot name="footer" :count="1" />
		<slot :extra="title" />
	</div>
</template>`, func(t *testing.T, f *fourslash.FourslashTest, version vueVersion) {
		f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{})
	})
}
