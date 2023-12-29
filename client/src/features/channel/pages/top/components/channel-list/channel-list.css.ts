import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

export const styles = {
  channelList: style({
    width: "100%",
    height: "100%",
    display: "flex",
    flexDirection: "column",
    gap: vars.spacing[2],
    padding: vars.spacing[4],
    boxSizing: "border-box",
    overflowY: "auto",
  }),
  channelListItem: style({
    display: "flex",
    gap: vars.spacing[2],
    alignItems: "center",
    cursor: "pointer",
    color: vars.color.gray[11],
    textDecoration: "none",
    transition: vars.transition.normal("background"),
    background: "transparent",
    padding: `${vars.spacing[1]} ${vars.spacing[2]}`,
    boxSizing: "border-box",
    borderRadius: vars.spacing[2],

    selectors: {
      "&:hover": {
        background: vars.color.gray[3],
      },
    },
  }),
  channelListItemActive: style({
    color: vars.color.gray[12],
    background: vars.color.gray[4],
    fontWeight: 600,
  }),
  channelListItemNotification: style({
    color: vars.color.gray[12],
  }),
  channelListItemIcon: style({
    strokeWidth: 3,
    width: vars.spacing[4],
    flexShrink: 0,
  }),
};
