import { vars } from "@/styles";
import { style } from "@vanilla-extract/css";

export const styles = {
  container: style({
    backgroundColor: vars.color.gray[1],
    height: vars.spacing.full,
    width: vars.spacing.full,
    position: "relative",
    flex: 3,
    "@media": {
      [`screen and (max-width: ${vars.breakpoint.mobile})`]: {
        flex: 1,
      },
    },
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
    "@media": {
      [`screen and (max-width: ${vars.breakpoint.mobile})`]: {
        margin: "auto auto 16px auto",
      },
    },
  }),
  title: style({
    fontSize: vars.font.size["4xl"],
    color: vars.color.gray[12],
    "@media": {
      [`screen and (max-width: ${vars.breakpoint.mobile})`]: {
        fontSize: vars.font.size["3xl"],
      },
    },
  }),
  description: style({
    fontSize: vars.font.size.base,
    color: vars.color.gray[11],
    textAlign: "center",
    "@media": {
      [`screen and (max-width: ${vars.breakpoint.mobile})`]: {
        fontSize: vars.font.size.sm,
      },
    },
  }),
  icon: style({
    position: "fixed",
    top: vars.spacing[8],
    left: vars.spacing[8],
    fontSize: vars.font.size.xl,
    color: vars.color.gray[11],
  }),
};
