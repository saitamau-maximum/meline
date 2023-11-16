import { vars } from "@/styles/theme.css";
import { style } from "@vanilla-extract/css";

export const styles = {
  container: style({
    backgroundColor: vars.color.gray[1],
    height: vars.spacing.full,
    width: vars.spacing.full,
    position: "relative",
    flex: 3,
  }),
  content: style({
    display: "flex",
    width: "fit-content",
    height: "fit-content",
    flexDirection: "column",
    alignItems: "center",
    gap: vars.spacing[6],
    margin: "auto 128px auto auto",
    position: "absolute",
    inset: 0,
  }),
  title: style({
    fontSize: vars.font.size["4xl"],
    color: vars.color.gray[12],
  }),
  description: style({
    fontSize: vars.font.size.base,
    color: vars.color.gray[11],
    textAlign: "center",
  }),
};
