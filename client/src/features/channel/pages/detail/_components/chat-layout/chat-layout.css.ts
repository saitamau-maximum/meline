import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

export const styles = {
  chatLayout: style({
    display: "grid",
    gridTemplateRows: "1fr auto",
    height: "100dvh",
  }),
  chatLayoutContent: style({
    minHeight: 0,
  }),
  chatLayoutFooter: style({
    padding: vars.spacing[4],
    paddingTop: 0,
  }),
};
