import { vars } from "@/styles/theme.css";
import { style, keyframes } from "@vanilla-extract/css";

const loadingAnimation = keyframes({
  "0%": { transform: "rotate(0deg)" },
  "100%": { transform: "rotate(360deg)" },
});

export const styles = {
  overlay: style({
    position: "fixed",
    top: 0,
    left: 0,
    width: "100%",
    height: "100%",
    zIndex: vars.zIndex.modal,
  }),
  loading: style({
    position: "absolute",
    top: "50%",
    left: "50%",
    width: "5rem",
    height: "5rem",
    margin: "-2.5rem 0 0 -2.5rem",
    border: `0.5rem solid ${vars.color.green[12]}`,
    borderRadius: "50%",
    borderTopColor: vars.color.green[3],
    animation: `${loadingAnimation} 1s linear infinite`,
  }),
};
