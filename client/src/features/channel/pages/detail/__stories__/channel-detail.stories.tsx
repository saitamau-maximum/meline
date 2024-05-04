import { Meta } from "@storybook/react";
import { ChannelDetailPageTemplate } from "../template";
import { DESKTOP_STORY_CONFIG } from "@/__stories__/config";
import { Message } from "@/repositories/message";

const meta = {
  title: "Features/Channel/Detail",
} satisfies Meta;

export default meta;

const MockMessages: Message[] = [
  {
    id: "1",
    content: "Hello, world!",
    user: {
      id: 1,
      name: "John Doe",
      image_url: "https://i.pravatar.cc/32?img=1",
    },
    created_at: "2021-09-01T00:00:00Z",
    updated_at: "2021-09-01T00:00:00Z",
    reply_to_message: null,
  },
  {
    id: "2",
    content: "Hi, there!",
    user: {
      id: 2,
      name: "Jane Doe",
      image_url: "https://i.pravatar.cc/32?img=2",
    },
    created_at: "2021-09-01T00:00:00Z",
    updated_at: "2021-09-01T00:00:00Z",
    reply_to_message: null,
  },
  {
    id: "3",
    content: "How are you?",
    user: {
      id: 3,
      name: "Alice",
      image_url: "https://i.pravatar.cc/32?img=3",
    },
    created_at: "2021-09-01T00:00:00Z",
    updated_at: "2021-09-01T00:00:00Z",
    reply_to_message: {
      id: "1",
      content: "Hello, world!",
      user: {
        id: 1,
        name: "John Doe",
        image_url: "https://i.pravatar.cc/32?img=1",
      },
    },
  },
];

export const Overview = () => (
  <ChannelDetailPageTemplate channelId={0} messages={MockMessages} />
);
Overview.story = DESKTOP_STORY_CONFIG;
