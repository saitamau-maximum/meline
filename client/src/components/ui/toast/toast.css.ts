import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

export const styles = {
  toastStackContainer: style({
    position: "fixed",
    top: "0",
    right: "0",
    padding: vars.spacing[4],
    zIndex: vars.zIndex.windowFloat,
    display: "flex",
    flexDirection: "column",
    gap: vars.spacing[2],
  }),
  toastItemContainer: style({
    display: "flex",
    width: "240px",
    padding: vars.spacing[2],
    borderRadius: vars.spacing[2],
    backgroundColor: vars.color.gray[3],
    boxShadow: `0 0 4px ${vars.color.gray[1]}`,
    border: `1px solid ${vars.color.gray[5]}`,
  }),
  link: style({
    textDecoration: "none",
    color: "inherit",
    transition: "background-color 0.2s",

    ":hover": {
      backgroundColor: vars.color.gray[4],
    },
  }),
};
