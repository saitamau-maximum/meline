import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

export const styles = {
  container: style({
    display: "flex",
    alignItems: "center",
    gap: vars.spacing[4],
  }),
  metaContainer: style({
    display: "flex",
    flexDirection: "column",
  }),
  avatar: style({
    width: "40px",
    height: "40px",
    borderRadius: "50%",
  }),
  username: style({
    fontWeight: "bold",
    fontSize: vars.font.size.sm,
    color: vars.color.gray[12],
  }),
  message: style({
    color: vars.color.gray[11],
    fontSize: vars.font.size.base,
  }),
};
