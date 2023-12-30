import { Meta } from "@storybook/react";
import { ChannelTopPageTemplate } from "../template";
import { DESKTOP_STORY_CONFIG } from "@/__stories__/config";
import { IChannelRepository } from "@/repositories/channel";

const meta = {
  title: "Features/Channel/Top",
} satisfies Meta;

export default meta;

const mockUser = {
  name: "test",
  imageURL: "https://example.com",
};

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

const mockChannelRepository: IChannelRepository = {
  createChannel: async () => {
    await new Promise((resolve) => setTimeout(resolve, 300));
    return new Response(null, { status: 200 });
  },
};

const fetchJoinedChannels = async () => {
  await new Promise((resolve) => setTimeout(resolve, 300));
};

export const Overview = () => (
  <ChannelTopPageTemplate
    user={mockUser}
    channels={mockChannels}
    channelRepository={mockChannelRepository}
    fetchJoinedChannels={fetchJoinedChannels}
  />
);
Overview.story = DESKTOP_STORY_CONFIG;
