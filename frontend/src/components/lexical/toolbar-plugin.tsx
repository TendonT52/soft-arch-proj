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
} from "@lexical/rich-text";
import {
  $isAtNodeEnd,
  $isParentElementRTL,
  $wrapNodes,
} from "@lexical/selection";
import { $getNearestNodeOfType, mergeRegister } from "@lexical/utils";
import {
  $createParagraphNode,
  $getNodeByKey,
  $getSelection,
  $isRangeSelection,
  CAN_REDO_COMMAND,
  CAN_UNDO_COMMAND,
  FORMAT_ELEMENT_COMMAND,
  FORMAT_TEXT_COMMAND,
  REDO_COMMAND,
  SELECTION_CHANGE_COMMAND,
  UNDO_COMMAND,
  type GridSelection,
  type LexicalEditor,
  type LexicalNode,
  type NodeSelection,
  type RangeSelection,
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
} from "lucide-react";
import { cn } from "@/lib/utils";
import { Button } from "../ui/button";
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuTrigger,
} from "../ui/dropdown-menu";
import { Input } from "../ui/input";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "../ui/select";
import { Separator } from "../ui/separator";
import { Toggle } from "../ui/toggle";

const LowPriority = 1;

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
  icon: JSX.Element;
};

const supportedBlocks: Record<BlockType, Block> = {
  paragraph: {
    name: "Normal",
    short: "Normal",
    icon: <TextIcon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  h1: {
    name: "Large Heading",
    short: "Large",
    icon: <Heading1Icon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  h2: {
    name: "Small Heading",
    short: "Small",
    icon: <Heading2Icon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  h3: {
    name: "Heading 3",
    short: "Heading 3",
    icon: <Heading3Icon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  h4: {
    name: "Heading 4",
    short: "Heading 4",
    icon: <Heading4Icon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  h5: {
    name: "Heading 5",
    short: "Heading 5",
    icon: <Heading5Icon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  h6: {
    name: "Heading 6",
    short: "Heading 6",
    icon: <Heading6Icon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  ul: {
    name: "Bulleted List",
    short: "Bulleted",
    icon: <ListIcon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  ol: {
    name: "Numbered List",
    short: "Numbered",
    icon: <ListOrderedIcon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  quote: {
    name: "Quote",
    short: "Quote",
    icon: <QuoteIcon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
  code: {
    name: "Code Block",
    short: "Code",
    icon: <CodeIcon className="mr-2 h-5 w-5" strokeWidth={1} />,
  },
};

const positionElement = (
  elem: HTMLElement,
  rect: DOMRect | null | undefined
) => {
  if (!rect) {
    elem.style.opacity = "0";
    elem.style.top = "-10000px";
    elem.style.left = "-10000px";
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
      !nativeSelection?.isCollapsed &&
      rootElement !== null &&
      rootElement.contains(nativeSelection?.anchorNode ?? null)
    ) {
      const domRange = nativeSelection?.getRangeAt(0);
      let rect: DOMRect | undefined;
      if (nativeSelection?.anchorNode === rootElement) {
        let inner = rootElement;
        while (inner.firstElementChild != null) {
          inner = inner.firstElementChild as HTMLElement;
        }
        rect = inner.getBoundingClientRect();
      } else {
        rect = domRange?.getBoundingClientRect();
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
    return () => {};
  }, []);

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
        LowPriority
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
        "prose prose-sky absolute left-[-10000px] top-[-10000px] z-50 inline-flex w-full max-w-xs items-center justify-between rounded-lg bg-white text-sm opacity-0 shadow transition-opacity duration-300",
        isEditMode ? "p-1" : "p-1 pl-4"
      )}
    >
      {isEditMode ? (
        <Input
          ref={inputRef}
          className="focus-visible:ring-0 focus-visible:ring-offset-0"
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
        <React.Fragment>
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
              variant="ghost"
              size="icon"
              tabIndex={0}
              onClick={() => void setEditMode(true)}
            >
              <PencilIcon className="h-5 w-10" strokeWidth={2} />
            </Button>
            <Button
              variant="ghost"
              size="icon"
              tabIndex={0}
              onClick={() => void navigator.clipboard.writeText(linkUrl)}
            >
              <CopyIcon className="h-5 w-10" strokeWidth={2} />
            </Button>
          </div>
        </React.Fragment>
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
        className="w-[10rem] min-w-[10rem] justify-between border-transparent px-3 py-1.5 font-normal"
        asChild
      >
        <SelectTrigger>
          <SelectValue />
        </SelectTrigger>
      </Button>
      <SelectContent
        className="h-[14.5rem] w-[10rem]"
        onCloseAutoFocus={() => void editor.focus()}
      >
        {options.map((option) => (
          <SelectItem key={option} value={option}>
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
  const anchorNode = selection.anchor.getNode() as LexicalNode;
  const focusNode = selection.focus.getNode() as LexicalNode;
  if (anchorNode === focusNode) {
    return anchorNode;
  }
  const isBackward = selection.isBackward();
  if (isBackward) {
    return $isAtNodeEnd(focus) ? anchorNode : focusNode;
  } else {
    return $isAtNodeEnd(anchor) ? focusNode : anchorNode;
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
  const { code, h1, h2, ol, paragraph, quote, ul } = supportedBlocks;
  const current = supportedBlocks[blockType];

  const formatParagraph = () => {
    if (blockType !== "paragraph") {
      editor.update(() => {
        const selection = $getSelection();

        if ($isRangeSelection(selection)) {
          $wrapNodes(selection, () => $createParagraphNode());
        }
      });
    }
  };

  const formatLargeHeading = () => {
    if (blockType !== "h1") {
      editor.update(() => {
        const selection = $getSelection();

        if ($isRangeSelection(selection)) {
          $wrapNodes(selection, () => $createHeadingNode("h1"));
        }
      });
    }
  };

  const formatSmallHeading = () => {
    if (blockType !== "h2") {
      editor.update(() => {
        const selection = $getSelection();

        if ($isRangeSelection(selection)) {
          $wrapNodes(selection, () => $createHeadingNode("h2"));
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

        if ($isRangeSelection(selection)) {
          $wrapNodes(selection, () => $createQuoteNode());
        }
      });
    }
  };

  const formatCode = () => {
    if (blockType !== "code") {
      editor.update(() => {
        const selection = $getSelection();

        if ($isRangeSelection(selection)) {
          $wrapNodes(selection, () => $createCodeNode());
        }
      });
    }
  };

  return (
    <DropdownMenu>
      <DropdownMenuTrigger asChild>
        <Button
          variant="ghost"
          className="w-[10rem] min-w-[10rem] justify-between px-3 py-1.5 font-normal"
        >
          <div className="flex">
            {current.icon}
            {current.short}
          </div>
          <ChevronDownIcon className="h-4 w-4 opacity-50" />
        </Button>
      </DropdownMenuTrigger>
      <DropdownMenuContent
        className="w-[10rem]"
        onCloseAutoFocus={() => void editor.focus()}
      >
        <DropdownMenuItem onSelect={formatParagraph}>
          {paragraph.icon}
          {paragraph.name}
        </DropdownMenuItem>
        <DropdownMenuItem onSelect={formatLargeHeading}>
          {h1.icon}
          {h1.name}
        </DropdownMenuItem>
        <DropdownMenuItem onSelect={formatSmallHeading}>
          {h2.icon}
          {h2.name}
        </DropdownMenuItem>
        <DropdownMenuItem onSelect={formatBulletList}>
          {ul.icon}
          {ul.name}
        </DropdownMenuItem>
        <DropdownMenuItem onSelect={formatNumberedList}>
          {ol.icon}
          {ol.name}
        </DropdownMenuItem>
        <DropdownMenuItem onSelect={formatQuote}>
          {quote.icon}
          {quote.name}
        </DropdownMenuItem>
        <DropdownMenuItem onSelect={formatCode}>
          {code.icon}
          {code.name}
        </DropdownMenuItem>
      </DropdownMenuContent>
    </DropdownMenu>
  );
};

export interface ToolbarPluginProps
  extends React.ComponentPropsWithoutRef<"div"> {}

const ToolbarPlugin = React.forwardRef<HTMLDivElement, ToolbarPluginProps>(
  ({ className, ...props }, ref) => {
    const [editor] = useLexicalComposerContext();
    const [canUndo, setCanUndo] = useState(false);
    const [canRedo, setCanRedo] = useState(false);
    const [blockType, setBlockType] = useState<BlockType>("paragraph");
    const [selectedElementKey, setSelectedElementKey] = useState<string | null>(
      null
    );
    const [, setIsRTL] = useState(false);
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
        const element =
          anchorNode.getKey() === "root"
            ? anchorNode
            : (anchorNode.getTopLevelElementOrThrow() as LexicalNode);
        const elementKey = element.getKey();
        const elementDOM = editor.getElementByKey(elementKey);
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
        // Update text format
        setIsBold(selection.hasFormat("bold"));
        setIsItalic(selection.hasFormat("italic"));
        setIsUnderline(selection.hasFormat("underline"));
        setIsStrikethrough(selection.hasFormat("strikethrough"));
        setIsCode(selection.hasFormat("code"));
        setIsRTL($isParentElementRTL(selection));

        // Update links
        const node = getSelectedNode(selection);
        const parent = node.getParent<LexicalNode>();
        if ($isLinkNode(parent) || $isLinkNode(node)) {
          setIsLink(true);
        } else {
          setIsLink(false);
        }
      }
    }, [editor]);

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
          LowPriority
        ),
        editor.registerCommand(
          CAN_UNDO_COMMAND,
          (payload) => {
            setCanUndo(payload);
            return false;
          },
          LowPriority
        ),
        editor.registerCommand(
          CAN_REDO_COMMAND,
          (payload) => {
            setCanRedo(payload);
            return false;
          },
          LowPriority
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
        ref={ref}
        className={cn(
          "relative flex flex-wrap justify-center gap-1 bg-background p-1",
          className
        )}
        {...props}
      >
        <div className="flex gap-1">
          <Button
            disabled={!canUndo}
            onClick={() => {
              editor.dispatchCommand(UNDO_COMMAND, undefined);
            }}
            variant="ghost"
            size="icon"
            aria-label="Undo"
          >
            <Undo2Icon className="h-5 w-10" strokeWidth={1} />
          </Button>
          <Button
            disabled={!canRedo}
            onClick={() => {
              editor.dispatchCommand(REDO_COMMAND, undefined);
            }}
            variant="ghost"
            size="icon"
            aria-label="Redo"
          >
            <Redo2Icon className="h-5 w-10" strokeWidth={1} />
          </Button>
        </div>
        <Separator orientation="vertical" className="max-lg:hidden" />
        <div className="flex gap-1">
          <BlockOptionsDropdownMenu editor={editor} blockType={blockType} />
        </div>
        <Separator orientation="vertical" className="max-lg:hidden" />

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
          <React.Fragment>
            <div className="flex gap-1">
              <Toggle
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "bold");
                }}
                size="icon"
                pressed={isBold}
                aria-label="Format Bold"
              >
                <BoldIcon className="h-5 w-10" strokeWidth={isBold ? 2 : 1} />
              </Toggle>
              <Toggle
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "italic");
                }}
                size="icon"
                pressed={isItalic}
                aria-label="Format Italics"
              >
                <ItalicIcon
                  className="h-5 w-10"
                  strokeWidth={isItalic ? 2 : 1}
                />
              </Toggle>
              <Toggle
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "underline");
                }}
                size="icon"
                pressed={isUnderline}
                aria-label="Format Underline"
              >
                <UnderlineIcon
                  className="h-5 w-10"
                  strokeWidth={isUnderline ? 2 : 1}
                />
              </Toggle>
              <Toggle
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "strikethrough");
                }}
                size="icon"
                pressed={isStrikethrough}
                aria-label="Format Strikethrough"
              >
                <StrikethroughIcon
                  className="h-5 w-10"
                  strokeWidth={isStrikethrough ? 2 : 1}
                />
              </Toggle>
              <Toggle
                onClick={() => {
                  editor.dispatchCommand(FORMAT_TEXT_COMMAND, "code");
                }}
                size="icon"
                pressed={isCode}
                aria-label="Insert Code"
              >
                <CodeIcon className="h-5 w-10" strokeWidth={isCode ? 2 : 1} />
              </Toggle>
              <Toggle
                onClick={insertLink}
                size="icon"
                pressed={isLink}
                aria-label="Insert Link"
              >
                <LinkIcon className="h-5 w-10" strokeWidth={isLink ? 2 : 1} />
              </Toggle>
              {isLink && <LinkEditor editor={editor} />}
            </div>
            <Separator orientation="vertical" className="max-lg:hidden" />
            <div className="flex gap-1">
              <Button
                onClick={() => {
                  editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, "left");
                }}
                variant="ghost"
                size="icon"
                aria-label="Left Align"
              >
                <AlignLeftIcon className="h-5 w-10" strokeWidth={1} />
              </Button>
              <Button
                onClick={() => {
                  editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, "center");
                }}
                variant="ghost"
                size="icon"
                aria-label="Center Align"
              >
                <AlignCenterIcon className="h-5 w-10" strokeWidth={1} />
              </Button>
              <Button
                onClick={() => {
                  editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, "right");
                }}
                variant="ghost"
                size="icon"
                aria-label="Right Align"
              >
                <AlignRightIcon className="h-5 w-10" strokeWidth={1} />
              </Button>
              <Button
                onClick={() => {
                  editor.dispatchCommand(FORMAT_ELEMENT_COMMAND, "justify");
                }}
                variant="ghost"
                size="icon"
                aria-label="Justify Align"
              >
                <AlignJustifyIcon className="h-5 w-10" strokeWidth={1} />
              </Button>
            </div>
          </React.Fragment>
        )}
      </div>
    );
  }
);
ToolbarPlugin.displayName = "ToolbarPlugin";

export { ToolbarPlugin };
