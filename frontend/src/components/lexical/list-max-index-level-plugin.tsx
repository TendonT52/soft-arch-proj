"use client";

import { useEffect } from "react";
import { $getListDepth, $isListItemNode, $isListNode } from "@lexical/list";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import {
  $getSelection,
  $isElementNode,
  $isRangeSelection,
  COMMAND_PRIORITY_HIGH,
  INDENT_CONTENT_COMMAND,
  type LexicalNode,
  type RangeSelection,
} from "lexical";

const getElementNodesInSelection = (
  selection: RangeSelection
): Set<LexicalNode> => {
  const nodesInSelection = selection.getNodes();

  if (nodesInSelection.length === 0) {
    return new Set<LexicalNode>([
      (selection.anchor.getNode() as LexicalNode).getParentOrThrow(),
      (selection.focus.getNode() as LexicalNode).getParentOrThrow(),
    ]);
  }

  return new Set<LexicalNode>(
    nodesInSelection.map((n) => ($isElementNode(n) ? n : n.getParentOrThrow()))
  );
};

const isIndentPermitted = (maxDepth: number) => {
  const selection = $getSelection();

  if (!$isRangeSelection(selection)) {
    return false;
  }

  const elementNodesInSelection = Array.from(
    getElementNodesInSelection(selection)
  );

  let totalDepth = 0;

  for (const elementNode of elementNodesInSelection) {
    if ($isListNode(elementNode)) {
      totalDepth = Math.max($getListDepth(elementNode) + 1, totalDepth);
    } else if ($isListItemNode(elementNode)) {
      const parent = elementNode.getParent<LexicalNode>();
      if (!$isListNode(parent)) {
        throw new Error(
          "ListMaxIndentLevelPlugin: A ListItemNode must have a ListNode for a parent."
        );
      }
      totalDepth = Math.max($getListDepth(parent) + 1, totalDepth);
    }
  }

  return totalDepth <= maxDepth;
};

type ListMaxIndentLevelPluginProps = {
  maxDepth?: number;
};

const ListMaxIndentLevelPlugin = ({
  maxDepth = 7,
}: ListMaxIndentLevelPluginProps) => {
  const [editor] = useLexicalComposerContext();

  useEffect(() => {
    return editor.registerCommand(
      INDENT_CONTENT_COMMAND,
      () => !isIndentPermitted(maxDepth),
      COMMAND_PRIORITY_HIGH
    );
  }, [editor, maxDepth]);

  return null;
};

export { ListMaxIndentLevelPlugin };
