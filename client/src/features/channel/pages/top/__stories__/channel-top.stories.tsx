import { Meta } from "@storybook/react";
import { ChannelTopPageTemplate } from "../template";
import { DESKTOP_STORY_CONFIG } from "@/__stories__/config";

const meta = {
  title: "Features/Channel/Top",
} satisfies Meta;

export default meta;

const mockUser = {
  name: "test",
  imageURL: "https://example.com",
};

export const Overview = () => <ChannelTopPageTemplate user={mockUser} />;
Overview.story = DESKTOP_STORY_CONFIG;
