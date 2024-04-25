import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

export const styles = {
  chatForm: style({
    display: "flex",
    alignItems: "flex-end",
    gap: vars.spacing[2],
  }),
};
