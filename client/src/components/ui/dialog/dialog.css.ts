import { keyframes, style } from "@vanilla-extract/css";
import { constants, vars } from "@/styles";

const overlayKF = keyframes({
  from: {
    opacity: 0,
  },
  to: {
    opacity: 0.9,
  },
});

const contentKF = keyframes({
  from: {
    opacity: 0,
    transform: "translate(-50%, -48%) scale(0.96)",
  },
  to: {
    opacity: 1,
    transform: "translate(-50%, -50%) scale(1)",
  },
});

export const styles = {
  dialogOverlay: style({
    backgroundColor: vars.semantic.background.secondary,
    position: "fixed",
    inset: 0,
    opacity: 0.9,
    animation: vars.transition.fastInteraction(overlayKF),
  }),
  dialogContent: style({
    backgroundColor: vars.semantic.background.primary,
    border: `1px solid ${vars.color.gray[6]}`,
    borderRadius: vars.spacing[2],
    position: "fixed",
    top: "50%",
    left: "50%",
    transform: "translate(-50%, -50%)",
    width: `calc(100vw - ${vars.spacing[4]} * 2)`,
    maxWidth: constants.sizes.dialogMaxWidth,
    maxHeight: "85vh",
    padding: vars.spacing[4],
    animation: vars.transition.fastInteraction(contentKF),
    boxSizing: "border-box",

    selectors: {
      "&:focus": {
        outline: "none",
      },
    },
  }),
  dialogTitle: style({
    fontSize: vars.font.size.lg,
    fontWeight: 600,
    marginBottom: vars.spacing[2],
    color: vars.color.gray[12],
  }),
  dialogDescription: style({
    fontSize: vars.font.size.sm,
    marginBottom: vars.spacing[4],
    color: vars.color.gray[11],
  }),
  dialogClose: style({
    position: "absolute",
    top: vars.spacing[3],
    right: vars.spacing[3],
    padding: vars.spacing[1],
    border: "none",
    borderRadius: vars.spacing[16],
    background: "transparent",
    cursor: "pointer",
    color: vars.color.gray[11],
    lineHeight: 0,
    transition: vars.transition.normal("background"),

    selectors: {
      "&:focus": {
        outline: "none",
      },
      "&:hover": {
        background: vars.color.gray[4],
      },
    },
  }),
  dialogRightTopCloseButton: style({
    position: "absolute",
    top: vars.spacing[3],
    right: vars.spacing[3],
  }),
  dialogFooter: style({
    display: "flex",
    justifyContent: "flex-end",
    marginTop: vars.spacing[4],
    gap: vars.spacing[4],
  }),
};
