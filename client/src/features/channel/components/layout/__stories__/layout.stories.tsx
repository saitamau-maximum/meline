import { Meta } from "@storybook/react";
import { ChannelLayout } from "..";
import { DESKTOP_STORY_CONFIG } from "@/__stories__/config";

const meta = {
  title: "Features/Channel/Components/Layout",
} satisfies Meta;

export default meta;

export const Overview = () => (
  <ChannelLayout
    sidePanel={<span>Channel List</span>}
    main={<span>Channel Chat</span>}
  />
);
Overview.story = DESKTOP_STORY_CONFIG;
