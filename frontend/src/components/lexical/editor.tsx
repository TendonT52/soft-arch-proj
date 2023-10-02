"use client";

import * as React from "react";
import { useRef } from "react";
import { CodeHighlightNode, CodeNode } from "@lexical/code";
import { AutoLinkNode, LinkNode } from "@lexical/link";
import { ListItemNode, ListNode } from "@lexical/list";
import { TRANSFORMERS } from "@lexical/markdown";
import { AutoFocusPlugin } from "@lexical/react/LexicalAutoFocusPlugin";
import {
  LexicalComposer,
  type InitialConfigType,
  type InitialEditorStateType,
} from "@lexical/react/LexicalComposer";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import LexicalErrorBoundary from "@lexical/react/LexicalErrorBoundary";
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin";
import { LinkPlugin } from "@lexical/react/LexicalLinkPlugin";
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { MarkdownShortcutPlugin } from "@lexical/react/LexicalMarkdownShortcutPlugin";
import { OnChangePlugin } from "@lexical/react/LexicalOnChangePlugin";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import { HeadingNode, QuoteNode } from "@lexical/rich-text";
import { TableCellNode, TableNode, TableRowNode } from "@lexical/table";
import { type EditorState } from "lexical";
import TextareaAutosize from "react-textarea-autosize";
import { cn } from "@/lib/utils";
import { CodeHighlightPlugin } from "./code-highlight-plugin";
import { ListMaxIndentLevelPlugin } from "./list-max-index-level-plugin";
import { ToolbarPlugin } from "./toolbar-plugin";

const initialConfig: InitialConfigType = {
  namespace: "Editor",
  // The editor theme
  theme: {
    text: {
      underline: "underline",
      strikethrough: "line-through",
      italic: "italic",
    },
    code: "block overflow-x-auto whitespace-pre rounded-sm bg-muted p-3 font-mono text-sm scrollbar-hide before:content-none after:content-none",
    quote: "font-normal not-italic text-muted-foreground",
    codeHighlight: {
      atrule: "text-code-attribute",
      attr: "text-code-attribute",
      boolean: "text-code-property",
      builtin: "text-code-selector",
      cdata: "text-code-comment",
      char: "text-code-selector",
      class: "text-code-function",
      "class-name": "text-code-function",
      comment: "text-code-comment",
      constant: "text-code-property",
      deleted: "text-code-property",
      doctype: "text-code-comment",
      entity: "text-code-operator",
      function: "text-code-function",
      important: "text-code-variable",
      inserted: "text-code-selector",
      keyword: "text-code-attribute",
      namespace: "text-code-variable",
      number: "text-code-property",
      operator: "text-code-operator",
      prolog: "text-code-comment",
      property: "text-code-property",
      punctuation: "text-code-punctuation",
      regex: "text-code-variable",
      selector: "text-code-selector",
      string: "text-code-selector",
      symbol: "text-code-property",
      tag: "text-code-property",
      url: "text-code-operator",
      variable: "text-code-variable",
    },
  },
  // Handling of errors during update
  onError: (error: Error) => {
    throw error;
  },
  // Any custom nodes go here
  nodes: [
    HeadingNode,
    ListNode,
    ListItemNode,
    QuoteNode,
    CodeNode,
    CodeHighlightNode,
    TableNode,
    TableCellNode,
    TableRowNode,
    AutoLinkNode,
    LinkNode,
  ],
};

type EditorProps = {
  title?: string;
  onTitleChange?: (title: string) => void;
  defaultTitle?: string;
  editable?: boolean;
  editorState?: InitialEditorStateType;
};

const Editor = ({
  title,
  onTitleChange,
  defaultTitle = "Untitled Post",
  editable = true,
  editorState,
}: EditorProps) => {
  const editorStateRef = useRef<EditorState>();

  return (
    <LexicalComposer
      initialConfig={{ ...initialConfig, editorState, editable }}
    >
      <div
        className={cn(
          "relative mx-auto flex min-h-[48rem] w-full max-w-3xl flex-col rounded-t-lg",
          editable && "border shadow"
        )}
      >
        {editable && (
          <div className="sticky top-0 z-10 rounded-t-lg border-b">
            <ToolbarPlugin />
          </div>
        )}
        {editable ? (
          <TextareaAutosize
            className="min-h-0 max-w-none resize-none appearance-none bg-transparent px-8 pt-12 text-4xl font-extrabold text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]"
            value={title}
            defaultValue={defaultTitle}
            spellCheck={false}
            onChange={(e) => {
              if (onTitleChange) {
                onTitleChange(e.target.value);
              }
            }}
            aria-label="title"
          />
        ) : (
          <div className="min-h-0 max-w-none resize-none appearance-none bg-transparent px-8 pt-12 text-4xl font-extrabold text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]">
            {title}
          </div>
        )}
        <RichTextPlugin
          contentEditable={
            <ContentEditable
              className="prose prose-green relative min-h-[6.8rem] max-w-none flex-1 overflow-auto bg-background px-8 pb-12 pt-8 dark:prose-invert focus:outline-none"
              spellCheck={false}
            />
          }
          placeholder={null}
          ErrorBoundary={LexicalErrorBoundary}
        />
      </div>
      <HistoryPlugin />
      <AutoFocusPlugin />
      <CodeHighlightPlugin />
      <ListPlugin />
      <LinkPlugin />
      <ListMaxIndentLevelPlugin maxDepth={1} />
      <MarkdownShortcutPlugin transformers={TRANSFORMERS} />
      <OnChangePlugin
        onChange={(editorState) => {
          editorStateRef.current = editorState;
        }}
      />
    </LexicalComposer>
  );
};
Editor.displayName = "Editor";

export { Editor };
