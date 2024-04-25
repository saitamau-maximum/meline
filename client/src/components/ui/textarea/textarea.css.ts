import { vars } from "@/styles";
import { globalStyle, style } from "@vanilla-extract/css";

export const styles = {
  textAreaContainer: style({
    display: "flex",
    flexDirection: "column",
    width: "100%",
  }),
  textAreaLabel: style({
    color: vars.semantic.text.primary,
    fontSize: "1rem",
    marginBottom: vars.spacing[2],
  }),
  textAreaLabelRequired: style({
    marginLeft: vars.spacing[1],
    color: vars.semantic.text.error,
  }),
  textArea: style({
    width: "100%",
    boxSizing: "border-box",
    padding: `${vars.spacing[2]} ${vars.spacing[3]}`,
    borderWidth: "1px",
    borderStyle: "solid",
    borderColor: vars.semantic.border.primary,
    borderRadius: vars.spacing[1],
    backgroundColor: vars.semantic.background.secondary,
    color: vars.semantic.text.primary,
    fontSize: vars.font.size.base,
    transition: vars.transition.normal("border-color", "box-shadow"),
    resize: "none",

    ":focus": {
      outline: "none",
      borderColor: "transparent",
      boxShadow: `0 0 0 2px ${vars.semantic.border.focus}`,
    },
    "::placeholder": {
      color: vars.semantic.text.weak,
    },
    ":disabled": {
      backgroundColor: vars.semantic.background.disabled,
      cursor: "not-allowed",
    },
  }),
  textAreaError: style({
    borderColor: vars.semantic.text.error,
    boxShadow: `0 0 0 0.5px ${vars.semantic.border.error}`,
  }),
  textAreaErrorText: style({
    color: vars.semantic.text.error,
    fontSize: vars.font.size.xs,
    marginTop: vars.spacing[2],
  }),
  textAreaDescription: style({
    color: vars.semantic.text.weak,
    fontSize: vars.font.size.xs,
    marginTop: vars.spacing[2],
  }),
};

globalStyle(`${styles.textArea}:disabled::placeholder`, {
  color: vars.semantic.text.weaker,
});
