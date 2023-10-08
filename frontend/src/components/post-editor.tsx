"use client";

import * as React from "react";
import { useRef, useState } from "react";
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
import { DatePickerWithRange } from "./date-range-picker";
import { CodeHighlightPlugin } from "./lexical/code-highlight-plugin";
import { ListMaxIndentLevelPlugin } from "./lexical/list-max-index-level-plugin";
import { ToolbarPlugin } from "./lexical/toolbar-plugin";
import { Card, CardContent, CardDescription, CardHeader } from "./ui/card";
import { Input } from "./ui/input";
import { Label } from "./ui/label";
import { Textarea } from "./ui/textarea";

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

type Post = {
  topic?: string;
  description?: InitialEditorStateType;
};

type EditorProps = {
  post?: Post;
  editable?: boolean;
};

const PostEditor = ({
  post: { topic: title, description: editorState } = {},
  editable = true,
}: EditorProps) => {
  const defaultTitle = "Untitled Post";
  const editorStateRef = useRef<EditorState>();
  const [topic, setTopic] = useState(defaultTitle);

  return (
    <LexicalComposer
      initialConfig={{ ...initialConfig, editorState, editable }}
    >
      <div className="relative mx-auto flex w-full max-w-3xl flex-col items-start">
        {editable && (
          <div className="pointer-events-none sticky top-0 z-10 flex w-full justify-center bg-background pt-4">
            <ToolbarPlugin className="pointer-events-auto rounded-full border bg-background px-4 shadow-sm" />
          </div>
        )}
        {editable ? (
          <TextareaAutosize
            className="flex h-[5.625rem] w-full max-w-none resize-none appearance-none items-end bg-transparent px-8 pb-0.5 pt-12 text-4xl font-extrabold text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]"
            value={title}
            defaultValue={defaultTitle}
            spellCheck={false}
            onChange={(e) => void setTopic(e.target.value)}
            aria-label="title"
          />
        ) : (
          <div className="w-full max-w-none bg-transparent px-8 pb-0.5 pt-12 text-4xl font-extrabold text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]">
            {title}
          </div>
        )}
        <div className="relative w-full">
          <RichTextPlugin
            contentEditable={
              <ContentEditable
                className="prose prose-green relative min-h-[28rem] max-w-none flex-1 bg-background p-8 dark:prose-invert focus:outline-none"
                spellCheck={false}
              />
            }
            placeholder={
              <div className="prose prose-green pointer-events-none absolute left-0 right-0 top-0 max-w-none p-8">
                <p className="text-muted-foreground">
                  Enter post description...
                </p>
              </div>
            }
            ErrorBoundary={LexicalErrorBoundary}
          />
        </div>
        {editable && (
          <div className="w-full px-8 pb-12">
            <Card className="flex w-full flex-col border">
              <CardHeader>
                <h2 className="text-2xl font-bold tracking-tight">
                  Additional information
                </h2>
                <CardDescription>
                  Please provide additional information about the internship
                </CardDescription>
              </CardHeader>
              <CardContent className="flex flex-col gap-6">
                <div className="flex w-full gap-6">
                  <div className="flex flex-1 flex-col gap-2">
                    <Label
                      className="w-full text-sm font-medium leading-none"
                      htmlFor="topic"
                    >
                      Topic
                    </Label>
                    <Input id="topic" value={topic} readOnly />
                  </div>
                  <div className="flex flex-1 flex-col gap-2">
                    <Label
                      htmlFor="period"
                      className="w-full text-sm font-medium leading-none"
                    >
                      Period
                    </Label>
                    <DatePickerWithRange id="period" />
                  </div>
                </div>
                <div className="flex flex-col gap-2">
                  <Label
                    htmlFor="howTo"
                    className="text-sm font-medium leading-none"
                  >
                    How to apply
                  </Label>
                  <Textarea
                    id="howTo"
                    placeholder="Run to the office like the flash ⚡"
                  />
                </div>
              </CardContent>
            </Card>
          </div>
        )}
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
PostEditor.displayName = "PostEditor";

export { PostEditor };
