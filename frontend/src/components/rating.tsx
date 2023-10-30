"use client";

import { useEffect, useState } from "react";
import { StarIcon } from "lucide-react";
import { cn } from "@/lib/utils";

type RatingProps = {
  rating?: number;
  onRatingChange?: (rating: number) => void;
  editable?: boolean;
};

const Rating = ({ rating, onRatingChange, editable = true }: RatingProps) => {
  const [_rating, _setRating] = useState<number | undefined>(rating);
  const [_displayRating, _setDisplayRating] = useState<number | undefined>();

  // prevent race conditions
  useEffect(() => {
    if (rating !== _rating) {
      _setRating(rating);
    }
  }, [_rating, rating]);

  return (
    <div className="inline-flex">
      {Array.from({ length: 5 }, (_, idx) => {
        const rating = 1 + idx;
        return (
          <StarIcon
            key={rating}
            className={cn(
              "h-5 w-5",
              editable && "cursor-pointer",
              (_displayRating ?? _rating ?? 0) >= rating && "fill-primary"
            )}
            strokeWidth={1}
            onMouseOver={() => {
              if (editable) {
                _setDisplayRating(rating);
              }
            }}
            onMouseOut={() => {
              if (editable) {
                _setDisplayRating(undefined);
              }
            }}
            onClick={() => {
              if (editable) {
                _setRating(rating);
                onRatingChange?.(rating);
              }
            }}
          />
        );
      })}
    </div>
  );
};

export { Rating };
