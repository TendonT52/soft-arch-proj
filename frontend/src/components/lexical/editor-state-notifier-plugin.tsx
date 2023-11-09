"use client";

import { useEffect } from "react";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import { type EditorState } from "lexical";

type EditorStateNotifierPluginProps = {
  editorState: string | EditorState;
};

const EditorStateNotifierPlugin = ({
  editorState,
}: EditorStateNotifierPluginProps) => {
  const [editor] = useLexicalComposerContext();

  useEffect(() => {
    editor.update(() => {
      if (typeof editorState === "string") {
        const parsed = editor.parseEditorState(editorState);
        editor.setEditorState(parsed);
      } else {
        editor.setEditorState(editorState);
      }
    });
  }, [editor, editorState]);

  return null;
};

export { EditorStateNotifierPlugin };
