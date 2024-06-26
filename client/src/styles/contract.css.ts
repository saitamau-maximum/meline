import { createThemeContract } from "@vanilla-extract/css";

export const colorVars = createThemeContract({
  gray: {
    1: null,
    2: null,
    3: null,
    4: null,
    5: null,
    6: null,
    7: null,
    8: null,
    9: null,
    10: null,
    11: null,
    12: null,
  },
  green: {
    1: null,
    2: null,
    3: null,
    4: null,
    5: null,
    6: null,
    7: null,
    8: null,
    9: null,
    10: null,
    11: null,
    12: null,
  },
  red: {
    1: null,
    2: null,
    3: null,
    4: null,
    5: null,
    6: null,
    7: null,
    8: null,
    9: null,
    10: null,
    11: null,
    12: null,
  },
  gradient: {
    primary: null,
  },
});

export const semanticVars = createThemeContract({
  background: {
    primary: null,
    primaryHover: null,
    secondary: null,
    disabled: null,
  },
  text: {
    primary: null,
    secondary: null,
    weak: null,
    weaker: null,
    error: null,
  },
  border: {
    primary: null,
    error: null,
    focus: null,
  },
  button: {
    primary: {
      backgroundColor: null,
      color: null,
    },
  },
});
