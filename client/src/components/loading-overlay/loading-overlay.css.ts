import { vars } from "@/styles/theme.css";
import { style, createVar } from "@vanilla-extract/css";

const loadingObjectSizeVar = createVar();

export const styles = {
  overlay: style({
    position: "fixed",
    top: 0,
    left: 0,
    width: "100%",
    height: "100%",
    zIndex: vars.zIndex.modal,
    backgroundColor: vars.color.gray[1],
    display: "flex",
    justifyContent: "center",
    alignItems: "center",
    flexDirection: "column",
    opacity: 0,
    pointerEvents: "none",
  }),
  loading: style({
    vars: {
      [loadingObjectSizeVar]: "3rem",
    },
    width: loadingObjectSizeVar,
    height: loadingObjectSizeVar,
  }),
  active: style({
    opacity: 1,
    pointerEvents: "auto",
  }),
  fade: style({
    transition: "opacity 0.5s ease-in-out",
  }),
};
