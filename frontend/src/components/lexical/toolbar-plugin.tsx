"use client";

import * as React from "react";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import Link from "next/link";
import {
  $createCodeNode,
  $isCodeNode,
  getCodeLanguages,
  getDefaultCodeLanguage,
} from "@lexical/code";
import { $isLinkNode, TOGGLE_LINK_COMMAND } from "@lexical/link";
import {
  $isListNode,
  INSERT_ORDERED_LIST_COMMAND,
  INSERT_UNORDERED_LIST_COMMAND,
  ListNode,
  REMOVE_LIST_COMMAND,
} from "@lexical/list";
import { useLexicalComposerContext } from "@lexical/react/LexicalComposerContext";
import {
  $createHeadingNode,
  $createQuoteNode,
  $isHeadingNode,
  type HeadingTagType,
} from "@lexical/rich-text";
import { $isAtNodeEnd, $setBlocksType } from "@lexical/selection";
import {
  $findMatchingParent,
  $getNearestNodeOfType,
  mergeRegister,
} from "@lexical/utils";
import {
  $createParagraphNode,
  $getNodeByKey,
  $getSelection,
  $isElementNode,
  $isRangeSelection,
  $isRootOrShadowRoot,
  CAN_REDO_COMMAND,
  CAN_UNDO_COMMAND,
  COMMAND_PRIORITY_CRITICAL,
  COMMAND_PRIORITY_LOW,
  DEPRECATED_$isGridSelection,
  FORMAT_ELEMENT_COMMAND,
  FORMAT_TEXT_COMMAND,
  REDO_COMMAND,
  SELECTION_CHANGE_COMMAND,
  UNDO_COMMAND,
  type ElementFormatType,
  type ElementNode,
  type GridSelection,
  type LexicalEditor,
  type LexicalNode,
  type NodeSelection,
  type RangeSelection,
  type TextNode,
} from "lexical";
import {
  AlignCenterIcon,
  AlignJustifyIcon,
  AlignLeftIcon,
  AlignRightIcon,
  BoldIcon,
  ChevronDownIcon,
  CodeIcon,
  CopyIcon,
  Heading1Icon,
  Heading2Icon,
  Heading3Icon,
  Heading4Icon,
  Heading5Icon,
  Heading6Icon,
  ItalicIcon,
  LinkIcon,
  ListIcon,
  ListOrderedIcon,
  PencilIcon,
  QuoteIcon,
  Redo2Icon,
  StrikethroughIcon,
  TextIcon,
  UnderlineIcon,
  Undo2Icon,
  type LucideIcon,
} from "lucide-react";
import { cn } from "@/lib/utils";
import { Button } from "../ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { Separator } from "../ui/separator";
import { Toggle } from "../ui/toggle";

const blockTypes = [
  "paragraph",
  "h1",
  "h2",
  "h3",
  "h4",
  "h5",
  "h6",
  "ul",
  "ol",
  "quote",
  "code",
] as const;

type BlockType = (typeof blockTypes)[number];

type Block = {
  name: string;
  short: string;
  Icon: LucideIcon;
};

const supportedBlocks: Record<BlockType, Block> = {
  paragraph: {
    name: "Normal",
    short: "Normal",
    Icon: TextIcon,
  },
  h1: {
    name: "Heading 1",
    short: "Heading 1",
    Icon: Heading1Icon,
  },
  h2: {
    name: "Heading 2",
    short: "Heading 2",
    Icon: Heading2Icon,
  },
  h3: {
    name: "Heading 3",
    short: "Heading 3",
    Icon: Heading3Icon,
  },
  h4: {
    name: "Heading 4",
    short: "Heading 4",
    Icon: Heading4Icon,
  },
  h5: {
    name: "Heading 5",
    short: "Heading 5",
    Icon: Heading5Icon,
  },
  h6: {
    name: "Heading 6",
    short: "Heading 6",
    Icon: Heading6Icon,
  },
  ul: {
    name: "Bulleted List",
    short: "Bulleted",
    Icon: ListIcon,
  },
  ol: {
    name: "Numbered List",
    short: "Numbered",
    Icon: ListOrderedIcon,
  },
  quote: {
    name: "Quote",
    short: "Quote",
    Icon: QuoteIcon,
  },
  code: {
    name: "Code Block",
    short: "Code",
    Icon: CodeIcon,
  },
};

