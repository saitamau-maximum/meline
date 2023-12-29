import { constants, vars } from "@/styles";
import { style } from "@vanilla-extract/css";

export const styles = {
  channelLayoutWrapper: style({
    display: "flex",
    width: "100%",
    height: "100dvh",
  }),
  channelLayoutSidePanel: style({
    width: constants.sizes.channelLayoutSidePanelWidth,
    flexShrink: 0,
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    gap: vars.spacing[4],
    padding: vars.spacing[4],
    boxSizing: "border-box",
  }),
  channelLayoutSidePanelLogo: style({
    color: vars.color.gray[11],
    fontSize: vars.font.size["xl"],
    padding: vars.spacing[4],
  }),
  channelLayoutMain: style({
    flexGrow: 1,
    background: vars.color.gray[2],
    borderLeft: `1px solid ${vars.color.gray[6]}`,
  }),
};
