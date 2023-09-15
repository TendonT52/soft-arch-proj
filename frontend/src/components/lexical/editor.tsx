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
import { cn } from "@/lib/utils";
import { CodeHighlightPlugin } from "./code-highlight-plugin";
import { ListMaxIndentLevelPlugin } from "./list-max-index-level-plugin";
import { ToolbarPlugin } from "./toolbar-plugin";

const initialConfig: InitialConfigType = {
  namespace: "Editor",
  // The editor theme
  theme: {
    ltr: "text-left",
    rtl: "text-right",
    text: {
      underline: "underline",
      strikethrough: "line-through",
      underlineStrikethrough: "line-through",
    },
    code: "block overflow-x-auto whitespace-pre rounded-sm bg-muted p-3 font-mono text-sm before:content-none after:content-none",
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

export interface EditorProps extends React.ComponentPropsWithoutRef<"div"> {}

const Editor = React.forwardRef<HTMLDivElement, EditorProps>(
  ({ className, ...props }, ref) => {
    const editorStateRef = useRef<EditorState>();

    return (
      <LexicalComposer initialConfig={initialConfig}>
        <div
          ref={ref}
          className={cn(
            "relative flex max-h-[32rem] w-full max-w-3xl flex-col rounded-[8px_8px_0_0] border shadow-md",
            className
          )}
          {...props}
        >
          <ToolbarPlugin className="rounded-[inherit] border-b" />
          <RichTextPlugin
            contentEditable={
              <ContentEditable
                className="prose prose-sky relative min-h-[7.8rem] max-w-none flex-1 overflow-auto bg-background px-8 py-12 focus-visible:outline-none"
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
  }
);
Editor.displayName = "Editor";

export { Editor };
