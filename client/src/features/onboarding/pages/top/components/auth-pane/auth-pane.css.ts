import { vars } from "@/styles";
import { style } from "@vanilla-extract/css";

export const styles = {
  container: style({
    height: vars.spacing.full,
    width: vars.spacing.full,
    position: "relative",
    borderLeft: `1px solid ${vars.color.gray[6]}`,
    flex: 2,
    "@media": {
      [`screen and (max-width: ${vars.breakpoint.mobile})`]: {
        flex: 1,
        backgroundColor: vars.color.gray[1],
        borderLeft: "none",
      },
    },
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
    borderRadius: vars.spacing[2],
    ":hover": {
      backgroundColor: vars.color.gray[4],
      transition: vars.transition.normal("background-color"),
    },
    "@media": {
      [`screen and (max-width: ${vars.breakpoint.mobile})`]: {
        margin: "16px auto auto auto",
      },
    },
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
