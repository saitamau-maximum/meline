import { style } from "@vanilla-extract/css";
import { vars } from "@/styles";

const avatarSize = "32px";
const avatarContentGap = vars.spacing[4];

export const styles = {
  container: style({
    display: "flex",
    flexDirection: "column",
    gap: vars.spacing[6],
    padding: vars.spacing[4],
    boxSizing: "border-box",
    overflowY: "auto",
    height: "100%",
  }),
  messageContainer: style({
    display: "flex",
    flexDirection: "column",
    gap: vars.spacing[2],
  }),
  replyMessage: style({
    display: "flex",
    alignItems: "center",
    gap: vars.spacing[2],
    marginLeft: `calc(${avatarSize} + ${avatarContentGap})`,
    borderRadius: vars.spacing[2],
  }),
  replyAvatar: style({
    width: "24px",
    height: "24px",
    borderRadius: vars.spacing.full,
  }),
  replyAuthor: style({
    color: vars.color.gray[10],
    fontSize: vars.font.size.sm,
  }),
  replyText: style({
    color: vars.color.gray[11],
    fontSize: vars.font.size.base,
  }),
  message: style({
    display: "flex",
    gap: avatarContentGap,
  }),
  avatar: style({
    width: avatarSize,
    height: avatarSize,
    borderRadius: vars.spacing.full,
  }),
  messageContent: style({
    display: "flex",
    flexDirection: "column",
    gap: vars.spacing[1],
  }),
  messageAuthor: style({
    color: vars.color.gray[11],
    fontSize: vars.font.size.sm,
  }),
  messageText: style({
    color: vars.color.gray[12],
    fontSize: vars.font.size.base,
    whiteSpace: "pre-wrap",
  }),
};
