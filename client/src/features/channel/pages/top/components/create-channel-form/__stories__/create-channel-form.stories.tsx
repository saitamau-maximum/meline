import { Meta } from "@storybook/react";
import { CreateChannelForm } from "..";
import { IChannelRepository } from "@/repositories/channel";

const meta = {
  title: "Features/Channel/Top/Components/CreateChannelForm",
} satisfies Meta;

export default meta;

const mockChannelRepository: IChannelRepository = {
  createChannel: async () => {
    await new Promise((resolve) => setTimeout(resolve, 300));
    return new Response(null, { status: 200 });
  },
};

export const Overview = () => (
  <CreateChannelForm
    channelRepository={mockChannelRepository}
    fetchJoinedChannels={async () => {
      await new Promise((resolve) => setTimeout(resolve, 300));
    }}
  />
);
