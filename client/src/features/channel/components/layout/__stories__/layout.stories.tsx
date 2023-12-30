import { Meta } from "@storybook/react";
import { ChannelLayout } from "..";

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
