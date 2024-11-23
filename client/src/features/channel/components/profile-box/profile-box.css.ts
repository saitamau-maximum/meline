import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

export const styles = {
  profileBox: style({
    display: "flex",
    alignItems: "center",
    width: "100%",
    gap: vars.spacing[4],
  }),
  iconUrl: style({
    width: "40px",
    height: "40px",
    borderRadius: "50%",
  }),
  profileName: style({
    fontSize: vars.font.size.base,
  }),
};
