import { Meta } from "@storybook/react";
import { ChannelList } from "..";

const meta = {
  title: "Features/Channel/Components/ChannelList",
} satisfies Meta;

export default meta;

const mockChannels = [
  {
    id: 1,
    name: "general",
  },
  {
    id: 2,
    name: "random",
  },
  {
    id: 3,
    name: "random2",
  },
];

export const Overview = () => <ChannelList channels={mockChannels} />;

export const WithActiveChannel = () => (
  <ChannelList
    channels={[
      mockChannels[0],
      { ...mockChannels[1], active: true },
      mockChannels[2],
    ]}
  />
);

export const WithNotification = () => (
  <ChannelList
    channels={[
      mockChannels[0],
      { ...mockChannels[1], hasNotification: true },
      mockChannels[2],
    ]}
  />
);
