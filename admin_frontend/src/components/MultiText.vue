<template>
	<editor-content :editor="editor"/>
</template>

<script setup lang="ts">
import {EditorContent, useEditor} from "@tiptap/vue-3";
import StarterKit from "@tiptap/starter-kit";
import {onBeforeUnmount, watch} from "vue";

const makeParagraphs = (lines: string[], divider: string) => {
	if (lines.length === 0) {
		return "";
	}

	// Convert lines to paragraphs
	const paragraphs = lines.map(line => {
		return `<p>${line}</p>`;
	})

	// Join paragraphs
	return paragraphs.join(divider);
};

const props = defineProps<{
	modelValue: string[];
	divider: string;
}>()

const emits = defineEmits<{
	(e: "update:modelValue", value: string[]): void;
}>()

// Helper function to initialize editor
const initEditor = () => {
	return useEditor({
		content: makeParagraphs(props.modelValue, props.divider),
		onUpdate: ({editor}) => {
			emits("update:modelValue", editor.getText({blockSeparator: props.divider}).split(props.divider));
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
}

const editor = initEditor()

watch(() => props.modelValue, (value) => {
	// Compare value with editor value
	const isSame = editor.value?.getText({blockSeparator: props.divider}) === value.join(props.divider);

	if (isSame) {
		return;
	}

	editor.value?.commands.setContent(value || "", false)
})

onBeforeUnmount(() => {
	editor.value?.destroy();
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