const positionElement = (
  elem: HTMLElement,
  rect: DOMRect | null | undefined
) => {
  if (!rect) {
    elem.style.opacity = "0";
    elem.style.top = "-100000px";
    elem.style.left = "-100000px";
  } else {
    elem.style.opacity = "1";
    elem.style.top = "0";
    elem.style.left = "0";

    const { top, left } = elem.getBoundingClientRect();

    elem.style.top = `calc(${-top + rect.bottom}px + 0.5rem)`;
    elem.style.left = `${
      -left + rect.left + rect.width / 2 - elem.offsetWidth / 2
    }px`;
  }
};

type LinkEditorProps = {
  editor: LexicalEditor;
};

const LinkEditor = ({ editor }: LinkEditorProps) => {
  const ref = useRef<HTMLDivElement | null>(null);
  const inputRef = useRef<HTMLInputElement | null>(null);
  const mouseDownRef = useRef(false);
  const [linkUrl, setLinkUrl] = useState("");
  const [isEditMode, setEditMode] = useState(false);
  const [lastSelection, setLastSelection] = useState<
    RangeSelection | NodeSelection | GridSelection | null
  >(null);

  const updateLinkEditor = useCallback(() => {
    const selection = $getSelection();
    if ($isRangeSelection(selection)) {
      const node = getSelectedNode(selection);
      const parent = node.getParent<LexicalNode>();
      if ($isLinkNode(parent)) {
        setLinkUrl(parent.getURL());
      } else if ($isLinkNode(node)) {
        setLinkUrl(node.getURL());
      } else {
        setLinkUrl("");
      }
    }
    const elem = ref.current;
    const nativeSelection = window.getSelection();

    if (elem === null) {
      return;
    }

    const rootElement = editor.getRootElement();
    if (
      selection !== null &&
      nativeSelection !== null &&
      rootElement !== null &&
      rootElement.contains(nativeSelection.anchorNode) &&
      editor.isEditable()
    ) {
      const domRange = nativeSelection.getRangeAt(0);
      let rect: DOMRect | undefined;
      if (nativeSelection.anchorNode === rootElement) {
        let inner = rootElement;
        while (inner.firstElementChild !== null) {
          inner = inner.firstElementChild as HTMLElement;
        }
        rect = inner.getBoundingClientRect();
      } else {
        rect = domRange.getBoundingClientRect();
      }

      if (!mouseDownRef.current) {
        positionElement(elem, rect);
      }
      setLastSelection(selection);
    } else {
      positionElement(elem, null);
      setLastSelection(null);
      setEditMode(false);
      setLinkUrl("");
    }

    return true;
  }, [editor]);

  useEffect(() => {
    return mergeRegister(
      editor.registerUpdateListener(({ editorState }) => {
        editorState.read(() => {
          updateLinkEditor();
        });
      }),

      editor.registerCommand(
        SELECTION_CHANGE_COMMAND,
        () => {
          updateLinkEditor();
          return true;
        },
        COMMAND_PRIORITY_LOW
      )
    );
  }, [editor, updateLinkEditor]);

  useEffect(() => {
    editor.getEditorState().read(() => {
      updateLinkEditor();
    });
  }, [editor, updateLinkEditor]);

  useEffect(() => {
    if (isEditMode && inputRef.current) {
      inputRef.current.focus();
    }
  }, [isEditMode]);

  return (
    <div
      ref={ref}
      className={cn(
        "prose prose-green absolute left-[-10000px] top-[-10000px] z-50 inline-flex w-full max-w-xs items-center justify-between gap-1 rounded-lg border bg-background text-sm text-foreground opacity-0 shadow transition-opacity duration-300",
        isEditMode ? "p-1" : "p-1 pl-4"
      )}
    >
      {isEditMode ? (
        <input
          ref={inputRef}
          className="h-8 w-full bg-background px-3 py-2 font-medium focus-visible:outline-none"
          value={linkUrl}
          onChange={(e) => {
            setLinkUrl(e.target.value);
          }}
          onKeyDown={(e) => {
            if (e.code === "Enter") {
              e.preventDefault();
              if (lastSelection !== null) {
                if (linkUrl !== "") {
                  editor.dispatchCommand(TOGGLE_LINK_COMMAND, linkUrl);
                }
                setEditMode(false);
              }
            } else if (e.code === "Escape") {
              e.preventDefault();
              setEditMode(false);
            }
          }}
        />
      ) : (
        <>
          <Link
            className="truncate"
            href={linkUrl}
            target="_blank"
            rel="noopener noreferrer"
          >
            {linkUrl}
          </Link>
          <div className="flex">
            <Button
              className="h-8 w-8 p-0"
              variant="ghost"
              tabIndex={0}
              onClick={() => void setEditMode(true)}
            >
              <PencilIcon className="h-4 w-8" strokeWidth={2} />
            </Button>
            <Button
              className="h-8 w-8 p-0"
              variant="ghost"
              tabIndex={0}
              onClick={() => void navigator.clipboard.writeText(linkUrl)}
            >
              <CopyIcon className="h-4 w-8" strokeWidth={2} />
            </Button>
          </div>
        </>
      )}
    </div>
  );
};

