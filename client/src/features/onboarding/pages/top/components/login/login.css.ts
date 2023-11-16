import { vars } from "@/styles/theme.css";
import { style } from "@vanilla-extract/css";

export const styles = {
  container: style({
    height: vars.spacing.full,
    width: vars.spacing.full,
    position: "relative",
    flex: 2,
  }),
  content: style({
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    gap: vars.spacing[4],
    padding: vars.spacing[4],
    textDecoration: "none",
    color: vars.color.gray[12],
    width: "fit-content",
    height: "fit-content",
    margin: "auto auto auto 48px",
    position: "absolute",
    inset: 0,
  }),
  logo: style({
    width: "100%",
    maxWidth: "24px",
    height: "auto",
  }),
  loginLabel: style({
    fontSize: vars.font.size["base"],
    lineHeight: 1,
  }),
};
