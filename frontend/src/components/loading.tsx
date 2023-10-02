type LoadingProps = {
  size?: number;
};

const Loading = ({ size = 4 }: LoadingProps) => {
  return (
    <div
      className="relative h-[var(--size)] w-[var(--size)] animate-[fade-in_0.2s_0.2s_both,_dot-elastic_1s_infinite_linear] rounded-[calc(var(--size)/2)] bg-primary ease-out before:absolute before:left-[calc(-3/2*var(--size))] before:inline-block before:h-[var(--size)] before:w-[var(--size)] before:animate-[dot-elastic-before_1s_infinite_linear] before:rounded-[calc(var(--size)/2)] before:bg-primary after:absolute after:left-[calc(3/2*var(--size))] after:inline-block after:h-[var(--size)] after:w-[var(--size)] after:animate-[dot-elastic-after_1s_infinite_linear] after:rounded-[calc(var(--size)/2)] after:bg-primary"
      style={{
        ["--size" as string]: `${size / 4}rem`,
      }}
    />
  );
};

export { Loading };