type LanguageOptionsSelectProps = {
  editor: LexicalEditor;
  options: string[];
  defaultValue: string;
  onValueChange: (value: string) => void;
};

const LanguageOptionsSelect = ({
  editor,
  options,
  defaultValue,
  onValueChange,
}: LanguageOptionsSelectProps) => {
  return (
    <Select defaultValue={defaultValue} onValueChange={onValueChange}>
      <Button
        variant="ghost"
        className="h-8 w-[9.5rem] justify-between border-transparent px-3 py-1.5 text-xs font-normal"
        asChild
      >
        <SelectTrigger>
          <SelectValue />
        </SelectTrigger>
      </Button>
      <SelectContent
        className="h-[12.875rem] w-[9.5rem]"
        onCloseAutoFocus={() => void editor.focus()}
      >
        {options.map((option) => (
          <SelectItem className="text-xs" key={option} value={option}>
            {option}
          </SelectItem>
        ))}
      </SelectContent>
    </Select>
  );
};

const getSelectedNode = (selection: RangeSelection) => {
  const anchor = selection.anchor;
  const focus = selection.focus;
  const anchorNode = selection.anchor.getNode() as TextNode | ElementNode;
  const focusNode = selection.focus.getNode() as TextNode | ElementNode;
  if (anchorNode === focusNode) {
    return anchorNode;
  }
  const isBackward = selection.isBackward();
  if (isBackward) {
    return $isAtNodeEnd(focus) ? anchorNode : focusNode;
  } else {
    return $isAtNodeEnd(anchor) ? anchorNode : focusNode;
  }
};

type BlockOptionsDropdownMenuProps = {
  editor: LexicalEditor;
  blockType: BlockType;
};

