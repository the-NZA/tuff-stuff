<template>
	<editor-content :editor="editor"/>
</template>

<script setup lang="ts">
import {useEditor, EditorContent} from "@tiptap/vue-3";
import StarterKit from "@tiptap/starter-kit";
import {watch} from "vue";

const makeParagraphs = (lines: string[]) => {
	// Convert lines to paragraphs
	const paragraphs = lines.map(line => {
		return `<p>${line}</p>`;
	})

	// Join paragraphs
	return paragraphs.join("\n");
};

const props = defineProps<{
	modelValue: string[];
}>()

const emits = defineEmits<{
	(e: "update:modelValue", value: string[]): void;
}>()

const editor = useEditor({
	content: makeParagraphs(props.modelValue),
	onUpdate: (value) => {
		emits("update:modelValue", value.editor.getText({blockSeparator: "\n"}).split("\n"));
	},
	editorProps: {
		attributes: {
			class: "multiTextEditor"
		}
	},
	extensions: [
		StarterKit
	]
});

watch(() => props.modelValue, (value) => {
	// Compare value with editor value
	const isSame = editor.value?.getText({blockSeparator: "\n"}) === value.join("\n");

	if (isSame) {
		return;
	}

	editor.value?.commands.setContent(value, false)
})
</script>

<style lang="postcss">
.multiTextEditor:focus-visible {
	outline: none;
}

.multiTextEditor p {
	margin: 0;
}

.multiTextEditor p:not(:last-child) {
	margin-bottom: var(--offset_half);
}

</style>