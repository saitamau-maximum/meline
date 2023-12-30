import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

export const styles = {
  createChannelForm: style({
    display: "flex",
    flexDirection: "column",
    width: "100%",
    gap: vars.spacing[4],
  }),
};