const BlockOptionsDropdownMenu = ({
  editor,
  blockType,
}: BlockOptionsDropdownMenuProps) => {
  const { code, h2, h3, ol, paragraph, quote, ul } = supportedBlocks;
  const current = supportedBlocks[blockType];

  const formatParagraph = () => {
    editor.update(() => {
      const selection = $getSelection();
      if (
        $isRangeSelection(selection) ||
        DEPRECATED_$isGridSelection(selection)
      ) {
        $setBlocksType(selection, () => $createParagraphNode());
      }
    });
  };

  const formatHeading = (headingSize: HeadingTagType) => {
    if (blockType !== headingSize) {
      editor.update(() => {
        const selection = $getSelection();
        if (
          $isRangeSelection(selection) ||
          DEPRECATED_$isGridSelection(selection)
        ) {
          $setBlocksType(selection, () => $createHeadingNode(headingSize));
        }
      });
    }
  };

  const formatBulletList = () => {
    if (blockType !== "ul") {
      editor.dispatchCommand(INSERT_UNORDERED_LIST_COMMAND, undefined);
    } else {
      editor.dispatchCommand(REMOVE_LIST_COMMAND, undefined);
    }
  };

  const formatNumberedList = () => {
    if (blockType !== "ol") {
      editor.dispatchCommand(INSERT_ORDERED_LIST_COMMAND, undefined);
    } else {
      editor.dispatchCommand(REMOVE_LIST_COMMAND, undefined);
    }
  };

  const formatQuote = () => {
    if (blockType !== "quote") {
      editor.update(() => {
        const selection = $getSelection();
        if (
          $isRangeSelection(selection) ||
          DEPRECATED_$isGridSelection(selection)
        ) {
          $setBlocksType(selection, () => $createQuoteNode());
        }
      });
    }
  };

  const formatCode = () => {
    if (blockType !== "code") {
      editor.update(() => {
        let selection = $getSelection();

        if (
          $isRangeSelection(selection) ||
          DEPRECATED_$isGridSelection(selection)
        ) {
          if (selection.isCollapsed()) {
            $setBlocksType(selection, () => $createCodeNode());
          } else {
            const textContent = selection.getTextContent();
            const codeNode = $createCodeNode();
            selection.insertNodes([codeNode]);
            selection = $getSelection();
            if ($isRangeSelection(selection))
              selection.insertRawText(textContent);
          }
        }
      });
    }
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="ghost"
          className="h-8 w-[9.5rem] items-center justify-between px-3 py-1 text-xs font-normal"
        >
          <div className="flex items-center">
            <current.Icon className="mr-2 h-4 w-4" strokeWidth={1} />
            {current.short}
          </div>
          <ChevronDownIcon className="h-4 w-4 opacity-50" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent
        className="h-[12.875rem] w-[9.5rem]"
        onCloseAutoFocus={() => void editor.focus()}
      >
        <DropdownMenuItem className="text-xs" onSelect={formatParagraph}>
          <paragraph.Icon className="mr-2 h-4 w-4" strokeWidth={1} />
          {paragraph.name}
        </DropdownMenuItem>
        <DropdownMenuItem
          className="text-xs"
          onSelect={() => void formatHeading("h2")}
        >
          <h2.Icon className="mr-2 h-4 w-4" strokeWidth={1} />
          {h2.name}
        </DropdownMenuItem>
        <DropdownMenuItem
          className="text-xs"
          onSelect={() => void formatHeading("h3")}
        >
          <h3.Icon className="mr-2 h-4 w-4" strokeWidth={1} />
          {h3.name}
        </DropdownMenuItem>
        <DropdownMenuItem className="text-xs" onSelect={formatBulletList}>
          <ul.Icon className="mr-2 h-4 w-4" strokeWidth={1} />
          {ul.name}
        </DropdownMenuItem>
        <DropdownMenuItem className="text-xs" onSelect={formatNumberedList}>
          <ol.Icon className="mr-2 h-4 w-4" strokeWidth={1} />
          {ol.name}
        </DropdownMenuItem>
        <DropdownMenuItem className="text-xs" onSelect={formatQuote}>
          <quote.Icon className="mr-2 h-4 w-4" strokeWidth={1} />
          {quote.name}
        </DropdownMenuItem>
        <DropdownMenuItem className="text-xs" onSelect={formatCode}>
          <code.Icon className="mr-2 h-4 w-4" strokeWidth={1} />
          {code.name}
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

type ToolbarPluginProps = React.HTMLAttributes<HTMLDivElement>;

const ToolbarPlugin = React.forwardRef<HTMLDivElement, ToolbarPluginProps>(
  ({ className, ...props }, ref) => {
    const [editor] = useLexicalComposerContext();
    const [canUndo, setCanUndo] = useState(false);
    const [canRedo, setCanRedo] = useState(false);
    const [blockType, setBlockType] = useState<BlockType>("paragraph");
    const [selectedElementKey, setSelectedElementKey] = useState<string | null>(
      null
    );
    const [elementFormat, setElementFormat] =
      useState<ElementFormatType>("start");
    const [isLink, setIsLink] = useState(false);
    const [isBold, setIsBold] = useState(false);
    const [isItalic, setIsItalic] = useState(false);
    const [isUnderline, setIsUnderline] = useState(false);
    const [isStrikethrough, setIsStrikethrough] = useState(false);
    const [isCode, setIsCode] = useState(false);

    const updateToolbar = useCallback(() => {
      const selection = $getSelection();
      if ($isRangeSelection(selection)) {
        const anchorNode = selection.anchor.getNode() as LexicalNode;
        let element =
          anchorNode.getKey() === "root"
            ? anchorNode
            : $findMatchingParent(anchorNode, (e) => {
                const parent = e.getParent<LexicalNode>();
                return parent !== null && $isRootOrShadowRoot(parent);
              });

        if (element === null) {
          element = anchorNode.getTopLevelElementOrThrow() as LexicalNode;
        }

        const elementKey = element.getKey();
        const elementDOM = editor.getElementByKey(elementKey);

        // Update text format
        setIsBold(selection.hasFormat("bold"));
        setIsItalic(selection.hasFormat("italic"));
        setIsUnderline(selection.hasFormat("underline"));
        setIsStrikethrough(selection.hasFormat("strikethrough"));
        setIsCode(selection.hasFormat("code"));

        const node = getSelectedNode(selection);
        const parent = node.getParent<ElementNode>();

        // Handle buttons
        setElementFormat(
          ($isElementNode(node)
            ? node.getFormatType()
            : parent?.getFormatType()) || "start"
        );

        // Update links
        if ($isLinkNode(parent) || $isLinkNode(node)) {
          setElementFormat("");
          setIsLink(true);
        } else {
          setIsLink(false);
        }

        if (elementDOM !== null) {
          setSelectedElementKey(elementKey);
          if ($isListNode(element)) {
            const parentList = $getNearestNodeOfType(anchorNode, ListNode);
            const type = parentList ? parentList.getTag() : element.getTag();
            setBlockType(type);
          } else {
            const type = (
              $isHeadingNode(element) ? element.getTag() : element.getType()
            ) as BlockType;
            setBlockType(type);
          }
        }
      }
    }, [editor]);

    useEffect(() => {
      return editor.registerCommand(
        SELECTION_CHANGE_COMMAND,
        () => {
          updateToolbar();
          return false;
        },
        COMMAND_PRIORITY_CRITICAL
      );
    }, [editor, updateToolbar]);

    useEffect(() => {
      return mergeRegister(
        editor.registerUpdateListener(({ editorState }) => {
          editorState.read(() => {
            updateToolbar();
          });
        }),
        editor.registerCommand(
          SELECTION_CHANGE_COMMAND,
          () => {
            updateToolbar();
            return false;
          },
          COMMAND_PRIORITY_CRITICAL
        ),
        editor.registerCommand(
          CAN_UNDO_COMMAND,
          (payload) => {
            setCanUndo(payload);
            return false;
          },
          COMMAND_PRIORITY_CRITICAL
        ),
        editor.registerCommand(
          CAN_REDO_COMMAND,
          (payload) => {
            setCanRedo(payload);
            return false;
          },
          COMMAND_PRIORITY_CRITICAL
        )
      );
    }, [editor, updateToolbar]);

    const codeLanguges = useMemo(() => getCodeLanguages(), []);
    const onCodeLanguageSelect = useCallback(
      (value: string) => {
        editor.update(() => {
          if (selectedElementKey !== null) {
            const node = $getNodeByKey(selectedElementKey);
            if ($isCodeNode(node)) {
              node.setLanguage(value);
            }
          }
        });
      },
      [editor, selectedElementKey]
    );

    const insertLink = useCallback(() => {
      if (!isLink) {
        editor.dispatchCommand(TOGGLE_LINK_COMMAND, "https://");
      } else {
        editor.dispatchCommand(TOGGLE_LINK_COMMAND, null);
      }
    }, [editor, isLink]);

    return (
      <div
        className={cn(
          "relative flex flex-wrap justify-center gap-1 rounded-[inherit] bg-background p-1",
          className
        )}
        ref={ref}
        {...props}
      >
        <div className="flex gap-1">
          <Button
            className="h-8 w-8 p-0"
            disabled={!canUndo}
            onClick={() => {
              editor.dispatchCommand(UNDO_COMMAND, undefined);
            }}
            variant="ghost"
            aria-label="Undo"
          >
            <Undo2Icon className="h-4 w-8" strokeWidth={1} />
          </Button>
          <Button
            className="h-8 w-8 p-0"
            disabled={!canRedo}
            onClick={() => {
              editor.dispatchCommand(REDO_COMMAND, undefined);
            }}
            variant="ghost"
            aria-label="Redo"
          >
            <Redo2Icon className="h-4 w-8" strokeWidth={1} />
          </Button>
        </div>
        <Separator orientation="vertical" className="hidden md:block" />
        <div className="flex gap-1">
          <BlockOptionsDropdownMenu editor={editor} blockType={blockType} />
        </div>
        <Separator orientation="vertical" className="hidden md:block" />
        {blockType === "code" ? (
          <div className="flex gap-1">
            <LanguageOptionsSelect
              editor={editor}
              options={codeLanguges}
              defaultValue={getDefaultCodeLanguage()}
              onValueChange={onCodeLanguageSelect}
            />
          </div>
        ) : (
          <>
            <div className="flex gap-1">
              <Toggle
                className="h-8 w-8 p-0"
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "bold");
                }}
                pressed={isBold}
                aria-label="Format Bold"
              >
                <BoldIcon className="h-4 w-8" strokeWidth={isBold ? 2 : 1} />
              </Toggle>
              <Toggle
                className="h-8 w-8 p-0"
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "italic");
                }}
                pressed={isItalic}
                aria-label="Format Italics"
              >
                <ItalicIcon
                  className="h-4 w-8"
                  strokeWidth={isItalic ? 2 : 1}
                />
              </Toggle>
              <Toggle
                className="h-8 w-8 p-0"
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "underline");
                }}
                pressed={isUnderline}
                aria-label="Format Underline"
              >
                <UnderlineIcon
                  className="h-4 w-8"
                  strokeWidth={isUnderline ? 2 : 1}
                />
              </Toggle>
              <Toggle
                className="h-8 w-8 p-0"
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "strikethrough");
                }}
                pressed={isStrikethrough}
                aria-label="Format Strikethrough"
              >
                <StrikethroughIcon
                  className="h-4 w-8"
                  strokeWidth={isStrikethrough ? 2 : 1}
                />
              </Toggle>
              <Toggle
                className="h-8 w-8 p-0"
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "code");
                }}
                pressed={isCode}
                aria-label="Insert Code"
              >
                <CodeIcon className="h-4 w-8" strokeWidth={isCode ? 2 : 1} />
              </Toggle>
              <Toggle
                className="h-8 w-8 p-0"
                onClick={insertLink}
                pressed={isLink}
                aria-label="Insert Link"
              >
                <LinkIcon className="h-4 w-8" strokeWidth={isLink ? 2 : 1} />
              </Toggle>
              {isLink && <LinkEditor editor={editor} />}
            </div>
            <Separator orientation="vertical" className="hidden md:block" />
            <div className="flex gap-1">
              <Toggle
                className="h-8 w-8 p-0"
                onClick={() => {
                  editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, "start");
                }}
                pressed={elementFormat === "start"}
                aria-label="Left Align"
              >
                <AlignLeftIcon
                  className="h-4 w-8"
                  strokeWidth={elementFormat === "start" ? 2 : 1}
                />
              </Toggle>
              <Toggle
                className="h-8 w-8 p-0"
                onClick={() => {
                  editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, "center");
                }}
                pressed={elementFormat === "center"}
                aria-label="Center Align"
              >
                <AlignCenterIcon
                  className="h-4 w-8"
                  strokeWidth={elementFormat === "center" ? 2 : 1}
                />
              </Toggle>
              <Toggle
                className="h-8 w-8 p-0"
                onClick={() => {
                  editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, "end");
                }}
                pressed={elementFormat === "end"}
                aria-label="Right Align"
              >
                <AlignRightIcon
                  className="h-4 w-8"
                  strokeWidth={elementFormat === "end" ? 2 : 1}
                />
              </Toggle>
              <Toggle
                className="h-8 w-8 p-0"
                onClick={() => {
                  editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, "justify");
                }}
                pressed={elementFormat === "justify"}
                aria-label="Justify Align"
              >
                <AlignJustifyIcon
                  className="h-4 w-8"
                  strokeWidth={elementFormat === "justify" ? 2 : 1}
                />
              </Toggle>
            </div>
          </>
        )}
      </div>
    );
  }
);
ToolbarPlugin.displayName = "ToolbarPlugin";

export { ToolbarPlugin };
