import { style } from "@vanilla-extract/css";
import { vars } from "./styles/theme.css";

export const styles = {
  container: style({
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "center",
    minHeight: "100dvh",
    height: vars.spacing.full,
    gap: vars.spacing[4],
  }),
  title: style({
    fontSize: vars.font.size["2xl"],
    color: vars.color.gray[12],
  }),
  subtitle: style({
    fontSize: vars.font.size.base,
    color: vars.color.gray[11],
  }),
};
