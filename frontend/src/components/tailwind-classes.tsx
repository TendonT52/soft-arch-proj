/**
 * To build styles of some class names
 */

const TailwindClasses = () => {
  return (
    <>
      <div className="underline" />
      <div className="line-through" />
      <div className="italic" />
      <div className="block overflow-x-auto whitespace-pre rounded-sm bg-muted p-3 font-mono text-sm scrollbar-hide before:content-none after:content-none" />
      <div className="font-normal not-italic text-muted-foreground" />
      <div className="text-code-attribute" />
      <div className="text-code-property" />
      <div className="text-code-selector" />
      <div className="text-code-comment" />
      <div className="text-code-function" />
      <div className="text-code-variable" />
      <div className="text-code-operator" />
      <div className="text-code-punctuation" />
    </>
  );
};

export { TailwindClasses };
