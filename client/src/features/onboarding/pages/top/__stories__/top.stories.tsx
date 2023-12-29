import { DESKTOP_STORY_CONFIG } from "@/__stories__/config";
import { OnboardingPage } from "../page";

import { Meta } from "@storybook/react";

const meta = {
  title: "Pages/Onboarding",
} satisfies Meta;

export default meta;

export const Overview = () => <OnboardingPage />;
Overview.story = DESKTOP_STORY_CONFIG;
