import { style, styleVariants } from "@vanilla-extract/css";
import { vars } from "@/styles";

export const styles = {
  button: style({
    display: "inline-flex",
    alignItems: "center",
    justifyContent: "center",
    fontSize: vars.font.size.base,
    fontWeight: 600,
    borderRadius: vars.spacing[2],
    cursor: "pointer",
    transition: vars.transition.fastInteraction("background-color"),
    userSelect: "none",
    whiteSpace: "nowrap",
    textOverflow: "ellipsis",
    textDecoration: "none",
    border: "none",
    outline: "none",
  }),
  buttonVariant: styleVariants({
    primary: {
      backgroundColor: vars.semantic.button.primary.backgroundColor,
      color: vars.semantic.button.primary.color,
      transition: vars.transition.normal("opacity"),

      selectors: {
        "&:hover": {
          opacity: 0.8,
        },
        "&:focus-visible": {
          outline: `2px solid ${vars.color.green[9]}`,
          outlineOffset: "2px",
        },
      },
    },
    secondary: {
      backgroundColor: vars.semantic.background.primary,
      color: vars.semantic.text.secondary,
      boxShadow: `0 0 0 1px ${vars.semantic.border.primary} inset`,

      selectors: {
        "&:hover": {
          backgroundColor: vars.semantic.background.primaryHover,
        },
        "&:focus-visible": {
          outline: `2px solid ${vars.color.green[9]}`,
          outlineOffset: "2px",
        },
      },
    },
    text: {
      backgroundColor: "transparent",
      color: vars.semantic.text.secondary,

      selectors: {
        "&:hover": {
          backgroundColor: vars.semantic.background.primaryHover,
        },
        "&:focus-visible": {
          outline: `2px solid ${vars.color.green[9]}`,
          outlineOffset: "2px",
        },
      },
    },
  }),
  buttonSize: styleVariants({
    sm: {
      fontSize: vars.font.size.sm,
      lineHeight: 1,
      padding: `${vars.spacing[1]} ${vars.spacing[1]}`,
    },
    md: {
      fontSize: vars.font.size.base,
      lineHeight: 1.25,
      padding: `${vars.spacing[2]} ${vars.spacing[4]}`,
    },
  }),
};
