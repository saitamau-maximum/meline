import { Meta } from "@storybook/react";
import { ChannelLayout } from "..";

const meta = {
  title: "Features/Channel/Components/Layout",
} satisfies Meta;

export default meta;

const mockChannels = [
  {
    id: 1,
    name: "Channel 1",
  },
  {
    id: 2,
    name: "Channel 2",
  },
];

export const Overview = () => <ChannelLayout channels={mockChannels} />;
