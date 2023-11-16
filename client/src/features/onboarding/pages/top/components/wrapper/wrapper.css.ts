import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

export const styles = {
  container: style({
    display: "flex",
    alignItems: "center",
    width: "100%",
    height: "100vh",
    justifyContent: "center",
    "@media": {
      [`screen and (max-width: ${vars.breakpoint.mobile})`]: {
        flexDirection: "column",
      },
    },
  }),
};
