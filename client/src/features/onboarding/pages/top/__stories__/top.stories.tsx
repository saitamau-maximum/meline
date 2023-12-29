import { DESKTOP_STORY_CONFIG } from "@/__stories__/config";
import { Meta } from "@storybook/react";
import { OnboardingPageTemplate } from "../template";

const meta = {
  title: "Features/Onboarding/Top",
} satisfies Meta;

export default meta;

export const Overview = () => <OnboardingPageTemplate />;
Overview.story = DESKTOP_STORY_CONFIG;
