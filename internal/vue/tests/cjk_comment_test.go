package vue_tests

import (
	"testing"

	"github.com/microsoft/typescript-go/shim/fourslash"
	"github.com/microsoft/typescript-go/shim/lsp/lsproto"
)

// Multi-byte (e.g. CJK) characters inside an HTML comment must not break the
// tokenizer. Before the fix the comment-end sequence `-->` was skipped because
// the parse loop advanced by the decoded rune size after fastForwardTo had
// already repositioned the index onto an ASCII delimiter, so the comment never
// closed — surfacing as "EOF in comment" + cascading "Missing end tag".
func TestCJKCommentDoesNotBreakParser(t *testing.T) {
	runFourslashTest(t, `// @filename: file.vue
<script setup lang="ts"></script>

<template>
	<!-- 层叠容器：图片与角标都放进同一个网格单元 -->
	<div>ok</div>
</template>`, func(t *testing.T, f *fourslash.FourslashTest, version vueVersion) {
		f.VerifyNonSuggestionDiagnostics(t, []*lsproto.Diagnostic{})
	})
}
