import { style } from "@vanilla-extract/css";

export const styles = {
  container: style({
    display: "flex",
    alignItems: "center",
    width: "100%",
    height: "100vh",
    justifyContent: "center",
    "@media": {
      "screen and (max-width: 768px)": {
        flexDirection: "column",
      },
    },
  }),
};
