import { CodeHighlightNode, CodeNode } from "@lexical/code";
import { AutoLinkNode, LinkNode } from "@lexical/link";
import { ListItemNode, ListNode } from "@lexical/list";
import { type InitialConfigType } from "@lexical/react/LexicalComposer";
import { HeadingNode, QuoteNode } from "@lexical/rich-text";
import { TableCellNode, TableNode, TableRowNode } from "@lexical/table";

export const editorConfig: InitialConfigType = {
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

export const initialEditorState = JSON.stringify({
  root: {
    children: [
      {
        children: [],
        direction: null,
        format: "",
        indent: 0,
        type: "paragraph",
        version: 1,
      },
    ],
    direction: null,
    format: "",
    indent: 0,
    type: "root",
    version: 1,
  },
});
