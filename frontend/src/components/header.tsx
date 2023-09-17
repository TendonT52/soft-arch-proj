import { Logo } from "./logo";

const Header = () => {
  return (
    <header className="sticky left-0 right-0 top-0 z-10 bg-background/50 backdrop-blur-xl backdrop-saturate-150">
      <div className="container flex h-20 items-center justify-between">
        <Logo />
      </div>
    </header>
  );
};

export { Header };
