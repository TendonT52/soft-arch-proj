"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { createReview } from "@/actions/create-review";
import { updateReview } from "@/actions/update-review";
import { LexicalComposer } from "@lexical/react/LexicalComposer";
import { ContentEditable } from "@lexical/react/LexicalContentEditable";
import LexicalErrorBoundary from "@lexical/react/LexicalErrorBoundary";
import { HistoryPlugin } from "@lexical/react/LexicalHistoryPlugin";
import { LinkPlugin } from "@lexical/react/LexicalLinkPlugin";
import { ListPlugin } from "@lexical/react/LexicalListPlugin";
import { OnChangePlugin } from "@lexical/react/LexicalOnChangePlugin";
import { RichTextPlugin } from "@lexical/react/LexicalRichTextPlugin";
import {
  Dialog,
  DialogContent,
  DialogOverlay,
  DialogPortal,
  DialogTrigger,
} from "@radix-ui/react-dialog";
import { Loader2Icon } from "lucide-react";
import { type Review } from "@/types/base/review";
import { editorConfig, initialEditorState } from "@/lib/lexical";
import { CodeHighlightPlugin } from "./lexical/code-highlight-plugin";
import { ListMaxIndentLevelPlugin } from "./lexical/list-max-index-level-plugin";
import { ToolbarPlugin } from "./lexical/toolbar-plugin";
import { Rating } from "./rating";
import { Button } from "./ui/button";
import { Label } from "./ui/label";
import { ScrollArea } from "./ui/scroll-area";
import { Switch } from "./ui/switch";
import { useToast } from "./ui/toaster";

type ReviewDialogProps = {
  review?: Review & {
    id: string;
  };
  companyId: string;
  open?: boolean;
  onOpenChange?: (open: boolean) => void;
};

const ReviewDialog = ({
  review,
  companyId,
  open,
  onOpenChange,
}: ReviewDialogProps) => {
  const router = useRouter();
  const { toast } = useToast();

  const [title, setTitle] = useState<string | undefined>(review?.title);
  const [description, setDescription] = useState<string | undefined>(
    review?.description
  );
  const [rating, setRating] = useState<number | undefined>(review?.rating);
  const [anonymous, setAnonymous] = useState(false);
  const [submitting, setSubmitting] = useState(false);

  const handleSubmit = async () => {
    setSubmitting(true);
    const response = review
      ? await updateReview(review.id, {
          review: {
            title: title ?? "Untitled Review",
            description: description ?? initialEditorState,
            rating: rating!,
            isAnonymous: anonymous,
          },
        })
      : await createReview({
          review: {
            cid: companyId,
            title: title ?? "Untitled Review",
            description: description ?? initialEditorState,
            rating: rating!,
            isAnonymous: anonymous,
          },
        });
    if (["200", "201"].includes(response.status)) {
      toast({
        title: "Success",
        description: response.message,
      });
      router.refresh();
    } else {
      toast({
        title: "Error",
        description: response.message,
        variant: "destructive",
      });
    }
    setSubmitting(false);
  };

  return (
    <Dialog
      open={open}
      onOpenChange={(open) => {
        if (open) {
          setRating((prev) => prev ?? 5);
        } else {
          setRating(0);
        }
        onOpenChange?.(open);
      }}
    >
      <DialogTrigger className="h-5">
        <Rating rating={rating} onRatingChange={setRating} />
      </DialogTrigger>
      <DialogPortal>
        <DialogOverlay className="fixed inset-0 z-50 bg-background/80 backdrop-blur-sm data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0" />
        <DialogContent className="fixed left-[50%] top-[50%] z-50 flex max-h-[90%] w-full max-w-lg translate-x-[-50%] translate-y-[-50%] flex-col border bg-background p-0 shadow-lg duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%] sm:rounded-lg md:w-full">
          <LexicalComposer
            initialConfig={{ ...editorConfig, editorState: description }}
          >
            <div className="relative flex flex-col items-start rounded-t-xl pb-6">
              <ToolbarPlugin
                className="w-full rounded-t-[inherit] border-b px-4"
                variant="minimal"
              />
              <div className="flex w-full gap-6 px-8 pt-6">
                <div className="flex flex-1 flex-col gap-4 overflow-x-hidden">
                  <Rating rating={rating} onRatingChange={setRating} />
                  <p className="truncate text-sm text-muted-foreground">
                    to Pizza Company
                  </p>
                </div>
                <div className="flex items-center gap-2">
                  <Switch
                    id="anonymous"
                    checked={anonymous}
                    onCheckedChange={setAnonymous}
                  />
                  <Label htmlFor="anonymous">Anonymous</Label>
                </div>
              </div>
              <input
                className="flex w-full max-w-none appearance-none items-end bg-transparent px-8 pb-[0.6em] pt-[1.6em] text-xl font-semibold leading-[1.6] tracking-tight text-[#111827] scrollbar-hide focus:outline-none focus-visible:outline-none dark:text-[#ffffff]"
                value={title}
                placeholder="Untitled Review"
                autoFocus
                spellCheck={false}
                onChange={(e) => void setTitle(e.target.value)}
                aria-label="title"
              />
              <div className="relative w-full">
                <RichTextPlugin
                  contentEditable={
                    <ScrollArea className="h-32">
                      <ContentEditable
                        className="prose prose-green relative max-w-none flex-1 overflow-auto px-8 dark:prose-invert focus:outline-none prose-p:m-0"
                        spellCheck={false}
                      />
                    </ScrollArea>
                  }
                  placeholder={
                    <div className="prose prose-green pointer-events-none absolute left-0 right-0 top-0 max-w-none px-8 pb-8 prose-p:m-0">
                      <p className="text-muted-foreground">
                        Enter review description...
                      </p>
                    </div>
                  }
                  ErrorBoundary={LexicalErrorBoundary}
                />
              </div>
            </div>
            <HistoryPlugin />
            <CodeHighlightPlugin />
            <ListPlugin />
            <LinkPlugin />
            <ListMaxIndentLevelPlugin maxDepth={1} />
            <OnChangePlugin
              onChange={(editorState) => {
                setDescription(JSON.stringify(editorState.toJSON()));
              }}
            />
          </LexicalComposer>
          <div className="flex justify-center px-8 pb-6 sm:space-x-2">
            <Button
              disabled={submitting}
              onClick={() => void handleSubmit()}
              size="sm"
            >
              {submitting && (
                <Loader2Icon className="mr-2 h-4 w-4 animate-spin" />
              )}
              Review
            </Button>
          </div>
        </DialogContent>
      </DialogPortal>
    </Dialog>
  );
};

export { ReviewDialog };
