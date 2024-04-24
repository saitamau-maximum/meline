import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

const channelListItemHeight = vars.spacing[8];

export const styles = {
  channelList: style({
    width: "100%",
    height: "100%",
    display: "flex",
    flexDirection: "column",
    gap: vars.spacing[2],
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
    padding: `0 ${vars.spacing[2]}`,
    boxSizing: "border-box",
    borderRadius: vars.spacing[2],
    height: channelListItemHeight,

    selectors: {
      "&:hover": {
        background: vars.color.gray[3],
      },
    },
  }),
  channelListItemSkeleton: style({
    width: "100%",
    height: vars.spacing[8],
    background: vars.color.gray[4],
    borderRadius: vars.spacing[2],
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
