import { createGlobalTheme } from "@vanilla-extract/css";

import { colorVars } from "./contract.css";

createGlobalTheme(".light", colorVars, {
  gray: {
    1: "hsl(0 0% 99.0%)",
    2: "hsl(0 0% 97.3%)",
    3: "hsl(0 0% 95.1%)",
    4: "hsl(0 0% 93.0%)",
    5: "hsl(0 0% 90.9%)",
    6: "hsl(0 0% 88.7%)",
    7: "hsl(0 0% 85.8%)",
    8: "hsl(0 0% 78.0%)",
    9: "hsl(0 0% 56.1%)",
    10: "hsl(0 0% 52.3%)",
    11: "hsl(0 0% 43.5%)",
    12: "hsl(0 0% 9.0%)",
  },
  green: {
    1: "hsl(134 44% 99.0%)",
    2: "hsl(134 44% 97.3%)",
    3: "hsl(134 44% 95.1%)",
    4: "hsl(134 44% 93.0%)",
    5: "hsl(134 44% 90.9%)",
    6: "hsl(134 44% 88.7%)",
    7: "hsl(134 44% 85.8%)",
    8: "hsl(134 44% 78.0%)",
    9: "hsl(134 44% 56.1%)",
    10: "hsl(134 44% 52.3%)",
    11: "hsl(134 44% 43.5%)",
    12: "hsl(134 44% 34.0%)",
  },
  gradient: {
    primary: "linear-gradient(291deg, #63C178 0%, #34AA8E 100%))",
  },
});

createGlobalTheme(".dark", colorVars, {
  gray: {
    1: "hsl(0 0% 8.5%)",
    2: "hsl(0 0% 11.0%)",
    3: "hsl(0 0% 13.6%)",
    4: "hsl(0 0% 15.8%)",
    5: "hsl(0 0% 17.9%)",
    6: "hsl(0 0% 20.5%)",
    7: "hsl(0 0% 24.3%)",
    8: "hsl(0 0% 31.2%)",
    9: "hsl(0 0% 43.9%)",
    10: "hsl(0 0% 49.4%)",
    11: "hsl(0 0% 62.8%)",
    12: "hsl(0 0% 93.0%)",
  },
  green: {
    1: "hsl(134 44% 8.5%)",
    2: "hsl(134 44% 11.0%)",
    3: "hsl(134 44% 13.6%)",
    4: "hsl(134 44% 15.8%)",
    5: "hsl(134 44% 17.9%)",
    6: "hsl(134 44% 20.5%)",
    7: "hsl(134 44% 24.3%)",
    8: "hsl(134 44% 31.2%)",
    9: "hsl(134 44% 43.9%)",
    10: "hsl(134 44% 49.4%)",
    11: "hsl(134 44% 62.8%)",
    12: "hsl(134 44% 93.0%)",
  },
  gradient: {
    primary: "linear-gradient(291deg, #34AA8E 0%, #63C178 100%))",
  },
});

const fontVars = createGlobalTheme(":root", {
  size: {
    xs: "0.75rem",
    sm: "0.875rem",
    base: "1rem",
    lg: "1.125rem",
    xl: "1.25rem",
    "2xl": "1.5rem",
    "3xl": "2rem",
    "4xl": "3rem",
  },
});

const spacingVars = createGlobalTheme(":root", {
  none: "0",
  1: "0.25rem",
  2: "0.5rem",
  3: "0.75rem",
  4: "1rem",
  6: "1.5rem",
  8: "2rem",
  10: "2.5rem",
  12: "3rem",
  16: "4rem",
  full: "100%",
});

const zIndexVars = createGlobalTheme(":root", {
  normal: "0",
  forward: "1",
  float: "10",
  windowFloat: "100",
  modal: "1000",
  overlay: "10000",
});

const breakpointVars =  {
  mobile: "768px",
  tablet: "1024px",
};


/**
 * @deprecated 
 * styles/theme.cssではなく、styles/indexをimportしてください
 */
export const vars = {
  color: colorVars,
  font: fontVars,
  spacing: spacingVars,
  zIndex: zIndexVars,
  breakpoint: breakpointVars,
};
