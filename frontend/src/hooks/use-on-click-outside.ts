import { useEffect, type RefObject } from "react";

type Handler = (event: MouseEvent) => void;

export function useOnClickOutside<T extends HTMLElement = HTMLElement>(
  ref: RefObject<T>,
  handler: Handler,
  mouseEvent: "mousedown" | "mouseup" = "mousedown"
) {
  useEffect(() => {
    const _handler = (event: MouseEvent) => {
      const elem = ref?.current;

      // Do nothing if clicking ref's element or descendent elements
      if (!elem || elem.contains(event.target as Node)) {
        return;
      }

      handler(event);
    };

    window.addEventListener(mouseEvent, _handler);

    return () => {
      window.removeEventListener(mouseEvent, _handler);
    };
  }, [ref, handler, mouseEvent]);
}
