"use client";

import { useEffect, useState } from "react";
import { StarIcon } from "lucide-react";
import { cn } from "@/lib/utils";

type RatingScore = 1 | 2 | 3 | 4 | 5;

type RatingProps = {
  rating?: RatingScore;
  onRatingChange?: (rating?: RatingScore) => void;
};

const Rating = ({ rating, onRatingChange }: RatingProps) => {
  const [_rating, _setRating] = useState<RatingScore | undefined>(rating);
  const [_displayRating, _setDisplayRating] = useState<
    RatingScore | undefined
  >();

  useEffect(() => {
    _setRating(rating);
  }, [rating]);

  useEffect(() => {
    if (onRatingChange) {
      onRatingChange(_rating);
    }
  }, [_rating, rating, onRatingChange]);

  return (
    <div className="inline-flex">
      {Array.from({ length: 5 }, (_, idx) => {
        const rating = (1 + idx) as RatingScore;
        return (
          <StarIcon
            key={`score${rating}`}
            className={cn(
              "h-5 w-5 cursor-pointer",
              (_displayRating ?? _rating ?? 0) >= rating && "fill-primary"
            )}
            strokeWidth={1}
            onMouseOver={() => void _setDisplayRating(rating)}
            onMouseOut={() => void _setDisplayRating(undefined)}
            onClick={() => void _setRating(rating)}
          />
        );
      })}
    </div>
  );
};

export { Rating, type RatingScore };
