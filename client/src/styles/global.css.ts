import { globalStyle } from "@vanilla-extract/css";

import { vars } from ".";

globalStyle("body", {
  backgroundColor: vars.semantic.background.secondary,
  fontFamily:
    'X, -apple-system, BlinkMacSystemFont, "Segoe UI", "Roboto", "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue", sans-serif',
  margin: 0,
  padding: 0,
  color: vars.semantic.text.primary,

  minHeight: "100dvh",
  height: "100%",
});

globalStyle("html", {
  scrollPaddingTop: "50dvh",
  scrollBehavior: "smooth",
});

globalStyle("*", {
  margin: 0,
});
