import { type Config } from "tailwindcss";
import { fontFamily } from "tailwindcss/defaultTheme";
import plugin from "tailwindcss/plugin";

const config: Config = {
  content: [
    "./src/app/**/*.{js,ts,jsx,tsx,mdx}",
    "./src/components/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  darkMode: ["class"],
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      colors: {
        border: "hsl(var(--border))",
        input: "hsl(var(--input))",
        ring: "hsl(var(--ring))",
        background: "hsl(var(--background))",
        foreground: "hsl(var(--foreground))",
        primary: {
          DEFAULT: "hsl(var(--primary))",
          foreground: "hsl(var(--primary-foreground))",
        },
        secondary: {
          DEFAULT: "hsl(var(--secondary))",
          foreground: "hsl(var(--secondary-foreground))",
        },
        destructive: {
          DEFAULT: "hsl(var(--destructive))",
          foreground: "hsl(var(--destructive-foreground))",
        },
        muted: {
          DEFAULT: "hsl(var(--muted))",
          foreground: "hsl(var(--muted-foreground))",
        },
        accent: {
          DEFAULT: "hsl(var(--accent))",
          foreground: "hsl(var(--accent-foreground))",
        },
        popover: {
          DEFAULT: "hsl(var(--popover))",
          foreground: "hsl(var(--popover-foreground))",
        },
        card: {
          DEFAULT: "hsl(var(--card))",
          foreground: "hsl(var(--card-foreground))",
        },
        code: {
          DEFAULT: "hsl(var(--code))",
          attribute: "hsl(var(--code-attribute))",
          comment: "hsl(var(--code-comment))",
          foreground: "hsl(var(--code-foreground))",
          function: "hsl(var(--code-function))",
          operator: "hsl(var(--code-operator))",
          property: "hsl(var(--code-property))",
          punctuation: "hsl(var(--code-punctuation))",
          selector: "hsl(var(--code-selector))",
          variable: "hsl(var(--code-variable))",
        },
      },
      borderRadius: {
        lg: `var(--radius)`,
        md: `calc(var(--radius) - 2px)`,
        sm: "calc(var(--radius) - 4px)",
      },
      fontFamily: {
        sans: [
          "var(--font-sans)",
          {
            ...fontFamily.sans,
            fontFeatureSettings: '"cv02", "cv03", "cv04", "cv11"',
          },
        ],
        mono: ["var(--font-mono)", ...fontFamily.mono],
      },
      keyframes: {
        "accordion-down": {
          from: { height: "0" },
          to: { height: "var(--radix-accordion-content-height)" },
        },
        "accordion-up": {
          from: { height: "var(--radix-accordion-content-height)" },
          to: { height: "0" },
        },
        "dot-elastic": {
          "0%": {
            transform: "scale(1, 1)",
          },
          "25%": {
            transform: "scale(1, 1)",
          },
          "50%": {
            transform: "scale(1, 1.5)",
          },
          "75%": {
            transform: "scale(1, 1)",
          },
          "100%": {
            transform: "scale(1, 1)",
          },
        },
        "dot-elastic-before": {
          "0%": {
            transform: "scale(1, 1)",
          },
          "25%": {
            transform: "scale(1, 1.5)",
          },
          "50%": {
            transform: "scale(1, 0.67)",
          },
          "75%": {
            transform: "scale(1, 1)",
          },
          "100%": {
            transform: "scale(1, 1)",
          },
        },
        "dot-elastic-after": {
          "0%": {
            transform: "scale(1, 1)",
          },
          "25%": {
            transform: "scale(1, 1)",
          },
          "50%": {
            transform: "scale(1, 0.67)",
          },
          "75%": {
            transform: "scale(1, 1.5)",
          },
          "100%": {
            transform: "scale(1, 1)",
          },
        },
        "fade-in": {
          "0%": {
            opacity: "0",
          },
          "100%": {
            opacity: "1",
          },
        },
      },
      animation: {
        "accordion-down": "accordion-down 0.2s ease-out",
        "accordion-up": "accordion-up 0.2s ease-out",
      },
    },
  },
  plugins: [
    require("tailwindcss-animate"),
    require("@tailwindcss/typography"),
    plugin(({ addUtilities }) => {
      addUtilities({
        ".scrollbar-hide": {
          /* IE and Edge */
          "-ms-overflow-style": "none",

          /* Firefox */
          "scrollbar-width": "none",

          /* Safari and Chrome */
          "&::-webkit-scrollbar": {
            display: "none",
          },
        },
      });
    }),
  ],
};

export default config;
